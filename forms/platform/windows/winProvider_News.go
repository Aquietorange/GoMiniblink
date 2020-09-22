package windows

import (
	plat "qq2564874169/goMiniblink/forms/platform"
)

func (_this *Provider) NewForm(param plat.FormParam) plat.Form {
	return new(winForm).init(_this, param)
}

func (_this *Provider) NewControl() plat.Control {
	return new(winControl).init(_this)
}
