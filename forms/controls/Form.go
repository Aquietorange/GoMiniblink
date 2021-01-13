package controls

import (
	fm "gitee.com/aochulai/GoMiniblink/forms"
	br "gitee.com/aochulai/GoMiniblink/forms/bridge"
)

type Form struct {
	BaseUI
	BaseContainer

	EvState map[string]func(s GUI, state fm.FormState)
	OnState func(state fm.FormState)

	impl          br.Form
	title         string
	showInTaskbar bool
	border        fm.FormBorder
	state         fm.FormState
	startPos      fm.FormStart
}

func (_this *Form) getFormImpl() br.Form {
	return _this.impl
}

func (_this *Form) InitEx(param br.FormParam) *Form {
	_this.impl = App.NewForm(param)
	_this.BaseUI.Init(_this, _this.impl)
	_this.BaseContainer.Init(_this)
	_this.EvState = make(map[string]func(GUI, fm.FormState))
	_this.title = ""
	_this.state = fm.FormState_Normal
	_this.border = fm.FormBorder_Default
	_this.startPos = fm.FormStart_Default

	_this.showInTaskbar = true
	_this.setOn()
	return _this
}

func (_this *Form) Init() *Form {
	return _this.InitEx(br.FormParam{})
}

func (_this *Form) SetTopMost(isTop bool) {
	_this.impl.SetTopMost(isTop)
}

func (_this *Form) NoneBorderResize() {
	_this.impl.NoneBorderResize()
}

func (_this *Form) toControls() br.Controls {
	return _this.impl
}

func (_this *Form) setOn() {
	_this.OnState = _this.defOnState
	var bakState br.FormStateProc
	bakState = _this.impl.SetOnState(func(state fm.FormState) {
		if bakState != nil {
			bakState(state)
		}
		_this.state = state
		_this.OnState(state)
	})
	bakLoad := _this.OnLoad
	_this.OnLoad = func() {
		switch _this.startPos {
		case fm.FormStart_Screen_Center:
			scr := App.GetScreen()
			size := _this.GetBound().Rect
			x, y := scr.WorkArea.Width/2-size.Width/2, scr.WorkArea.Height/2-size.Height/2
			_this.impl.SetLocation(x, y)
		case fm.FormStart_Default:
			_this.impl.SetLocation(100, 100)
		}
		bakLoad()
	}
}

func (_this *Form) defOnState(state fm.FormState) {
	for _, v := range _this.EvState {
		v(_this, state)
	}
}

func (_this *Form) SetTitle(title string) {
	_this.title = title
	_this.impl.SetTitle(_this.title)
}

func (_this *Form) SetBorderStyle(style fm.FormBorder) {
	_this.border = style
	_this.impl.SetBorderStyle(_this.border)
}

func (_this *Form) GetBorderStyle() fm.FormBorder {
	return _this.border
}

func (_this *Form) SetState(state fm.FormState) {
	if _this.state == state {
		return
	}
	switch state {
	case fm.FormState_Max:
		_this.impl.ShowToMax()
	case fm.FormState_Min:
		_this.impl.ShowToMin()
	case fm.FormState_Normal:
		_this.impl.Show()
	}
}

func (_this *Form) GetState() fm.FormState {
	return _this.state
}

func (_this *Form) SetStartPosition(pos fm.FormStart) {
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
	_this.SetStartPosition(fm.FormStart_Screen_Center)
	_this.impl.ShowDialog()
}
