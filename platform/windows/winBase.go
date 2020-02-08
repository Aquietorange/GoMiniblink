package windows

import (
	"image"
	"image/draw"
	mb "qq.2564874169/goMiniblink"
	"qq.2564874169/goMiniblink/platform"
	"qq.2564874169/goMiniblink/platform/windows/win32"
	"time"
	"unsafe"
)

type winBase struct {
	provider     *Provider
	idName       string
	handle       win32.HWND
	isCreated    bool
	thisIsDialog bool
	invokeMap    map[string]*InvokeContext
	onWndProc    windowsMsgProc

	onCreate     platform.WindowCreateProc
	onDestroy    platform.WindowDestroyProc
	onResize     platform.WindowResizeProc
	onMove       platform.WindowMoveProc
	onMouseMove  platform.WindowMouseMoveProc
	onMouseDown  platform.WindowMouseDownProc
	onMouseUp    platform.WindowMouseUpProc
	onMouseWheel platform.WindowMouseWheelProc
	onMouseClick platform.WindowMouseClickProc
	onPaint      platform.WindowPaintProc
	onKeyDown    platform.WindowKeyDownProc
	onKeyUp      platform.WindowKeyUpProc
	onKeyPress   platform.WindowKeyPressProc

	bgColor int
}

func (_this *winBase) init(provider *Provider, id string) *winBase {
	_this.provider = provider
	_this.idName = id
	_this.SetBgColor(provider.defBgColor)
	_this.invokeMap = make(map[string]*InvokeContext)
	return _this
}

func (_this *winBase) SetBgColor(color int) {
	_this.bgColor = color
	//lbp := win32.LOGBRUSH{
	//	LbStyle: win32.BS_SOLID,
	//	LbColor: win32.COLORREF(color),
	//}
	//_this.bgColor = win32.CreateBrushIndirect(&lbp)
}

func (_this *winBase) isDialog() bool {
	return _this.thisIsDialog
}

func (_this *winBase) IsCreate() bool {
	return _this.isCreated
}

func (_this *winBase) getCreateProc() windowsCreateProc {
	return _this.createProc
}

func (_this *winBase) createProc(hWnd win32.HWND) {
	_this.isCreated = true
	_this.handle = hWnd
	if _this.onCreate != nil {
		_this.onCreate(uintptr(hWnd))
	}
}

func (_this *winBase) getWindowMsgProc() windowsMsgProc {
	return _this.msgProc
}

func (_this *winBase) msgProc(hWnd win32.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	ret := _this.execCmd(hWnd, msg, wParam, lParam)
	if ret != 0 {
		return ret
	}
	switch msg {
	case win32.WM_SIZE:
		if _this.onResize != nil {
			w, h := win32.GET_X_LPARAM(lParam), win32.GET_Y_LPARAM(lParam)
			_this.onResize(mb.Rect{
				Width:  int(w),
				Height: int(h),
			})
		}
	case win32.WM_MOVE:
		if _this.onMove != nil {
			x, y := win32.GET_X_LPARAM(lParam), win32.GET_Y_LPARAM(lParam)
			_this.onMove(mb.Point{
				X: int(x),
				Y: int(y),
			})
		}
	case win32.WM_DESTROY:
		if _this.onDestroy != nil {
			_this.onDestroy()
		}
		win32.DeleteObject(win32.HGDIOBJ(_this.bgColor))
		_this.provider.remove(_this.hWnd(), true)
	case win32.WM_SYSKEYDOWN, win32.WM_KEYDOWN:
		key := vkToKey(int(wParam))
		if _this.onKeyDown != nil && key != mb.Keys_Error {
			e := mb.KeyEvArgs{
				Key:        key,
				KeysIsDown: _this.provider.keysIsDown,
			}
			_this.onKeyDown(&e)
			if e.IsHandle {
				return 1
			}
		}
		return 0
	case win32.WM_SYSKEYUP, win32.WM_KEYUP:
		key := vkToKey(int(wParam))
		if _this.onKeyUp != nil && key != mb.Keys_Error {
			e := mb.KeyEvArgs{
				Key:        key,
				KeysIsDown: _this.provider.keysIsDown,
			}
			_this.onKeyUp(&e)
			if e.IsHandle {
				return 1
			}
		}
		return 0
	case win32.WM_SYSCHAR, win32.WM_CHAR:
		if _this.onKeyPress != nil {
			e := mb.KeyPressEvArgs{
				KeyChar:    string(wParam),
				KeysIsDown: _this.provider.keysIsDown,
			}
			_this.onKeyPress(&e)
			if e.IsHandle {
				return 1
			}
		}
		return 0
	case win32.WM_ERASEBKGND:
		//hdc := win32.HDC(wParam)
		//rect := new(win32.RECT)
		//win32.GetClientRect(hWnd, rect)
		//win32.FillRect(hdc, rect, _this.bgColor)
		return 1
	case win32.WM_PAINT:
		pt := win32.PAINTSTRUCT{}
		hdc := win32.BeginPaint(hWnd, &pt)
		e := mb.PaintEvArgs{
			Clip: mb.Bound{
				Point: mb.Point{
					X: int(pt.RcPaint.Left),
					Y: int(pt.RcPaint.Top),
				},
				Rect: mb.Rect{
					Width:  int(pt.RcPaint.Right - pt.RcPaint.Left),
					Height: int(pt.RcPaint.Bottom - pt.RcPaint.Top),
				},
			},
			Graphics: new(winGraphics).init(hdc),
		}
		if e.Clip.IsEmpty() == false {
			w := e.Clip.Width + (4 - e.Clip.Width%4)
			bg := image.NewRGBA(image.Rect(0, 0, w, e.Clip.Height))
			draw.Draw(bg, bg.Rect, image.NewUniform(mb.IntToRGBA(_this.bgColor)), image.Pt(0, 0), draw.Src)
			e.Graphics.DrawImage(bg, 0, 0, e.Clip.Width, e.Clip.Height, e.Clip.X, e.Clip.Y)
			//if _this.onPaint != nil {
			//	_this.onPaint(e)
			//}
		}
		win32.EndPaint(hWnd, &pt)
	case win32.WM_MOUSEMOVE:
		if _this.onMouseMove != nil {
			e := mb.MouseEvArgs{
				X:            int(win32.GET_X_LPARAM(lParam)),
				Y:            int(win32.GET_Y_LPARAM(lParam)),
				ButtonIsDown: make(map[mb.MouseButtons]bool),
				Time:         time.Now(),
			}
			wp := int(wParam)
			if wp&win32.MK_LBUTTON != 0 {
				e.ButtonIsDown[mb.MouseButtons_Left] = true
			}
			if wp&win32.MK_MBUTTON != 0 {
				e.ButtonIsDown[mb.MouseButtons_Middle] = true
			}
			if wp&win32.MK_RBUTTON != 0 {
				e.ButtonIsDown[mb.MouseButtons_Right] = true
			}
			_this.onMouseMove(e)
		}
	case win32.WM_LBUTTONDOWN, win32.WM_RBUTTONDOWN, win32.WM_MBUTTONDOWN:
		if _this.onMouseDown != nil {
			e := mb.MouseEvArgs{
				X:            int(win32.GET_X_LPARAM(lParam)),
				Y:            int(win32.GET_Y_LPARAM(lParam)),
				ButtonIsDown: make(map[mb.MouseButtons]bool),
				Time:         time.Now(),
			}
			switch msg {
			case win32.WM_LBUTTONDOWN:
				e.ButtonIsDown[mb.MouseButtons_Left] = true
			case win32.WM_RBUTTONDOWN:
				e.ButtonIsDown[mb.MouseButtons_Right] = true
			case win32.WM_MBUTTONDOWN:
				e.ButtonIsDown[mb.MouseButtons_Middle] = true
			}
			_this.onMouseDown(e)
		}
	case win32.WM_LBUTTONUP, win32.WM_RBUTTONUP, win32.WM_MBUTTONUP:
		if _this.onMouseUp != nil {
			e := mb.MouseEvArgs{
				X:            int(win32.GET_X_LPARAM(lParam)),
				Y:            int(win32.GET_Y_LPARAM(lParam)),
				ButtonIsDown: make(map[mb.MouseButtons]bool),
				Time:         time.Now(),
			}
			switch msg {
			case win32.WM_LBUTTONUP:
				e.ButtonIsDown[mb.MouseButtons_Left] = true
			case win32.WM_RBUTTONUP:
				e.ButtonIsDown[mb.MouseButtons_Right] = true
			case win32.WM_MBUTTONUP:
				e.ButtonIsDown[mb.MouseButtons_Middle] = true
			}
			_this.onMouseUp(e)
		}
	case win32.WM_MOUSEWHEEL:
		if _this.onMouseWheel != nil {
			lp, hp := win32.LOWORD(int32(wParam)), win32.HIWORD(int32(wParam))
			e := mb.MouseEvArgs{
				X:            int(win32.GET_X_LPARAM(lParam)),
				Y:            int(win32.GET_Y_LPARAM(lParam)),
				Delta:        int(hp),
				ButtonIsDown: make(map[mb.MouseButtons]bool),
				Time:         time.Now(),
			}
			if lp&win32.MK_LBUTTON != 0 {
				e.ButtonIsDown[mb.MouseButtons_Left] = true
			}
			if lp&win32.MK_MBUTTON != 0 {
				e.ButtonIsDown[mb.MouseButtons_Middle] = true
			}
			if lp&win32.MK_RBUTTON != 0 {
				e.ButtonIsDown[mb.MouseButtons_Right] = true
			}
			_this.onMouseWheel(e)
		}
	default:
		return 0
	}
	return 1
}

func (_this *winBase) CreateGraphics() mb.Graphics {
	dc := win32.GetDC(_this.hWnd())
	gdi := new(winGraphics).init(dc)
	gdi.onClose = func() {
		win32.ReleaseDC(_this.hWnd(), dc)
	}
	return gdi
}

func (_this *winBase) Invoke(fn func(state interface{}), state interface{}) {
	ctx := InvokeContext{
		fn:    fn,
		state: state,
		key:   mb.NewUUID(),
	}
	_this.invokeMap[ctx.key] = &ctx
	win32.PostMessage(_this.hWnd(), uint32(win32.WM_COMMAND), uintptr(cmd_invoke), uintptr(unsafe.Pointer(&ctx)))
}

func (_this *winBase) execCmd(hWnd win32.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	if msg != win32.WM_COMMAND {
		return 0
	}
	switch int(wParam) {
	case cmd_invoke:
		ctx := *((*InvokeContext)(unsafe.Pointer(lParam)))
		delete(_this.invokeMap, ctx.key)
		ctx.fn(ctx.state)
	case cmd_mouse_click:
		e := *((*mb.MouseEvArgs)(unsafe.Pointer(lParam)))
		_this.onMouseClick(e)
	}
	return 1
}

func (_this *winBase) hWnd() win32.HWND {
	return _this.handle
}

func (_this *winBase) Id() string {
	return _this.idName
}

func (_this *winBase) id() string {
	return _this.Id()
}
