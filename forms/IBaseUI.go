package forms

import "qq.2564874169/miniblink"

type IBaseUI interface {
	GetHandle() uintptr
	GetSize() miniblink.Rect
	GetLocation() miniblink.Point
	SetSize(width, height int)
	SetLocation(x, y int)
	SetBgColor(color int)
	Invoke(fn func(state interface{}), state interface{})
}
