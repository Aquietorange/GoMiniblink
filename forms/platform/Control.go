package platform

type Control interface {
	Window

	SetParent(window Window)
}
