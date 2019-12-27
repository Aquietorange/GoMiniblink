package CrossPlatform

type IProvider interface {
	RunMain(form IForm, show func())
	Exit(code int)
	NewForm() IForm
}
