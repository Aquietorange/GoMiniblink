package Forms

import plat "GoMiniblink/CrossPlatform"

type Control struct {
	BaseUI

	impl   plat.IControl
	isInit bool
}

func (_this *Control) Init() *Control {
	_this.impl = Provider.NewControl()
	_this.BaseUI.init(_this, _this.impl)
	_this.isInit = true
	return _this
}

func (_this *Control) getImpl() plat.IControl {
	if _this.isInit == false {
		panic("必须使用Init()初始化 ")
	}
	return _this.impl
}

func (_this *Control) toChild() plat.IControl {
	return _this.impl
}

func (_this *Control) Show() {
	_this.getImpl().Show()
}

func (_this *Control) Hide() {
	_this.getImpl().Hide()
}
