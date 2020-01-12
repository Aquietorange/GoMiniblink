package Forms

import plat "GoMiniblink/CrossPlatform"

type BaseControl struct {
	BaseUI

	impl   plat.IControl
	isInit bool
}

func (_this *BaseControl) Init(impl plat.IControl) *BaseControl {
	_this.impl = impl
	_this.BaseUI.init(_this, _this.impl)
	_this.isInit = true
	return _this
}

func (_this *BaseControl) getImpl() plat.IControl {
	if _this.isInit == false {
		panic("必须使用Init()初始化 ")
	}
	return _this.impl
}

func (_this *BaseControl) toChild() plat.IControl {
	return _this.impl
}

func (_this *BaseControl) Show() {
	_this.getImpl().Show()
}

func (_this *BaseControl) Hide() {
	_this.getImpl().Hide()
}
