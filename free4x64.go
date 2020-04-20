package GoMiniblink

import (
	"image"
	"math"
	f "qq2564874169/goMiniblink/forms"
	c "qq2564874169/goMiniblink/forms/controls"
	"qq2564874169/goMiniblink/forms/platform/windows/win32"
	"unsafe"
)

func init() {
	mbApi = new(freeApiForWindows).init()
}

type free4x64 struct {
	_view *c.Control
	_wke  wkeHandle
}

func (_this *free4x64) init(control *c.Control) *free4x64 {
	_this._view = control
	_this.mbInit()
	return _this
}

func (_this *free4x64) BindFunc(fn GoFunc) {

}

func (_this *free4x64) SetOnRequest(func(e RequestEvArgs)) {

}

func (_this *free4x64) mbInit() {
	_this._view.OnResize = _this.viewResize
	_this._view.OnPaint = _this.viewPaint
	_this._view.OnMouseMove = _this.viewMouseMove
	_this._view.OnMouseDown = _this.viewMouseDown
	_this._view.OnMouseUp = _this.viewMouseUp
	_this._view.OnMouseWheel = _this.viewMouseWheel
	_this._view.OnSetCursor = _this.viewSetCursor

	_this._wke = createWebView(_this)
	mbApi.wkeSetHandle(_this._wke, _this._view.GetHandle())
	mbApi.wkeOnPaintBitUpdated(_this._wke, _this.onPaintBitUpdated, nil)
	_this.viewResize(_this._view.GetSize())
}

func (_this *free4x64) LoadUri(uri string) {
	mbApi.wkeLoadURL(_this._wke, uri)
}

func (_this *free4x64) viewSetCursor() bool {

}

func (_this *free4x64) viewMouseWheel(e f.MouseEvArgs) {
	flags := wkeMouseFlags_None
	keys := c.App.ModifierKeys()
	if s, ok := keys[f.Keys_Ctrl]; ok && s {
		flags |= wkeMouseFlags_CONTROL
	}
	if s, ok := keys[f.Keys_Shift]; ok && s {
		flags |= wkeMouseFlags_SHIFT
	}
	if e.Button&f.MouseButtons_Left != 0 {
		flags |= wkeMouseFlags_LBUTTON
	}
	if e.Button&f.MouseButtons_Right != 0 {
		flags |= wkeMouseFlags_RBUTTON
	}
	if e.Button&f.MouseButtons_Middle != 0 {
		flags |= wkeMouseFlags_MBUTTON
	}
	if mbApi.wkeFireMouseWheelEvent(_this._wke, int32(e.X), int32(e.Y), int32(e.Delta), int32(flags)) {
		e.IsHandle = true
	}
}

func (_this *free4x64) viewMouseUp(e f.MouseEvArgs) {
	_this.viewMouseClick(e, false)
}

func (_this *free4x64) viewMouseDown(e f.MouseEvArgs) {
	_this.viewMouseClick(e, true)
}

func (_this *free4x64) viewMouseClick(e f.MouseEvArgs, isDown bool) {
	flags := wkeMouseFlags_None
	keys := c.App.ModifierKeys()
	if s, ok := keys[f.Keys_Ctrl]; ok && s {
		flags |= wkeMouseFlags_CONTROL
	}
	if s, ok := keys[f.Keys_Shift]; ok && s {
		flags |= wkeMouseFlags_SHIFT
	}
	msg := 0
	if e.Button&f.MouseButtons_Left != 0 {
		flags |= wkeMouseFlags_LBUTTON
		if e.IsDouble {
			msg = win32.WM_LBUTTONDBLCLK
		} else if isDown {
			msg = win32.WM_LBUTTONDOWN
		} else {
			msg = win32.WM_LBUTTONUP
		}
	}
	if e.Button&f.MouseButtons_Right != 0 {
		flags |= wkeMouseFlags_RBUTTON
		if e.IsDouble {
			msg = win32.WM_RBUTTONDBLCLK
		} else if isDown {
			msg = win32.WM_RBUTTONDOWN
		} else {
			msg = win32.WM_RBUTTONUP
		}
	}
	if e.Button&f.MouseButtons_Middle != 0 {
		flags |= wkeMouseFlags_MBUTTON
		if e.IsDouble {
			msg = win32.WM_MBUTTONDBLCLK
		} else if isDown {
			msg = win32.WM_MBUTTONDOWN
		} else {
			msg = win32.WM_MBUTTONUP
		}
	}
	if msg != 0 && mbApi.wkeFireMouseEvent(_this._wke, int32(msg), int32(e.X), int32(e.Y), int32(flags)) {
		e.IsHandle = true
	}
}

func (_this *free4x64) viewMouseMove(e f.MouseEvArgs) {
	flags := wkeMouseFlags_None
	if e.Button&f.MouseButtons_Left != 0 {
		flags |= wkeMouseFlags_LBUTTON
	}
	if e.Button&f.MouseButtons_Right != 0 {
		flags |= wkeMouseFlags_RBUTTON
	}
	if mbApi.wkeFireMouseEvent(_this._wke, int32(win32.WM_MOUSEMOVE), int32(e.X), int32(e.Y), int32(flags)) {
		e.IsHandle = true
	}
}

func (_this *free4x64) viewPaint(e f.PaintEvArgs) {
	w := mbApi.wkeGetWidth(_this._wke)
	h := mbApi.wkeGetHeight(_this._wke)
	if w > 0 && h > 0 {
		view := image.NewRGBA(image.Rect(0, 0, int(w), int(h)))
		mbApi.wkePaint(_this._wke, view.Pix, 0)
		e.Graphics.DrawImage(view, e.Clip.X, e.Clip.Y, e.Clip.Width, e.Clip.Height, e.Clip.X, e.Clip.Y)
	}
}

func (_this *free4x64) onPaintBitUpdated(wke wkeHandle, param, bits uintptr, rect *wkeRect, width, height int32) uintptr {
	bx, by := int(rect.x), int(rect.y)
	bw, bh := int(math.Min(float64(rect.w), float64(width))), int(math.Min(float64(rect.h), float64(mbApi.wkeGetHeight(wke))))
	bmp := image.NewRGBA(image.Rect(0, 0, bw, bh))
	stride := int(width) * 4
	pixs := (*[1 << 30]byte)(unsafe.Pointer(bits))
	for y := 0; y < bh; y++ {
		for x := 0; x < bw*4; x++ {
			sp := bmp.Stride*y + x
			dp := stride*(by+y) + bx*4 + x
			bmp.Pix[sp] = pixs[dp]
		}
	}
	_this._view.CreateGraphics().DrawImage(bmp, 0, 0, bw, bh, bx, by).Close()
	return 0
}

func (_this *free4x64) viewResize(e f.Rect) {
	mbApi.wkeResize(_this._wke, uint32(e.Width), uint32(e.Height))
}

func (_this *free4x64) GetHandle() wkeHandle {
	return _this._wke
}
