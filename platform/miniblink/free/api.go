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
	wkeHandle uintptr
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
)

func init() {
	showError = true
	is64 := unsafe.Sizeof(uintptr(0)) == 8
	if is64 {
		lib = windows.NewLazyDLL(file_x64_dll)
	} else {
		lib = windows.NewLazyDLL(file_x86_dll)
	}
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

func wkeNetCancelRequest(job wkeNetJob) {
	r, _, err := _wkeNetCancelRequest.Call(uintptr(job))
	if r == 0 && showError {
		fmt.Println("wkeNetCancelRequest", err)
	}
}

func wkeNetOnResponse(wke wkeHandle, callback wkeNetResponseCallback, param uintptr) {
	r, _, err := _wkeNetOnResponse.Call(uintptr(wke), syscall.NewCallbackCDecl(callback), param)
	if r == 0 && showError {
		fmt.Println("wkeNetOnResponse", err)
	}
}

func wkeOnLoadUrlBegin(wke wkeHandle, callback wkeLoadUrlBeginCallback, param uintptr) {
	r, _, err := _wkeOnLoadUrlBegin.Call(uintptr(wke), syscall.NewCallbackCDecl(callback), param)
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

func wkeFireKeyDownEvent(wke wkeHandle, code uintptr, flags uint32, isSysKey bool) bool {
	ret, _, _ := _wkeFireKeyDownEvent.Call(
		uintptr(wke),
		code,
		uintptr(flags),
		uintptr(toBool(isSysKey)))
	return byte(ret) != 0
}

func wkeFireKeyUpEvent(wke wkeHandle, code uintptr, flags uint32, isSysKey bool) bool {
	ret, _, _ := _wkeFireKeyUpEvent.Call(
		uintptr(wke),
		code,
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

func wkeOnPaintBitUpdated(wke wkeHandle, callback wkePaintBitUpdatedCallback, param uintptr) {
	r, _, err := _wkeOnPaintBitUpdated.Call(uintptr(wke), syscall.NewCallbackCDecl(callback), param)
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
