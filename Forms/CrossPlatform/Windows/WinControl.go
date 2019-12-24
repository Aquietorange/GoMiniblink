package Windows

import (
	"GoMiniblink/Forms/CrossPlatform/Windows/win32"
)

type winControl struct {
	provider  *Provider
	className string
	idName    string
	hwnd      win32.HWND
	onWndProc func(hWnd win32.HWND, msg uint32, wParam uintptr, lParam uintptr) uintptr
	owner     *winForm
	x         int
	y         int
	width     int
	height    int
}
