package controls

import (
	f "qq2564874169/goMiniblink/forms"
	p "qq2564874169/goMiniblink/forms/platform"
)

type Control struct {
	BaseUI

	impl   p.Control
	anchor f.AnchorStyle
}

func (_this *Control) Init() *Control {
	_this.impl = App.NewControl()
	_this.BaseUI.Init(_this, _this.impl)
	_this.anchor = f.AnchorStyle_Left | f.AnchorStyle_Top
	return _this
}

func (_this *Control) getAnchor() f.AnchorStyle {
	return _this.anchor
}

func (_this *Control) SetAnchor(style f.AnchorStyle) {
	_this.anchor = style
}

func (_this *Control) toControl() p.Control {
	return _this.impl
}

func (_this *Control) Show() {
	_this.impl.Show()
}

func (_this *Control) Hide() {
	_this.impl.Hide()
}
