package windows

import (
	"qq2564874169/goMiniblink/platform"
)

func (_this *Provider) NewForm() platform.IForm {
	return new(winForm).init(_this)
}

func (_this *Provider) NewControl() platform.IControl {
	return new(winControl).init(_this)
}

func (_this *Provider) NewMiniblink() platform.IMiniblink {
	return new(winMiniblink).init(_this)
}
