package controls

import f "qq2564874169/goMiniblink/forms"

type GUI interface {
	GetHandle() uintptr
	GetSize() f.Rect
	GetLocation() f.Point
	SetSize(width, height int)
	SetLocation(x, y int)
	SetBgColor(color int32)
	IsInvoke() bool
	Invoke(fn func(state interface{}), state interface{})
	Enable(b bool)
	IsEnable() bool
}
