package platform

type Controls interface {
	Window

	AddControl(control Control)
	RemoveControl(control Control)
}
