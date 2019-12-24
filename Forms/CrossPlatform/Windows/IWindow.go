package Windows

import "GoMiniblink/Forms/CrossPlatform/Windows/win32"

type IWindow interface {
	class() string
	name() string
	hWnd() win32.HWND
	wndProc(hWnd win32.HWND, msg uint32, wParam, lParam uintptr) uintptr
}
