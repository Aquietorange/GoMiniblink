package GoMiniblink

type IMiniblink interface {
	BindGoFunc(fn GoFunc)
	SetOnRequest(func(e RequestEvArgs))
	LoadUri(uri string)
	GetHandle() wkeHandle
}
