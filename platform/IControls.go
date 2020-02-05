package platform

type IControls interface {
	IWindow

	AddControl(control IControl)
	RemoveControl(control IControl)
}
