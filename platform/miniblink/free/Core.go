package free

import (
	"image"
	"image/draw"
	mb "qq.2564874169/goMiniblink"
	plat "qq.2564874169/goMiniblink/platform"
	core "qq.2564874169/goMiniblink/platform/miniblink"
	"qq.2564874169/goMiniblink/platform/windows/win32"
	"unsafe"
)

type Core struct {
	owner plat.IWindow
	wke   wkeHandle

	onPaint core.PaintCallback
}

func (_this *Core) Init(window plat.IWindow) *Core {
	_this.owner = window
	_this.wke = wkeCreateWebView()
	if _this.wke == 0 {
		panic("创建失败")
	}
	wkeSetHandle(_this.wke, _this.owner.GetHandle())
	wkeOnPaintBitUpdated(_this.wke, _this.onPaintBitUpdated, 0)
	return _this
}

func (_this *Core) SetFocus() {
	wkeSetFocus(_this.wke)
}

func (_this *Core) GetCaretPos() mb.Point {
	//rect := wkeGetCaretRect(_this.wke)
	//return mb.Point{X: int(rect.x), Y: int(rect.y)}
	var p win32.POINT
	win32.GetCaretPos(&p)
	return mb.Point{X: int(p.X), Y: int(p.Y)}
}

func (_this *Core) FireKeyPressEvent(charCode int, isRepeat, isExtend, isSys bool) {
	flags := 0
	if isRepeat {
		flags |= int(wkeKeyFlags_Repeat)
	}
	if isExtend {
		flags |= int(wkeKeyFlags_Extend)
	}
	wkeFireKeyPressEvent(_this.wke, charCode, uint32(flags), isSys)
}

func (_this *Core) FireKeyEvent(keyCode uintptr, isRepeat, isExtend, isDown, isSys bool) {
	flags := 0
	if isRepeat {
		flags |= int(wkeKeyFlags_Repeat)
	}
	if isExtend {
		flags |= int(wkeKeyFlags_Extend)
	}
	if isDown {
		wkeFireKeyDownEvent(_this.wke, keyCode, uint32(flags), isSys)
	} else {
		wkeFireKeyUpEvent(_this.wke, keyCode, uint32(flags), isSys)
	}
}

func (_this *Core) GetCursor() mb.CursorType {
	cur := wkeGetCursorInfoType(_this.wke)
	switch cur {
	case wkeCursorType_Hand:
		return mb.CursorType_HAND
	case wkeCursorType_IBeam:
		return mb.CursorType_IBEAM
	case wkeCursorType_ColumnResize:
		return mb.CursorType_SIZEWE
	case wkeCursorType_RowResize:
		return mb.CursorType_SIZENS
	default:
		return mb.CursorType_ARROW
	}
}

func (_this *Core) FireMouseWheelEvent(app plat.IProvider, button mb.MouseButtons, delta, x, y int) {
	flags := wkeMouseFlags_None
	if app.KeyIsDown(mb.Keys_Ctrl) {
		flags |= wkeMouseFlags_CONTROL
	}
	if app.KeyIsDown(mb.Keys_Shift) {
		flags |= wkeMouseFlags_SHIFT
	}
	if button&mb.MouseButtons_Left != 0 {
		flags |= wkeMouseFlags_LBUTTON
	}
	if button&mb.MouseButtons_Right != 0 {
		flags |= wkeMouseFlags_RBUTTON
	}
	if button&mb.MouseButtons_Middle != 0 {
		flags |= wkeMouseFlags_MBUTTON
	}
	wkeFireMouseWheelEvent(_this.wke, int32(x), int32(y), int32(delta), int32(flags))
}

func (_this *Core) FireMouseMoveEvent(app plat.IProvider, button mb.MouseButtons, x, y int) {
	flags := wkeMouseFlags_None
	if app.KeyIsDown(mb.Keys_Ctrl) {
		flags |= wkeMouseFlags_CONTROL
	}
	if app.KeyIsDown(mb.Keys_Shift) {
		flags |= wkeMouseFlags_SHIFT
	}
	if button&mb.MouseButtons_Left != 0 {
		flags |= wkeMouseFlags_LBUTTON
	}
	if button&mb.MouseButtons_Right != 0 {
		flags |= wkeMouseFlags_RBUTTON
	}
	wkeFireMouseEvent(_this.wke, int32(win32.WM_MOUSEMOVE), int32(x), int32(y), int32(flags))
}

func (_this *Core) FireMouseClickEvent(app plat.IProvider, button mb.MouseButtons, isDown, isDb bool, x, y int) {
	flags := wkeMouseFlags_None
	if app.KeyIsDown(mb.Keys_Ctrl) {
		flags |= wkeMouseFlags_CONTROL
	}
	if app.KeyIsDown(mb.Keys_Shift) {
		flags |= wkeMouseFlags_SHIFT
	}
	msg := 0
	if button&mb.MouseButtons_Left != 0 {
		flags |= wkeMouseFlags_LBUTTON
		if isDb {
			msg = win32.WM_LBUTTONDBLCLK
		} else if isDown {
			msg = win32.WM_LBUTTONDOWN
		} else {
			msg = win32.WM_LBUTTONUP
		}
	}
	if button&mb.MouseButtons_Right != 0 {
		flags |= wkeMouseFlags_RBUTTON
		if isDb {
			msg = win32.WM_RBUTTONDBLCLK
		} else if isDown {
			msg = win32.WM_RBUTTONDOWN
		} else {
			msg = win32.WM_RBUTTONUP
		}
	}
	if button&mb.MouseButtons_Middle != 0 {
		flags |= wkeMouseFlags_MBUTTON
		if isDb {
			msg = win32.WM_MBUTTONDBLCLK
		} else if isDown {
			msg = win32.WM_MBUTTONDOWN
		} else {
			msg = win32.WM_MBUTTONUP
		}
	}
	if msg != 0 {
		wkeFireMouseEvent(_this.wke, int32(msg), int32(x), int32(y), int32(flags))
	}
}

func (_this *Core) GetImage(bound mb.Bound) *image.RGBA {
	w := wkeGetWidth(_this.wke)
	h := wkeGetHeight(_this.wke)
	view := image.NewRGBA(image.Rect(0, 0, int(w), int(h)))
	wkePaint(_this.wke, view.Pix, 0)
	bmp := image.NewRGBA(image.Rect(0, 0, bound.Width, bound.Height))
	draw.Draw(bmp, image.Rect(0, 0, bound.Width, bound.Height), view, image.Pt(bound.X, bound.Y), draw.Src)
	return bmp
}

func (_this *Core) onPaintBitUpdated(wke wkeHandle, param, bits uintptr, rect *wkeRect, width, height int32) uintptr {
	if width == 0 || height == 0 {
		return 0
	}
	w, h := int(rect.w), int(rect.h)
	e := core.PaintUpdateArgs{
		Wke: uintptr(wke),
		Clip: mb.Bound{
			Point: mb.Point{
				X: int(rect.x),
				Y: int(rect.y),
			},
			Rect: mb.Rect{
				Width:  w,
				Height: h,
			},
		},
		Size: mb.Rect{
			Width:  int(width),
			Height: int(height),
		},
		Param: param,
	}
	bmp := image.NewRGBA(image.Rect(0, 0, w, h))
	stride := e.Size.Width * 4
	pixs := (*[1 << 30]byte)(unsafe.Pointer(bits))
	for y := 0; y < h; y++ {
		for x := 0; x < w*4; x++ {
			sp := bmp.Stride*y + x
			dp := stride*(e.Clip.Y+y) + e.Clip.X*4 + x
			bmp.Pix[sp] = pixs[dp]
		}
	}
	e.Image = bmp
	_this.onPaint(e)
	return 0
}

func (_this *Core) Resize(width, height int) {
	wkeResize(_this.wke, uint32(width), uint32(height))
}

func (_this *Core) SetOnPaint(callback core.PaintCallback) {
	_this.onPaint = callback
}

func (_this *Core) LoadUri(uri string) {
	wkeLoadURL(_this.wke, uri)
}
