package forms

import (
	mb "qq2564874169/goMiniblink"
	p "qq2564874169/goMiniblink/platform"
)

type BaseUI struct {
	EvKeyDown map[string]func(target interface{}, e *mb.KeyEvArgs)
	OnKeyDown func(e *mb.KeyEvArgs)

	EvKeyUp map[string]func(target interface{}, e *mb.KeyEvArgs)
	OnKeyUp func(e *mb.KeyEvArgs)

	EvKeyPress map[string]func(target interface{}, e *mb.KeyPressEvArgs)
	OnKeyPress func(e *mb.KeyPressEvArgs)

	EvLoad map[string]func(target interface{})
	OnLoad func()

	EvResize map[string]func(target interface{}, e mb.Rect)
	OnResize func(e mb.Rect)

	EvMove map[string]func(target interface{}, e mb.Point)
	OnMove func(e mb.Point)

	EvMouseMove map[string]func(target interface{}, e mb.MouseEvArgs)
	OnMouseMove func(e mb.MouseEvArgs)

	EvMouseDown map[string]func(target interface{}, e mb.MouseEvArgs)
	OnMouseDown func(e mb.MouseEvArgs)

	EvMouseUp map[string]func(target interface{}, e mb.MouseEvArgs)
	OnMouseUp func(e mb.MouseEvArgs)

	EvMouseWheel map[string]func(target interface{}, e mb.MouseEvArgs)
	OnMouseWheel func(e mb.MouseEvArgs)

	EvMouseClick map[string]func(target interface{}, e mb.MouseEvArgs)
	OnMouseClick func(e mb.MouseEvArgs)

	EvPaint map[string]func(target interface{}, e mb.PaintEvArgs)
	OnPaint func(e mb.PaintEvArgs)

	Handle uintptr

	instance interface{}
	impl     p.IWindow
	size     mb.Rect
	pos      mb.Point
}

func (_this *BaseUI) init(instance interface{}, impl p.IWindow) *BaseUI {
	_this.instance = instance
	_this.impl = impl

	_this.EvKeyPress = make(map[string]func(target interface{}, e *mb.KeyPressEvArgs))
	_this.EvKeyDown = make(map[string]func(target interface{}, e *mb.KeyEvArgs))
	_this.EvKeyUp = make(map[string]func(target interface{}, e *mb.KeyEvArgs))
	_this.EvPaint = make(map[string]func(target interface{}, e mb.PaintEvArgs))
	_this.EvLoad = make(map[string]func(target interface{}))
	_this.EvResize = make(map[string]func(target interface{}, e mb.Rect))
	_this.EvMove = make(map[string]func(target interface{}, e mb.Point))
	_this.EvMouseMove = make(map[string]func(target interface{}, args mb.MouseEvArgs))
	_this.EvMouseDown = make(map[string]func(target interface{}, args mb.MouseEvArgs))
	_this.EvMouseUp = make(map[string]func(target interface{}, args mb.MouseEvArgs))
	_this.EvMouseWheel = make(map[string]func(target interface{}, args mb.MouseEvArgs))
	_this.EvMouseClick = make(map[string]func(target interface{}, args mb.MouseEvArgs))

	_this.OnKeyPress = _this.defOnKeyPress
	_this.OnKeyUp = _this.defOnKeyUp
	_this.OnKeyDown = _this.defOnKeyDown
	_this.OnPaint = _this.defOnPaint
	_this.OnLoad = _this.defOnLoad
	_this.OnResize = _this.defOnResize
	_this.OnMove = _this.defOnMove
	_this.OnMouseMove = _this.defOnMouseMove
	_this.OnMouseDown = _this.defOnMouseDown
	_this.OnMouseUp = _this.defOnMouseUp
	_this.OnMouseWheel = _this.defOnMouseWheel
	_this.OnMouseClick = _this.defOnMouseClick

	var bakKeyPress p.WindowKeyPressProc
	bakKeyPress = _this.impl.SetOnKeyPress(func(e *mb.KeyPressEvArgs) bool {
		b := false
		if bakKeyPress != nil {
			b = bakKeyPress(e)
		}
		if !b {
			_this.OnKeyPress(e)
		}
		return b
	})

	var bakKeyUp p.WindowKeyUpProc
	bakKeyUp = _this.impl.SetOnKeyUp(func(e *mb.KeyEvArgs) bool {
		b := false
		if bakKeyUp != nil {
			b = bakKeyUp(e)
		}
		if !b {
			_this.OnKeyUp(e)
		}
		return b
	})

	var bakKeyDown p.WindowKeyDownProc
	bakKeyDown = _this.impl.SetOnKeyDown(func(e *mb.KeyEvArgs) bool {
		b := false
		if bakKeyDown != nil {
			b = bakKeyDown(e)
		}
		if !b {
			_this.OnKeyDown(e)
		}
		return b
	})

	var bakPaint p.WindowPaintProc
	bakPaint = _this.impl.SetOnPaint(func(e mb.PaintEvArgs) bool {
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
	bakMouseClick = _this.impl.SetOnMouseClick(func(e mb.MouseEvArgs) bool {
		b := false
		if bakMouseClick != nil {
			b = bakMouseClick(e)
		}
		if !b {
			_this.OnMouseClick(e)
		}
		return b
	})

	var bakMouseWheel p.WindowMouseWheelProc
	bakMouseWheel = _this.impl.SetOnMouseWheel(func(e mb.MouseEvArgs) bool {
		b := false
		if bakMouseWheel != nil {
			b = bakMouseWheel(e)
		}
		if !b {
			_this.OnMouseWheel(e)
		}
		return b
	})

	var bakMouseUp p.WindowMouseUpProc
	bakMouseUp = _this.impl.SetOnMouseUp(func(e mb.MouseEvArgs) bool {
		b := false
		if bakMouseUp != nil {
			b = bakMouseUp(e)
		}
		if !b {
			_this.OnMouseUp(e)
		}
		return b
	})

	var bakMouseDown p.WindowMouseDownProc
	bakMouseDown = _this.impl.SetOnMouseDown(func(e mb.MouseEvArgs) bool {
		b := false
		if bakMouseDown != nil {
			b = bakMouseDown(e)
		}
		if !b {
			_this.OnMouseDown(e)
		}
		return b
	})

	var bakMouseMove p.WindowMouseMoveProc
	bakMouseMove = _this.impl.SetOnMouseMove(func(e mb.MouseEvArgs) bool {
		b := false
		if bakMouseMove != nil {
			b = bakMouseMove(e)
		}
		if !b {
			_this.OnMouseMove(e)
		}
		return b
	})

	var bakResize p.WindowResizeProc
	bakResize = _this.impl.SetOnResize(func(e mb.Rect) bool {
		b := false
		if bakResize != nil {
			b = bakResize(e)
		}
		if !b {
			_this.size = e
			_this.OnResize(e)
		}
		return b
	})

	var bakMove p.WindowMoveProc
	bakMove = _this.impl.SetOnMove(func(e mb.Point) bool {
		b := false
		if bakMove != nil {
			b = bakMove(e)
		}
		if !b {
			_this.pos = e
			_this.OnMove(e)
		}
		return b
	})

	var bakCreate p.WindowCreateProc
	bakCreate = _this.impl.SetOnCreate(func(handle uintptr) bool {
		b := false
		if bakCreate != nil {
			b = bakCreate(handle)
		}
		if !b {
			_this.Handle = handle
			_this.OnLoad()
		}
		return b
	})
	return _this
}

func (_this *BaseUI) SetCursor(cursor mb.CursorType) {
	_this.impl.SetCursor(cursor)
}

func (_this *BaseUI) GetHandle() uintptr {
	return _this.Handle
}

func (_this *BaseUI) SetLocation(x, y int) {
	_this.pos = mb.Point{
		X: x,
		Y: y,
	}
	_this.impl.SetLocation(x, y)
}

func (_this *BaseUI) GetLocation() mb.Point {
	return _this.pos
}

func (_this *BaseUI) GetSize() mb.Rect {
	return _this.size
}

func (_this *BaseUI) SetSize(width, height int) {
	_this.size = mb.Rect{
		Width:  width,
		Height: height,
	}
	_this.impl.SetSize(width, height)
}

func (_this *BaseUI) SetBgColor(color int) {
	_this.impl.SetBgColor(color)
}

func (_this *BaseUI) Invoke(fn func(state interface{}), state interface{}) {
	_this.impl.Invoke(fn, state)
}
