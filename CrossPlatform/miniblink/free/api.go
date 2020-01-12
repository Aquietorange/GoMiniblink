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

	_wkeInitialize        *windows.LazyProc
	_wkeCreateWebView     *windows.LazyProc
	_wkeSetHandle         *windows.LazyProc
	_wkeOnPaintBitUpdated *windows.LazyProc
	_wkeLoadURL           *windows.LazyProc
	_wkeCreateWebWindow   *windows.LazyProc
	_wkeResize            *windows.LazyProc
)

func init() {
	is64 := unsafe.Sizeof(uintptr(0)) == 8
	if is64 {
		lib = windows.NewLazyDLL(file_x64_dll)
	} else {
		lib = windows.NewLazyDLL(file_x86_dll)
	}
	_wkeInitialize = lib.NewProc("wkeInitialize")
	_wkeCreateWebView = lib.NewProc("wkeCreateWebView")
	_wkeSetHandle = lib.NewProc("wkeSetHandle")
	_wkeOnPaintBitUpdated = lib.NewProc("wkeOnPaintBitUpdated")
	_wkeLoadURL = lib.NewProc("wkeLoadURLW")
	_wkeCreateWebWindow = lib.NewProc("wkeCreateWebWindow")
	_wkeResize = lib.NewProc("wkeResize")
}

func wkeResize(wke wkeHandle, w, h int32) {
	r, _, err := _wkeResize.Call(uintptr(wke), uintptr(w), uintptr(h))
	if r == 0 {
		fmt.Println("wkeResize", err)
	}
}

func wkeCreateWebWindow() wkeHandle {
	r, _, err := _wkeCreateWebWindow.Call(0, 0, 100, 100, 500, 500)
	if r == 0 {
		fmt.Println("wkeCreateWebWindow", err)
	}
	return wkeHandle(r)
}

func wkeLoadURL(wke wkeHandle, url string) {
	ptr, _ := syscall.UTF16PtrFromString(url)
	//ptr := []rune(url)
	r, _, err := _wkeLoadURL.Call(uintptr(wke), uintptr(unsafe.Pointer(&ptr)))
	if r == 0 {
		fmt.Println("wkeLoadURL", err)
	}
}

func wkeOnPaintBitUpdated(wke wkeHandle, callback wkePaintBitUpdatedCallback, param uintptr) {
	r, _, err := _wkeOnPaintBitUpdated.Call(uintptr(wke), syscall.NewCallback(callback), param)
	if r == 0 {
		fmt.Println("wkeOnPaintBitUpdated", err)
	}
}

func wkeSetHandle(wke wkeHandle, handle uintptr) {
	r, _, err := _wkeSetHandle.Call(uintptr(wke), handle)
	if r == 0 {
		fmt.Println("wkeSetHandle", err)
	}
}

func wkeCreateWebView() wkeHandle {
	r, _, err := _wkeCreateWebView.Call()
	if r == 0 {
		fmt.Println("wkeCreateWebView", err)
	}
	return wkeHandle(r)
}

func wkeInitialize() bool {
	r, _, err := _wkeInitialize.Call()
	if r == 0 {
		fmt.Println("wkeInitialize", err)
	}
	return r != 0
}
