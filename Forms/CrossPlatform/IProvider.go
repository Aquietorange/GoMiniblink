package CrossPlatform

type IProvider interface {
	RunMain(form IForm)
	NewForm() IForm
}
