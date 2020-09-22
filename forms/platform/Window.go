package platform

import (
	f "qq2564874169/goMiniblink/forms"
)

type WindowLostFocusProc func() bool
type WindowFocusProc func() bool
type WindowImeStartCompositionProc func() bool
type WindowSetCursorProc func() bool
type WindowShowProc func()
type WindowCreateProc func(handle uintptr) bool
type WindowDestroyProc func()
type WindowResizeProc func(e f.Rect) bool
type WindowMoveProc func(e f.Point) bool
type WindowMouseMoveProc func(e *f.MouseEvArgs)
type WindowMouseDownProc func(e *f.MouseEvArgs)
type WindowMouseUpProc func(e *f.MouseEvArgs)
type WindowMouseWheelProc func(e *f.MouseEvArgs)
type WindowMouseClickProc func(e *f.MouseEvArgs)
type WindowPaintProc func(e f.PaintEvArgs) bool
type WindowKeyDownProc func(e *f.KeyEvArgs)
type WindowKeyUpProc func(e *f.KeyEvArgs)
type WindowKeyPressProc func(e *f.KeyPressEvArgs)

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
	GetSize() (w, h int)
	SetLocation(x int, y int)
	GetLocation() (x, y int)
	Show()
	Hide()
	SetBgColor(color int32)
	CreateGraphics() f.Graphics
	SetCursor(cursor f.CursorType)
	GetCursor() f.CursorType
	GetParent() Control
	GetOwner() Form
	MousePosition() f.Point
	IsEnable() bool
	Enable(b bool)
}
