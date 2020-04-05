package win32

import (
	"fmt"
	"golang.org/x/sys/windows"
	"unsafe"
)

type (
	HIMC HWND
)

type COMPOSITIONFORM struct {
	DwStyle uint32
	Pos     POINT
	Rect    RECT
}

const (
	CFS_POINT          = 2
	CFS_FORCE_POSITION = 32
)

var (
	lib *windows.LazyDLL

	_ImmGetContext           *windows.LazyProc
	_ImmSetCompositionWindow *windows.LazyProc
	_ImmReleaseContext       *windows.LazyProc
)

func init() {
	lib = windows.NewLazyDLL("imm32.dll")
	_ImmGetContext = lib.NewProc("ImmGetContext")
	_ImmSetCompositionWindow = lib.NewProc("ImmSetCompositionWindow")
	_ImmReleaseContext = lib.NewProc("ImmReleaseContext")
}

func ImmReleaseContext(hWnd HWND, himc HIMC) bool {
	ret, _, _ := _ImmReleaseContext.Call(uintptr(hWnd), uintptr(himc))
	return ret != 0
}

func ImmSetCompositionWindow(himc HIMC, comp *COMPOSITIONFORM) bool {
	ret, _, _ := _ImmSetCompositionWindow.Call(uintptr(himc), uintptr(unsafe.Pointer(comp)))
	return ret != 0
}

func ImmGetContext(hWnd HWND) HIMC {
	ret, _, err := _ImmGetContext.Call(uintptr(hWnd))
	if ret == 0 {
		fmt.Println("ImmGetContext", err)
	}
	return HIMC(ret)
}
