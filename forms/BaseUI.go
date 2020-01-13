package forms

import (
	mb "qq.2564874169/miniblink"
	plat "qq.2564874169/miniblink/platform"
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

	real interface{}
	impl plat.IWindow
	size mb.Rect
	pos  mb.Point
}

func (_this *BaseUI) init(instance interface{}, impl plat.IWindow) *BaseUI {
	_this.real = instance
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

	_this.impl.SetOnKeyPress(func(e *mb.KeyPressEvArgs) {
		_this.OnKeyPress(e)
	})
	_this.impl.SetOnKeyUp(func(e *mb.KeyEvArgs) {
		_this.OnKeyUp(e)
	})
	_this.impl.SetOnKeyDown(func(e *mb.KeyEvArgs) {
		_this.OnKeyDown(e)
	})
	_this.impl.SetOnPaint(func(e mb.PaintEvArgs) {
		_this.OnPaint(e)
	})
	_this.impl.SetOnMouseClick(func(e mb.MouseEvArgs) {
		_this.OnMouseClick(e)
	})
	_this.impl.SetOnMouseWheel(func(e mb.MouseEvArgs) {
		_this.OnMouseWheel(e)
	})
	_this.impl.SetOnMouseUp(func(e mb.MouseEvArgs) {
		_this.OnMouseUp(e)
	})
	_this.impl.SetOnMouseDown(func(e mb.MouseEvArgs) {
		_this.OnMouseDown(e)
	})
	_this.impl.SetOnMouseMove(func(e mb.MouseEvArgs) {
		_this.OnMouseMove(e)
	})
	_this.impl.SetOnResize(func(e mb.Rect) {
		_this.size = e
		_this.OnResize(e)
	})
	_this.impl.SetOnMove(func(e mb.Point) {
		_this.pos = e
		_this.OnMove(e)
	})
	_this.impl.SetOnCreate(func(handle uintptr) {
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
		Wdith:  width,
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
