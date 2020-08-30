package controls

import p "qq2564874169/goMiniblink/forms/platform"

type Control struct {
	BaseControl
	core p.Control
}

func (_this *Control) Init() *Control {
	_this.impl = App.NewControl()
	_this.BaseControl.Init(_this.impl)
	_this.BaseControl.SetBgColor(-1)
	return _this
}
