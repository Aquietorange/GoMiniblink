package forms

import (
	mb "qq.2564874169/goMiniblink"
	p "qq.2564874169/goMiniblink/platform"
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
	bakKeyPress = _this.impl.SetOnKeyPress(func(e *mb.KeyPressEvArgs) {
		if bakKeyPress != nil {
			bakKeyPress(e)
		}
		_this.OnKeyPress(e)
	})

	var bakKeyUp p.WindowKeyUpProc
	bakKeyUp = _this.impl.SetOnKeyUp(func(e *mb.KeyEvArgs) {
		if bakKeyUp != nil {
			bakKeyUp(e)
		}
		_this.OnKeyUp(e)
	})

	var bakKeyDown p.WindowKeyDownProc
	bakKeyDown = _this.impl.SetOnKeyDown(func(e *mb.KeyEvArgs) {
		if bakKeyDown != nil {
			bakKeyDown(e)
		}
		_this.OnKeyDown(e)
	})

	var bakPaint p.WindowPaintProc
	bakPaint = _this.impl.SetOnPaint(func(e mb.PaintEvArgs) {
		if bakPaint != nil {
			bakPaint(e)
		}
		_this.OnPaint(e)
	})

	var bakMouseClick p.WindowMouseClickProc
	bakMouseClick = _this.impl.SetOnMouseClick(func(e mb.MouseEvArgs) {
		if bakMouseClick != nil {
			bakMouseClick(e)
		}
		_this.OnMouseClick(e)
	})

	var bakMouseWheel p.WindowMouseWheelProc
	bakMouseWheel = _this.impl.SetOnMouseWheel(func(e mb.MouseEvArgs) {
		if bakMouseWheel != nil {
			bakMouseWheel(e)
		}
		_this.OnMouseWheel(e)
	})

	var bakMouseUp p.WindowMouseUpProc
	bakMouseUp = _this.impl.SetOnMouseUp(func(e mb.MouseEvArgs) {
		if bakMouseUp != nil {
			bakMouseUp(e)
		}
		_this.OnMouseUp(e)
	})

	var bakMouseDown p.WindowMouseDownProc
	_this.impl.SetOnMouseDown(func(e mb.MouseEvArgs) {
		if bakMouseDown != nil {
			bakMouseDown(e)
		}
		_this.OnMouseDown(e)
	})

	var bakMouseMove p.WindowMouseMoveProc
	bakMouseMove = _this.impl.SetOnMouseMove(func(e mb.MouseEvArgs) {
		if bakMouseMove != nil {
			bakMouseMove(e)
		}
		_this.OnMouseMove(e)
	})

	var bakResize p.WindowResizeProc
	bakResize = _this.impl.SetOnResize(func(e mb.Rect) {
		if bakResize != nil {
			bakResize(e)
		}
		_this.size = e
		_this.OnResize(e)
	})

	var bakMove p.WindowMoveProc
	bakMove = _this.impl.SetOnMove(func(e mb.Point) {
		if bakMove != nil {
			bakMove(e)
		}
		_this.pos = e
		_this.OnMove(e)
	})

	var bakCreate p.WindowCreateProc
	bakCreate = _this.impl.SetOnCreate(func(handle uintptr) {
		if bakCreate != nil {
			bakCreate(handle)
		}
		_this.Handle = handle
		_this.OnLoad()
	})
	return _this
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
