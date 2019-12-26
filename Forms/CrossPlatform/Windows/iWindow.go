package Windows

import "GoMiniblink/Forms/CrossPlatform/Windows/win32"

type IWindow interface {
	class() string
	name() string
	onWndCreate(hWnd win32.HWND)
	hWnd() win32.HWND
	fireWndProc(hWnd win32.HWND, msg uint32, wParam, lParam uintptr) uintptr
}
