package Forms

import (
	plat "GoMiniblink/Forms/CrossPlatform"
)

type Form struct {
	Controls *controlList
	EvLoad   map[string]func(form *Form)
	OnLoad   func()
	EvResize []func(f *Form, w, h int)
	OnResize func(w, h int)
	EvMove   []func(f *Form, x, y int)
	OnMove   func(x, y int)

	impl          plat.IForm
	isInit        bool
	x             int
	y             int
	weight        int
	height        int
	title         string
	borderStyle   FormBorder
	showInTaskbar bool
	isFirstShow   bool
}

func (_this *Form) RunMain(provider plat.IProvider) {
	provider.RunMain(_this.impl, func() {
		_this.Show()
	})
}

func (_this *Form) getImpl() plat.IForm {
	if _this.isInit == false {
		panic("窗体在使用前必须先调用 Init() ")
	}
	return _this.impl
}

func (_this *Form) Init() *Form {
	_this.EvLoad = make(map[string]func(*Form))
	_this.impl = Provider.NewForm()
	_this.Controls = new(controlList).Init()
	_this.title = ""
	_this.borderStyle = FormBorder_Default
	_this.showInTaskbar = true
	_this.OnLoad = _this.defOnLoad
	_this.OnResize = _this.defOnResize
	_this.OnMove = _this.defOnMove
	_this.registerEvents()
	_this.isInit = true
	return _this
}

func (_this *Form) registerEvents() {
	_this.impl.SetOnCreate(func() {
		_this.OnLoad()
	})
	_this.impl.SetOnResize(func(w, h int) {
		_this.weight, _this.height = w, h
		_this.OnResize(w, h)
	})
	_this.impl.SetOnMove(func(x, y int) {
		_this.x, _this.y = x, y
		_this.OnMove(x, y)
	})
}

func (_this *Form) Invoke(fn func(state interface{}), state interface{}) {
	_this.getImpl().Invoke(fn, state)
}

func (_this *Form) Show() {
	if _this.impl.IsCreate() == false {
		_this.impl.Create()
	}
	_this.impl.Show()
}

func (_this *Form) SetSize(w, h int) {
	_this.weight, _this.height = w, h
	_this.getImpl().SetSize(_this.weight, _this.height)
}

func (_this *Form) GetSize() (w, h int) {
	return _this.weight, _this.height
}

func (_this *Form) SetLocation(x, y int) {
	_this.x, _this.y = x, y
	_this.getImpl().SetLocation(_this.x, _this.y)
}

func (_this *Form) SetTitle(title string) {
	_this.title = title
	_this.getImpl().SetTitle(_this.title)
}

func (_this *Form) SetBorderStyle(style FormBorder) {
	_this.borderStyle = style
	switch _this.borderStyle {
	case FormBorder_Default:
		_this.getImpl().SetBorderStyle(plat.IFormBorder_Default)
	case FormBorder_Disable_Resize:
		_this.getImpl().SetBorderStyle(plat.IFormBorder_Disable_Resize)
	case FormBorder_None:
		_this.getImpl().SetBorderStyle(plat.IFormBorder_None)
	}
}

func (_this *Form) GetBorderStyle() FormBorder {
	return _this.borderStyle
}

func (_this *Form) ShowInTaskbar(isShow bool) {
	_this.showInTaskbar = isShow
	_this.getImpl().ShowInTaskbar(_this.showInTaskbar)
}
