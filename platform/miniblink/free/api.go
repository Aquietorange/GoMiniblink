package free

import (
	"fmt"
	"golang.org/x/sys/windows"
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

type jsType int

const (
	jsType_NUMBER    jsType = 0
	jsType_STRING    jsType = 1
	jsType_BOOLEAN   jsType = 2
	jsType_OBJECT    jsType = 3
	jsType_FUNCTION  jsType = 4
	jsType_UNDEFINED jsType = 5
	jsType_ARRAY     jsType = 6
	jsType_NULL      jsType = 7
)

var (
	lib *windows.LazyDLL

	showError               bool
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
	_jsToDouble             *windows.LazyProc
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
	showError = true
	is64 := unsafe.Sizeof(uintptr(0)) == 8
	if is64 {
		lib = windows.NewLazyDLL(file_x64_dll)
	} else {
		lib = windows.NewLazyDLL(file_x86_dll)
	}
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
	_jsToDouble = lib.NewProc("jsToDouble")
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
	if ret == 0 && showError {
		fmt.Println(err)
	}
}

func jsSet(es jsExecState, obj jsValue, name string, value jsValue) {
	ptr := []byte(name)
	_jsSet.Call(uintptr(es), uintptr(obj), uintptr(unsafe.Pointer(&ptr[0])), uintptr(value))
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
	_jsSetAt.Call(uintptr(es), uintptr(array), uintptr(index), uintptr(value))
}

func jsSetLength(es jsExecState, array jsValue, length uint32) {
	_jsSetLength.Call(uintptr(es), uintptr(array), uintptr(length))
}

func jsEmptyArray(es jsExecState) jsValue {
	r, _, _ := _jsEmptyArray.Call(uintptr(es))
	return jsValue(r)
}

func jsString(es jsExecState, value string) jsValue {
	ptr := []byte(value)
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
	r, _, _ := _jsInt.Call(uintptr(value))
	return jsValue(r)
}

func jsCall(es jsExecState, fn, thisObject jsValue, args []jsValue, argCount int) jsValue {
	var ptr = uintptr(0)
	if len(args) > 0 {
		ptr = uintptr(unsafe.Pointer(&args[0]))
	}
	r, _, _ := _jsCall.Call(uintptr(es), uintptr(fn), uintptr(thisObject), ptr, uintptr(argCount))
	return jsValue(r)
}

func wkeGlobalExec(wke wkeHandle) jsExecState {
	r, _, _ := _wkeGlobalExec.Call(uintptr(wke))
	return jsExecState(r)
}

func jsGetGlobal(es jsExecState, name string) jsValue {
	ptr := []byte(name)
	r, _, _ := _jsGetGlobal.Call(uintptr(es), uintptr(unsafe.Pointer(&ptr[0])))
	return jsValue(r)
}

func jsSetGlobal(es jsExecState, name string, value jsValue) {
	ptr := []byte(name)
	_jsSetGlobal.Call(uintptr(es), uintptr(unsafe.Pointer(&ptr[0])), uintptr(value))
}

func jsGetKeys(es jsExecState, value jsValue) []string {
	r, _, _ := _jsGetKeys.Call(uintptr(es), uintptr(value))
	keys := *((*jsKeys)(unsafe.Pointer(r)))
	items := make([]string, keys.length)
	for i := 0; i < int(keys.length); i++ {
		items[i] = string(keys.first)
		keys.first += unsafe.Sizeof(keys.first)
	}
	return items
}

func jsGet(es jsExecState, value jsValue, name string) jsValue {
	r, _, _ := _jsGet.Call(uintptr(es), uintptr(value))
	return jsValue(r)
}

func jsGetAt(es jsExecState, value jsValue, index uint32) jsValue {
	r, _, _ := _jsGetAt.Call(uintptr(es), uintptr(value), uintptr(index))
	return jsValue(r)
}

func jsGetLength(es jsExecState, value jsValue) int {
	r, _, _ := _jsGetLength.Call(uintptr(es), uintptr(value))
	return int(r)
}

func jsUndefined() jsValue {
	r, _, _ := _jsUndefined.Call()
	return jsValue(r)
}

func jsToBoolean(es jsExecState, value jsValue) bool {
	r, _, _ := _jsToBoolean.Call(uintptr(es), uintptr(value))
	return byte(r) != 0
}

func jsToDouble(es jsExecState, value jsValue) float64 {
	r, _, _ := _jsToDouble.Call(uintptr(es), uintptr(value))
	return float64(r)
}

func jsToTempString(es jsExecState, value jsValue) string {
	r, _, _ := _jsToTempString.Call(uintptr(es), uintptr(value))
	return wkePtrToUtf8(r)
}

func jsTypeOf(value jsValue) jsType {
	r, _, _ := _jsTypeOf.Call(uintptr(value))
	return jsType(r)
}

func jsArg(es jsExecState, index uint32) jsValue {
	r, _, _ := _jsArg.Call(uintptr(es), uintptr(index))
	return jsValue(r)
}

func jsArgCount(es jsExecState) uint32 {
	r, _, _ := _jsArgCount.Call(uintptr(es))
	return uint32(r)
}

func wkeJsBindFunction(name string, fn wkeJsNativeFunction, param unsafe.Pointer, argCount uint32) {
	ptr := []byte(name)
	r, _, err := _wkeJsBindFunction.Call(uintptr(unsafe.Pointer(&ptr[0])), syscall.NewCallbackCDecl(fn), uintptr(param), uintptr(argCount))
	if r == 0 && showError {
		fmt.Println("wkeJsBindFunction", err)
	}
}

func wkeNetCancelRequest(job wkeNetJob) {
	r, _, err := _wkeNetCancelRequest.Call(uintptr(job))
	if r == 0 && showError {
		fmt.Println("wkeNetCancelRequest", err)
	}
}

func wkeNetOnResponse(wke wkeHandle, callback wkeNetResponseCallback, param unsafe.Pointer) {
	r, _, err := _wkeNetOnResponse.Call(uintptr(wke), syscall.NewCallbackCDecl(callback), uintptr(param))
	if r == 0 && showError {
		fmt.Println("wkeNetOnResponse", err)
	}
}

func wkeOnLoadUrlBegin(wke wkeHandle, callback wkeLoadUrlBeginCallback, param unsafe.Pointer) {
	r, _, err := _wkeOnLoadUrlBegin.Call(uintptr(wke), syscall.NewCallbackCDecl(callback), uintptr(param))
	if r == 0 && showError {
		fmt.Println("wkeOnLoadUrlBegin", err)
	}
}

func wkeNetGetRequestMethod(job wkeNetJob) wkeRequestType {
	r, _, err := _wkeNetGetRequestMethod.Call(uintptr(job))
	if r == 0 && showError {
		fmt.Println("wkeNetGetRequestMethod", err)
	}
	return wkeRequestType(r)
}

func wkeNetSetData(job wkeNetJob, buf *byte, len uint32) {
	r, _, err := _wkeNetSetData.Call(uintptr(job), uintptr(unsafe.Pointer(buf)), uintptr(len))
	if r == 0 && showError {
		fmt.Println("wkeNetSetData", err)
	}
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

func wkePaint(wke wkeHandle, bits *uint8, pitch uint32) {
	r, _, err := _wkePaint.Call(uintptr(wke), uintptr(unsafe.Pointer(bits)), uintptr(pitch))
	if r == 0 && showError {
		fmt.Println("wkePaint", err)
	}
}

func wkeGetHeight(wke wkeHandle) uint32 {
	r, _, err := _wkeGetHeight.Call(uintptr(wke))
	if r == 0 && showError {
		fmt.Println("wkeGetHeight", err)
	}
	return uint32(r)
}

func wkeGetWidth(wke wkeHandle) uint32 {
	r, _, err := _wkeGetWidth.Call(uintptr(wke))
	if r == 0 && showError {
		fmt.Println("wkeGetWidth", err)
	}
	return uint32(r)
}

func wkeResize(wke wkeHandle, w, h uint32) {
	r, _, err := _wkeResize.Call(uintptr(wke), uintptr(w), uintptr(h))
	if r == 0 && showError {
		fmt.Println("wkeResize", err)
	}
}

func wkeLoadURL(wke wkeHandle, url string) {
	ptr := []byte(url)
	r, _, err := _wkeLoadURL.Call(uintptr(wke), uintptr(unsafe.Pointer(&ptr[0])))
	if r == 0 && showError {
		fmt.Println("wkeLoadURL", err)
	}
}

func wkeOnPaintBitUpdated(wke wkeHandle, callback wkePaintBitUpdatedCallback, param unsafe.Pointer) {
	r, _, err := _wkeOnPaintBitUpdated.Call(uintptr(wke), syscall.NewCallbackCDecl(callback), uintptr(param))
	if r == 0 && showError {
		fmt.Println("wkeOnPaintBitUpdated", err)
	}
}

func wkeSetHandle(wke wkeHandle, handle uintptr) {
	r, _, err := _wkeSetHandle.Call(uintptr(wke), handle)
	if r == 0 && showError {
		fmt.Println("wkeSetHandle", err)
	}
}

func wkeCreateWebView() wkeHandle {
	r, _, err := _wkeCreateWebView.Call()
	if r == 0 && showError {
		fmt.Println("wkeCreateWebView", err)
	}
	return wkeHandle(r)
}

func toBool(b bool) byte {
	if b {
		return 1
	} else {
		return 0
	}
}
