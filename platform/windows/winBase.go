package windows

import (
	"image"
	"image/draw"
	mb "qq2564874169/goMiniblink"
	plat "qq2564874169/goMiniblink/platform"
	"qq2564874169/goMiniblink/platform/windows/win32"
	"sync"
	"time"
	"unsafe"
)

type winBase struct {
	app          *Provider
	idName       string
	handle       win32.HWND
	isCreated    bool
	thisIsDialog bool
	invokeMap    sync.Map
	onWndProc    windowsMsgProc

	onCreate              plat.WindowCreateProc
	onDestroy             plat.WindowDestroyProc
	onResize              plat.WindowResizeProc
	onMove                plat.WindowMoveProc
	onMouseMove           plat.WindowMouseMoveProc
	onMouseDown           plat.WindowMouseDownProc
	onMouseUp             plat.WindowMouseUpProc
	onMouseWheel          plat.WindowMouseWheelProc
	onMouseClick          plat.WindowMouseClickProc
	onPaint               plat.WindowPaintProc
	onKeyDown             plat.WindowKeyDownProc
	onKeyUp               plat.WindowKeyUpProc
	onKeyPress            plat.WindowKeyPressProc
	onSetCursor           plat.WindowSetCursorProc
	onImeStartComposition plat.WindowImeStartCompositionProc
	onFocus               plat.WindowFocusProc
	onLostFocus           plat.WindowLostFocusProc

	bgColor int
}

func (_this *winBase) init(provider *Provider, id string) *winBase {
	_this.app = provider
	_this.idName = id
	_this.SetBgColor(provider.defBgColor)
	return _this
}

func (_this *winBase) SetBgColor(color int) {
	_this.bgColor = color
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

func (_this *winBase) SetCursor(cursor mb.CursorType) {
	res := win32.MAKEINTRESOURCE(uintptr(toWinCursor(cursor)))
	cur := win32.LoadCursor(0, res)
	win32.SetCursor(cur)
}

func (_this *winBase) msgProc(hWnd win32.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	ret := int(_this.execCmd(hWnd, msg, wParam, lParam))
	if ret != 0 {
		return uintptr(ret)
	}
	switch msg {
	case win32.WM_KILLFOCUS:
		if _this.onLostFocus != nil && _this.onLostFocus() {
			ret = 1
		}
	case win32.WM_SETFOCUS:
		if _this.onFocus != nil && _this.onFocus() {
			ret = 1
		}
	case win32.WM_IME_STARTCOMPOSITION:
		if _this.onImeStartComposition != nil && _this.onImeStartComposition() {
			ret = 1
		}
	case win32.WM_SETCURSOR:
		if _this.onSetCursor != nil && _this.onSetCursor() {
			ret = 1
		}
	case win32.WM_SIZE:
		if _this.onResize != nil {
			w, h := win32.GET_X_LPARAM(lParam), win32.GET_Y_LPARAM(lParam)
			if _this.onResize(mb.Rect{
				Width:  int(w),
				Height: int(h),
			}) {
				ret = 1
			}
		}
	case win32.WM_MOVE:
		if _this.onMove != nil {
			x, y := win32.GET_X_LPARAM(lParam), win32.GET_Y_LPARAM(lParam)
			if _this.onMove(mb.Point{
				X: int(x),
				Y: int(y),
			}) {
				ret = 1
			}
		}
	case win32.WM_DESTROY:
		if _this.onDestroy != nil {
			_this.onDestroy()
		}
		_this.app.remove(_this.hWnd(), true)
	case win32.WM_SYSKEYDOWN, win32.WM_KEYDOWN:
		key := vkToKey(int(wParam))
		if _this.onKeyDown != nil && key != mb.Keys_Error {
			e := mb.KeyEvArgs{
				Key:   key,
				Value: wParam,
				IsSys: msg == win32.WM_SYSKEYDOWN,
			}
			if _this.onKeyDown(&e) || e.IsHandle {
				ret = 1
			}
		}
	case win32.WM_SYSKEYUP, win32.WM_KEYUP:
		key := vkToKey(int(wParam))
		if _this.onKeyUp != nil && key != mb.Keys_Error {
			e := mb.KeyEvArgs{
				Key:   key,
				Value: wParam,
				IsSys: msg == win32.WM_SYSKEYUP,
			}
			if _this.onKeyUp(&e) || e.IsHandle {
				ret = 1
			}
		}
	case win32.WM_SYSCHAR, win32.WM_CHAR:
		if _this.onKeyPress != nil {
			e := mb.KeyPressEvArgs{
				KeyChar: string(wParam),
				Value:   wParam,
				IsSys:   msg == win32.WM_SYSCHAR,
			}
			if _this.onKeyPress(&e) || e.IsHandle {
				ret = 1
			}
		}
	case win32.WM_ERASEBKGND:
		//lbp := win32.LOGBRUSH{
		//	LbStyle: win32.BS_SOLID,
		//	LbColor: win32.COLORREF(_this.bgColor),
		//}
		//b := win32.CreateBrushIndirect(&lbp)
		//rect := new(win32.RECT)
		//win32.GetClientRect(hWnd, rect)
		//win32.FillRect(win32.HDC(wParam), rect, b)
		//win32.DeleteObject(win32.HGDIOBJ(b))
		return 1
	case win32.WM_PAINT:
		ps := new(win32.PAINTSTRUCT)
		hdc := win32.BeginPaint(hWnd, ps)
		e := mb.PaintEvArgs{
			Clip: mb.Bound{
				Point: mb.Point{
					X: int(ps.RcPaint.Left),
					Y: int(ps.RcPaint.Top),
				},
				Rect: mb.Rect{
					Width:  int(ps.RcPaint.Right - ps.RcPaint.Left),
					Height: int(ps.RcPaint.Bottom - ps.RcPaint.Top),
				},
			},
		}
		if e.Clip.IsEmpty() == false {
			e.Graphics = new(winGraphics).init(hdc)
			if _this.bgColor >= 0 {
				bg := image.NewRGBA(image.Rect(0, 0, e.Clip.Width, e.Clip.Height))
				draw.Draw(bg, bg.Rect, image.NewUniform(mb.IntToRGBA(_this.bgColor)), image.Pt(0, 0), draw.Src)
				e.Graphics.DrawImage(bg, 0, 0, e.Clip.Width, e.Clip.Height, e.Clip.X, e.Clip.Y)
			}
			if _this.onPaint != nil {
				if _this.onPaint(e) {
					ret = 1
				}
			}
			e.Graphics.Close()
		}
		win32.EndPaint(hWnd, ps)
	case win32.WM_MOUSEMOVE:
		if _this.onMouseMove != nil {
			e := mb.MouseEvArgs{
				X:    int(win32.GET_X_LPARAM(lParam)),
				Y:    int(win32.GET_Y_LPARAM(lParam)),
				Time: time.Now(),
			}
			wp := int(wParam)
			if wp&win32.MK_LBUTTON != 0 {
				e.Button |= mb.MouseButtons_Left
			}
			if wp&win32.MK_MBUTTON != 0 {
				e.Button |= mb.MouseButtons_Middle
			}
			if wp&win32.MK_RBUTTON != 0 {
				e.Button |= mb.MouseButtons_Right
			}
			if _this.onMouseMove(e) {
				ret = 1
			}
		}
	case win32.WM_LBUTTONDOWN, win32.WM_RBUTTONDOWN, win32.WM_MBUTTONDOWN:
		if _this.onMouseDown != nil {
			e := mb.MouseEvArgs{
				X:    int(win32.GET_X_LPARAM(lParam)),
				Y:    int(win32.GET_Y_LPARAM(lParam)),
				Time: time.Now(),
			}
			switch msg {
			case win32.WM_LBUTTONDOWN:
				e.Button |= mb.MouseButtons_Left
			case win32.WM_RBUTTONDOWN:
				e.Button |= mb.MouseButtons_Right
			case win32.WM_MBUTTONDOWN:
				e.Button |= mb.MouseButtons_Middle
			}
			if _this.onMouseDown(e) {
				ret = 1
			}
		}
	case win32.WM_LBUTTONUP, win32.WM_RBUTTONUP, win32.WM_MBUTTONUP:
		if _this.onMouseUp != nil {
			e := mb.MouseEvArgs{
				X:    int(win32.GET_X_LPARAM(lParam)),
				Y:    int(win32.GET_Y_LPARAM(lParam)),
				Time: time.Now(),
			}
			switch msg {
			case win32.WM_LBUTTONDOWN:
				e.Button |= mb.MouseButtons_Left
			case win32.WM_RBUTTONDOWN:
				e.Button |= mb.MouseButtons_Right
			case win32.WM_MBUTTONDOWN:
				e.Button |= mb.MouseButtons_Middle
			}
			if _this.onMouseUp(e) {
				ret = 1
			}
		}
	case win32.WM_MOUSEWHEEL:
		if _this.onMouseWheel != nil {
			lp, hp := win32.LOWORD(int32(wParam)), win32.HIWORD(int32(wParam))
			e := mb.MouseEvArgs{
				X:     int(win32.GET_X_LPARAM(lParam)),
				Y:     int(win32.GET_Y_LPARAM(lParam)),
				Delta: int(hp),
				Time:  time.Now(),
			}
			if lp&win32.MK_LBUTTON != 0 {
				e.Button |= mb.MouseButtons_Left
			}
			if lp&win32.MK_MBUTTON != 0 {
				e.Button |= mb.MouseButtons_Middle
			}
			if lp&win32.MK_RBUTTON != 0 {
				e.Button |= mb.MouseButtons_Right
			}
			if _this.onMouseWheel(e) {
				ret = 1
			}
		}
	default:
		return 0
	}
	return uintptr(ret)
}

func (_this *winBase) CreateGraphics() mb.Graphics {
	hdc := win32.GetDC(_this.hWnd())
	g := new(winGraphics).init(hdc)
	g.onClose = func(e *winGraphics) {
		win32.ReleaseDC(_this.hWnd(), e.dc)
	}
	return g
}

func (_this *winBase) Invoke(fn func(state interface{}), state interface{}) {
	ctx := InvokeContext{
		fn:    fn,
		state: state,
	}
	ptr := uintptr(unsafe.Pointer(&ctx))
	_this.invokeMap.Store(ptr, &ctx)
	win32.PostMessage(_this.hWnd(), uint32(win32.WM_COMMAND), uintptr(cmd_invoke), ptr)
}

func (_this *winBase) execCmd(hWnd win32.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	if msg != win32.WM_COMMAND {
		return 0
	}
	switch int(wParam) {
	case cmd_invoke:
		ctx := *((*InvokeContext)(unsafe.Pointer(lParam)))
		ctx.fn(ctx.state)
		_this.invokeMap.Delete(lParam)
	case cmd_mouse_click:
		e := *((*mb.MouseEvArgs)(unsafe.Pointer(lParam)))
		_this.onMouseClick(e)
	default:
		return 0
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

func (_this *winBase) GetProvider() plat.IProvider {
	return _this.app
}
