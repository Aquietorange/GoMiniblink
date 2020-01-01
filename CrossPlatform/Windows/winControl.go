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
	evMouseMove  map[string]func(target interface{}, e MB.MouseEvArgs)
	evMouseDown  map[string]func(target interface{}, e MB.MouseEvArgs)
	evMouseUp    map[string]func(target interface{}, e MB.MouseEvArgs)
	evMouseWheel map[string]func(target interface{}, e MB.MouseEvArgs)
	evMouseClick map[string]func(target interface{}, e MB.MouseEvArgs)

	onMouseMove  func(MB.MouseEvArgs)
	onMouseDown  func(MB.MouseEvArgs)
	onMouseUp    func(MB.MouseEvArgs)
	onMouseWheel func(MB.MouseEvArgs)
	onMouseClick func(MB.MouseEvArgs)
}

func (_this *winControl) init() {
	_this.evWndCreate = make(map[string]func(win32.HWND))
	_this.invokeMap = make(map[string]*InvokeContext)
	_this.evWndProc = make(map[string]func(win32.HWND, uint32, uintptr, uintptr) uintptr)
	_this.evMouseMove = make(map[string]func(interface{}, MB.MouseEvArgs))
	_this.evMouseDown = make(map[string]func(interface{}, MB.MouseEvArgs))
	_this.evMouseUp = make(map[string]func(interface{}, MB.MouseEvArgs))
	_this.evMouseWheel = make(map[string]func(interface{}, MB.MouseEvArgs))
	_this.evMouseClick = make(map[string]func(interface{}, MB.MouseEvArgs))
	_this.SetOnMouseMove(_this.defOnMouseMove)
	_this.SetOnMouseDown(_this.defOnMouseDown)
	_this.SetOnMouseUp(_this.defOnMouseUp)
	_this.SetOnMouseWheel(_this.defOnMouseWheel)
	_this.SetOnMouseClick(_this.defOnMouseClick)
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
	case win32.WM_MOUSEMOVE:
		var btns MB.MouseButtons
		wp := int(wParam)
		if wp&win32.MK_LBUTTON != 0 {
			btns |= MB.MouseButtons_Left
		}
		if wp&win32.MK_MBUTTON != 0 {
			btns |= MB.MouseButtons_Middle
		}
		if wp&win32.MK_RBUTTON != 0 {
			btns |= MB.MouseButtons_Right
		}
		x, y := win32.GET_X_LPARAM(lParam), win32.GET_Y_LPARAM(lParam)
		_this.onMouseMove(MB.MouseEvArgs{
			X:       int(x),
			Y:       int(y),
			Buttons: btns,
			Time:    time.Now(),
		})
	case win32.WM_LBUTTONDOWN, win32.WM_RBUTTONDOWN, win32.WM_MBUTTONDOWN:
		x, y := int(win32.GET_X_LPARAM(lParam)), int(win32.GET_Y_LPARAM(lParam))
		var btns MB.MouseButtons
		switch msg {
		case win32.WM_LBUTTONDOWN:
			btns |= MB.MouseButtons_Left
		case win32.WM_RBUTTONDOWN:
			btns |= MB.MouseButtons_Right
		case win32.WM_MBUTTONDOWN:
			btns |= MB.MouseButtons_Middle
		}
		_this.onMouseDown(MB.MouseEvArgs{
			X:       x,
			Y:       y,
			Buttons: btns,
			Time:    time.Now(),
		})
	case win32.WM_LBUTTONUP, win32.WM_RBUTTONUP, win32.WM_MBUTTONUP:
		x, y := int(win32.GET_X_LPARAM(lParam)), int(win32.GET_Y_LPARAM(lParam))
		var btns MB.MouseButtons
		switch msg {
		case win32.WM_LBUTTONUP:
			btns |= MB.MouseButtons_Left
		case win32.WM_RBUTTONUP:
			btns |= MB.MouseButtons_Right
		case win32.WM_MBUTTONUP:
			btns |= MB.MouseButtons_Middle
		}
		_this.onMouseUp(MB.MouseEvArgs{
			X:       x,
			Y:       y,
			Buttons: btns,
			Time:    time.Now(),
		})
	case win32.WM_MOUSEWHEEL:
		x, y := win32.GET_X_LPARAM(lParam), win32.GET_Y_LPARAM(lParam)
		lp, hp := win32.LOWORD(int32(wParam)), win32.HIWORD(int32(wParam))
		var btns MB.MouseButtons
		if lp&win32.MK_LBUTTON != 0 {
			btns |= MB.MouseButtons_Left
		}
		if lp&win32.MK_MBUTTON != 0 {
			btns |= MB.MouseButtons_Middle
		}
		if lp&win32.MK_RBUTTON != 0 {
			btns |= MB.MouseButtons_Right
		}
		_this.onMouseWheel(MB.MouseEvArgs{
			Buttons: btns,
			X:       int(x),
			Y:       int(y),
			Delta:   int(hp),
			Time:    time.Now(),
		})
	}
	return 0
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
