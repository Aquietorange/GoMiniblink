package bridge

type Controls interface {
	Control

	AddControl(control Control)
	RemoveControl(control Control)
	GetChilds() []Control
}
