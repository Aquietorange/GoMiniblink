package CrossPlatform

import MB "GoMiniblink"

type IWindow interface {
	Create()
	IsCreate() bool
	SetOnCreate(func())
	SetOnMouseMove(func(MB.MouseEvArgs))
	SetOnMouseDown(func(MB.MouseEvArgs))
	SetOnMouseUp(func(MB.MouseEvArgs))
	SetOnMouseWheel(func(MB.MouseEvArgs))
	SetOnMouseClick(func(MB.MouseEvArgs))

	Invoke(fn func(state interface{}), state interface{})
	Show()
	Hide()
}
