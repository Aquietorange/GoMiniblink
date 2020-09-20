package windows

import (
	w "qq2564874169/goMiniblink/forms/platform/windows/win32"
	"unsafe"
)

type winControl struct {
	winBase
}

func (_this *winControl) init(provider *Provider) *winControl {
	_this.winBase.init(provider)
	w.CreateWindowEx(
		0,
		sto16(provider.className),
		sto16(""),
		w.WS_CHILD, 0, 0, 100, 100, 0, 0,
		provider.hInstance, unsafe.Pointer(_this))
	return _this
}
