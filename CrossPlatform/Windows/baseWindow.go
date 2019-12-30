package Windows

import (
	win322 "GoMiniblink/CrossPlatform/Windows/win32"
)

type baseWindow interface {
	class() string
	name() string
	hWnd() win322.HWND
	fireWndCreate(hWnd win322.HWND)
	fireWndProc(hWnd win322.HWND, msg uint32, wParam, lParam uintptr) uintptr
}
