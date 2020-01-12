package miniblink

import MB "GoMiniblink"

type ICore interface {
	LoadUri(uri string)

	SetOnPaint(callback PaintCallback)
	Resize(width, height int)
}

type PaintArgs struct {
	Wke   uintptr
	Clip  MB.Bound
	Size  MB.Rect
	Bits  []byte
	Param uintptr
}
type PaintCallback func(args PaintArgs)
