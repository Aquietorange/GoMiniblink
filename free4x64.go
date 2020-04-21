package GoMiniblink

import (
	"fmt"
	"image"
	"math"
	f "qq2564874169/goMiniblink/forms"
	c "qq2564874169/goMiniblink/forms/controls"
	w "qq2564874169/goMiniblink/forms/platform/windows/win32"
	"strconv"
	"unsafe"
)

func init() {
	mbApi = new(freeApiForWindows).init()
}

type free4x64 struct {
	_view      *c.Control
	_wke       wkeHandle
	_onRequest RequestCallback
}

func (_this *free4x64) init(control *c.Control) *free4x64 {
	_this._view = control
	_this.setView()
	_this.mbInit()
	return _this
}

func (_this *free4x64) BindFunc(fn GoFunc) {

}

func (_this *free4x64) SetOnRequest(callback RequestCallback) {
	_this._onRequest = callback
}

func (_this *free4x64) mbInit() {
	_this._wke = createWebView(_this)
	_this.viewResize(_this._view.GetSize())
	mbApi.wkeSetHandle(_this._wke, _this._view.GetHandle())
	mbApi.wkeOnPaintBitUpdated(_this._wke, _this.onPaintBitUpdated, nil)
	mbApi.wkeOnLoadUrlBegin(_this._wke, _this.onUrlBegin, nil)
}

func (_this *free4x64) onUrlBegin(_ wkeHandle, _, utf8ptr uintptr, job wkeNetJob) uintptr {
	if _this._onRequest == nil {
		return uintptr(toBool(false))
	}
	url := wkePtrToUtf8(utf8ptr)
	e := new(freeRequestEvArgs).init(_this, url, job)
	_this._onRequest(e)
	return uintptr(toBool(e.onBegin()))
}

func (_this *free4x64) setView() {
	bakFocus := _this._view.OnFocus
	_this._view.OnFocus = func() {
		_this.viewFocus()
		if bakFocus != nil {
			bakFocus()
		}
	}
	bakLostFocus := _this._view.OnLostFocus
	_this._view.OnLostFocus = func() {
		_this.viewLostFocus()
		if bakLostFocus != nil {
			bakLostFocus()
		}
	}
	bakResize := _this._view.OnResize
	_this._view.OnResize = func(e f.Rect) {
		_this.viewResize(e)
		if bakResize != nil {
			bakResize(e)
		}
	}
	bakPaint := _this._view.OnPaint
	_this._view.OnPaint = func(e f.PaintEvArgs) {
		_this.viewPaint(e)
		if bakPaint != nil {
			bakPaint(e)
		}
	}
	bakMouseMove := _this._view.OnMouseMove
	_this._view.OnMouseMove = func(e *f.MouseEvArgs) {
		_this.viewMouseMove(e)
		if bakMouseMove != nil {
			bakMouseMove(e)
		}
	}
	bakMouseDown := _this._view.OnMouseDown
	_this._view.OnMouseDown = func(e *f.MouseEvArgs) {
		_this.viewMouseDown(e)
		if bakMouseDown != nil {
			bakMouseDown(e)
		}
	}
	bakMouseUp := _this._view.OnMouseUp
	_this._view.OnMouseUp = func(e *f.MouseEvArgs) {
		_this.viewMouseUp(e)
		if bakMouseUp != nil {
			bakMouseUp(e)
		}
	}
	bakMouseWheel := _this._view.OnMouseWheel
	_this._view.OnMouseWheel = func(e *f.MouseEvArgs) {
		_this.viewMouseWheel(e)
		if bakMouseWheel != nil {
			bakMouseWheel(e)
		}
	}
	bakSetCursor := _this._view.OnSetCursor
	_this._view.OnSetCursor = func() bool {
		b := _this.viewSetCursor()
		if !b && bakSetCursor != nil {
			b = bakSetCursor()
		}
		return b
	}
	bakKeyDown := _this._view.OnKeyDown
	_this._view.OnKeyDown = func(e *f.KeyEvArgs) {
		_this.viewKeyDown(e)
		if bakKeyDown != nil {
			bakKeyDown(e)
		}
	}
	bakKeyUp := _this._view.OnKeyUp
	_this._view.OnKeyUp = func(e *f.KeyEvArgs) {
		_this.viewKeyUp(e)
		if bakKeyUp != nil {
			bakKeyUp(e)
		}
	}
	bakKeyPress := _this._view.OnKeyPress
	_this._view.OnKeyPress = func(e *f.KeyPressEvArgs) {
		_this.viewKeyPress(e)
		if bakKeyPress != nil {
			bakKeyPress(e)
		}
	}
	bakImeStart := _this._view.OnImeStartComposition
	_this._view.OnImeStartComposition = func() bool {
		b := _this.viewImeStart()
		if !b && bakImeStart != nil {
			b = bakImeStart()
		}
		return b
	}
}

func (_this *free4x64) viewImeStart() bool {
	rect := mbApi.wkeGetCaretRect(_this._wke)
	comp := w.COMPOSITIONFORM{
		DwStyle: w.CFS_POINT | w.CFS_FORCE_POSITION,
		Pos: w.POINT{
			X: rect.x,
			Y: rect.y,
		},
	}
	h := w.HWND(_this._view.GetHandle())
	imc := w.ImmGetContext(h)
	w.ImmSetCompositionWindow(imc, &comp)
	w.ImmReleaseContext(h, imc)
	return true
}

func (_this *free4x64) viewKeyPress(e *f.KeyPressEvArgs) {
	if mbApi.wkeFireKeyPressEvent(_this._wke, int([]rune(e.KeyChar)[0]), uint32(wkeKeyFlags_Repeat), e.IsSys) {
		e.IsHandle = true
	}
}

func (_this *free4x64) viewKeyUp(e *f.KeyEvArgs) {
	if _this.viewKeyEvent(e, false) {
		e.IsHandle = true
	}
}

func (_this *free4x64) viewKeyDown(e *f.KeyEvArgs) {
	if _this.viewKeyEvent(e, true) {
		e.IsHandle = true
	}
}

func (_this *free4x64) viewKeyEvent(e *f.KeyEvArgs, isDown bool) bool {
	flags := int(wkeKeyFlags_Repeat)
	switch e.Key {
	case f.Keys_Insert, f.Keys_Delete, f.Keys_Home, f.Keys_End, f.Keys_PageUp,
		f.Keys_PageDown, f.Keys_Left, f.Keys_Right, f.Keys_Up, f.Keys_Down:
		flags |= int(wkeKeyFlags_Extend)
	}
	if isDown {
		return mbApi.wkeFireKeyDownEvent(_this._wke, uint32(e.Value), uint32(flags), e.IsSys)
	} else {
		return mbApi.wkeFireKeyUpEvent(_this._wke, uint32(e.Value), uint32(flags), e.IsSys)
	}
}

func (_this *free4x64) LoadUri(uri string) {
	mbApi.wkeLoadURL(_this._wke, uri)
}

func (_this *free4x64) viewSetCursor() bool {
	cur := mbApi.wkeGetCursorInfoType(_this._wke)
	newCur := f.CursorType_Default
	switch cur {
	case wkeCursorType_Pointer:
		newCur = f.CursorType_ARROW
	case wkeCursorType_Hand:
		newCur = f.CursorType_HAND
	case wkeCursorType_IBeam:
		newCur = f.CursorType_IBEAM
	case wkeCursorType_ColumnResize:
		newCur = f.CursorType_SIZEWE
	case wkeCursorType_RowResize:
		newCur = f.CursorType_SIZENS
	case wkeCursorType_Cross:
		newCur = f.CursorType_CROSS
	default:
		fmt.Println("未实现的鼠标指针类型：" + strconv.Itoa(int(cur)))
	}
	if newCur != f.CursorType_Default {
		_this._view.SetCursor(newCur)
		return true
	}
	return false
}

func (_this *free4x64) viewMouseWheel(e *f.MouseEvArgs) {
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

func (_this *free4x64) viewMouseUp(e *f.MouseEvArgs) {
	_this.viewMouseEvent(e, false)
}

func (_this *free4x64) viewMouseDown(e *f.MouseEvArgs) {
	_this.viewMouseEvent(e, true)
}

func (_this *free4x64) viewMouseEvent(e *f.MouseEvArgs, isDown bool) {
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
			msg = w.WM_LBUTTONDBLCLK
		} else if isDown {
			msg = w.WM_LBUTTONDOWN
		} else {
			msg = w.WM_LBUTTONUP
		}
	}
	if e.Button&f.MouseButtons_Right != 0 {
		flags |= wkeMouseFlags_RBUTTON
		if e.IsDouble {
			msg = w.WM_RBUTTONDBLCLK
		} else if isDown {
			msg = w.WM_RBUTTONDOWN
		} else {
			msg = w.WM_RBUTTONUP
		}
	}
	if e.Button&f.MouseButtons_Middle != 0 {
		flags |= wkeMouseFlags_MBUTTON
		if e.IsDouble {
			msg = w.WM_MBUTTONDBLCLK
		} else if isDown {
			msg = w.WM_MBUTTONDOWN
		} else {
			msg = w.WM_MBUTTONUP
		}
	}
	if msg != 0 && mbApi.wkeFireMouseEvent(_this._wke, int32(msg), int32(e.X), int32(e.Y), int32(flags)) {
		e.IsHandle = true
	}
}

func (_this *free4x64) viewMouseMove(e *f.MouseEvArgs) {
	flags := wkeMouseFlags_None
	if e.Button&f.MouseButtons_Left != 0 {
		flags |= wkeMouseFlags_LBUTTON
	}
	if e.Button&f.MouseButtons_Right != 0 {
		flags |= wkeMouseFlags_RBUTTON
	}
	if mbApi.wkeFireMouseEvent(_this._wke, int32(w.WM_MOUSEMOVE), int32(e.X), int32(e.Y), int32(flags)) {
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

func (_this *free4x64) viewLostFocus() {
	mbApi.wkeKillFocus(_this._wke)
}

func (_this *free4x64) viewFocus() {
	mbApi.wkeSetFocus(_this._wke)
}

func (_this *free4x64) GetHandle() wkeHandle {
	return _this._wke
}
