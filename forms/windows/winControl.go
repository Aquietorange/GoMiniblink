package windows

import (
	win "qq2564874169/goMiniblink/forms/windows/win32"
	"unsafe"
)

type winControl struct {
	winBase
}

func (_this *winControl) init(provider *Provider) *winControl {
	_this.winBase.init(provider)
	win.CreateWindowEx(
		0,
		sto16(provider.className),
		sto16(""),
		win.WS_CHILD, 0, 0, 100, 100, _this.app.defOwner, 0,
		provider.hInstance, unsafe.Pointer(_this))
	return _this
}
