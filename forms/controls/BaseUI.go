package controls

import (
	f "qq2564874169/goMiniblink/forms"
	p "qq2564874169/goMiniblink/forms/platform"
)

type BaseUI struct {
	EvKeyDown map[string]func(s GUI, e *f.KeyEvArgs)
	OnKeyDown func(e *f.KeyEvArgs)

	EvKeyUp map[string]func(s GUI, e *f.KeyEvArgs)
	OnKeyUp func(e *f.KeyEvArgs)

	EvKeyPress map[string]func(s GUI, e *f.KeyPressEvArgs)
	OnKeyPress func(e *f.KeyPressEvArgs)

	EvShow map[string]func(s GUI)
	OnShow func()

	EvResize map[string]func(s GUI, e f.Rect)
	OnResize func(e f.Rect)

	EvMove map[string]func(s GUI, e f.Point)
	OnMove func(e f.Point)

	EvMouseMove map[string]func(s GUI, e *f.MouseEvArgs)
	OnMouseMove func(e *f.MouseEvArgs)

	EvMouseDown map[string]func(s GUI, e *f.MouseEvArgs)
	OnMouseDown func(e *f.MouseEvArgs)

	EvMouseUp map[string]func(s GUI, e *f.MouseEvArgs)
	OnMouseUp func(e *f.MouseEvArgs)

	EvMouseWheel map[string]func(s GUI, e *f.MouseEvArgs)
	OnMouseWheel func(e *f.MouseEvArgs)

	EvMouseClick map[string]func(s GUI, e *f.MouseEvArgs)
	OnMouseClick func(e *f.MouseEvArgs)

	EvPaint map[string]func(s GUI, e f.PaintEvArgs)
	OnPaint func(e f.PaintEvArgs)

	EvFocus map[string]func(s GUI)
	OnFocus func()

	EvLostFocus map[string]func(s GUI)
	OnLostFocus func()

	OnSetCursor           func() bool
	OnImeStartComposition func() bool

	instance GUI
	impl     p.Window
}

func (_this *BaseUI) Init(instance GUI, impl p.Window) *BaseUI {
	_this.instance = instance
	_this.impl = impl

	_this.EvKeyPress = make(map[string]func(s GUI, e *f.KeyPressEvArgs))
	_this.OnKeyPress = _this.defOnKeyPress

	_this.EvKeyDown = make(map[string]func(s GUI, e *f.KeyEvArgs))
	_this.OnKeyDown = _this.defOnKeyDown

	_this.EvKeyUp = make(map[string]func(s GUI, e *f.KeyEvArgs))
	_this.OnKeyUp = _this.defOnKeyUp

	_this.EvPaint = make(map[string]func(s GUI, e f.PaintEvArgs))
	_this.OnPaint = _this.defOnPaint

	_this.EvShow = make(map[string]func(s GUI))
	_this.OnShow = _this.defOnLoad

	_this.EvResize = make(map[string]func(s GUI, e f.Rect))
	_this.OnResize = _this.defOnResize

	_this.EvMove = make(map[string]func(s GUI, e f.Point))
	_this.OnMove = _this.defOnMove

	_this.EvMouseMove = make(map[string]func(s GUI, e *f.MouseEvArgs))
	_this.OnMouseMove = _this.defOnMouseMove

	_this.EvMouseDown = make(map[string]func(s GUI, e *f.MouseEvArgs))
	_this.OnMouseDown = _this.defOnMouseDown

	_this.EvMouseUp = make(map[string]func(s GUI, e *f.MouseEvArgs))
	_this.OnMouseUp = _this.defOnMouseUp

	_this.EvMouseWheel = make(map[string]func(s GUI, e *f.MouseEvArgs))
	_this.OnMouseWheel = _this.defOnMouseWheel

	_this.EvMouseClick = make(map[string]func(s GUI, e *f.MouseEvArgs))
	_this.OnMouseClick = _this.defOnMouseClick

	_this.EvFocus = make(map[string]func(s GUI))
	_this.OnFocus = _this.defOnFocus

	_this.EvLostFocus = make(map[string]func(s GUI))
	_this.OnLostFocus = _this.defOnLostFocus

	var bakImeStart p.WindowImeStartCompositionProc
	bakImeStart = _this.impl.SetOnImeStartComposition(func() bool {
		b := false
		if bakImeStart != nil {
			b = bakImeStart()
		}
		if !b && _this.OnImeStartComposition != nil && _this.OnImeStartComposition() {
			b = true
		}
		return b
	})

	var bakLostFocus p.WindowLostFocusProc
	bakLostFocus = _this.impl.SetOnLostFocus(func() bool {
		b := false
		if bakLostFocus != nil {
			b = bakLostFocus()
		}
		if !b && _this.OnLostFocus != nil {
			_this.OnLostFocus()
		}
		return b
	})

	var bakFocus p.WindowFocusProc
	bakFocus = _this.impl.SetOnFocus(func() bool {
		b := false
		if bakFocus != nil {
			b = bakFocus()
		}
		if !b && _this.OnFocus != nil {
			_this.OnFocus()
		}
		return b
	})

	var bakOnCursor p.WindowSetCursorProc
	bakOnCursor = _this.impl.SetOnCursor(func() bool {
		b := false
		if bakOnCursor != nil {
			b = bakOnCursor()
		}
		if !b && _this.OnSetCursor != nil && _this.OnSetCursor() {
			b = true
		}
		return b
	})

	var bakKeyPress p.WindowKeyPressProc
	bakKeyPress = _this.impl.SetOnKeyPress(func(e *f.KeyPressEvArgs) {
		if bakKeyPress != nil {
			bakKeyPress(e)
		}
		if !e.IsHandle {
			_this.OnKeyPress(e)
		}
	})

	var bakKeyUp p.WindowKeyUpProc
	bakKeyUp = _this.impl.SetOnKeyUp(func(e *f.KeyEvArgs) {
		if bakKeyUp != nil {
			bakKeyUp(e)
		}
		if !e.IsHandle {
			_this.OnKeyUp(e)
		}
	})

	var bakKeyDown p.WindowKeyDownProc
	bakKeyDown = _this.impl.SetOnKeyDown(func(e *f.KeyEvArgs) {
		if bakKeyDown != nil {
			bakKeyDown(e)
		}
		if !e.IsHandle {
			_this.OnKeyDown(e)
		}
	})

	var bakPaint p.WindowPaintProc
	bakPaint = _this.impl.SetOnPaint(func(e f.PaintEvArgs) bool {
		b := false
		if bakPaint != nil {
			b = bakPaint(e)
		}
		if !b {
			_this.OnPaint(e)
		}
		return b
	})

	var bakMouseClick p.WindowMouseClickProc
	bakMouseClick = _this.impl.SetOnMouseClick(func(e *f.MouseEvArgs) {
		if bakMouseClick != nil {
			bakMouseClick(e)
		}
		if !e.IsHandle {
			_this.OnMouseClick(e)
		}
	})

	var bakMouseWheel p.WindowMouseWheelProc
	bakMouseWheel = _this.impl.SetOnMouseWheel(func(e *f.MouseEvArgs) {
		if bakMouseWheel != nil {
			bakMouseWheel(e)
		}
		if !e.IsHandle {
			_this.OnMouseWheel(e)
		}
	})

	var bakMouseUp p.WindowMouseUpProc
	bakMouseUp = _this.impl.SetOnMouseUp(func(e *f.MouseEvArgs) {
		if bakMouseUp != nil {
			bakMouseUp(e)
		}
		if !e.IsHandle {
			_this.OnMouseUp(e)
		}
	})

	var bakMouseDown p.WindowMouseDownProc
	bakMouseDown = _this.impl.SetOnMouseDown(func(e *f.MouseEvArgs) {
		if bakMouseDown != nil {
			bakMouseDown(e)
		}
		if !e.IsHandle {
			_this.OnMouseDown(e)
		}
	})

	var bakMouseMove p.WindowMouseMoveProc
	bakMouseMove = _this.impl.SetOnMouseMove(func(e *f.MouseEvArgs) {
		if bakMouseMove != nil {
			bakMouseMove(e)
		}
		if !e.IsHandle {
			_this.OnMouseMove(e)
		}
	})

	var bakResize p.WindowResizeProc
	bakResize = _this.impl.SetOnResize(func(e f.Rect) bool {
		b := false
		if bakResize != nil {
			b = bakResize(e)
		}
		if !b {
			_this.OnResize(e)
		}
		return b
	})

	var bakMove p.WindowMoveProc
	bakMove = _this.impl.SetOnMove(func(e f.Point) bool {
		b := false
		if bakMove != nil {
			b = bakMove(e)
		}
		if !b {
			_this.OnMove(e)
		}
		return b
	})

	var bakShow p.WindowShowProc
	bakShow = _this.impl.SetOnShow(func() {
		if bakShow != nil {
			bakShow()
		}
		_this.OnShow()
	})
	return _this
}

func (_this *BaseUI) Enable(enable bool) {
	_this.impl.Enable(enable)
}

func (_this *BaseUI) IsEnable() bool {
	return _this.impl.IsEnable()
}

func (_this *BaseUI) CreateGraphics() f.Graphics {
	return _this.impl.CreateGraphics()
}

func (_this *BaseUI) SetCursor(cursor f.CursorType) {
	_this.impl.SetCursor(cursor)
}

func (_this *BaseUI) GetCursor() f.CursorType {
	return _this.impl.GetCursor()
}

func (_this *BaseUI) GetHandle() uintptr {
	return _this.impl.GetHandle()
}

func (_this *BaseUI) SetLocation(x, y int) {
	_this.impl.SetLocation(x, y)
}

func (_this *BaseUI) GetLocation() f.Point {
	x, y := _this.impl.GetLocation()
	return f.Point{
		X: x,
		Y: y,
	}
}

func (_this *BaseUI) GetSize() f.Rect {
	w, h := _this.impl.GetSize()
	return f.Rect{
		Width:  w,
		Height: h,
	}
}

func (_this *BaseUI) SetSize(width, height int) {
	_this.impl.SetSize(width, height)
}

func (_this *BaseUI) SetBgColor(color int32) {
	_this.impl.SetBgColor(color)
}

func (_this *BaseUI) Invoke(fn func(state interface{}), state interface{}) {
	_this.impl.Invoke(fn, state)
}

func (_this *BaseUI) IsInvoke() bool {
	return _this.impl.IsInvoke()
}

func (_this *BaseUI) Show() {
	_this.impl.Show()
}

func (_this *BaseUI) Hide() {
	_this.impl.Hide()
}
