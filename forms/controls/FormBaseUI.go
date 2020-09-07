package controls

import (
	f "qq2564874169/goMiniblink/forms"
	p "qq2564874169/goMiniblink/forms/platform"
)

type FormBaseUI struct {
	EvKeyDown map[string]func(sender interface{}, e *f.KeyEvArgs)
	OnKeyDown func(e *f.KeyEvArgs)

	EvKeyUp map[string]func(sender interface{}, e *f.KeyEvArgs)
	OnKeyUp func(e *f.KeyEvArgs)

	EvKeyPress map[string]func(sender interface{}, e *f.KeyPressEvArgs)
	OnKeyPress func(e *f.KeyPressEvArgs)

	EvLoad map[string]func(sender interface{})
	OnLoad func()

	EvResize map[string]func(sender interface{}, e f.Rect)
	OnResize func(e f.Rect)

	EvMove map[string]func(sender interface{}, e f.Point)
	OnMove func(e f.Point)

	EvMouseMove map[string]func(sender interface{}, e *f.MouseEvArgs)
	OnMouseMove func(e *f.MouseEvArgs)

	EvMouseDown map[string]func(sender interface{}, e *f.MouseEvArgs)
	OnMouseDown func(e *f.MouseEvArgs)

	EvMouseUp map[string]func(sender interface{}, e *f.MouseEvArgs)
	OnMouseUp func(e *f.MouseEvArgs)

	EvMouseWheel map[string]func(sender interface{}, e *f.MouseEvArgs)
	OnMouseWheel func(e *f.MouseEvArgs)

	EvMouseClick map[string]func(sender interface{}, e *f.MouseEvArgs)
	OnMouseClick func(e *f.MouseEvArgs)

	EvPaint map[string]func(sender interface{}, e f.PaintEvArgs)
	OnPaint func(e f.PaintEvArgs)

	EvFocus map[string]func(sender interface{})
	OnFocus func()

	EvLostFocus map[string]func(sender interface{})
	OnLostFocus func()

	OnSetCursor           func() bool
	OnImeStartComposition func() bool

	instance interface{}
	impl     p.Window
}

func (_this *FormBaseUI) init(instance interface{}, impl p.Window) *FormBaseUI {
	_this.instance = instance
	_this.impl = impl

	_this.EvKeyPress = make(map[string]func(target interface{}, e *f.KeyPressEvArgs))
	_this.OnKeyPress = _this.defOnKeyPress

	_this.EvKeyDown = make(map[string]func(target interface{}, e *f.KeyEvArgs))
	_this.OnKeyDown = _this.defOnKeyDown

	_this.EvKeyUp = make(map[string]func(target interface{}, e *f.KeyEvArgs))
	_this.OnKeyUp = _this.defOnKeyUp

	_this.EvPaint = make(map[string]func(target interface{}, e f.PaintEvArgs))
	_this.OnPaint = _this.defOnPaint

	_this.EvLoad = make(map[string]func(target interface{}))
	_this.OnLoad = _this.defOnLoad

	_this.EvResize = make(map[string]func(target interface{}, e f.Rect))
	_this.OnResize = _this.defOnResize

	_this.EvMove = make(map[string]func(target interface{}, e f.Point))
	_this.OnMove = _this.defOnMove

	_this.EvMouseMove = make(map[string]func(target interface{}, e *f.MouseEvArgs))
	_this.OnMouseMove = _this.defOnMouseMove

	_this.EvMouseDown = make(map[string]func(target interface{}, e *f.MouseEvArgs))
	_this.OnMouseDown = _this.defOnMouseDown

	_this.EvMouseUp = make(map[string]func(target interface{}, e *f.MouseEvArgs))
	_this.OnMouseUp = _this.defOnMouseUp

	_this.EvMouseWheel = make(map[string]func(target interface{}, e *f.MouseEvArgs))
	_this.OnMouseWheel = _this.defOnMouseWheel

	_this.EvMouseClick = make(map[string]func(target interface{}, e *f.MouseEvArgs))
	_this.OnMouseClick = _this.defOnMouseClick

	_this.EvFocus = make(map[string]func(target interface{}))
	_this.OnFocus = _this.defOnFocus

	_this.EvLostFocus = make(map[string]func(target interface{}))
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

	var bakLoad p.WindowLoadProc
	bakLoad = _this.impl.SetOnLoad(func() {
		if bakLoad != nil {
			bakLoad()
		}
		_this.OnLoad()
	})
	return _this
}

func (_this *FormBaseUI) SetMouseEnable(enable bool) {
	_this.impl.SetMouseEnable(enable)
}
func (_this *FormBaseUI) GetMouseEnable() bool {
	return _this.impl.GetMouseEnable()
}

func (_this *FormBaseUI) CreateGraphics() f.Graphics {
	return _this.impl.CreateGraphics()
}

func (_this *FormBaseUI) SetCursor(cursor f.CursorType) {
	_this.impl.SetCursor(cursor)
}

func (_this *FormBaseUI) GetHandle() uintptr {
	return _this.impl.GetHandle()
}

func (_this *FormBaseUI) SetLocation(x, y int) {
	_this.impl.SetLocation(x, y)
}

func (_this *FormBaseUI) GetLocation() f.Point {
	x, y := _this.impl.GetLocation()
	return f.Point{
		X: x,
		Y: y,
	}
}

func (_this *FormBaseUI) GetSize() f.Rect {
	w, h := _this.impl.GetSize()
	return f.Rect{
		Width:  w,
		Height: h,
	}
}

func (_this *FormBaseUI) SetSize(width, height int) {
	_this.impl.SetSize(width, height)
}

func (_this *FormBaseUI) SetBgColor(color int) {
	_this.impl.SetBgColor(color)
}

func (_this *FormBaseUI) Invoke(fn func(state interface{}), state interface{}) {
	_this.impl.Invoke(fn, state)
}

func (_this *FormBaseUI) IsInvoke() bool {
	return _this.impl.IsInvoke()
}
