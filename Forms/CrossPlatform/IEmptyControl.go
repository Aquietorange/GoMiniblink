package CrossPlatform

type IEmptyControl interface {
	Create()
	IsCreate() bool
	SetOnCreate(func())

	SetSize(w int, h int)
	SetOnResize(func(w int, h int))

	SetLocation(x int, y int)
	SetOnMove(func(x int, y int))

	Invoke(fn func(state interface{}), state interface{})
	Show()
	Hide()
}
