package GoMiniblink

type FrameContext interface {
	Id() uintptr
	IsMain() bool
	Url() string
	IsRemote() bool
	RunJs(script string) interface{}
}

type freeFrameContext struct {
	id       uintptr
	isMain   bool
	url      string
	isRemote bool
	core     Miniblink
}

func (_this *freeFrameContext) init(mb Miniblink, frame wkeFrame) *freeFrameContext {
	_this.core = mb
	_this.id = uintptr(frame)
	_this.isMain = mbApi.wkeIsMainFrame(_this.core.GetHandle(), frame)
	_this.isRemote = mbApi.wkeIsWebRemoteFrame(_this.core.GetHandle(), frame)
	_this.url = mbApi.wkeGetFrameUrl(_this.core.GetHandle(), frame)
	return _this
}

func (_this *freeFrameContext) RunJs(script string) interface{} {
	if len(script) > 0 {
		es := mbApi.wkeGetGlobalExecByFrame(_this.core.GetHandle(), wkeFrame(_this.id))
		rs := mbApi.jsEval(es, script)
		return toGoValue(_this.core, es, rs)
	}
	return nil
}

func (_this *freeFrameContext) IsRemote() bool {
	return _this.isRemote
}

func (_this *freeFrameContext) Url() string {
	return _this.url
}

func (_this *freeFrameContext) IsMain() bool {
	return _this.isMain
}

func (_this *freeFrameContext) Id() uintptr {
	return _this.id
}

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
