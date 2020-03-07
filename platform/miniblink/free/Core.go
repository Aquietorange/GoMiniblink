package free

import (
	"image"
	"image/draw"
	"math"
	mb "qq2564874169/goMiniblink"
	plat "qq2564874169/goMiniblink/platform"
	core "qq2564874169/goMiniblink/platform/miniblink"
	"qq2564874169/goMiniblink/platform/windows/win32"
	"unsafe"
)

type Core struct {
	app   plat.IProvider
	owner plat.IWindow
	wke   wkeHandle

	onPaint core.PaintCallback
}

func (_this *Core) Init(window plat.IWindow) *Core {
	_this.app = window.GetProvider()
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
	rect := wkeGetCaretRect(_this.wke)
	return mb.Point{X: int(rect.x), Y: int(rect.y)}
}

func (_this *Core) FireKeyPressEvent(charCode int, isSys bool) bool {
	return wkeFireKeyPressEvent(_this.wke, charCode, uint32(wkeKeyFlags_Repeat), isSys)
}

func (_this *Core) FireKeyEvent(e mb.KeyEvArgs, isDown, isSys bool) bool {
	flags := int(wkeKeyFlags_Repeat)
	if isExtKey(e.Key) {
		flags |= int(wkeKeyFlags_Extend)
	}
	if isDown {
		return wkeFireKeyDownEvent(_this.wke, e.Value, uint32(flags), isSys)
	} else {
		return wkeFireKeyUpEvent(_this.wke, e.Value, uint32(flags), isSys)
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

func (_this *Core) FireMouseWheelEvent(button mb.MouseButtons, delta, x, y int) bool {
	flags := wkeMouseFlags_None
	keys := _this.app.ModifierKeys()
	if s, ok := keys[mb.Keys_Ctrl]; ok && s {
		flags |= wkeMouseFlags_CONTROL
	}
	if s, ok := keys[mb.Keys_Shift]; ok && s {
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
	return wkeFireMouseWheelEvent(_this.wke, int32(x), int32(y), int32(delta), int32(flags))
}

func (_this *Core) FireMouseMoveEvent(button mb.MouseButtons, x, y int) bool {
	flags := wkeMouseFlags_None
	if button&mb.MouseButtons_Left != 0 {
		flags |= wkeMouseFlags_LBUTTON
	}
	if button&mb.MouseButtons_Right != 0 {
		flags |= wkeMouseFlags_RBUTTON
	}
	return wkeFireMouseEvent(_this.wke, int32(win32.WM_MOUSEMOVE), int32(x), int32(y), int32(flags))
}

func (_this *Core) FireMouseClickEvent(button mb.MouseButtons, isDown, isDb bool, x, y int) bool {
	flags := wkeMouseFlags_None
	keys := _this.app.ModifierKeys()
	if s, ok := keys[mb.Keys_Ctrl]; ok && s {
		flags |= wkeMouseFlags_CONTROL
	}
	if s, ok := keys[mb.Keys_Shift]; ok && s {
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
		return wkeFireMouseEvent(_this.wke, int32(msg), int32(x), int32(y), int32(flags))
	}
	return false
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
	e := core.PaintUpdateArgs{
		Wke: uintptr(wke),
		Clip: mb.Bound{
			Point: mb.Point{
				X: int(rect.x),
				Y: int(rect.y),
			},
			Rect: mb.Rect{
				Width:  int(math.Min(float64(rect.w), float64(width))),
				Height: int(math.Min(float64(rect.h), float64(wkeGetHeight(wke)))),
			},
		},
		Size: mb.Rect{
			Width:  int(width),
			Height: int(height),
		},
		Param: param,
	}
	bmp := image.NewRGBA(image.Rect(0, 0, e.Clip.Width, e.Clip.Height))
	stride := e.Size.Width * 4
	pixs := (*[1 << 30]byte)(unsafe.Pointer(bits))
	for y := 0; y < e.Clip.Height; y++ {
		for x := 0; x < e.Clip.Width*4; x++ {
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

func isExtKey(key mb.Keys) bool {
	switch key {
	case mb.Keys_Insert, mb.Keys_Delete, mb.Keys_Home, mb.Keys_End, mb.Keys_PageUp,
		mb.Keys_PageDown, mb.Keys_Left, mb.Keys_Right, mb.Keys_Up, mb.Keys_Down:
		return true
	default:
		return false
	}
}
