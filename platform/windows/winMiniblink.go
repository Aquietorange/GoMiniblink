package windows

import (
	mb "qq.2564874169/miniblink"
	core "qq.2564874169/miniblink/platform/miniblinkc-core"
	"qq.2564874169/miniblink/platform/miniblinkc-core/free"
	"qq.2564874169/miniblink/platform/miniblinkc-core/vip"
	"qq.2564874169/miniblink/platform/windows/win32"
)

type winMiniblink struct {
	winControl

	wke      core.ICore
	initUri  string
	initSize mb.Rect
}

func (_this *winMiniblink) init(provider *Provider) *winMiniblink {
	_this.winControl.init(provider)
	_this.evWndCreate["__initWke"] = _this.initWke
	return _this
}

func (_this *winMiniblink) initWke(hWnd win32.HWND) {
	if vip.Exists() {
		//todo
	} else {
		_this.wke = new(free.Core).Init(_this)
	}
	_this.wke.SetOnPaint(_this.paint)
	if _this.initUri != "" {
		_this.LoadUri(_this.initUri)
	}
	if _this.initSize.Wdith > 0 && _this.initSize.Height > 0 {
		_this.Resize(_this.initSize.Wdith, _this.initSize.Height)
	}
}

func (_this *winMiniblink) paint(args core.PaintArgs) {

}

func (_this *winMiniblink) Resize(width, height int) {
	if _this.IsCreate() {
		_this.wke.Resize(width, height)
	} else {
		_this.initSize = mb.Rect{
			Wdith:  width,
			Height: height,
		}
	}
}

func (_this *winMiniblink) LoadUri(uri string) {
	if _this.IsCreate() {
		_this.wke.LoadUri(uri)
	} else {
		_this.initUri = uri
	}
}
