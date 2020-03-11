package platform

import mb "qq2564874169/goMiniblink"

type IMiniblink interface {
	IControl

	SetOnRequest(func(e mb.RequestEvArgs))
	LoadUri(uri string)
	BindFunc(fn mb.GoFunc)
}
