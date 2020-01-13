package miniblinkc_core

import mb "qq.2564874169/miniblink"

type ICore interface {
	LoadUri(uri string)

	SetOnPaint(callback PaintCallback)
	Resize(width, height int)
}

type PaintArgs struct {
	Wke   uintptr
	Clip  mb.Bound
	Size  mb.Rect
	Bits  []byte
	Param uintptr
}
type PaintCallback func(args PaintArgs)
