package CrossPlatform

type IProvider interface {
	RunMain(form IForm)
	Exit(code int)
	NewForm() IForm
}
