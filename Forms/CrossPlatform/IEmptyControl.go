package CrossPlatform

type IEmptyControl interface {
	Create()
	SetSize(w int, h int)
	OnResize(func(w int, h int))
	SetLocation(x int, y int)
	OnMove(func(x int, y int))

	Show()
	Hide()
}
