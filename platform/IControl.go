package platform

type IControl interface {
	IWindow

	SetParent(window IWindow)
}
