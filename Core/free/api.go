package free

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

const (
	free_x86_dll = "miniblink_x86.dll"
	free_x64_dll = "miniblink_x64.dll"
)

type (
	WKE uintptr
)

var (
	lib *windows.LazyDLL

	_wkeIsInitialize      *windows.LazyProc
	_wkeCreateWebView     *windows.LazyProc
	_wkeSetHandle         *windows.LazyProc
	_wkeOnPaintBitUpdated *windows.LazyProc
	_wkeLoadURL           *windows.LazyProc
)

func init() {
	is64 := unsafe.Sizeof(uintptr(0)) == 8
	if is64 {
		lib = windows.NewLazyDLL(free_x64_dll)
	} else {
		lib = windows.NewLazyDLL(free_x86_dll)
	}
	_wkeIsInitialize = lib.NewProc("wkeIsInitialize")
	_wkeCreateWebView = lib.NewProc("wkeCreateWebView")
	_wkeSetHandle = lib.NewProc("wkeSetHandle")
	_wkeOnPaintBitUpdated = lib.NewProc("wkeOnPaintBitUpdated")
	_wkeLoadURL = lib.NewProc("wkeLoadURL")
}

func wkeLoadURL(wke WKE, url string) {
	r, _, err := _wkeLoadURL.Call(uintptr(wke), uintptr(unsafe.Pointer(&url)))
	if r == 0 {
		println("wkeLoadURL", err)
	}
}

func wkeOnPaintBitUpdated(wke WKE, callback wkePaintBitUpdatedCallback, param uintptr) {
	r, _, err := _wkeOnPaintBitUpdated.Call(uintptr(wke), syscall.NewCallback(callback), param)
	if r == 0 {
		println("wkeOnPaintBitUpdated", err)
	}
}

func wkeSetHandle(wke WKE, handle uintptr) {
	r, _, err := _wkeSetHandle.Call(uintptr(wke), handle)
	if r == 0 {
		println("wkeSetHandle", err)
	}
}

func wkeCreateWebView() WKE {
	r, _, err := _wkeCreateWebView.Call()
	if r == 0 {
		println("wkeCreateWebView", err)
	} else {
		println("success")
	}
	return WKE(r)
}

func wkeIsInitialize() bool {
	r, _, err := _wkeIsInitialize.Call()
	if r == 0 {
		println("wkeIsInitialize", err)
	}
	return r != 0
}
