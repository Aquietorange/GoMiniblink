package Windows

import "GoMiniblink/CrossPlatform"

func (_this *Provider) NewForm() CrossPlatform.IForm {
	return new(winForm).init(_this)
}

func (_this *Provider) NewControl() CrossPlatform.IControl {
	return new(winControl).init(_this)
}
