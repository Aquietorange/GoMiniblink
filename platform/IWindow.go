package platform

import mb "qq.2564874169/miniblink"

type IWindow interface {
	Id() string
	Create()
	IsCreate() bool
	GetHandle() uintptr
	SetOnCreate(fn func(handle uintptr))
	SetOnResize(func(e mb.Rect))
	SetOnMove(func(e mb.Point))
	SetOnMouseMove(func(e mb.MouseEvArgs))
	SetOnMouseDown(func(e mb.MouseEvArgs))
	SetOnMouseUp(func(e mb.MouseEvArgs))
	SetOnMouseWheel(func(e mb.MouseEvArgs))
	SetOnMouseClick(func(e mb.MouseEvArgs))
	SetOnPaint(func(e mb.PaintEvArgs))
	SetOnKeyDown(func(e *mb.KeyEvArgs))
	SetOnKeyUp(func(e *mb.KeyEvArgs))
	SetOnKeyPress(func(e *mb.KeyPressEvArgs))

	Invoke(fn func(state interface{}), state interface{})
	SetSize(w int, h int)
	SetLocation(x int, y int)
	Show()
	Hide()
	SetBgColor(color int)
	CreateGraphics() mb.Graphics
}
