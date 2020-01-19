package driver

import (
	"image"
	mb "qq.2564874169/miniblink"
)

type ICore interface {
	LoadUri(uri string)

	GetView(bound mb.Bound) image.Image
	SetOnPaint(callback PaintCallback)
	Resize(width, height int)
}

type PaintArgs struct {
	Wke   uintptr
	Clip  mb.Bound
	Size  mb.Rect
	Image image.Image
	Param uintptr
}
type PaintCallback func(args PaintArgs)
