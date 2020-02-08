package platform

import mb "qq.2564874169/goMiniblink"

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

	Invoke(fn func(state interface{}), state interface{})
	SetSize(w int, h int)
	SetLocation(x int, y int)
	Show()
	Hide()
	SetBgColor(color int)
	CreateGraphics() mb.Graphics
}

type WindowCreateProc func(handle uintptr)
type WindowDestroyProc func()
type WindowResizeProc func(e mb.Rect)
type WindowMoveProc func(e mb.Point)
type WindowMouseMoveProc func(e mb.MouseEvArgs)
type WindowMouseDownProc func(e mb.MouseEvArgs)
type WindowMouseUpProc func(e mb.MouseEvArgs)
type WindowMouseWheelProc func(e mb.MouseEvArgs)
type WindowMouseClickProc func(e mb.MouseEvArgs)
type WindowPaintProc func(e mb.PaintEvArgs)
type WindowKeyDownProc func(e *mb.KeyEvArgs)
type WindowKeyUpProc func(e *mb.KeyEvArgs)
type WindowKeyPressProc func(e *mb.KeyPressEvArgs)
