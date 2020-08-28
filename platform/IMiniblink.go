package platform

import mb "qq2564874169/goMiniblink"

type IMiniblink interface {
	IControl

	SetOnRequest(func(e mb.RequestBeforeEvArgs))
	SetOnJsReady(func(e mb.JsReadyEvArgs))
	LoadUri(uri string)
	BindJsFunc(fn mb.JsFuncBinding)
	RunJs(script string) interface{}
	SetWindowProp(name string, value interface{})
}
