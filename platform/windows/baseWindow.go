package windows

import (
	"qq.2564874169/miniblink/platform/windows/win32"
)

type baseWindow interface {
	id() string
	class() string
	hWnd() win32.HWND
	isDialog() bool
	fireWndCreate(hWnd win32.HWND)
	fireWndProc(hWnd win32.HWND, msg uint32, wParam, lParam uintptr) uintptr
}
