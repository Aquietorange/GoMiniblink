package Windows

import (
	win322 "GoMiniblink/CrossPlatform/Windows/win32"
)

type baseWindow interface {
	id() string
	class() string
	hWnd() win322.HWND
	isDialog() bool
	fireWndCreate(hWnd win322.HWND)
	fireWndProc(hWnd win322.HWND, msg uint32, wParam, lParam uintptr) uintptr
}
