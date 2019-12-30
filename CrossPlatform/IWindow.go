package CrossPlatform

import MB "GoMiniblink"

type IWindow interface {
	Create()
	IsCreate() bool
	SetOnCreate(fn func())
	SetOnMouseMove(func(e MB.MouseEvArgs))
	SetOnMouseDown(func(e MB.MouseEvArgs))
	SetOnMouseUp(func(e MB.MouseEvArgs))
	SetOnMouseWheel(func(e MB.MouseEvArgs))
	SetOnMouseClick(func(e MB.MouseEvArgs))

	Invoke(fn func(state interface{}), state interface{})
	Show()
	Hide()
}
