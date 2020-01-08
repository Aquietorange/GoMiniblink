package Windows

import (
	MB "GoMiniblink"
	"GoMiniblink/CrossPlatform/Windows/win32"
	"GoMiniblink/Utils"
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
	evWndProc    map[string]func(hWnd win32.HWND, msg uint32, wParam uintptr, lParam uintptr) uintptr
	evWndCreate  map[string]func(hWnd win32.HWND)
	evWndDestroy map[string]func()

	onCreate     func()
	onDestroy    func()
	onResize     func(e MB.Rect)
	onMove       func(e MB.Point)
	onMouseMove  func(e MB.MouseEvArgs)
	onMouseDown  func(e MB.MouseEvArgs)
	onMouseUp    func(e MB.MouseEvArgs)
	onMouseWheel func(e MB.MouseEvArgs)
	onMouseClick func(e MB.MouseEvArgs)
	onPaint      func(e MB.PaintEvArgs)
	onKeyDown    func(e *MB.KeyEvArgs)
	onKeyUp      func(e *MB.KeyEvArgs)
	onKeyPress   func(e *MB.KeyPressEvArgs)

	bgColor win32.HBRUSH
}

func (_this *winBase) init(provider *Provider, idName string) *winBase {
	_this.provider = provider
	_this.idName = idName
	_this.evWndCreate = make(map[string]func(win32.HWND))
	_this.invokeMap = make(map[string]*InvokeContext)
	_this.evWndProc = make(map[string]func(win32.HWND, uint32, uintptr, uintptr) uintptr)
	_this.evWndDestroy = make(map[string]func())
	_this.evWndProc["__exec_cmd"] = _this.execCmd
	return _this
}

func (_this *winBase) SetBgColor(color int) {
	if _this.bgColor != 0 {
		win32.DeleteObject(win32.HGDIOBJ(_this.bgColor))
	}
	lbp := win32.LOGBRUSH{
		LbStyle: win32.BS_SOLID,
		LbColor: win32.COLORREF(color),
	}
	_this.bgColor = win32.CreateBrushIndirect(&lbp)
}

func (_this *winBase) isDialog() bool {
	return _this.thisIsDialog
}

func (_this *winBase) IsCreate() bool {
	return _this.isCreated
}

func (_this *winBase) fireWndCreate(hWnd win32.HWND) {
	_this.isCreated = true
	_this.handle = hWnd
	for _, v := range _this.evWndCreate {
		v(hWnd)
	}
	if _this.onCreate != nil {
		_this.onCreate()
	}
}

func (_this *winBase) fireWndProc(hWnd win32.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	for _, v := range _this.evWndProc {
		ret := v(hWnd, msg, wParam, lParam)
		if ret != 0 {
			return ret
		}
	}
	switch msg {
	case win32.WM_DESTROY:
		if _this.onDestroy != nil {
			_this.onDestroy()
		}
		if _this.bgColor != 0 {
			win32.DeleteObject(win32.HGDIOBJ(_this.bgColor))
		}
		for _, v := range _this.evWndDestroy {
			v()
		}
		_this.provider.remove(_this.hWnd(), true)
	case win32.WM_SYSKEYDOWN, win32.WM_KEYDOWN:
		key := vkToKey(int(wParam))
		if _this.onKeyDown != nil && key != MB.Keys_Error {
			e := MB.KeyEvArgs{
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
		if _this.onKeyUp != nil && key != MB.Keys_Error {
			e := MB.KeyEvArgs{
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
			e := MB.KeyPressEvArgs{
				KeyChar:    string(wParam),
				KeysIsDown: _this.provider.keysIsDown,
			}
			_this.onKeyPress(&e)
			if e.IsHandle {
				return 1
			}
		}
		return 0
	case win32.WM_PAINT:
		pt := win32.PAINTSTRUCT{}
		win32.BeginPaint(hWnd, &pt)
		e := MB.PaintEvArgs{
			Update: MB.Bound{
				Point: MB.Point{
					X: int(pt.RcPaint.Left),
					Y: int(pt.RcPaint.Top),
				},
				Rect: MB.Rect{
					Wdith:  int(pt.RcPaint.Right - pt.RcPaint.Left),
					Height: int(pt.RcPaint.Bottom - pt.RcPaint.Top),
				},
			},
			Context: uintptr(pt.Hdc),
		}
		if _this.bgColor != 0 {
			win32.FillRect(pt.Hdc, &pt.RcPaint, _this.bgColor)
		}
		if _this.onPaint != nil {
			_this.onPaint(e)
		}
		win32.EndPaint(hWnd, &pt)
	case win32.WM_MOUSEMOVE:
		if _this.onMouseMove != nil {
			e := MB.MouseEvArgs{
				X:            int(win32.GET_X_LPARAM(lParam)),
				Y:            int(win32.GET_Y_LPARAM(lParam)),
				ButtonIsDown: make(map[MB.MouseButtons]bool),
				Time:         time.Now(),
			}
			wp := int(wParam)
			if wp&win32.MK_LBUTTON != 0 {
				e.ButtonIsDown[MB.MouseButtons_Left] = true
			}
			if wp&win32.MK_MBUTTON != 0 {
				e.ButtonIsDown[MB.MouseButtons_Middle] = true
			}
			if wp&win32.MK_RBUTTON != 0 {
				e.ButtonIsDown[MB.MouseButtons_Right] = true
			}
			_this.onMouseMove(e)
		}
	case win32.WM_LBUTTONDOWN, win32.WM_RBUTTONDOWN, win32.WM_MBUTTONDOWN:
		if _this.onMouseDown != nil {
			e := MB.MouseEvArgs{
				X:            int(win32.GET_X_LPARAM(lParam)),
				Y:            int(win32.GET_Y_LPARAM(lParam)),
				ButtonIsDown: make(map[MB.MouseButtons]bool),
				Time:         time.Now(),
			}
			switch msg {
			case win32.WM_LBUTTONDOWN:
				e.ButtonIsDown[MB.MouseButtons_Left] = true
			case win32.WM_RBUTTONDOWN:
				e.ButtonIsDown[MB.MouseButtons_Right] = true
			case win32.WM_MBUTTONDOWN:
				e.ButtonIsDown[MB.MouseButtons_Middle] = true
			}
			_this.onMouseDown(e)
		}
	case win32.WM_LBUTTONUP, win32.WM_RBUTTONUP, win32.WM_MBUTTONUP:
		if _this.onMouseUp != nil {
			e := MB.MouseEvArgs{
				X:            int(win32.GET_X_LPARAM(lParam)),
				Y:            int(win32.GET_Y_LPARAM(lParam)),
				ButtonIsDown: make(map[MB.MouseButtons]bool),
				Time:         time.Now(),
			}
			switch msg {
			case win32.WM_LBUTTONUP:
				e.ButtonIsDown[MB.MouseButtons_Left] = true
			case win32.WM_RBUTTONUP:
				e.ButtonIsDown[MB.MouseButtons_Right] = true
			case win32.WM_MBUTTONUP:
				e.ButtonIsDown[MB.MouseButtons_Middle] = true
			}
			_this.onMouseUp(e)
		}
	case win32.WM_MOUSEWHEEL:
		if _this.onMouseWheel != nil {
			lp, hp := win32.LOWORD(int32(wParam)), win32.HIWORD(int32(wParam))
			e := MB.MouseEvArgs{
				X:            int(win32.GET_X_LPARAM(lParam)),
				Y:            int(win32.GET_Y_LPARAM(lParam)),
				Delta:        int(hp),
				ButtonIsDown: make(map[MB.MouseButtons]bool),
				Time:         time.Now(),
			}
			if lp&win32.MK_LBUTTON != 0 {
				e.ButtonIsDown[MB.MouseButtons_Left] = true
			}
			if lp&win32.MK_MBUTTON != 0 {
				e.ButtonIsDown[MB.MouseButtons_Middle] = true
			}
			if lp&win32.MK_RBUTTON != 0 {
				e.ButtonIsDown[MB.MouseButtons_Right] = true
			}
			_this.onMouseWheel(e)
		}
	default:
		return 0
	}
	return 1
}

func (_this *winBase) Invoke(fn func(state interface{}), state interface{}) {
	ctx := InvokeContext{
		fn:    fn,
		state: state,
		key:   Utils.NewUUID(),
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
		e := *((*MB.MouseEvArgs)(unsafe.Pointer(lParam)))
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
