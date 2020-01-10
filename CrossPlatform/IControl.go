package CrossPlatform

type IControl interface {
	IWindow

	SetParent(window IWindow)
}
