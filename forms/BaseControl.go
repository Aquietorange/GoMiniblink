package forms

import (
	mb "qq.2564874169/goMiniblink"
	p "qq.2564874169/goMiniblink/platform"
)

type BaseControl struct {
	BaseUI

	impl   p.IControl
	isInit bool
	anchor mb.AnchorStyle
}

func (_this *BaseControl) Init(impl p.IControl) *BaseControl {
	_this.impl = impl
	_this.BaseUI.init(_this, _this.impl)
	_this.anchor = mb.AnchorStyle_Left | mb.AnchorStyle_Top
	_this.isInit = true
	return _this
}

func (_this *BaseControl) getAnchor() mb.AnchorStyle {
	return _this.anchor
}

func (_this *BaseControl) SetAnchor(style mb.AnchorStyle) {
	_this.anchor = style
}

func (_this *BaseControl) toControl() p.IControl {
	return _this.impl
}

func (_this *BaseControl) Show() {
	_this.getImpl().Show()
}

func (_this *BaseControl) Hide() {
	_this.getImpl().Hide()
}

func (_this *BaseControl) getImpl() p.IControl {
	if _this.isInit == false {
		panic("必须使用Init()初始化 ")
	}
	return _this.impl
}
