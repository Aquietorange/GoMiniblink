package Windows

import "GoMiniblink/CrossPlatform"

func (_this *Provider) NewForm() CrossPlatform.IForm {
	return new(winForm).init(_this)
}
