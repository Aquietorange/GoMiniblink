package CrossPlatform

import MB "GoMiniblink"

type IWindow interface {
	Id() string
	Create()
	IsCreate() bool
	SetOnCreate(fn func())
	SetOnResize(func(e MB.Rect))
	SetOnMove(func(e MB.Point))
	SetOnMouseMove(func(e MB.MouseEvArgs))
	SetOnMouseDown(func(e MB.MouseEvArgs))
	SetOnMouseUp(func(e MB.MouseEvArgs))
	SetOnMouseWheel(func(e MB.MouseEvArgs))
	SetOnMouseClick(func(e MB.MouseEvArgs))
	SetOnPaint(func(e MB.PaintEvArgs))
	SetOnKeyDown(func(e *MB.KeyEvArgs))
	SetOnKeyUp(func(e *MB.KeyEvArgs))
	SetOnKeyPress(func(e *MB.KeyPressEvArgs))

	Invoke(fn func(state interface{}), state interface{})
	SetSize(w int, h int)
	SetLocation(x int, y int)
	Show()
	Hide()
}
