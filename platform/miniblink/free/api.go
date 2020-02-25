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

	showError             bool
	_wkeIsInitialize      *windows.LazyProc
	_wkeInitialize        *windows.LazyProc
	_wkeCreateWebView     *windows.LazyProc
	_wkeSetHandle         *windows.LazyProc
	_wkeOnPaintBitUpdated *windows.LazyProc
	_wkeLoadURL           *windows.LazyProc
	_wkeResize            *windows.LazyProc
	_wkeNetOnResponse     *windows.LazyProc
	_wkeOnLoadUrlBegin    *windows.LazyProc
	_wkePaint             *windows.LazyProc
	_wkeGetWidth          *windows.LazyProc
	_wkeGetHeight         *windows.LazyProc
	_wkePaint2            *windows.LazyProc
)

func init() {
	showError = true
	is64 := unsafe.Sizeof(uintptr(0)) == 8
	if is64 {
		lib = windows.NewLazyDLL(file_x64_dll)
	} else {
		lib = windows.NewLazyDLL(file_x86_dll)
	}
	_wkePaint2 = lib.NewProc("wkePaint2")
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
	_wkeIsInitialize = lib.NewProc("wkeIsInitialize")

	ret, _, err := _wkeInitialize.Call()
	if ret == 0 && showError {
		fmt.Println(err)
	}
}

func wkePaint2(wke wkeHandle, buf []uint8, bufWdith, bufHeight, xDst, yDst, w, h, xSrc, ySrc uint32, copyAlpha bool) {
	var b byte = 0
	if copyAlpha {
		b = 1
	}
	r, _, err := _wkePaint2.Call(
		uintptr(wke),
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(bufWdith),
		uintptr(bufHeight),
		uintptr(xDst),
		uintptr(yDst),
		uintptr(w),
		uintptr(h),
		uintptr(xSrc),
		uintptr(ySrc),
		uintptr(b))
	if r == 0 && showError {
		fmt.Println("wkePaint2", err)
	}
}

func wkePaint(wke wkeHandle, bits []uint8, pitch uint32) {
	r, _, err := _wkePaint.Call(uintptr(wke), uintptr(unsafe.Pointer(&bits[0])), uintptr(pitch))
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

func wkeOnLoadUrlBegin(wke wkeHandle, callback wkeLoadUrlBeginCallback, param uintptr) {
	r, _, err := _wkeOnLoadUrlBegin.Call(uintptr(wke), syscall.NewCallback(callback), param)
	if r == 0 && showError {
		fmt.Println("wkeOnLoadUrlBegin", err)
	}
}

func wkeNetOnResponse(wke wkeHandle, callback wkeNetResponseCallback, param uintptr) {
	r, _, err := _wkeNetOnResponse.Call(uintptr(wke), syscall.NewCallback(callback), param)
	if r == 0 && showError {
		fmt.Println("wkeNetOnResponse", err)
	}
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
	r, _, err := _wkeOnPaintBitUpdated.Call(uintptr(wke), syscall.NewCallback(callback), param)
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

func wkeInitialize() bool {
	r, _, err := _wkeInitialize.Call()
	if r == 0 && showError {
		fmt.Println("wkeInitialize", err)
	}
	return r != 0
}

func wkeIsInitialize() bool {
	r, _, err := _wkeIsInitialize.Call()
	if r == 0 && showError {
		fmt.Println("wkeIsInitialize", err)
	}
	return r != 0
}
