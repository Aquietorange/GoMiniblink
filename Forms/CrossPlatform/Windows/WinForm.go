package Windows

import (
	"GoMiniblink/Forms/CrossPlatform/Windows/win32"
	"unsafe"
)

type winForm struct {
	winControl
	onResize func(int, int)
	onMove   func(int, int)
	hCaption int
}

func (_this *winForm) hWnd() win32.HWND {
	return _this.hwnd
}

func (_this *winForm) class() string {
	return _this.className
}

func (_this *winForm) name() string {
	return _this.idName
}

func (_this *winForm) init(provider *Provider) *winForm {
	_this.idName = newUUID()
	_this.className = provider.className
	_this.onWndProc = _this.OnWndProc
	provider.add(_this)
	_this.owner = _this
	_this.hwnd = win32.CreateWindowEx(
		win32.WS_EX_CLIENTEDGE,
		sto16(provider.className),
		sto16(newUUID()),
		win32.WS_OVERLAPPEDWINDOW,
		int32(_this.x),
		int32(_this.y),
		int32(_this.width),
		int32(_this.height),
		0, 0,
		provider.hInstance,
		unsafe.Pointer(nil))
	return _this
}

func (_this *winForm) OnWndProc(hWnd win32.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {
	case win32.WM_CREATE:
		_this.hCaption = int(win32.GetSystemMetrics(win32.SM_CYCAPTION))
		break
	case win32.WM_DESTROY:
		win32.PostQuitMessage(0)
		break
	case win32.WM_SIZE:
		w, h := win32.LOWORD(uint32(lParam)), win32.HIWORD(uint32(lParam))
		if _this.onResize != nil {
			_this.onResize(int(w), int(h))
		}
		break
	case win32.WM_MOVE:
		x, y := win32.LOWORD(uint32(lParam)), win32.HIWORD(uint32(lParam))
		if _this.onMove != nil {
			_this.onMove(int(x), int(y))
		}
		break
	}
	return 0
}

func (_this *winForm) wndProc(hWnd win32.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	return _this.OnWndProc(hWnd, msg, wParam, lParam)
}

func (_this *winForm) Show() {
	win32.ShowWindow(_this.hwnd, win32.SW_SHOW)
	win32.UpdateWindow(_this.hwnd)
}

func (_this *winForm) Hide() {
	win32.ShowWindow(_this.hwnd, win32.SW_SHOW)
	win32.UpdateWindow(_this.hwnd)
}

func (_this *winForm) SetSize(w, h int) {
	_this.width, _this.height = w, h
	win32.MoveWindow(_this.hwnd, int32(_this.x), int32(_this.y), int32(_this.width), int32(_this.height), false)
}

func (_this *winForm) OnResize(fn func(w, h int)) {
	_this.onResize = fn
}

func (_this *winForm) SetLocation(x, y int) {
	_this.x, _this.y = x, y
	win32.MoveWindow(_this.hwnd, int32(_this.x), int32(_this.y), int32(_this.width), int32(_this.height), false)
}

func (_this *winForm) OnMove(fn func(x, y int)) {
	_this.onMove = fn
}

func (_this *winForm) SetTitle(title string) {

}
func (_this *winForm) ShowDialog() {

}
