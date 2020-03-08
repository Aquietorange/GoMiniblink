package windows

import (
	"qq2564874169/goMiniblink/platform/windows/win32"
)

type baseWindow interface {
	hWnd() win32.HWND
	wndMsgProc(hWnd win32.HWND, msg uint32, wParam uintptr, lParam uintptr) uintptr
}
