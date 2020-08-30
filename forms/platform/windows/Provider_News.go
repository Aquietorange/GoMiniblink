package windows

import (
	"qq2564874169/goMiniblink/forms/platform"
)

func (_this *Provider) NewForm() platform.Form {
	return new(winForm).init(_this)
}

func (_this *Provider) NewControl() platform.Control {
	return new(winControl).init(_this)
}
