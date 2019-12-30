package Windows

import MB "GoMiniblink"

func (_this *winControl) defOnMouseMove(e MB.MouseEvArgs) {
	for _, v := range _this.evMouseMove {
		v(_this, e)
	}
}

func (_this *winControl) defOnMouseDown(e MB.MouseEvArgs) {
	for _, v := range _this.evMouseDown {
		v(_this, e)
	}
}

func (_this *winControl) defOnMouseUp(e MB.MouseEvArgs) {
	for _, v := range _this.evMouseUp {
		v(_this, e)
	}
}

func (_this *winControl) defOnMouseWheel(e MB.MouseEvArgs) {
	for _, v := range _this.evMouseWheel {
		v(_this, e)
	}
}

func (_this *winControl) defOnMouseClick(e MB.MouseEvArgs) {
	for _, v := range _this.evMouseClick {
		v(_this, e)
	}
}

func (_this *winControl) SetOnMouseClick(fn func(MB.MouseEvArgs)) {
	_this.onMouseClick = fn
}

func (_this *winControl) SetOnMouseWheel(fn func(MB.MouseEvArgs)) {
	_this.onMouseWheel = fn
}

func (_this *winControl) SetOnMouseUp(fn func(MB.MouseEvArgs)) {
	_this.onMouseUp = fn
}

func (_this *winControl) SetOnMouseDown(fn func(MB.MouseEvArgs)) {
	_this.onMouseDown = fn
}

func (_this *winControl) SetOnMouseMove(fn func(MB.MouseEvArgs)) {
	_this.onMouseMove = fn
}
