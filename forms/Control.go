package forms

import "qq.2564874169/miniblink/platform"

type BaseControl struct {
	BaseUI

	impl   platform.IControl
	isInit bool
}

func (_this *BaseControl) Init(impl platform.IControl) *BaseControl {
	_this.impl = impl
	_this.BaseUI.init(_this, _this.impl)
	_this.isInit = true
	return _this
}

func (_this *BaseControl) getImpl() platform.IControl {
	if _this.isInit == false {
		panic("必须使用Init()初始化 ")
	}
	return _this.impl
}

func (_this *BaseControl) toChild() platform.IControl {
	return _this.impl
}

func (_this *BaseControl) Show() {
	_this.getImpl().Show()
}

func (_this *BaseControl) Hide() {
	_this.getImpl().Hide()
}
