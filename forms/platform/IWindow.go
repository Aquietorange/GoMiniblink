package platform

import (
	f "qq2564874169/goMiniblink/forms"
)

type IWindow interface {
	Create()
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
	SetOnLoad(proc WindowLoadProc) WindowLoadProc

	GetProvider() IProvider
	Invoke(fn func(state interface{}), state interface{})
	SetSize(w int, h int)
	SetLocation(x int, y int)
	Show()
	Hide()
	SetBgColor(color int)
	CreateGraphics() f.Graphics
	SetCursor(cursor f.CursorType)
}

type WindowLostFocusProc func() bool
type WindowFocusProc func() bool
type WindowImeStartCompositionProc func() bool
type WindowSetCursorProc func() bool
type WindowLoadProc func()
type WindowCreateProc func(handle uintptr) bool
type WindowDestroyProc func()
type WindowResizeProc func(e f.Rect) bool
type WindowMoveProc func(e f.Point) bool
type WindowMouseMoveProc func(e f.MouseEvArgs) bool
type WindowMouseDownProc func(e f.MouseEvArgs) bool
type WindowMouseUpProc func(e f.MouseEvArgs) bool
type WindowMouseWheelProc func(e f.MouseEvArgs) bool
type WindowMouseClickProc func(e f.MouseEvArgs) bool
type WindowPaintProc func(e f.PaintEvArgs) bool
type WindowKeyDownProc func(e *f.KeyEvArgs) bool
type WindowKeyUpProc func(e *f.KeyEvArgs) bool
type WindowKeyPressProc func(e *f.KeyPressEvArgs) bool
