package windows

import (
	plat "qq.2564874169/goMiniblink/platform"
)

func (_this *winBase) SetOnFocus(proc plat.WindowFocusProc) plat.WindowFocusProc {
	pre := _this.onFocus
	_this.onFocus = proc
	return pre
}

func (_this *winBase) SetOnImeStartComposition(proc plat.WindowImeStartCompositionProc) plat.WindowImeStartCompositionProc {
	pre := _this.onImeStartComposition
	_this.onImeStartComposition = proc
	return pre
}

func (_this *winBase) SetOnCreate(proc plat.WindowCreateProc) plat.WindowCreateProc {
	pre := _this.onCreate
	_this.onCreate = proc
	return pre
}

func (_this *winBase) SetOnDestroy(proc plat.WindowDestroyProc) plat.WindowDestroyProc {
	pre := _this.onDestroy
	_this.onDestroy = proc
	return pre
}

func (_this *winBase) SetOnKeyPress(proc plat.WindowKeyPressProc) plat.WindowKeyPressProc {
	pre := _this.onKeyPress
	_this.onKeyPress = proc
	return pre
}

func (_this *winBase) SetOnKeyUp(proc plat.WindowKeyUpProc) plat.WindowKeyUpProc {
	pre := _this.onKeyUp
	_this.onKeyUp = proc
	return pre
}

func (_this *winBase) SetOnKeyDown(proc plat.WindowKeyDownProc) plat.WindowKeyDownProc {
	pre := _this.onKeyDown
	_this.onKeyDown = proc
	return pre
}

func (_this *winBase) SetOnMove(proc plat.WindowMoveProc) plat.WindowMoveProc {
	pre := _this.onMove
	_this.onMove = proc
	return pre
}

func (_this *winBase) SetOnResize(proc plat.WindowResizeProc) plat.WindowResizeProc {
	pre := _this.onResize
	_this.onResize = proc
	return pre
}

func (_this *winBase) SetOnPaint(proc plat.WindowPaintProc) plat.WindowPaintProc {
	pre := _this.onPaint
	_this.onPaint = proc
	return pre
}

func (_this *winBase) SetOnMouseMove(proc plat.WindowMouseMoveProc) plat.WindowMouseMoveProc {
	pre := _this.onMouseMove
	_this.onMouseMove = proc
	return pre
}

func (_this *winBase) SetOnMouseDown(proc plat.WindowMouseDownProc) plat.WindowMouseDownProc {
	pre := _this.onMouseDown
	_this.onMouseDown = proc
	return pre
}

func (_this *winBase) SetOnMouseUp(proc plat.WindowMouseUpProc) plat.WindowMouseUpProc {
	pre := _this.onMouseUp
	_this.onMouseUp = proc
	return pre
}

func (_this *winBase) SetOnMouseWheel(proc plat.WindowMouseWheelProc) plat.WindowMouseWheelProc {
	pre := _this.onMouseWheel
	_this.onMouseWheel = proc
	return pre
}

func (_this *winBase) SetOnMouseClick(proc plat.WindowMouseClickProc) plat.WindowMouseClickProc {
	pre := _this.onMouseClick
	_this.onMouseClick = proc
	return pre
}

func (_this *winBase) SetOnCursor(proc plat.WindowSetCursorProc) plat.WindowSetCursorProc {
	pre := _this.onSetCursor
	_this.onSetCursor = proc
	return pre
}
