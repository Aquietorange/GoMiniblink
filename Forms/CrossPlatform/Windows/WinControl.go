package Windows

import (
	"GoMiniblink/Forms/CrossPlatform/Windows/win32"
)

type winControl struct {
	provider   *Provider
	className  string
	idName     string
	handle     win32.HWND
	isCreated  bool
	defWndProc func(hWnd win32.HWND, msg uint32, wParam uintptr, lParam uintptr) uintptr
	evCreate   []func(IWindow)
	owner      *winForm
	x          int
	y          int
	width      int
	height     int
}

func (_this *winControl) onCreate(hWnd win32.HWND) {
	_this.handle = hWnd
	for _, v := range _this.evCreate {
		v(_this)
	}
	_this.evCreate = nil
	_this.isCreated = true
}

func (_this *winControl) fireWndProc(hWnd win32.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	return 0
}

func (_this *winControl) addEvCreate(fn func(IWindow)) {
	_this.evCreate = append(_this.evCreate, fn)
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
