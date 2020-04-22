package GoMiniblink

type JsFunc func(param ...interface{}) interface{}

type GoFuncContext struct {
	Miniblink IMiniblink
	Name      string
	State     interface{}
	Param     []interface{}
}

type GoFunc struct {
	Name      string
	BindToSub bool
	Func      func(context GoFuncContext) interface{}
	State     interface{}
	core      wkeJsNativeFunction
}

func (_this *GoFunc) Call(miniblink IMiniblink, param []interface{}) interface{} {
	ctx := GoFuncContext{
		Miniblink: miniblink,
		Name:      _this.Name,
		State:     _this.State,
		Param:     param,
	}
	return _this.Func(ctx)
}
