package windows

import mb "qq.2564874169/miniblink"

func (_this *winBase) SetOnCreate(fn func(handle uintptr)) {
	_this.onCreate = fn
}

func (_this *winBase) SetOnKeyPress(fn func(e *mb.KeyPressEvArgs)) {
	_this.onKeyPress = fn
}

func (_this *winBase) SetOnKeyUp(fn func(e *mb.KeyEvArgs)) {
	_this.onKeyUp = fn
}

func (_this *winBase) SetOnKeyDown(fn func(e *mb.KeyEvArgs)) {
	_this.onKeyDown = fn
}

func (_this *winBase) SetOnMove(fn func(point mb.Point)) {
	_this.onMove = fn
}

func (_this *winBase) SetOnResize(fn func(rect mb.Rect)) {
	_this.onResize = fn
}

func (_this *winBase) SetOnPaint(fn func(mb.PaintEvArgs)) {
	_this.onPaint = fn
}

func (_this *winBase) SetOnMouseMove(fn func(mb.MouseEvArgs)) {
	_this.onMouseMove = fn
}

func (_this *winBase) SetOnMouseDown(fn func(mb.MouseEvArgs)) {
	_this.onMouseDown = fn
}

func (_this *winBase) SetOnMouseUp(fn func(mb.MouseEvArgs)) {
	_this.onMouseUp = fn
}

func (_this *winBase) SetOnMouseWheel(fn func(mb.MouseEvArgs)) {
	_this.onMouseWheel = fn
}

func (_this *winBase) SetOnMouseClick(fn func(mb.MouseEvArgs)) {
	_this.onMouseClick = fn
}
