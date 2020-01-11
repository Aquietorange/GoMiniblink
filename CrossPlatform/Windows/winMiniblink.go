package Windows

import (
	"GoMiniblink/CrossPlatform/Windows/win32"
	"GoMiniblink/CrossPlatform/miniblink"
	"GoMiniblink/CrossPlatform/miniblink/free"
	"GoMiniblink/CrossPlatform/miniblink/vip"
)

type winMiniblink struct {
	winControl

	wke miniblink.ICore
}

func (_this *winMiniblink) init(provider *Provider) *winMiniblink {
	_this.winControl.init(provider)
	_this.evWndCreate["__initWke"] = _this.initWke
	return _this
}

func (_this *winMiniblink) initWke(hWnd win32.HWND) {
	if vip.Exists() {

	} else {
		_this.wke = new(free.Core).Init(_this)
	}
	_this.wke.SetOnPaint(_this.paint)
}

func (_this *winMiniblink) paint(args miniblink.PaintArgs) {

}

func (_this *winMiniblink) LoadUri(uri string) {
	_this.wke.LoadUri(uri)
}
