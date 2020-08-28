package miniblink

import "qq2564874169/goMiniblink"

type wkeJsReadyEvArgs struct {
	ctx *wkeFrameContext
}

func (_this *wkeJsReadyEvArgs) init() *wkeJsReadyEvArgs {
	return _this
}

func (_this *wkeJsReadyEvArgs) Frame() goMiniblink.FrameContext {
	return _this.ctx
}
