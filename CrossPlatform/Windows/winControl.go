package Windows

import (
	MB "GoMiniblink"
	"GoMiniblink/CrossPlatform/Windows/win32"
	"GoMiniblink/Utils"
	"time"
	"unsafe"
)

type winControl struct {
	provider     *Provider
	className    string
	idName       string
	handle       win32.HWND
	isCreated    bool
	thisIsDialog bool
	invokeMap    map[string]*InvokeContext
	evWndProc    map[string]func(hWnd win32.HWND, msg uint32, wParam uintptr, lParam uintptr) uintptr
	evWndCreate  map[string]func(hWnd win32.HWND)

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
}

func (_this *winControl) init() {
	_this.evWndCreate = make(map[string]func(win32.HWND))
	_this.invokeMap = make(map[string]*InvokeContext)
	_this.evWndProc = make(map[string]func(win32.HWND, uint32, uintptr, uintptr) uintptr)
	_this.evWndProc["__exec_cmd"] = _this.execCmd
}

func (_this *winControl) isDialog() bool {
	return _this.thisIsDialog
}

func (_this *winControl) IsCreate() bool {
	return _this.isCreated
}

func (_this *winControl) fireWndCreate(hWnd win32.HWND) {
	_this.isCreated = true
	_this.handle = hWnd
	for _, v := range _this.evWndCreate {
		v(hWnd)
	}
}

func (_this *winControl) fireWndProc(hWnd win32.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	for _, v := range _this.evWndProc {
		ret := v(hWnd, msg, wParam, lParam)
		if ret != 0 {
			return ret
		}
	}
	switch msg {
	case win32.WM_SYSKEYDOWN, win32.WM_KEYDOWN:
		if _this.onKeyDown != nil {
			key := vkToKey(int(wParam))
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
		if _this.onKeyUp != nil {
			key := vkToKey(int(wParam))
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
	case win32.WM_CHAR:
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
		if _this.onPaint != nil {
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
				State: uintptr(pt.Hdc),
			}
			_this.onPaint(e)
		}
		win32.EndPaint(hWnd, &pt)
	case win32.WM_MOUSEMOVE:
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
	case win32.WM_LBUTTONDOWN, win32.WM_RBUTTONDOWN, win32.WM_MBUTTONDOWN:
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
	case win32.WM_LBUTTONUP, win32.WM_RBUTTONUP, win32.WM_MBUTTONUP:
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
	case win32.WM_MOUSEWHEEL:
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
	default:
		return 0
	}
	return 1
}

func (_this *winControl) Invoke(fn func(state interface{}), state interface{}) {
	ctx := InvokeContext{
		fn:    fn,
		state: state,
		key:   Utils.NewUUID(),
	}
	_this.invokeMap[ctx.key] = &ctx
	win32.PostMessage(_this.hWnd(), uint32(win32.WM_COMMAND), uintptr(cmd_invoke), uintptr(unsafe.Pointer(&ctx)))
}

func (_this *winControl) execCmd(hWnd win32.HWND, msg uint32, wParam, lParam uintptr) uintptr {
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

func (_this *winControl) hWnd() win32.HWND {
	return _this.handle
}

func (_this *winControl) class() string {
	return _this.className
}

func (_this *winControl) name() string {
	return _this.idName
}
