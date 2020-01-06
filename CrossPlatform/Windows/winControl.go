package Windows

import (
	"GoMiniblink/CrossPlatform/Windows/win32"
	"GoMiniblink/Utils"
)

type winControl struct {
	winBase

	getParent func() win32.HWND
}

func (_this *winControl) init(provider *Provider) *winControl {
	_this.winBase.init(provider, Utils.NewUUID())
	_this.provider.add(_this)
	return _this
}

func (_this *winControl) Create() {

}

func (_this *winControl) class() string {
	return _this.provider.className
}
