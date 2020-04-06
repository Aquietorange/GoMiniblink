package GoMiniblink

import (
	c "qq2564874169/goMiniblink/forms/controls"
)

func init() {
	mbApi = new(freeApiForWindows).init()
}

type free4x64 struct {
	view *c.Control
	wke  wkeHandle
}

func (_this *free4x64) init(control *c.Control) *free4x64 {
	_this.view = control
	bakOnLoad := _this.view.OnLoad
	_this.view.OnLoad = func() {
		_this.mbInit()
		bakOnLoad()
	}
	return _this
}

func (_this *free4x64) mbInit() {
	_this.wke = createWebView(_this)
}

func (_this *free4x64) BindGoFunc(fn GoFunc) {

}

func (_this *free4x64) SetOnRequest(func(e RequestEvArgs)) {

}

func (_this *free4x64) LoadUri(uri string) {

}

func (_this *free4x64) GetHandle() wkeHandle {
	return _this.wke
}
