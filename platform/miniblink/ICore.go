package miniblink

import (
	"image"
	mb "qq2564874169/goMiniblink"
)

type ICore interface {
	LoadUri(uri string)

	SetWindowProp(name string, value interface{})
	RunJs(script string) interface{}
	BindJsFunc(fn mb.JsFuncBinding)
	SetOnRequest(callback RequestCallback)
	SetOnJsReady(callback JsReadyCallback)

	SetFocus()
	GetCaretPos() mb.Point
	GetCursor() mb.CursorType
	FireKeyPressEvent(charCode int, isSys bool) bool
	FireKeyEvent(e mb.KeyEvArgs, isDown, isSys bool) bool
	FireMouseWheelEvent(button mb.MouseButtons, delta, x, y int) bool
	FireMouseMoveEvent(button mb.MouseButtons, x, y int) bool
	FireMouseClickEvent(button mb.MouseButtons, isDown, isDb bool, x, y int) bool
	GetImage(bound mb.Bound) *image.RGBA
	SetOnPaint(callback PaintCallback)
	Resize(width, height int)
	Invoke(fn func(interface{}), state interface{})
	//todo 差一个AsyncInvoke
	GetHandle() uintptr
}

type RequestCallback func(args mb.RequestBeforeEvArgs)

type JsReadyCallback func(args mb.JsReadyEvArgs)

type PaintUpdateArgs struct {
	Wke   uintptr
	Clip  mb.Bound
	Size  mb.Rect
	Image *image.RGBA
	Param uintptr
}
type PaintCallback func(args PaintUpdateArgs)
