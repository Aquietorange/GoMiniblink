package GoMiniblink

type IMiniblink interface {
	BindFunc(fn GoFunc)
	SetOnRequest(func(e RequestEvArgs))
	LoadUri(uri string)
	GetHandle() wkeHandle
}
