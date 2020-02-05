package windows

import (
	"qq.2564874169/miniblink/platform"
)

func (_this *winBase) SetOnCreate(proc platform.WindowCreateProc) platform.WindowCreateProc {
	pre := _this.onCreate
	_this.onCreate = proc
	return pre
}

func (_this *winBase) SetOnDestroy(proc platform.WindowDestroyProc) platform.WindowDestroyProc {
	pre := _this.onDestroy
	_this.onDestroy = proc
	return pre
}

func (_this *winBase) SetOnKeyPress(proc platform.WindowKeyPressProc) platform.WindowKeyPressProc {
	pre := _this.onKeyPress
	_this.onKeyPress = proc
	return pre
}

func (_this *winBase) SetOnKeyUp(proc platform.WindowKeyUpProc) platform.WindowKeyUpProc {
	pre := _this.onKeyUp
	_this.onKeyUp = proc
	return pre
}

func (_this *winBase) SetOnKeyDown(proc platform.WindowKeyDownProc) platform.WindowKeyDownProc {
	pre := _this.onKeyDown
	_this.onKeyDown = proc
	return pre
}

func (_this *winBase) SetOnMove(proc platform.WindowMoveProc) platform.WindowMoveProc {
	pre := _this.onMove
	_this.onMove = proc
	return pre
}

func (_this *winBase) SetOnResize(proc platform.WindowResizeProc) platform.WindowResizeProc {
	pre := _this.onResize
	_this.onResize = proc
	return pre
}

func (_this *winBase) SetOnPaint(proc platform.WindowPaintProc) platform.WindowPaintProc {
	pre := _this.onPaint
	_this.onPaint = proc
	return pre
}

func (_this *winBase) SetOnMouseMove(proc platform.WindowMouseMoveProc) platform.WindowMouseMoveProc {
	pre := _this.onMouseMove
	_this.onMouseMove = proc
	return pre
}

func (_this *winBase) SetOnMouseDown(proc platform.WindowMouseDownProc) platform.WindowMouseDownProc {
	pre := _this.onMouseDown
	_this.onMouseDown = proc
	return pre
}

func (_this *winBase) SetOnMouseUp(proc platform.WindowMouseUpProc) platform.WindowMouseUpProc {
	pre := _this.onMouseUp
	_this.onMouseUp = proc
	return pre
}

func (_this *winBase) SetOnMouseWheel(proc platform.WindowMouseWheelProc) platform.WindowMouseWheelProc {
	pre := _this.onMouseWheel
	_this.onMouseWheel = proc
	return pre
}

func (_this *winBase) SetOnMouseClick(proc platform.WindowMouseClickProc) platform.WindowMouseClickProc {
	pre := _this.onMouseClick
	_this.onMouseClick = proc
	return pre
}
