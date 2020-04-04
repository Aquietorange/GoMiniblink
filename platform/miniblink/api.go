package miniblink

import (
	"fmt"
	"golang.org/x/sys/windows"
	"strconv"
	"syscall"
	"unsafe"
)

const (
	file_x86_dll = "miniblink_x86.dll"
	file_x64_dll = "miniblink_x64.dll"
)

type (
	wkeHandle   uintptr
	jsExecState uintptr
	jsValue     int64
)

type jsType uint32

const (
	jsType_NUMBER jsType = iota
	jsType_STRING
	jsType_BOOLEAN
	jsType_OBJECT
	jsType_FUNCTION
	jsType_UNDEFINED
	jsType_ARRAY
	jsType_NULL
)

var (
	is64 bool
	lib  *windows.LazyDLL

	_wkeInitialize          *windows.LazyProc
	_wkeCreateWebView       *windows.LazyProc
	_wkeSetHandle           *windows.LazyProc
	_wkeOnPaintBitUpdated   *windows.LazyProc
	_wkeLoadURL             *windows.LazyProc
	_wkeResize              *windows.LazyProc
	_wkeNetOnResponse       *windows.LazyProc
	_wkeOnLoadUrlBegin      *windows.LazyProc
	_wkePaint               *windows.LazyProc
	_wkeGetWidth            *windows.LazyProc
	_wkeGetHeight           *windows.LazyProc
	_wkeFireMouseEvent      *windows.LazyProc
	_wkeFireMouseWheelEvent *windows.LazyProc
	_wkeGetCursorInfoType   *windows.LazyProc
	_wkeFireKeyUpEvent      *windows.LazyProc
	_wkeFireKeyDownEvent    *windows.LazyProc
	_wkeFireKeyPressEvent   *windows.LazyProc
	_wkeGetCaretRect        *windows.LazyProc
	_wkeSetFocus            *windows.LazyProc
	_wkeNetGetRequestMethod *windows.LazyProc
	_wkeNetSetData          *windows.LazyProc
	_wkeNetCancelRequest    *windows.LazyProc
	_wkeJsBindFunction      *windows.LazyProc
	_jsArgCount             *windows.LazyProc
	_jsArg                  *windows.LazyProc
	_jsTypeOf               *windows.LazyProc
	_jsToTempString         *windows.LazyProc
	_jsToDoubleString       *windows.LazyProc
	_jsToInt                *windows.LazyProc
	_jsToBoolean            *windows.LazyProc
	_jsGetLength            *windows.LazyProc
	_jsGetAt                *windows.LazyProc
	_jsGetKeys              *windows.LazyProc
	_jsGet                  *windows.LazyProc
	_jsSetGlobal            *windows.LazyProc
	_jsGetGlobal            *windows.LazyProc
	_wkeGlobalExec          *windows.LazyProc
	_jsCall                 *windows.LazyProc
	_jsUndefined            *windows.LazyProc
	_jsInt                  *windows.LazyProc
	_jsBoolean              *windows.LazyProc
	_jsDouble               *windows.LazyProc
	_jsFloat                *windows.LazyProc
	_jsString               *windows.LazyProc
	_jsEmptyArray           *windows.LazyProc
	_jsSetLength            *windows.LazyProc
	_jsSetAt                *windows.LazyProc
	_jsFunction             *windows.LazyProc
	_jsEmptyObject          *windows.LazyProc
	_jsSet                  *windows.LazyProc
)

func init() {
	is64 = unsafe.Sizeof(uintptr(0)) == 8
	if is64 {
		lib = windows.NewLazyDLL(file_x64_dll)
	} else {
		lib = windows.NewLazyDLL(file_x86_dll)
	}
	_jsToInt = lib.NewProc("jsToInt")
	_jsSet = lib.NewProc("jsSet")
	_jsEmptyObject = lib.NewProc("jsEmptyObject")
	_jsFunction = lib.NewProc("jsFunction")
	_jsSetAt = lib.NewProc("jsSetAt")
	_jsSetLength = lib.NewProc("jsSetLength")
	_jsEmptyArray = lib.NewProc("jsEmptyArray")
	_jsString = lib.NewProc("jsString")
	_jsFloat = lib.NewProc("jsFloat")
	_jsDouble = lib.NewProc("jsDouble")
	_jsBoolean = lib.NewProc("jsBoolean")
	_jsInt = lib.NewProc("jsInt")
	_jsUndefined = lib.NewProc("jsUndefined")
	_jsCall = lib.NewProc("jsCall")
	_wkeGlobalExec = lib.NewProc("wkeGlobalExec")
	_jsGetGlobal = lib.NewProc("jsGetGlobal")
	_jsSetGlobal = lib.NewProc("jsSetGlobal")
	_jsGet = lib.NewProc("jsGet")
	_jsGetKeys = lib.NewProc("jsGetKeys")
	_jsGetAt = lib.NewProc("jsGetAt")
	_jsGetLength = lib.NewProc("jsGetLength")
	_jsToBoolean = lib.NewProc("jsToBoolean")
	_jsToDoubleString = lib.NewProc("jsToDoubleString")
	_jsToTempString = lib.NewProc("jsToTempString")
	_jsTypeOf = lib.NewProc("jsTypeOf")
	_jsArg = lib.NewProc("jsArg")
	_jsArgCount = lib.NewProc("jsArgCount")
	_wkeJsBindFunction = lib.NewProc("wkeJsBindFunction")
	_wkeNetCancelRequest = lib.NewProc("wkeNetCancelRequest")
	_wkeNetSetData = lib.NewProc("wkeNetSetData")
	_wkeNetGetRequestMethod = lib.NewProc("wkeNetGetRequestMethod")
	_wkeFireKeyPressEvent = lib.NewProc("wkeFireKeyPressEvent")
	_wkeFireKeyUpEvent = lib.NewProc("wkeFireKeyUpEvent")
	_wkeFireKeyDownEvent = lib.NewProc("wkeFireKeyDownEvent")
	_wkeGetCursorInfoType = lib.NewProc("wkeGetCursorInfoType")
	_wkeFireMouseWheelEvent = lib.NewProc("wkeFireMouseWheelEvent")
	_wkeFireMouseEvent = lib.NewProc("wkeFireMouseEvent")
	_wkeGetHeight = lib.NewProc("wkeGetHeight")
	_wkeGetWidth = lib.NewProc("wkeGetWidth")
	_wkePaint = lib.NewProc("wkePaint")
	_wkeOnLoadUrlBegin = lib.NewProc("wkeOnLoadUrlBegin")
	_wkeNetOnResponse = lib.NewProc("wkeNetOnResponse")
	_wkeLoadURL = lib.NewProc("wkeLoadURL")
	_wkeResize = lib.NewProc("wkeResize")
	_wkeOnPaintBitUpdated = lib.NewProc("wkeOnPaintBitUpdated")
	_wkeSetHandle = lib.NewProc("wkeSetHandle")
	_wkeCreateWebView = lib.NewProc("wkeCreateWebView")
	_wkeInitialize = lib.NewProc("wkeInitialize")
	_wkeGetCaretRect = lib.NewProc("wkeGetCaretRect2")
	_wkeSetFocus = lib.NewProc("wkeSetFocus")

	ret, _, err := _wkeInitialize.Call()
	if ret == 0 {
		fmt.Println(err)
	}
}

func _toInt64(low, high int32) int64 {
	var l = int64(high)<<32 + int64(low)
	return *((*int64)(unsafe.Pointer(&l)))
}

func _toLH(value jsValue) (low, high int32) {
	if is64 {
		return 0, 0
	}
	return int32(int64(value)), int32(int64(value) >> 32 & 0xffffffff)
}

func jsSet(es jsExecState, obj jsValue, name string, value jsValue) {
	ptr := []byte(name)
	if is64 {
		_jsSet.Call(uintptr(es), uintptr(obj), uintptr(unsafe.Pointer(&ptr[0])), uintptr(value))
	} else {
		l1, h1 := _toLH(obj)
		l2, h2 := _toLH(value)
		_jsSet.Call(uintptr(es), uintptr(l1), uintptr(h1), uintptr(unsafe.Pointer(&ptr[0])), uintptr(l2), uintptr(h2))
	}
}

func jsEmptyObject(es jsExecState) jsValue {
	r, _, _ := _jsEmptyObject.Call(uintptr(es))
	return jsValue(r)
}

func jsFunction(es jsExecState, data *jsData) jsValue {
	r, _, _ := _jsFunction.Call(uintptr(es), uintptr(unsafe.Pointer(data)))
	return jsValue(r)
}

func jsSetAt(es jsExecState, array jsValue, index uint32, value jsValue) {
	if is64 {
		_jsSetAt.Call(uintptr(es), uintptr(array), uintptr(index), uintptr(value))
	} else {
		l1, h1 := _toLH(array)
		l2, h2 := _toLH(value)
		_jsSetAt.Call(uintptr(es), uintptr(l1), uintptr(h1), uintptr(index), uintptr(l2), uintptr(h2))
	}
}

func jsSetLength(es jsExecState, array jsValue, length uint32) {
	if is64 {
		_jsSetLength.Call(uintptr(es), uintptr(array), uintptr(length))
	} else {
		l, h := _toLH(array)
		_jsSetLength.Call(uintptr(es), uintptr(l), uintptr(h), uintptr(length))
	}
}

func jsEmptyArray(es jsExecState) jsValue {
	r, _, _ := _jsEmptyArray.Call(uintptr(es))
	return jsValue(r)
}

func jsString(es jsExecState, value string) jsValue {
	ptr := toCallStr(value)
	r, _, _ := _jsString.Call(uintptr(es), uintptr(unsafe.Pointer(&ptr[0])))
	return jsValue(r)
}

func jsFloat(value float32) jsValue {
	r, _, _ := _jsFloat.Call(uintptr(value))
	return jsValue(r)
}

func jsDouble(value float64) jsValue {
	r, _, _ := _jsDouble.Call(uintptr(value))
	return jsValue(r)
}

func jsBoolean(value bool) jsValue {
	r, _, _ := _jsBoolean.Call(uintptr(toBool(value)))
	return jsValue(r)
}

func jsInt(value int32) jsValue {
	fmt.Println(value)
	l, h, _ := _jsInt.Call(uintptr(value))
	fmt.Println(l, h)
	n := _toInt64(int32(l), int32(h))
	fmt.Println(n)
	return jsValue(n)
}

func jsCall(es jsExecState, fn, thisObject jsValue, args []jsValue) jsValue {
	var ptr = uintptr(0)
	if len(args) > 0 {
		ptr = uintptr(unsafe.Pointer(&args[0]))
	}
	r, _, _ := _jsCall.Call(uintptr(es), uintptr(fn), uintptr(thisObject), ptr, uintptr(len(args)))
	return jsValue(r)
}

func wkeGlobalExec(wke wkeHandle) jsExecState {
	r, _, _ := _wkeGlobalExec.Call(uintptr(wke))
	return jsExecState(r)
}

func jsGetGlobal(es jsExecState, name string) jsValue {
	ptr := toCallStr(name)
	r, _, _ := _jsGetGlobal.Call(uintptr(es), uintptr(unsafe.Pointer(&ptr[0])))
	return jsValue(r)
}

func jsSetGlobal(es jsExecState, name string, value jsValue) {
	ptr := toCallStr(name)
	if is64 {
		_jsSetGlobal.Call(uintptr(es), uintptr(unsafe.Pointer(&ptr[0])), uintptr(value))
	} else {
		l, h := _toLH(value)
		_jsSetGlobal.Call(uintptr(es), uintptr(unsafe.Pointer(&ptr[0])), uintptr(l), uintptr(h))
	}
}

func jsGetKeys(es jsExecState, value jsValue) []string {
	var rs uintptr
	if is64 {
		rs, _, _ = _jsGetKeys.Call(uintptr(es), uintptr(value))
	} else {
		l, h := _toLH(value)
		rs, _, _ = _jsGetKeys.Call(uintptr(es), uintptr(l), uintptr(h))
	}
	keys := *((*jsKeys)(unsafe.Pointer(rs)))
	items := make([]string, keys.length)
	for i := 0; i < int(keys.length); i++ {
		items[i] = string(keys.first)
		keys.first += unsafe.Sizeof(keys.first)
	}
	return items
}

func jsGet(es jsExecState, value jsValue, name string) jsValue {
	ptr := toCallStr(name)
	if is64 {
		r, _, _ := _jsGet.Call(uintptr(es), uintptr(value), uintptr(unsafe.Pointer(&ptr[0])))
		return jsValue(r)
	} else {
		l, h := _toLH(value)
		r, _, _ := _jsGet.Call(uintptr(es), uintptr(l), uintptr(h), uintptr(unsafe.Pointer(&ptr[0])))
		return jsValue(r)
	}
}

func jsGetAt(es jsExecState, value jsValue, index uint32) jsValue {
	if is64 {
		r, _, _ := _jsGetAt.Call(uintptr(es), uintptr(value), uintptr(index))
		return jsValue(r)
	} else {
		l, h := _toLH(value)
		r, _, _ := _jsGetAt.Call(uintptr(es), uintptr(l), uintptr(h), uintptr(index))
		return jsValue(r)
	}
}

func jsGetLength(es jsExecState, value jsValue) int {
	if is64 {
		r, _, _ := _jsGetLength.Call(uintptr(es), uintptr(value))
		return int(r)
	} else {
		l, h := _toLH(value)
		r, _, _ := _jsGetLength.Call(uintptr(es), uintptr(l), uintptr(h))
		return int(r)
	}
}

func jsUndefined() jsValue {
	r, _, _ := _jsUndefined.Call()
	return jsValue(r)
}

func jsToBoolean(es jsExecState, value jsValue) bool {
	if is64 {
		r, _, _ := _jsToBoolean.Call(uintptr(es), uintptr(value))
		return byte(r) != 0
	} else {
		l, h := _toLH(value)
		r, _, _ := _jsToBoolean.Call(uintptr(es), uintptr(l), uintptr(h))
		return byte(r) != 0
	}
}

func jsToDouble(es jsExecState, value jsValue) float64 {
	var r uintptr
	if is64 {
		r, _, _ = _jsToDoubleString.Call(uintptr(es), uintptr(value))
	} else {
		l, h := _toLH(value)
		r, _, _ = _jsToDoubleString.Call(uintptr(es), uintptr(l), uintptr(h))
	}
	str := wkePtrToUtf8(r)
	n, _ := strconv.ParseFloat(str, 10)
	return n
}

func jsToTempString(es jsExecState, value jsValue) string {
	if is64 {
		r, _, _ := _jsToTempString.Call(uintptr(es), uintptr(value))
		return wkePtrToUtf8(r)
	} else {
		l, h := _toLH(value)
		r, _, _ := _jsToTempString.Call(uintptr(es), uintptr(l), uintptr(h))
		return wkePtrToUtf8(r)
	}
}

func jsTypeOf(value jsValue) jsType {
	if is64 {
		r, _, _ := _jsTypeOf.Call(uintptr(value))
		return jsType(r)
	} else {
		l, h := _toLH(value)
		r, _, _ := _jsTypeOf.Call(uintptr(l), uintptr(h))
		return jsType(r)
	}
}

func jsArg(es jsExecState, argIdx uint32) jsValue {
	r, _, _ := _jsArg.Call(uintptr(es), uintptr(argIdx))
	return jsValue(r)
}

func jsArgCount(es jsExecState) uint32 {
	r, _, _ := _jsArgCount.Call(uintptr(es))
	return uint32(r)
}

func wkeJsBindFunction(name string, fn wkeJsNativeFunction, param uintptr, argCount uint32) {
	ptr := toCallStr(name)
	_wkeJsBindFunction.Call(uintptr(unsafe.Pointer(&ptr[0])), syscall.NewCallbackCDecl(fn), param, uintptr(argCount))
}

func wkeNetCancelRequest(job wkeNetJob) {
	_wkeNetCancelRequest.Call(uintptr(job))
}

func wkeNetOnResponse(wke wkeHandle, callback wkeNetResponseCallback, param unsafe.Pointer) {
	_wkeNetOnResponse.Call(uintptr(wke), syscall.NewCallbackCDecl(callback), uintptr(param))
}

func wkeOnLoadUrlBegin(wke wkeHandle, callback wkeLoadUrlBeginCallback, param unsafe.Pointer) {
	_wkeOnLoadUrlBegin.Call(uintptr(wke), syscall.NewCallbackCDecl(callback), uintptr(param))
}

func wkeNetGetRequestMethod(job wkeNetJob) wkeRequestType {
	r, _, _ := _wkeNetGetRequestMethod.Call(uintptr(job))
	return wkeRequestType(r)
}

func wkeNetSetData(job wkeNetJob, buf []byte, len uint32) {
	_wkeNetSetData.Call(uintptr(job), uintptr(unsafe.Pointer(&buf[0])), uintptr(len))
}

func wkeGetCaretRect(wke wkeHandle) wkeRect {
	r, _, _ := _wkeGetCaretRect.Call(uintptr(wke))
	return *((*wkeRect)(unsafe.Pointer(r)))
}

func wkeSetFocus(wke wkeHandle) {
	_wkeSetFocus.Call(uintptr(wke))
}

func wkeFireKeyPressEvent(wke wkeHandle, code int, flags uint32, isSysKey bool) bool {
	ret, _, _ := _wkeFireKeyPressEvent.Call(
		uintptr(wke),
		uintptr(code),
		uintptr(flags),
		uintptr(toBool(isSysKey)))
	return byte(ret) != 0
}

func wkeFireKeyDownEvent(wke wkeHandle, code, flags uint32, isSysKey bool) bool {
	ret, _, _ := _wkeFireKeyDownEvent.Call(
		uintptr(wke),
		uintptr(code),
		uintptr(flags),
		uintptr(toBool(isSysKey)))
	return byte(ret) != 0
}

func wkeFireKeyUpEvent(wke wkeHandle, code, flags uint32, isSysKey bool) bool {
	ret, _, _ := _wkeFireKeyUpEvent.Call(
		uintptr(wke),
		uintptr(code),
		uintptr(flags),
		uintptr(toBool(isSysKey)))
	return byte(ret) != 0
}

func wkeGetCursorInfoType(wke wkeHandle) wkeCursorType {
	r, _, _ := _wkeGetCursorInfoType.Call(uintptr(wke))
	return wkeCursorType(r)
}

func wkeFireMouseWheelEvent(wke wkeHandle, x, y, delta, flags int32) bool {
	r, _, _ := _wkeFireMouseWheelEvent.Call(
		uintptr(wke),
		uintptr(x),
		uintptr(y),
		uintptr(delta),
		uintptr(flags))
	return byte(r) != 0
}

func wkeFireMouseEvent(wke wkeHandle, message, x, y, flags int32) bool {
	r, _, _ := _wkeFireMouseEvent.Call(
		uintptr(wke),
		uintptr(message),
		uintptr(x),
		uintptr(y),
		uintptr(flags))
	return byte(r) != 0
}

func wkePaint(wke wkeHandle, bits []byte, pitch uint32) {
	_wkePaint.Call(uintptr(wke), uintptr(unsafe.Pointer(&bits[0])), uintptr(pitch))
}

func wkeGetHeight(wke wkeHandle) uint32 {
	r, _, _ := _wkeGetHeight.Call(uintptr(wke))
	return uint32(r)
}

func wkeGetWidth(wke wkeHandle) uint32 {
	r, _, _ := _wkeGetWidth.Call(uintptr(wke))
	return uint32(r)
}

func wkeResize(wke wkeHandle, w, h uint32) {
	_wkeResize.Call(uintptr(wke), uintptr(w), uintptr(h))
}

func wkeLoadURL(wke wkeHandle, url string) {
	ptr := toCallStr(url)
	_wkeLoadURL.Call(uintptr(wke), uintptr(unsafe.Pointer(&ptr[0])))
}

func wkeOnPaintBitUpdated(wke wkeHandle, callback wkePaintBitUpdatedCallback, param unsafe.Pointer) {
	_wkeOnPaintBitUpdated.Call(uintptr(wke), syscall.NewCallbackCDecl(callback), uintptr(param))
}

func wkeSetHandle(wke wkeHandle, handle uintptr) {
	_wkeSetHandle.Call(uintptr(wke), handle)
}

func wkeCreateWebView() wkeHandle {
	r, _, _ := _wkeCreateWebView.Call()
	return wkeHandle(r)
}

func toBool(b bool) byte {
	if b {
		return 1
	} else {
		return 0
	}
}

func toCallStr(str string) []byte {
	buf := []byte(str)
	rs := make([]byte, len(str)+1)
	for i, v := range buf {
		rs[i] = v
	}
	return rs
}
