package miniblink

type wkeFrameContext struct {
	id       uintptr
	isMain   bool
	url      string
	isRemote bool
	core     ICore
}

func (_this *wkeFrameContext) init(core ICore, frame wkeFrame) *wkeFrameContext {
	_this.core = core
	wke := wkeHandle(_this.core.GetHandle())
	_this.id = uintptr(frame)
	_this.isMain = wkeIsMainFrame(wke, frame)
	_this.isRemote = wkeIsWebRemoteFrame(wke, frame)
	_this.url = wkeGetFrameUrl(wke, frame)
	return _this
}

func (_this *wkeFrameContext) RunJs(script string) interface{} {
	es := wkeGetGlobalExecByFrame(wkeHandle(_this.core.GetHandle()), wkeFrame(_this.id))
	rs := jsEvalExW(es, script, true)
	return toGoValue(_this.core, es, rs)
}

func (_this *wkeFrameContext) IsRemote() bool {
	return _this.isRemote
}

func (_this *wkeFrameContext) Url() string {
	return _this.url
}

func (_this *wkeFrameContext) IsMain() bool {
	return _this.isMain
}

func (_this *wkeFrameContext) Id() uintptr {
	return _this.id
}
