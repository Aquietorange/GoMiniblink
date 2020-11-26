package bridge

import (
	fm "GoMiniblink/forms"
)

type WindowLostFocusProc func() bool
type WindowFocusProc func() bool
type WindowImeStartCompositionProc func() bool
type WindowSetCursorProc func() bool
type WindowShowProc func()
type WindowCreateProc func(handle uintptr)
type WindowDestroyProc func()
type WindowResizeProc func(e fm.Rect)
type WindowMoveProc func(e fm.Point) bool
type WindowMouseMoveProc func(e *fm.MouseEvArgs)
type WindowMouseDownProc func(e *fm.MouseEvArgs)
type WindowMouseUpProc func(e *fm.MouseEvArgs)
type WindowMouseWheelProc func(e *fm.MouseEvArgs)
type WindowMouseClickProc func(e *fm.MouseEvArgs)
type WindowPaintProc func(e fm.PaintEvArgs) bool
type WindowKeyDownProc func(e *fm.KeyEvArgs)
type WindowKeyUpProc func(e *fm.KeyEvArgs)
type WindowKeyPressProc func(e *fm.KeyPressEvArgs)

type Window interface {
	GetHandle() uintptr
	SetOnCreate(proc WindowCreateProc) WindowCreateProc
	SetOnDestroy(proc WindowDestroyProc) WindowDestroyProc
	SetOnResize(proc WindowResizeProc) WindowResizeProc
	SetOnMove(proc WindowMoveProc) WindowMoveProc
	SetOnMouseMove(proc WindowMouseMoveProc) WindowMouseMoveProc
	SetOnMouseDown(proc WindowMouseDownProc) WindowMouseDownProc
	SetOnMouseUp(proc WindowMouseUpProc) WindowMouseUpProc
	SetOnMouseWheel(proc WindowMouseWheelProc) WindowMouseWheelProc
	SetOnMouseClick(proc WindowMouseClickProc) WindowMouseClickProc
	SetOnPaint(proc WindowPaintProc) WindowPaintProc
	SetOnKeyDown(proc WindowKeyDownProc) WindowKeyDownProc
	SetOnKeyUp(proc WindowKeyUpProc) WindowKeyUpProc
	SetOnKeyPress(proc WindowKeyPressProc) WindowKeyPressProc
	SetOnCursor(proc WindowSetCursorProc) WindowSetCursorProc
	SetOnImeStartComposition(proc WindowImeStartCompositionProc) WindowImeStartCompositionProc
	SetOnFocus(proc WindowFocusProc) WindowFocusProc
	SetOnLostFocus(proc WindowLostFocusProc) WindowLostFocusProc
	SetOnShow(proc WindowShowProc) WindowShowProc

	GetProvider() Provider
	Invoke(fn func(state interface{}), state interface{})
	IsInvoke() bool
	SetSize(w int, h int)
	SetLocation(x int, y int)
	GetBound() fm.Bound
	Show()
	Hide()
	SetBgColor(color int32)
	CreateGraphics() fm.Graphics
	SetCursor(cursor fm.CursorType)
	GetCursor() fm.CursorType
	GetParent() Control
	GetOwner() Form
	ToClientPoint(p fm.Point) fm.Point
	IsEnable() bool
	Enable(b bool)
}
