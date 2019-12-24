package Windows

import "GoMiniblink/Forms/CrossPlatform"

func (_this *Provider) NewForm() CrossPlatform.IForm {
	return new(winForm).init(_this)
}
