package Windows

import (
	"GoMiniblink/CrossPlatform/Windows/win32"
	"GoMiniblink/Utils"
	"unsafe"
)

type winControl struct {
	winBase

	createParams *win32.CREATESTRUCT
	bgColor      win32.HBRUSH
}

func (_this *winControl) init(provider *Provider) *winControl {
	_this.winBase.init(provider, Utils.NewUUID())
	_this.createParams = &win32.CREATESTRUCT{
		Instance:  provider.hInstance,
		ClassName: uintptr(unsafe.Pointer(sto16(provider.className))),
		Name:      uintptr(unsafe.Pointer(sto16(""))),
		Style:     win32.WS_CHILD,
		ExStyle:   0,
	}
	_this.provider.add(_this)
	return _this
}

func (_this *winControl) setParent(hWnd win32.HWND) {
	_this.createParams.Parent = hWnd
}

func (_this *winControl) Hide() {
	if _this.IsCreate() {
		win32.ShowWindow(_this.hWnd(), win32.SW_HIDE)
	} else {
		_this.createParams.Style &= ^win32.WS_VISIBLE
	}
}

func (_this *winControl) Show() {
	if _this.IsCreate() {
		win32.ShowWindow(_this.hWnd(), win32.SW_SHOW)
	} else {
		_this.createParams.Style |= win32.WS_VISIBLE
	}
}

func (_this *winControl) SetSize(width, height int) {
	if _this.IsCreate() {
		win32.SetWindowPos(_this.hWnd(), 0, 0, 0, int32(width), int32(height), win32.SWP_NOMOVE|win32.SWP_NOZORDER)
	} else {
		_this.createParams.Cx, _this.createParams.Cy = int32(width), int32(height)
	}
}

func (_this *winControl) SetLocation(x, y int) {
	if _this.IsCreate() {
		win32.SetWindowPos(_this.hWnd(), 0, int32(x), int32(y), 0, 0, win32.SWP_NOSIZE|win32.SWP_NOZORDER)
	} else {
		_this.createParams.X, _this.createParams.Y = int32(x), int32(y)
	}
}

func (_this *winControl) Create() {
	if _this.isCreated {
		return
	}
	if _this.createParams.Parent == 0 {
		panic("身为一个控件，必须有父窗口")
	}
	win32.CreateWindowEx(
		_this.createParams.ExStyle,
		(*uint16)(unsafe.Pointer(_this.createParams.ClassName)),
		(*uint16)(unsafe.Pointer(_this.createParams.Name)),
		_this.createParams.Style,
		_this.createParams.X,
		_this.createParams.Y,
		_this.createParams.Cx,
		_this.createParams.Cy,
		_this.createParams.Parent, 0,
		_this.createParams.Instance,
		unsafe.Pointer(&_this.idName))
}

func (_this *winControl) class() string {
	return _this.provider.className
}
