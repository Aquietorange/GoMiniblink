package Forms

import (
	gm "GoMiniblink"
	plat "GoMiniblink/CrossPlatform"
)

type Form struct {
	BaseEvents
	onLoad       func()
	onResize     func(w, h int)
	onMove       func(x, y int)
	onMouseMove  func(e gm.MouseEvArgs)
	onMouseDown  func(e gm.MouseEvArgs)
	onMouseUp    func(e gm.MouseEvArgs)
	onMouseWheel func(e gm.MouseEvArgs)
	onMouseClick func(e gm.MouseEvArgs)

	EvState map[string]func(target interface{}, state gm.FormState)
	onState func(gm.FormState)

	impl          plat.IForm
	isInit        bool
	pos           gm.Point
	size          gm.Rect
	title         string
	isFirstShow   bool
	showInTaskbar bool
	border        gm.FormBorder
	state         gm.FormState
	startPos      gm.FormStartPosition
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
	_this.EvState = make(map[string]func(interface{}, gm.FormState))
	_this.impl = Provider.NewForm()
	_this.title = ""
	_this.border = gm.FormBorder_Default
	_this.state = gm.FormState_Normal
	_this.startPos = gm.FormStartPosition_Screen_Center
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

	_this.impl.SetOnMouseClick(func(e gm.MouseEvArgs) {
		_this.onMouseClick(e)
	})
	_this.impl.SetOnMouseWheel(func(e gm.MouseEvArgs) {
		_this.onMouseWheel(e)
	})
	_this.impl.SetOnMouseUp(func(e gm.MouseEvArgs) {
		_this.onMouseUp(e)
	})
	_this.impl.SetOnMouseDown(func(e gm.MouseEvArgs) {
		_this.onMouseDown(e)
	})
	_this.impl.SetOnMouseMove(func(e gm.MouseEvArgs) {
		_this.onMouseMove(e)
	})
	_this.impl.SetOnResize(func(w, h int) {
		_this.size = gm.Rect{Wdith: w, Height: h}
		_this.onResize(w, h)
	})
	_this.impl.SetOnMove(func(x, y int) {
		_this.pos = gm.Point{X: x, Y: y}
		_this.onMove(x, y)
	})
	_this.impl.SetOnState(func(state gm.FormState) {
		_this.onState(_this.state)
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
		_this.impl.Create()
	}
	switch _this.startPos {
	case gm.FormStartPosition_Screen_Center:
		scr := Provider.GetScreen()
		x, y := scr.WorkArea.Wdith/2-_this.size.Wdith/2, scr.WorkArea.Height/2-_this.size.Height/2
		_this.getImpl().SetLocation(x, y)
	}
	_this.getImpl().Show()
}

func (_this *Form) SetSize(rect gm.Rect) {
	_this.size = rect
	_this.getImpl().SetSize(_this.size.Wdith, _this.size.Height)
}

func (_this *Form) GetSize() gm.Rect {
	return _this.size
}

func (_this *Form) SetLocation(pos gm.Point) {
	_this.pos = pos
	_this.getImpl().SetLocation(_this.pos.X, _this.pos.Y)
}

func (_this *Form) GetLocation() gm.Point {
	return _this.pos
}

func (_this *Form) SetTitle(title string) {
	_this.title = title
	_this.getImpl().SetTitle(_this.title)
}

func (_this *Form) SetBorderStyle(style gm.FormBorder) {
	_this.border = style
	_this.getImpl().SetBorderStyle(_this.border)
}

func (_this *Form) GetBorderStyle() gm.FormBorder {
	return _this.border
}

func (_this *Form) ShowInTaskbar(isShow bool) {
	_this.showInTaskbar = isShow
	_this.getImpl().ShowInTaskbar(_this.showInTaskbar)
}

func (_this *Form) SetState(state gm.FormState) {
	_this.state = state
	_this.getImpl().SetState(_this.state)
}

func (_this *Form) GetState() gm.FormState {
	return _this.state
}

func (_this *Form) SetStartPosition(startPos gm.FormStartPosition) {
	_this.startPos = startPos
}

func (_this *Form) SetMaximizeBox(isShow bool) {
	_this.getImpl().SetMaximizeBox(isShow)
}

func (_this *Form) SetMinimizeBox(isShow bool) {
	_this.getImpl().SetMinimizeBox(isShow)
}
