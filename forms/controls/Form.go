package controls

import (
	f "qq2564874169/goMiniblink/forms"
	br "qq2564874169/goMiniblink/forms/bridge"
)

type Form struct {
	BaseUI
	BaseContainer

	EvState map[string]func(target interface{}, state f.FormState)
	OnState func(state f.FormState)

	impl          br.Form
	isInit        bool
	title         string
	showInTaskbar bool
	border        f.FormBorder
	state         f.FormState
	startPos      f.FormStartPosition
}

func (_this *Form) getFormImpl() br.Form {
	return _this.impl
}

func (_this *Form) InitEx(param br.FormParam) *Form {
	_this.impl = App.NewForm(param)
	_this.BaseUI.Init(_this, _this.impl)
	_this.BaseContainer.Init(_this)
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

func (_this *Form) Init() *Form {
	return _this.InitEx(br.FormParam{})
}

func (_this *Form) NoneBorderResize() {
	_this.impl.NoneBorderResize()
}

func (_this *Form) toControls() br.Controls {
	return _this.impl
}

func (_this *Form) registerEvents() {
	_this.OnState = _this.defOnState
	var bakState br.FormStateProc
	bakState = _this.impl.SetOnState(func(state f.FormState) {
		if bakState != nil {
			bakState(state)
		}
		_this.state = state
		_this.OnState(state)
	})
	var bakShow br.WindowShowProc
	bakShow = _this.impl.SetOnShow(func() {
		switch _this.startPos {
		case f.FormStartPosition_Screen_Center:
			scr := App.GetScreen()
			size := _this.GetSize()
			x, y := scr.WorkArea.Width/2-size.Width/2, scr.WorkArea.Height/2-size.Height/2
			_this.impl.SetLocation(x, y)
		}
		if bakShow != nil {
			bakShow()
		}
	})
}

func (_this *Form) defOnState(state f.FormState) {
	for _, v := range _this.EvState {
		v(_this, state)
	}
}

func (_this *Form) SetTitle(title string) {
	_this.title = title
	_this.impl.SetTitle(_this.title)
}

func (_this *Form) SetBorderStyle(style f.FormBorder) {
	_this.border = style
	_this.impl.SetBorderStyle(_this.border)
}

func (_this *Form) GetBorderStyle() f.FormBorder {
	return _this.border
}

func (_this *Form) SetState(state f.FormState) {
	if _this.state == state {
		return
	}
	switch state {
	case f.FormState_Max:
		_this.impl.ShowToMax()
	case f.FormState_Min:
		_this.impl.ShowToMin()
	case f.FormState_Normal:
		_this.impl.Show()
	}
}

func (_this *Form) GetState() f.FormState {
	return _this.state
}

func (_this *Form) SetStartPosition(pos f.FormStartPosition) {
	_this.startPos = pos
}

func (_this *Form) SetMaximizeBox(isShow bool) {
	_this.impl.SetMaximizeBox(isShow)
}

func (_this *Form) SetMinimizeBox(isShow bool) {
	_this.impl.SetMinimizeBox(isShow)
}

func (_this *Form) Close() {
	_this.impl.Close()
}

func (_this *Form) SetIcon(file string) {
	_this.impl.SetIcon(file)
}

func (_this *Form) ShowDialog() {
	_this.SetStartPosition(f.FormStartPosition_Screen_Center)
	_this.impl.ShowDialog()
}
