package windows

import (
	"qq.2564874169/goMiniblink/platform/windows/win32"
)

type baseWindow interface {
	id() string
	hWnd() win32.HWND
	isDialog() bool
	getCreateProc() windowsCreateProc
	getWindowMsgProc() windowsMsgProc
}
