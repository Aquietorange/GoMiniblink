package windows

import (
	p "qq2564874169/goMiniblink/forms/platform"
	"qq2564874169/goMiniblink/forms/platform/windows/win32"
	"unsafe"
)

type winControl struct {
	winBase

	createParams *win32.CREATESTRUCT
}

func (_this *winControl) init(provider *Provider) *winControl {
	_this.winBase.init(provider)
	win32.CreateWindowEx(
		_this.createParams.ExStyle,
		(*uint16)(unsafe.Pointer(sto16(provider.className))),
		(*uint16)(unsafe.Pointer(sto16(""))),
		win32.WS_CHILD, 0, 0, 0, 0, 0, 0,
		provider.hInstance, unsafe.Pointer(_this))
	return _this
}

func (_this *winControl) SetParent(window p.Window) {
	_this.createParams.Parent = win32.HWND(window.GetHandle())
}

func (_this *winControl) GetHandle() uintptr {
	return uintptr(_this.handle)
}

func (_this *winControl) Hide() {
	win32.ShowWindow(_this.hWnd(), win32.SW_HIDE)
}

func (_this *winControl) Show() {
	win32.ShowWindow(_this.hWnd(), win32.SW_SHOW)
	win32.UpdateWindow(_this.hWnd())
}

func (_this *winControl) SetSize(width, height int) {
	win32.SetWindowPos(_this.hWnd(), 0, 0, 0, int32(width), int32(height), win32.SWP_NOMOVE|win32.SWP_NOZORDER)
}

func (_this *winControl) SetLocation(x, y int) {
	win32.SetWindowPos(_this.hWnd(), 0, int32(x), int32(y), 0, 0, win32.SWP_NOSIZE|win32.SWP_NOZORDER)
}
