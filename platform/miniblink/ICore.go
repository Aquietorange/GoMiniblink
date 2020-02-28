package miniblink

import (
	"image"
	mb "qq.2564874169/goMiniblink"
	plat "qq.2564874169/goMiniblink/platform"
)

type ICore interface {
	LoadUri(uri string)

	FireMouseWheelEvent(provider plat.IProvider, button mb.MouseButtons, delta, x, y int)
	FireMouseEvent(provider plat.IProvider, button mb.MouseButtons, isDown, isMove bool, x, y int)
	GetImage(bound mb.Bound) *image.RGBA
	SetOnPaint(callback PaintCallback)
	Resize(width, height int)
	GetCursor() mb.CursorType
}

type PaintUpdateArgs struct {
	Wke   uintptr
	Clip  mb.Bound
	Size  mb.Rect
	Image *image.RGBA
	Param uintptr
}
type PaintCallback func(args PaintUpdateArgs)
