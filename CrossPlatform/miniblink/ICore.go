package miniblink

import MB "GoMiniblink"

type ICore interface {
	LoadUri(uri string)

	SetOnPaint(callback PaintCallback)
}

type PaintArgs struct {
	Wke    uintptr
	Update MB.Bound
	Size   MB.Rect
	Bits   []byte
	Param  uintptr
}
type PaintCallback func(args PaintArgs)
