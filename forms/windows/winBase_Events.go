package windows

import (
	br "qq2564874169/goMiniblink/forms/bridge"
)

func (_this *winBase) SetOnShow(proc br.WindowShowProc) br.WindowShowProc {
	pre := _this.onShow
	_this.onShow = proc
	return pre
}

func (_this *winBase) SetOnLostFocus(proc br.WindowLostFocusProc) br.WindowLostFocusProc {
	pre := _this.onLostFocus
	_this.onLostFocus = proc
	return pre
}

func (_this *winBase) SetOnFocus(proc br.WindowFocusProc) br.WindowFocusProc {
	pre := _this.onFocus
	_this.onFocus = proc
	return pre
}

func (_this *winBase) SetOnImeStartComposition(proc br.WindowImeStartCompositionProc) br.WindowImeStartCompositionProc {
	pre := _this.onImeStartComposition
	_this.onImeStartComposition = proc
	return pre
}

func (_this *winBase) SetOnCreate(proc br.WindowCreateProc) br.WindowCreateProc {
	pre := _this.onCreate
	_this.onCreate = proc
	return pre
}

func (_this *winBase) SetOnDestroy(proc br.WindowDestroyProc) br.WindowDestroyProc {
	pre := _this.onDestroy
	_this.onDestroy = proc
	return pre
}

func (_this *winBase) SetOnKeyPress(proc br.WindowKeyPressProc) br.WindowKeyPressProc {
	pre := _this.onKeyPress
	_this.onKeyPress = proc
	return pre
}

func (_this *winBase) SetOnKeyUp(proc br.WindowKeyUpProc) br.WindowKeyUpProc {
	pre := _this.onKeyUp
	_this.onKeyUp = proc
	return pre
}

func (_this *winBase) SetOnKeyDown(proc br.WindowKeyDownProc) br.WindowKeyDownProc {
	pre := _this.onKeyDown
	_this.onKeyDown = proc
	return pre
}

func (_this *winBase) SetOnMove(proc br.WindowMoveProc) br.WindowMoveProc {
	pre := _this.onMove
	_this.onMove = proc
	return pre
}

func (_this *winBase) SetOnResize(proc br.WindowResizeProc) br.WindowResizeProc {
	pre := _this.onResize
	_this.onResize = proc
	return pre
}

func (_this *winBase) SetOnPaint(proc br.WindowPaintProc) br.WindowPaintProc {
	pre := _this.onPaint
	_this.onPaint = proc
	return pre
}

func (_this *winBase) SetOnMouseMove(proc br.WindowMouseMoveProc) br.WindowMouseMoveProc {
	pre := _this.onMouseMove
	_this.onMouseMove = proc
	return pre
}

func (_this *winBase) SetOnMouseDown(proc br.WindowMouseDownProc) br.WindowMouseDownProc {
	pre := _this.onMouseDown
	_this.onMouseDown = proc
	return pre
}

func (_this *winBase) SetOnMouseUp(proc br.WindowMouseUpProc) br.WindowMouseUpProc {
	pre := _this.onMouseUp
	_this.onMouseUp = proc
	return pre
}

func (_this *winBase) SetOnMouseWheel(proc br.WindowMouseWheelProc) br.WindowMouseWheelProc {
	pre := _this.onMouseWheel
	_this.onMouseWheel = proc
	return pre
}

func (_this *winBase) SetOnMouseClick(proc br.WindowMouseClickProc) br.WindowMouseClickProc {
	pre := _this.onMouseClick
	_this.onMouseClick = proc
	return pre
}

func (_this *winBase) SetOnCursor(proc br.WindowSetCursorProc) br.WindowSetCursorProc {
	pre := _this.onSetCursor
	_this.onSetCursor = proc
	return pre
}
