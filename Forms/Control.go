package Forms

import "GoMiniblink/CrossPlatform"

type Control struct {
	BaseUI

	impl CrossPlatform.IControl
}

func (_this *Control) Init() *Control {
	_this.impl = Provider.NewControl()
	return _this
}
