package platform

type IControls interface {
	AddControl(control IControl)
	RemoveControl(control IControl)
}
