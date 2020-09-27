package miniblink

import (
	"image"
	"image/draw"
	"math"
	mb "qq2564874169/goMiniblink"
	plat "qq2564874169/goMiniblink/platform"
	"qq2564874169/goMiniblink/platform/windows/win32"
	"unsafe"
)

type FreeVer struct {
	app   plat.IProvider
	owner plat.IWindow
	wke   wkeHandle

	onPaint   PaintCallback
	onRequest RequestCallback
	onJsReady JsReadyCallback
}

func (_this *FreeVer) Init(window plat.IWindow) *FreeVer {
	_this.app = window.GetProvider()
	_this.owner = window
	_this.wke = createWke(_this)
	if _this.wke == 0 {
		panic("创建失败")
	}

	wkeSetHandle(_this.wke, _this.owner.GetHandle())
	wkeOnPaintBitUpdated(_this.wke, _this.onPaintBitUpdated, nil)
	wkeOnLoadUrlBegin(_this.wke, _this.onUrlBegin, nil)
	wkeOnDidCreateScriptContext(_this.wke, _this.onDidCreateScriptContext, nil)
	return _this
}

func (_this *FreeVer) SetWindowProp(name string, value interface{}) {
	es := wkeGlobalExec(_this.wke)
	v := toJsValue(_this, es, value)
	jsSetGlobal(es, name, v)
}

func (_this *FreeVer) onDidCreateScriptContext(wke wkeHandle, param uintptr, frame wkeFrame, context uintptr, exGroup, worldId int) uintptr {
	if _this.onJsReady == nil {
		return 0
	}
	args := new(wkeJsReadyEvArgs).init()
	args.ctx = new(wkeFrameContext).init(_this, frame)
	_this.onJsReady(args)
	return 0
}

func (_this *FreeVer) RunJs(script string) interface{} {
	es := wkeGlobalExec(wkeHandle(_this.GetHandle()))
	rs := jsEvalExW(es, script, true)
	return toGoValue(_this, es, rs)
}

func (_this *FreeVer) jsFuncCallback(es jsExecState, state uintptr) jsValue {
	count := jsArgCount(es)
	ps := make([]interface{}, count)
	for i := 0; i < int(count); i++ {
		value := jsArg(es, uint32(i))
		ps[i] = toGoValue(_this, es, value)
	}
	if fn, ok := _jsFns[state]; ok {
		rs := fn.OnExecute(ps)
		return toJsValue(_this, es, rs)
	}
	return 0
}

func (_this *FreeVer) BindJsFunc(fn mb.JsFuncBinding) {
	id := uintptr(unsafe.Pointer(&fn))
	_jsFns[id] = &fn
	wkeJsBindFunction(fn.Name, _this.jsFuncCallback, id, 0)
}

func (_this *FreeVer) onUrlBegin(_ wkeHandle, _, utf8ptr uintptr, job wkeNetJob) uintptr {
	if _this.onRequest == nil {
		return 0
	}
	url := wkePtrToUtf8(utf8ptr)
	e := new(wkeRequestEvArgs).init(_this, url, job)
	_this.onRequest(e)
	e.OnBegin()
	return 0
}

func (_this *FreeVer) SetOnJsReady(callback JsReadyCallback) {
	_this.onJsReady = callback
}

func (_this *FreeVer) SetOnRequest(callback RequestCallback) {
	_this.onRequest = callback
}

func (_this *FreeVer) SetFocus() {
	wkeSetFocus(_this.wke)
}

func (_this *FreeVer) GetCaretPos() mb.Point {
	rect := wkeGetCaretRect(_this.wke)
	return mb.Point{X: int(rect.x), Y: int(rect.y)}
}

func (_this *FreeVer) FireKeyPressEvent(charCode int, isSys bool) bool {
	return wkeFireKeyPressEvent(_this.wke, charCode, uint32(wkeKeyFlags_Repeat), isSys)
}

func (_this *FreeVer) FireKeyEvent(e mb.KeyEvArgs, isDown, isSys bool) bool {
	flags := int(wkeKeyFlags_Repeat)
	if isExtKey(e.Key) {
		flags |= int(wkeKeyFlags_Extend)
	}
	if isDown {
		return wkeFireKeyDownEvent(_this.wke, uint32(e.Value), uint32(flags), isSys)
	} else {
		return wkeFireKeyUpEvent(_this.wke, uint32(e.Value), uint32(flags), isSys)
	}
}

func (_this *FreeVer) GetCursor() mb.CursorType {
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

func (_this *FreeVer) FireMouseWheelEvent(button mb.MouseButtons, delta, x, y int) bool {
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

func (_this *FreeVer) FireMouseMoveEvent(button mb.MouseButtons, x, y int) bool {
	flags := wkeMouseFlags_None
	if button&mb.MouseButtons_Left != 0 {
		flags |= wkeMouseFlags_LBUTTON
	}
	if button&mb.MouseButtons_Right != 0 {
		flags |= wkeMouseFlags_RBUTTON
	}
	return wkeFireMouseEvent(_this.wke, int32(win32.WM_MOUSEMOVE), int32(x), int32(y), int32(flags))
}

func (_this *FreeVer) FireMouseClickEvent(button mb.MouseButtons, isDown, isDb bool, x, y int) bool {
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

func (_this *FreeVer) GetImage(bound mb.Bound) *image.RGBA {
	bmp := image.NewRGBA(image.Rect(0, 0, bound.Width, bound.Height))
	w := wkeGetWidth(_this.wke)
	h := wkeGetHeight(_this.wke)
	if w > 0 && h > 0 {
		view := image.NewRGBA(image.Rect(0, 0, int(w), int(h)))
		wkePaint(_this.wke, view.Pix, 0)
		draw.Draw(bmp, image.Rect(0, 0, bound.Width, bound.Height), view, image.Pt(bound.X, bound.Y), draw.Src)
	}
	return bmp
}

func (_this *FreeVer) onPaintBitUpdated(wke wkeHandle, param, bits uintptr, rect *wkeRect, width, height int32) uintptr {
	e := PaintUpdateArgs{
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

func (_this *FreeVer) Resize(width, height int) {
	wkeResize(_this.wke, uint32(width), uint32(height))
}

func (_this *FreeVer) SetOnPaint(callback PaintCallback) {
	_this.onPaint = callback
}

func (_this *FreeVer) LoadUri(uri string) {
	wkeLoadURL(_this.wke, uri)
}

func (_this *FreeVer) Invoke(fn func(interface{}), state interface{}) {
	_this.owner.Invoke(fn, state)
}

func (_this *FreeVer) GetHandle() uintptr {
	return uintptr(_this.wke)
}
