package CrossPlatform

type IProvider interface {
	RunMain(form IForm, show func())
	Exit(code int)
	SetIcon(file string)
	GetScreen() Screen

	NewForm() IForm
}
