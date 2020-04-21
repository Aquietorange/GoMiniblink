package GoMiniblink

type IMiniblink interface {
	BindFunc(fn GoFunc)
	SetOnRequest(callback RequestCallback)
	LoadUri(uri string)
	GetHandle() wkeHandle
}

type RequestCallback func(args RequestEvArgs)
