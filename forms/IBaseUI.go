package forms

import "qq.2564874169/goMiniblink"

type IBaseUI interface {
	GetHandle() uintptr
	GetSize() goMiniblink.Rect
	GetLocation() goMiniblink.Point
	SetSize(width, height int)
	SetLocation(x, y int)
	SetBgColor(color int)
	Invoke(fn func(state interface{}), state interface{})
}
