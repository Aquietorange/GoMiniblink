package free

import (
	"image"
	"image/draw"
	"math"
	mb "qq2564874169/goMiniblink"
	plat "qq2564874169/goMiniblink/platform"
	core "qq2564874169/goMiniblink/platform/miniblink"
	"qq2564874169/goMiniblink/platform/windows/win32"
	"reflect"
	"strconv"
	"syscall"
	"time"
	"unsafe"
)

var _jsFns = make(map[int64]*mb.GoFunc)
var _ref = make(map[int64]interface{})

type Core struct {
	app   plat.IProvider
	owner plat.IWindow
	wke   wkeHandle

	onPaint   core.PaintCallback
	onRequest core.RequestCallback

	//_jsFunc map[string]wkeJsNativeFunction
}

func (_this *Core) Init(window plat.IWindow) *Core {
	_this.app = window.GetProvider()
	_this.owner = window
	_this.wke = wkeCreateWebView()
	if _this.wke == 0 {
		panic("创建失败")
	}
	//_this._jsFunc = make(map[string]wkeJsNativeFunction)

	wkeSetHandle(_this.wke, _this.owner.GetHandle())
	wkeOnPaintBitUpdated(_this.wke, _this.onPaintBitUpdated, nil)
	wkeOnLoadUrlBegin(_this.wke, _this.onUrlBegin, nil)
	return _this
}

func toJsValue(core *Core, value interface{}, es jsExecState) jsValue {
	if value == nil {
		return jsUndefined()
	}
	switch value.(type) {
	case int:
		return jsInt(int32(value.(int)))
	case int8:
		return jsInt(int32(value.(int8)))
	case int16:
		return jsInt(int32(value.(int16)))
	case int32:
		return jsInt(value.(int32))
	case int64:
		return jsDouble(float64(value.(int64)))
	case uint:
		return jsInt(int32(value.(uint)))
	case uint8:
		return jsInt(int32(value.(uint8)))
	case uint16:
		return jsInt(int32(value.(uint16)))
	case uint32:
		return jsInt(int32(value.(uint32)))
	case uint64:
		return jsDouble(float64(value.(uint64)))
	case float32:
		return jsFloat(value.(float32))
	case float64:
		return jsDouble(value.(float64))
	case bool:
		return jsBoolean(value.(bool))
	case string:
		return jsString(es, value.(string))
	case time.Time:
		return jsDouble(float64(value.(time.Time).Unix()))
	default:
		break
	}
	rv := reflect.ValueOf(value)
	switch rv.Kind() {
	case reflect.Slice, reflect.Array:
		length := rv.Len()
		arr := jsEmptyArray(es)
		jsSetLength(es, arr, uint32(length))
		for i := 0; i < length; i++ {
			v := toJsValue(core, rv.Index(i).Interface(), es)
			jsSetAt(es, arr, uint32(i), v)
		}
		return arr
	case reflect.Map:
		obj := jsEmptyObject(es)
		kv := rv.MapRange()
		for kv.Next() && kv.Key().Kind() == reflect.String {
			k := kv.Key().Interface().(string)
			v := toJsValue(core, kv.Value().Interface(), es)
			jsSet(es, obj, k, v)
		}
		return obj
	case reflect.Struct:
		obj := jsEmptyObject(es)
		for i := 0; i < rv.NumField(); i++ {
			f := rv.Field(i).Type().Name()
			v := toJsValue(core, rv.Field(i).Interface(), es)
			jsSet(es, obj, f, v)
		}
		return obj
	case reflect.Func:
		jsFn := jsData{}
		name, _ := syscall.UTF16FromString("function")
		for i := 0; i < len(name); i++ {
			jsFn.name[i] = name[i]
		}
		var call = func(fnes jsExecState, obj, args jsValue, count uint32) jsValue {
			arr := make([]reflect.Value, count)
			for i := uint32(0); i < count; i++ {
				jv := jsGetAt(fnes, args, i)
				gv := toGoValue(core, jv, fnes)
				arr[i] = reflect.ValueOf(gv)
			}
			rs := rv.Call(arr)
			if len(rs) > 0 {
				return toJsValue(core, rs[0].Interface(), fnes)
			}
			return 0
		}
		jsFn.callAsFunction = syscall.NewCallbackCDecl(call)
		jsFn.finalize = syscall.NewCallbackCDecl(func(ptr uintptr) {
			delete(_ref, int64(ptr))
		})
		_ref[int64(jsFn.callAsFunction)] = call
		return jsFunction(es, &jsFn)
	}
	panic("不支持的go类型：" + rv.Kind().String())
}

func toGoValue(core *Core, value jsValue, es jsExecState) interface{} {
	switch jsTypeOf(value) {
	case jsType_NULL, jsType_UNDEFINED:
		return nil
	case jsType_NUMBER:
		return jsToDouble(es, value)
	case jsType_BOOLEAN:
		return jsToBoolean(es, value)
	case jsType_STRING:
		return jsToTempString(es, value)
	case jsType_ARRAY:
		length := jsGetLength(es, value)
		ps := make([]interface{}, length)
		for i := 0; i < length; i++ {
			v := jsGetAt(es, value, uint32(i))
			ps[i] = toGoValue(core, v, es)
		}
		return ps
	case jsType_OBJECT:
		ps := make(map[string]interface{})
		keys := jsGetKeys(es, value)
		for _, k := range keys {
			v := jsGet(es, value, k)
			ps[k] = toGoValue(core, v, es)
		}
		return ps
	case jsType_FUNCTION:
		name := "func" + strconv.FormatInt(mb.NewId(), 10)
		jsSetGlobal(es, name, value)
		return mb.JsFunc(func(param ...interface{}) interface{} {
			jses := wkeGlobalExec(core.wke)
			ps := make([]jsValue, len(param))
			for i, v := range param {
				ps[i] = toJsValue(core, v, jses)
			}
			fn := jsGetGlobal(jses, name)
			ret := jsCall(jses, fn, jsUndefined(), ps, len(ps))
			jsSetGlobal(jses, name, jsUndefined())
			return toGoValue(core, ret, jses)
		})
	default:
		panic("不支持的js类型：" + strconv.Itoa(int(value)))
	}
}

func (_this *Core) jsFuncCallback(es jsExecState, state uintptr) jsValue {
	count := jsArgCount(es)
	ps := make([]interface{}, count)
	for i := 0; i < int(count); i++ {
		value := jsArg(es, uint32(i))
		ps[count] = toGoValue(_this, value, es)
	}
	if fn, ok := _jsFns[*(*int64)(unsafe.Pointer(state))]; ok {
		ret := fn.OnExecute(ps)
		return toJsValue(_this, ret, es)
	}
	return 0
}

func (_this *Core) BindFunc(fn mb.GoFunc) {
	id := mb.NewId()
	_jsFns[id] = &fn
	wkeJsBindFunction(fn.Name, _this.jsFuncCallback, unsafe.Pointer(&id), 0)
}

func (_this *Core) onUrlBegin(_ wkeHandle, _, utf8ptr uintptr, job wkeNetJob) uintptr {
	if _this.onRequest == nil {
		return uintptr(toBool(false))
	}
	url := wkePtrToUtf8(utf8ptr)
	e := new(wkeRequestEvArgs).init(_this, url, job)
	_this.onRequest(e)
	return uintptr(toBool(e.OnBegin()))
}

func (_this *Core) SetOnRequest(callback core.RequestCallback) {
	_this.onRequest = callback
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
		return wkeFireKeyDownEvent(_this.wke, uint32(e.Value), uint32(flags), isSys)
	} else {
		return wkeFireKeyUpEvent(_this.wke, uint32(e.Value), uint32(flags), isSys)
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
	bmp := image.NewRGBA(image.Rect(0, 0, bound.Width, bound.Height))
	w := wkeGetWidth(_this.wke)
	h := wkeGetHeight(_this.wke)
	if w > 0 && h > 0 {
		view := image.NewRGBA(image.Rect(0, 0, int(w), int(h)))
		wkePaint(_this.wke, &view.Pix[0], 0)
		draw.Draw(bmp, image.Rect(0, 0, bound.Width, bound.Height), view, image.Pt(bound.X, bound.Y), draw.Src)
	}
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

func (_this *Core) SafeInvoke(fn func(interface{}), state interface{}) {
	_this.owner.Invoke(fn, state)
}

func (_this *Core) GetHandle() uintptr {
	return uintptr(_this.wke)
}
