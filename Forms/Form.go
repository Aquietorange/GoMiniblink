package Forms

import (
	MB "GoMiniblink"
	plat "GoMiniblink/CrossPlatform"
)

type Form struct {
	BaseEvents
	onLoad       func()
	onResize     func(w, h int)
	onMove       func(x, y int)
	onMouseMove  func(e MB.MouseEvArgs)
	onMouseDown  func(e MB.MouseEvArgs)
	onMouseUp    func(e MB.MouseEvArgs)
	onMouseWheel func(e MB.MouseEvArgs)
	onMouseClick func(e MB.MouseEvArgs)

	EvState map[string]func(target interface{}, state MB.FormState)
	onState func(MB.FormState)

	impl          plat.IForm
	isInit        bool
	pos           MB.Point
	size          MB.Rect
	title         string
	isFirstShow   bool
	showInTaskbar bool
	border        MB.FormBorder
	state         MB.FormState
	startPos      MB.FormStartPosition
}

func (_this *Form) runMain(provider plat.IProvider) {
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
	_this.BaseEvents.init()
	_this.EvState = make(map[string]func(interface{}, MB.FormState))
	_this.impl = Provider.NewForm()
	_this.title = ""
	_this.border = MB.FormBorder_Default
	_this.state = MB.FormState_Normal
	_this.startPos = MB.FormStartPosition_Screen_Center
	_this.showInTaskbar = true
	_this.registerEvents()
	_this.isInit = true
	return _this
}

func (_this *Form) registerEvents() {
	_this.onLoad = _this.defOnLoad
	_this.onResize = _this.defOnResize
	_this.onMove = _this.defOnMove
	_this.onState = _this.defOnState
	_this.onMouseMove = _this.defOnMouseMove
	_this.onMouseDown = _this.defOnMouseDown
	_this.onMouseUp = _this.defOnMouseUp
	_this.onMouseWheel = _this.defOnMouseWheel
	_this.onMouseClick = _this.defOnMouseClick

	_this.impl.SetOnMouseClick(func(e MB.MouseEvArgs) {
		_this.onMouseClick(e)
	})
	_this.impl.SetOnMouseWheel(func(e MB.MouseEvArgs) {
		_this.onMouseWheel(e)
	})
	_this.impl.SetOnMouseUp(func(e MB.MouseEvArgs) {
		_this.onMouseUp(e)
	})
	_this.impl.SetOnMouseDown(func(e MB.MouseEvArgs) {
		_this.onMouseDown(e)
	})
	_this.impl.SetOnMouseMove(func(e MB.MouseEvArgs) {
		_this.onMouseMove(e)
	})
	_this.impl.SetOnResize(func(w, h int) {
		_this.size = MB.Rect{Wdith: w, Height: h}
		_this.onResize(w, h)
	})
	_this.impl.SetOnMove(func(x, y int) {
		_this.pos = MB.Point{X: x, Y: y}
		_this.onMove(x, y)
	})
	_this.impl.SetOnState(func(state MB.FormState) {
		_this.state = state
		_this.onState(state)
	})
	_this.impl.SetOnCreate(func() {
		_this.onLoad()
	})
}

func (_this *Form) Invoke(fn func(state interface{}), state interface{}) {
	_this.getImpl().Invoke(fn, state)
}

func (_this *Form) Show() {
	if _this.impl.IsCreate() == false {
		switch _this.startPos {
		case MB.FormStartPosition_Screen_Center:
			scr := Provider.GetScreen()
			x, y := scr.WorkArea.Wdith/2-_this.size.Wdith/2, scr.WorkArea.Height/2-_this.size.Height/2
			_this.impl.SetLocation(x, y)
		}
		_this.impl.Create()
	}
	_this.getImpl().Show()
}

func (_this *Form) SetSize(rect MB.Rect) {
	_this.size = rect
	_this.getImpl().SetSize(_this.size.Wdith, _this.size.Height)
}

func (_this *Form) GetSize() MB.Rect {
	return _this.size
}

func (_this *Form) SetLocation(pos MB.Point) {
	_this.pos = pos
	_this.getImpl().SetLocation(_this.pos.X, _this.pos.Y)
}

func (_this *Form) GetLocation() MB.Point {
	return _this.pos
}

func (_this *Form) SetTitle(title string) {
	_this.title = title
	_this.getImpl().SetTitle(_this.title)
}

func (_this *Form) SetBorderStyle(style MB.FormBorder) {
	_this.border = style
	_this.getImpl().SetBorderStyle(_this.border)
}

func (_this *Form) GetBorderStyle() MB.FormBorder {
	return _this.border
}

func (_this *Form) ShowInTaskbar(isShow bool) {
	_this.showInTaskbar = isShow
	_this.getImpl().ShowInTaskbar(_this.showInTaskbar)
}

func (_this *Form) SetState(state MB.FormState) {
	if _this.state == state {
		return
	}
	switch state {
	case MB.FormState_Max:
		_this.getImpl().ShowToMax()
	case MB.FormState_Min:
		_this.getImpl().ShowToMin()
	case MB.FormState_Normal:
		_this.getImpl().Show()
	}
}

func (_this *Form) GetState() MB.FormState {
	return _this.state
}

func (_this *Form) SetStartPosition(pos MB.FormStartPosition) {
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
