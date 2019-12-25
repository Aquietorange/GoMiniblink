package Forms

import (
	"GoMiniblink/Forms/CrossPlatform"
	"GoMiniblink/Utils"
)

type Form struct {
	Controls *controlList
	EvResize []func(f *Form, w, h int)
	OnResize func(w, h int)
	EvMove   []func(f *Form, x, y int)
	OnMove   func(x, y int)

	isInit bool
	impl   CrossPlatform.IForm
	x      int
	y      int
	weight int
	height int
}

func (_this *Form) RunMain(provider CrossPlatform.IProvider) {
	provider.RunMain(_this.impl)
}

func (_this *Form) getImpl() CrossPlatform.IForm {
	if _this.isInit == false {
		panic("窗体在使用前必须先调用 Init() ")
	}
	return _this.impl
}

func (_this *Form) Init() *Form {
	_this.impl = Provider.NewForm()
	_this.Controls = new(controlList).Init()
	_this.OnResize = Utils.IfNull(_this.OnResize, _this.onResize).(func(int, int))
	_this.OnMove = Utils.IfNull(_this.OnMove, _this.onMove).(func(int, int))
	_this.registerEvents()
	_this.isInit = true
	return _this
}

func (_this *Form) registerEvents() {
	_this.impl.OnResize(func(w, h int) {
		_this.weight, _this.height = w, h
		_this.OnResize(w, h)
	})
	_this.impl.OnMove(func(x int, y int) {
		_this.x, _this.y = x, y
		_this.OnMove(x, y)
	})
}

func (_this *Form) Show() {
	_this.impl.Show()
}

func (_this *Form) SetSize(w, h int) {
	_this.getImpl().SetSize(w, h)
	_this.weight, _this.height = w, h
	_this.OnResize(w, h)
}

func (_this *Form) onResize(w, h int) {
	for _, v := range _this.EvResize[:] {
		v(_this, w, h)
	}
}

func (_this *Form) onMove(x, y int) {
	for _, v := range _this.EvMove[:] {
		v(_this, x, y)
	}
}

func (_this *Form) GetSize() (w, h int) {
	return _this.weight, _this.height
}

func (_this *Form) SetLocation(x, y int) {
	_this.getImpl().SetLocation(x, y)
	_this.x, _this.y = x, y
}
