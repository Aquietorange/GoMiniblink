package platform

import mb "qq2564874169/goMiniblink"

type IWindow interface {
	Id() string
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

	GetProvider() IProvider
	Invoke(fn func(state interface{}), state interface{})
	SetSize(w int, h int)
	SetLocation(x int, y int)
	Show()
	Hide()
	SetBgColor(color int)
	CreateGraphics() mb.Graphics
	SetCursor(cursor mb.CursorType)
}

type WindowLostFocusProc func() bool
type WindowFocusProc func() bool
type WindowImeStartCompositionProc func() bool
type WindowSetCursorProc func() bool
type WindowCreateProc func(handle uintptr) bool
type WindowDestroyProc func()
type WindowResizeProc func(e mb.Rect) bool
type WindowMoveProc func(e mb.Point) bool
type WindowMouseMoveProc func(e mb.MouseEvArgs) bool
type WindowMouseDownProc func(e mb.MouseEvArgs) bool
type WindowMouseUpProc func(e mb.MouseEvArgs) bool
type WindowMouseWheelProc func(e mb.MouseEvArgs) bool
type WindowMouseClickProc func(e mb.MouseEvArgs) bool
type WindowPaintProc func(e mb.PaintEvArgs) bool
type WindowKeyDownProc func(e *mb.KeyEvArgs) bool
type WindowKeyUpProc func(e *mb.KeyEvArgs) bool
type WindowKeyPressProc func(e *mb.KeyPressEvArgs) bool
