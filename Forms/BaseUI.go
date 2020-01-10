package Forms

import (
	MB "GoMiniblink"
	plat "GoMiniblink/CrossPlatform"
)

type BaseUI struct {
	EvKeyDown map[string]func(target interface{}, e *MB.KeyEvArgs)
	OnKeyDown func(e *MB.KeyEvArgs)

	EvKeyUp map[string]func(target interface{}, e *MB.KeyEvArgs)
	OnKeyUp func(e *MB.KeyEvArgs)

	EvKeyPress map[string]func(target interface{}, e *MB.KeyPressEvArgs)
	OnKeyPress func(e *MB.KeyPressEvArgs)

	EvLoad map[string]func(target interface{})
	OnLoad func()

	EvResize map[string]func(target interface{}, e MB.Rect)
	OnResize func(e MB.Rect)

	EvMove map[string]func(target interface{}, e MB.Point)
	OnMove func(e MB.Point)

	EvMouseMove map[string]func(target interface{}, e MB.MouseEvArgs)
	OnMouseMove func(e MB.MouseEvArgs)

	EvMouseDown map[string]func(target interface{}, e MB.MouseEvArgs)
	OnMouseDown func(e MB.MouseEvArgs)

	EvMouseUp map[string]func(target interface{}, e MB.MouseEvArgs)
	OnMouseUp func(e MB.MouseEvArgs)

	EvMouseWheel map[string]func(target interface{}, e MB.MouseEvArgs)
	OnMouseWheel func(e MB.MouseEvArgs)

	EvMouseClick map[string]func(target interface{}, e MB.MouseEvArgs)
	OnMouseClick func(e MB.MouseEvArgs)

	EvPaint map[string]func(target interface{}, e MB.PaintEvArgs)
	OnPaint func(e MB.PaintEvArgs)

	Handle uintptr

	real interface{}
	impl plat.IWindow
	size MB.Rect
	pos  MB.Point
}

func (_this *BaseUI) init(instance interface{}, impl plat.IWindow) *BaseUI {
	_this.real = instance
	_this.impl = impl

	_this.EvKeyPress = make(map[string]func(target interface{}, e *MB.KeyPressEvArgs))
	_this.EvKeyDown = make(map[string]func(target interface{}, e *MB.KeyEvArgs))
	_this.EvKeyUp = make(map[string]func(target interface{}, e *MB.KeyEvArgs))
	_this.EvPaint = make(map[string]func(target interface{}, e MB.PaintEvArgs))
	_this.EvLoad = make(map[string]func(target interface{}))
	_this.EvResize = make(map[string]func(target interface{}, e MB.Rect))
	_this.EvMove = make(map[string]func(target interface{}, e MB.Point))
	_this.EvMouseMove = make(map[string]func(target interface{}, args MB.MouseEvArgs))
	_this.EvMouseDown = make(map[string]func(target interface{}, args MB.MouseEvArgs))
	_this.EvMouseUp = make(map[string]func(target interface{}, args MB.MouseEvArgs))
	_this.EvMouseWheel = make(map[string]func(target interface{}, args MB.MouseEvArgs))
	_this.EvMouseClick = make(map[string]func(target interface{}, args MB.MouseEvArgs))

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

	_this.impl.SetOnKeyPress(func(e *MB.KeyPressEvArgs) {
		_this.OnKeyPress(e)
	})
	_this.impl.SetOnKeyUp(func(e *MB.KeyEvArgs) {
		_this.OnKeyUp(e)
	})
	_this.impl.SetOnKeyDown(func(e *MB.KeyEvArgs) {
		_this.OnKeyDown(e)
	})
	_this.impl.SetOnPaint(func(e MB.PaintEvArgs) {
		_this.OnPaint(e)
	})
	_this.impl.SetOnMouseClick(func(e MB.MouseEvArgs) {
		_this.OnMouseClick(e)
	})
	_this.impl.SetOnMouseWheel(func(e MB.MouseEvArgs) {
		_this.OnMouseWheel(e)
	})
	_this.impl.SetOnMouseUp(func(e MB.MouseEvArgs) {
		_this.OnMouseUp(e)
	})
	_this.impl.SetOnMouseDown(func(e MB.MouseEvArgs) {
		_this.OnMouseDown(e)
	})
	_this.impl.SetOnMouseMove(func(e MB.MouseEvArgs) {
		_this.OnMouseMove(e)
	})
	_this.impl.SetOnResize(func(e MB.Rect) {
		_this.size = e
		_this.OnResize(e)
	})
	_this.impl.SetOnMove(func(e MB.Point) {
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
	_this.pos = MB.Point{
		X: x,
		Y: y,
	}
	_this.impl.SetLocation(x, y)
}

func (_this *BaseUI) GetLocation() MB.Point {
	return _this.pos
}

func (_this *BaseUI) GetSize() MB.Rect {
	return _this.size
}

func (_this *BaseUI) SetSize(width, height int) {
	_this.size = MB.Rect{
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
