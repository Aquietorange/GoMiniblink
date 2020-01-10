package Windows

import MB "GoMiniblink"

func (_this *winBase) SetOnCreate(fn func(handle uintptr)) {
	_this.onCreate = fn
}

func (_this *winBase) SetOnKeyPress(fn func(e *MB.KeyPressEvArgs)) {
	_this.onKeyPress = fn
}

func (_this *winBase) SetOnKeyUp(fn func(e *MB.KeyEvArgs)) {
	_this.onKeyUp = fn
}

func (_this *winBase) SetOnKeyDown(fn func(e *MB.KeyEvArgs)) {
	_this.onKeyDown = fn
}

func (_this *winBase) SetOnMove(fn func(point MB.Point)) {
	_this.onMove = fn
}

func (_this *winBase) SetOnResize(fn func(rect MB.Rect)) {
	_this.onResize = fn
}

func (_this *winBase) SetOnPaint(fn func(MB.PaintEvArgs)) {
	_this.onPaint = fn
}

func (_this *winBase) SetOnMouseMove(fn func(MB.MouseEvArgs)) {
	_this.onMouseMove = fn
}

func (_this *winBase) SetOnMouseDown(fn func(MB.MouseEvArgs)) {
	_this.onMouseDown = fn
}

func (_this *winBase) SetOnMouseUp(fn func(MB.MouseEvArgs)) {
	_this.onMouseUp = fn
}

func (_this *winBase) SetOnMouseWheel(fn func(MB.MouseEvArgs)) {
	_this.onMouseWheel = fn
}

func (_this *winBase) SetOnMouseClick(fn func(MB.MouseEvArgs)) {
	_this.onMouseClick = fn
}
