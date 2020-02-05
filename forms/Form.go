package forms

import (
	mb "qq.2564874169/miniblink"
	plat "qq.2564874169/miniblink/platform"
)

type Form struct {
	BaseUI
	ChildContainer

	EvState map[string]func(target interface{}, state mb.FormState)
	OnState func(mb.FormState)

	impl          plat.IForm
	isInit        bool
	title         string
	showInTaskbar bool
	border        mb.FormBorder
	state         mb.FormState
	startPos      mb.FormStartPosition
}

func (_this *Form) runMain(provider plat.IProvider) {
	provider.RunMain(_this.impl, func() {
		_this.Show()
	})
}

func (_this *Form) getImpl() plat.IForm {
	if _this.isInit == false {
		panic("必须使用Init()初始化 ")
	}
	return _this.impl
}

func (_this *Form) Init() *Form {
	_this.impl = Provider.NewForm()
	_this.BaseUI.init(_this, _this.impl)
	_this.ChildContainer.init(_this)
	_this.EvState = make(map[string]func(interface{}, mb.FormState))
	_this.title = ""
	_this.border = mb.FormBorder_Default
	_this.state = mb.FormState_Normal
	_this.startPos = mb.FormStartPosition_Screen_Center
	_this.showInTaskbar = true
	_this.registerEvents()
	_this.isInit = true
	return _this
}

func (_this *Form) toControls() plat.IControls {
	return _this.impl
}

func (_this *Form) registerEvents() {
	_this.OnState = _this.defOnState
	_this.impl.SetOnState(func(state mb.FormState) {
		_this.state = state
		_this.OnState(state)
	})
}

func (_this *Form) defOnState(state mb.FormState) {
	for _, v := range _this.EvState {
		v(_this, state)
	}
}

func (_this *Form) Show() {
	if _this.Handle == 0 {
		switch _this.startPos {
		case mb.FormStartPosition_Screen_Center:
			scr := Provider.GetScreen()
			x, y := scr.WorkArea.Wdith/2-_this.size.Wdith/2, scr.WorkArea.Height/2-_this.size.Height/2
			_this.impl.SetLocation(x, y)
		}
		_this.impl.Create()
	}
	_this.impl.Show()
}

func (_this *Form) SetTitle(title string) {
	_this.title = title
	_this.getImpl().SetTitle(_this.title)
}

func (_this *Form) SetBorderStyle(style mb.FormBorder) {
	_this.border = style
	_this.getImpl().SetBorderStyle(_this.border)
}

func (_this *Form) GetBorderStyle() mb.FormBorder {
	return _this.border
}

func (_this *Form) ShowInTaskbar(isShow bool) {
	_this.showInTaskbar = isShow
	_this.getImpl().ShowInTaskbar(_this.showInTaskbar)
}

func (_this *Form) SetState(state mb.FormState) {
	if _this.state == state {
		return
	}
	switch state {
	case mb.FormState_Max:
		_this.getImpl().ShowToMax()
	case mb.FormState_Min:
		_this.getImpl().ShowToMin()
	case mb.FormState_Normal:
		_this.getImpl().Show()
	}
}

func (_this *Form) GetState() mb.FormState {
	return _this.state
}

func (_this *Form) SetStartPosition(pos mb.FormStartPosition) {
	_this.startPos = pos
}

func (_this *Form) SetMaximizeBox(isShow bool) {
	_this.getImpl().SetMaximizeBox(isShow)
}

func (_this *Form) SetMinimizeBox(isShow bool) {
	_this.getImpl().SetMinimizeBox(isShow)
}

func (_this *Form) SetIconVisable(isShow bool) {
	_this.getImpl().SetIconVisable(isShow)
}
