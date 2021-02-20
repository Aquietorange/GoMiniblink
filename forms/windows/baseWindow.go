package windows

import (
	"github.com/hujun528/GoMiniblink/forms/windows/win32"
)

type baseWindow interface {
	hWnd() win32.HWND
	onWndMsg(hWnd win32.HWND, msg uint32, wParam uintptr, lParam uintptr) uintptr
}
