package GoMiniblink

type JsFunc func(param ...interface{}) interface{}

type GoFnContext struct {
	Miniblink Miniblink
	Name      string
	State     interface{}
	Param     []interface{}
}

type GoFn func(context GoFnContext) interface{}

type JsFnBinding struct {
	Name  string
	Fn    GoFn
	State interface{}
	core  wkeJsNativeFunction
}

func (_this *JsFnBinding) Call(mb Miniblink, param []interface{}) interface{} {
	ctx := GoFnContext{
		Miniblink: mb,
		Name:      _this.Name,
		State:     _this.State,
		Param:     param,
	}
	return _this.Fn(ctx)
}
