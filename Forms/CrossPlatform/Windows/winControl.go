package Windows

import (
	"GoMiniblink/Forms/CrossPlatform/Windows/win32"
	"GoMiniblink/Utils"
	"unsafe"
)

type winControl struct {
	provider     *Provider
	className    string
	idName       string
	handle       win32.HWND
	isCreated    bool
	evWndProc    map[string]func(hWnd win32.HWND, msg uint32, wParam uintptr, lParam uintptr) uintptr
	evWndCreate  map[string]func(hWnd win32.HWND)
	invokeCtxMap map[string]*InvokeContext
}

func (_this *winControl) init() {
	_this.evWndCreate = make(map[string]func(hWnd win32.HWND))
	_this.invokeCtxMap = make(map[string]*InvokeContext)
	_this.evWndProc = make(map[string]func(hWnd win32.HWND, msg uint32, wParam uintptr, lParam uintptr) uintptr)
	_this.evWndProc["__execInvoke"] = _this.execInvoke
}

func (_this *winControl) IsCreate() bool {
	return _this.isCreated
}

func (_this *winControl) onWndCreate(hWnd win32.HWND) {
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
	return 0
}

func (_this *winControl) Invoke(fn func(state interface{}), state interface{}) {
	ctx := InvokeContext{
		fn:    fn,
		state: state,
		key:   Utils.NewUUID(),
	}
	_this.invokeCtxMap[ctx.key] = &ctx
	win32.PostMessage(_this.hWnd(), uint32(win32.WM_COMMAND), uintptr(cmd_invoke), uintptr(unsafe.Pointer(&ctx)))
}

func (_this *winControl) execInvoke(hWnd win32.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	if msg != win32.WM_COMMAND || uint(wParam) != cmd_invoke {
		return 0
	}
	ctx := *((*InvokeContext)(unsafe.Pointer(lParam)))
	ctx.fn(ctx.state)
	delete(_this.invokeCtxMap, ctx.key)
	return 0
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
