package miniblink

import (
	"image"
	mb "qq.2564874169/goMiniblink"
	plat "qq.2564874169/goMiniblink/platform"
)

type ICore interface {
	LoadUri(uri string)

	SetFocus()
	GetCaretPos() mb.Point
	FireKeyPressEvent(charCode int, isRepeat, isExtend, isSys bool)
	FireKeyEvent(keyCode uintptr, isRepeat, isExtend, isDown, isSys bool)
	GetCursor() mb.CursorType
	FireMouseWheelEvent(provider plat.IProvider, button mb.MouseButtons, delta, x, y int)
	FireMouseMoveEvent(provider plat.IProvider, button mb.MouseButtons, x, y int)
	FireMouseClickEvent(provider plat.IProvider, button mb.MouseButtons, isDown, isDb bool, x, y int)
	GetImage(bound mb.Bound) *image.RGBA
	SetOnPaint(callback PaintCallback)
	Resize(width, height int)
}

type PaintUpdateArgs struct {
	Wke   uintptr
	Clip  mb.Bound
	Size  mb.Rect
	Image *image.RGBA
	Param uintptr
}
type PaintCallback func(args PaintUpdateArgs)
