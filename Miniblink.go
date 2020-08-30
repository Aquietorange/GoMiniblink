package GoMiniblink

type RequestCallback func(args RequestEvArgs)
type JsReadyCallback func(args JsReadyEvArgs)
type ConsoleCallback func(args ConsoleEvArgs)

type Miniblink interface {
	JsFunc(name string, fn GoFn, state interface{})
	RunJs(script string) interface{}
	SetOnConsole(callback ConsoleCallback)
	SetOnJsReady(callback JsReadyCallback)
	SetOnRequest(callback RequestCallback)
	LoadUri(uri string)
	GetHandle() wkeHandle
}
