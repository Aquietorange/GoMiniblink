package driver

import (
	"image"
	mb "qq.2564874169/miniblink"
)

type ICore interface {
	LoadUri(uri string)

	GetView(bound mb.Bound) *image.RGBA
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
