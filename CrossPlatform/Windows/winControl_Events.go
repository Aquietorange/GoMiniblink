package Windows

import MB "GoMiniblink"

func (_this *winControl) SetOnKeyPress(fn func(e *MB.KeyPressEvArgs)) {
	_this.onKeyPress = fn
}

func (_this *winControl) SetOnKeyUp(fn func(e *MB.KeyEvArgs)) {
	_this.onKeyUp = fn
}

func (_this *winControl) SetOnKeyDown(fn func(e *MB.KeyEvArgs)) {
	_this.onKeyDown = fn
}

func (_this *winControl) SetOnMove(fn func(point MB.Point)) {
	_this.onMove = fn
}

func (_this *winControl) SetOnResize(fn func(rect MB.Rect)) {
	_this.onResize = fn
}

func (_this *winControl) SetOnPaint(fn func(MB.PaintEvArgs)) {
	_this.onPaint = fn
}

func (_this *winControl) SetOnMouseMove(fn func(MB.MouseEvArgs)) {
	_this.onMouseMove = fn
}

func (_this *winControl) SetOnMouseDown(fn func(MB.MouseEvArgs)) {
	_this.onMouseDown = fn
}

func (_this *winControl) SetOnMouseUp(fn func(MB.MouseEvArgs)) {
	_this.onMouseUp = fn
}
func (_this *winControl) SetOnMouseWheel(fn func(MB.MouseEvArgs)) {
	_this.onMouseWheel = fn
}

func (_this *winControl) SetOnMouseClick(fn func(MB.MouseEvArgs)) {
	_this.onMouseClick = fn
}
