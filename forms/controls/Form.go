package controls

import (
	f "qq2564874169/goMiniblink/forms"
	p "qq2564874169/goMiniblink/forms/platform"
)

type Form struct {
	FormBaseUI
	DefChildContainer

	EvState map[string]func(target interface{}, state f.FormState)
	OnState func(state f.FormState)

	impl          p.Form
	isInit        bool
	title         string
	showInTaskbar bool
	border        f.FormBorder
	state         f.FormState
	startPos      f.FormStartPosition
}

func (_this *Form) getImpl() p.Form {
	if _this.isInit == false {
		panic("必须使用Init()初始化 ")
	}
	return _this.impl
}

func (_this *Form) Init() *Form {
	_this.impl = App.NewForm()
	_this.FormBaseUI.init(_this, _this.impl)
	_this.DefChildContainer.init(_this)
	_this.EvState = make(map[string]func(interface{}, f.FormState))
	_this.title = ""
	_this.border = f.FormBorder_Default
	_this.state = f.FormState_Normal
	_this.startPos = f.FormStartPosition_Screen_Center
	_this.showInTaskbar = true
	_this.registerEvents()
	_this.isInit = true
	return _this
}

func (_this *Form) toControls() p.Controls {
	return _this.impl
}

func (_this *Form) registerEvents() {
	_this.OnState = _this.defOnState
	var bakState p.FormStateProc
	bakState = _this.impl.SetOnState(func(state f.FormState) {
		if bakState != nil {
			bakState(state)
		}
		_this.state = state
		_this.OnState(state)
	})
	var bakCreate p.WindowCreateProc
	bakCreate = _this.impl.SetOnCreate(func(handle uintptr) bool {
		switch _this.startPos {
		case f.FormStartPosition_Screen_Center:
			scr := App.GetScreen()
			size := _this.GetSize()
			x, y := scr.WorkArea.Width/2-size.Width/2, scr.WorkArea.Height/2-size.Height/2
			_this.impl.SetLocation(x, y)
		}
		if bakCreate != nil {
			return bakCreate(handle)
		}
		return false
	})
}

func (_this *Form) defOnState(state f.FormState) {
	for _, v := range _this.EvState {
		v(_this, state)
	}
}

func (_this *Form) SetTitle(title string) {
	_this.title = title
	_this.getImpl().SetTitle(_this.title)
}

func (_this *Form) SetBorderStyle(style f.FormBorder) {
	_this.border = style
	_this.getImpl().SetBorderStyle(_this.border)
}

func (_this *Form) GetBorderStyle() f.FormBorder {
	return _this.border
}

func (_this *Form) ShowInTaskbar(isShow bool) {
	_this.showInTaskbar = isShow
	_this.getImpl().ShowInTaskbar(_this.showInTaskbar)
}

func (_this *Form) SetState(state f.FormState) {
	if _this.state == state {
		return
	}
	switch state {
	case f.FormState_Max:
		_this.getImpl().ShowToMax()
	case f.FormState_Min:
		_this.getImpl().ShowToMin()
	case f.FormState_Normal:
		_this.getImpl().Show()
	}
}

func (_this *Form) GetState() f.FormState {
	return _this.state
}

func (_this *Form) SetStartPosition(pos f.FormStartPosition) {
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

func (_this *Form) Close() {
	_this.impl.Close()
}
