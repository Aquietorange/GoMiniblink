package controls

import (
	fm "gitee.com/aochulai/GoMiniblink/forms"
	br "gitee.com/aochulai/GoMiniblink/forms/bridge"
)

type BaseUI struct {
	EvLoad map[string]func(s GUI)
	OnLoad func()

	EvDestroy map[string]func(s GUI)
	OnDestroy func()

	EvKeyDown map[string]func(s GUI, e *fm.KeyEvArgs)
	OnKeyDown func(e *fm.KeyEvArgs)

	EvKeyUp map[string]func(s GUI, e *fm.KeyEvArgs)
	OnKeyUp func(e *fm.KeyEvArgs)

	EvKeyPress map[string]func(s GUI, e *fm.KeyPressEvArgs)
	OnKeyPress func(e *fm.KeyPressEvArgs)

	EvShow map[string]func(s GUI)
	OnShow func()

	EvResize map[string]func(s GUI, e fm.Rect)
	OnResize func(e fm.Rect)

	EvMove map[string]func(s GUI, e fm.Point)
	OnMove func(e fm.Point)

	EvMouseMove map[string]func(s GUI, e *fm.MouseEvArgs)
	OnMouseMove func(e *fm.MouseEvArgs)

	EvMouseDown map[string]func(s GUI, e *fm.MouseEvArgs)
	OnMouseDown func(e *fm.MouseEvArgs)

	EvMouseUp map[string]func(s GUI, e *fm.MouseEvArgs)
	OnMouseUp func(e *fm.MouseEvArgs)

	EvMouseWheel map[string]func(s GUI, e *fm.MouseEvArgs)
	OnMouseWheel func(e *fm.MouseEvArgs)

	EvMouseClick map[string]func(s GUI, e *fm.MouseEvArgs)
	OnMouseClick func(e *fm.MouseEvArgs)

	EvPaint map[string]func(s GUI, e fm.PaintEvArgs)
	OnPaint func(e fm.PaintEvArgs)

	EvFocus map[string]func(s GUI)
	OnFocus func()

	EvLostFocus map[string]func(s GUI)
	OnLostFocus func()

	OnSetCursor           func() bool
	OnImeStartComposition func() bool

	instance GUI
	impl     br.Window
	skipLoad bool
	parent   GUI
	owner    GUI
}

func (_this *BaseUI) Init(instance GUI, impl br.Window) *BaseUI {
	_this.instance = instance
	_this.impl = impl

	_this.EvLoad = make(map[string]func(GUI))
	_this.OnLoad = _this.defOnLoad

	_this.EvDestroy = make(map[string]func(GUI))
	_this.OnDestroy = _this.defOnDestroy

	_this.EvKeyPress = make(map[string]func(GUI, *fm.KeyPressEvArgs))
	_this.OnKeyPress = _this.defOnKeyPress

	_this.EvKeyDown = make(map[string]func(GUI, *fm.KeyEvArgs))
	_this.OnKeyDown = _this.defOnKeyDown

	_this.EvKeyUp = make(map[string]func(GUI, *fm.KeyEvArgs))
	_this.OnKeyUp = _this.defOnKeyUp

	_this.EvPaint = make(map[string]func(GUI, fm.PaintEvArgs))
	_this.OnPaint = _this.defOnPaint

	_this.EvShow = make(map[string]func(GUI))
	_this.OnShow = _this.defOnShow

	_this.EvResize = make(map[string]func(GUI, fm.Rect))
	_this.OnResize = _this.defOnResize

	_this.EvMove = make(map[string]func(GUI, fm.Point))
	_this.OnMove = _this.defOnMove

	_this.EvMouseMove = make(map[string]func(GUI, *fm.MouseEvArgs))
	_this.OnMouseMove = _this.defOnMouseMove

	_this.EvMouseDown = make(map[string]func(GUI, *fm.MouseEvArgs))
	_this.OnMouseDown = _this.defOnMouseDown

	_this.EvMouseUp = make(map[string]func(GUI, *fm.MouseEvArgs))
	_this.OnMouseUp = _this.defOnMouseUp

	_this.EvMouseWheel = make(map[string]func(GUI, *fm.MouseEvArgs))
	_this.OnMouseWheel = _this.defOnMouseWheel

	_this.EvMouseClick = make(map[string]func(GUI, *fm.MouseEvArgs))
	_this.OnMouseClick = _this.defOnMouseClick

	_this.EvFocus = make(map[string]func(GUI))
	_this.OnFocus = _this.defOnFocus

	_this.EvLostFocus = make(map[string]func(GUI))
	_this.OnLostFocus = _this.defOnLostFocus

	var bakDestroy br.WindowDestroyProc
	bakDestroy = _this.impl.SetOnDestroy(func() {
		if bakDestroy != nil {
			bakDestroy()
		}
		if _this.OnDestroy != nil {
			_this.OnDestroy()
		}
	})

	var bakImeStart br.WindowImeStartCompositionProc
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

	var bakLostFocus br.WindowLostFocusProc
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

	var bakFocus br.WindowFocusProc
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

	var bakOnCursor br.WindowSetCursorProc
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

	var bakKeyPress br.WindowKeyPressProc
	bakKeyPress = _this.impl.SetOnKeyPress(func(e *fm.KeyPressEvArgs) {
		println("press=", e.KeyChar)
		if bakKeyPress != nil {
			bakKeyPress(e)
		}
		if !e.IsHandle && _this.OnKeyPress != nil {
			_this.OnKeyPress(e)
		}
	})

	var bakKeyUp br.WindowKeyUpProc
	bakKeyUp = _this.impl.SetOnKeyUp(func(e *fm.KeyEvArgs) {
		if bakKeyUp != nil {
			bakKeyUp(e)
		}
		if !e.IsHandle && _this.OnKeyUp != nil {
			_this.OnKeyUp(e)
		}
	})

	var bakKeyDown br.WindowKeyDownProc
	bakKeyDown = _this.impl.SetOnKeyDown(func(e *fm.KeyEvArgs) {
		if bakKeyDown != nil {
			bakKeyDown(e)
		}
		if !e.IsHandle && _this.OnKeyDown != nil {
			_this.OnKeyDown(e)
		}
	})

	var bakPaint br.WindowPaintProc
	bakPaint = _this.impl.SetOnPaint(func(e fm.PaintEvArgs) bool {
		b := false
		if bakPaint != nil {
			b = bakPaint(e)
		}
		if !b && _this.OnPaint != nil {
			_this.OnPaint(e)
		}
		return b
	})

	var bakMouseClick br.WindowMouseClickProc
	bakMouseClick = _this.impl.SetOnMouseClick(func(e *fm.MouseEvArgs) {
		if bakMouseClick != nil {
			bakMouseClick(e)
		}
		if !e.IsHandle && _this.OnMouseClick != nil {
			_this.OnMouseClick(e)
		}
	})

	var bakMouseWheel br.WindowMouseWheelProc
	bakMouseWheel = _this.impl.SetOnMouseWheel(func(e *fm.MouseEvArgs) {
		if bakMouseWheel != nil {
			bakMouseWheel(e)
		}
		if !e.IsHandle && _this.OnMouseWheel != nil {
			_this.OnMouseWheel(e)
		}
	})

	var bakMouseUp br.WindowMouseUpProc
	bakMouseUp = _this.impl.SetOnMouseUp(func(e *fm.MouseEvArgs) {
		if bakMouseUp != nil {
			bakMouseUp(e)
		}
		if !e.IsHandle && _this.OnMouseUp != nil {
			_this.OnMouseUp(e)
		}
	})

	var bakMouseDown br.WindowMouseDownProc
	bakMouseDown = _this.impl.SetOnMouseDown(func(e *fm.MouseEvArgs) {
		if bakMouseDown != nil {
			bakMouseDown(e)
		}
		if !e.IsHandle && _this.OnMouseDown != nil {
			_this.OnMouseDown(e)
		}
	})

	var bakMouseMove br.WindowMouseMoveProc
	bakMouseMove = _this.impl.SetOnMouseMove(func(e *fm.MouseEvArgs) {
		if bakMouseMove != nil {
			bakMouseMove(e)
		}
		if !e.IsHandle && _this.OnMouseMove != nil {
			_this.OnMouseMove(e)
		}
	})

	var bakResize br.WindowResizeProc
	bakResize = _this.impl.SetOnResize(func(e fm.Rect) {
		if bakResize != nil {
			bakResize(e)
		}
		if _this.OnResize != nil {
			_this.OnResize(e)
		}
	})

	var bakMove br.WindowMoveProc
	bakMove = _this.impl.SetOnMove(func(e fm.Point) bool {
		b := false
		if bakMove != nil {
			b = bakMove(e)
		}
		if !b && _this.OnMove != nil {
			_this.OnMove(e)
		}
		return b
	})

	var bakShow br.WindowShowProc
	bakShow = _this.impl.SetOnShow(func() {
		if _this.skipLoad == false {
			_this.skipLoad = true
			if _this.OnLoad != nil {
				_this.OnLoad()
			}
		}
		if bakShow != nil {
			bakShow()
		}
		if _this.OnShow != nil {
			_this.OnShow()
		}
	})
	return _this
}

func (_this *BaseUI) Enable(enable bool) {
	_this.impl.Enable(enable)
}

func (_this *BaseUI) IsEnable() bool {
	return _this.impl.IsEnable()
}

func (_this *BaseUI) CreateGraphics() fm.Graphics {
	return _this.impl.CreateGraphics()
}

func (_this *BaseUI) SetCursor(cursor fm.CursorType) {
	_this.impl.SetCursor(cursor)
}

func (_this *BaseUI) GetCursor() fm.CursorType {
	return _this.impl.GetCursor()
}

func (_this *BaseUI) GetHandle() uintptr {
	return _this.impl.GetHandle()
}

func (_this *BaseUI) SetLocation(x, y int) {
	_this.impl.SetLocation(x, y)
}

func (_this *BaseUI) GetBound() fm.Bound {
	return _this.impl.GetBound()
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

func (_this *BaseUI) GetParent() GUI {
	return _this.parent
}

func (_this *BaseUI) GetOwner() GUI {
	return _this.owner
}
