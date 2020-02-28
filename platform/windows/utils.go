package windows

import (
	mb "qq.2564874169/goMiniblink"
	"qq.2564874169/goMiniblink/platform/windows/win32"
	"syscall"
)

func isExtendKey(key mb.Keys) bool {
	switch key {
	case mb.Keys_Insert, mb.Keys_Delete, mb.Keys_Home, mb.Keys_End, mb.Keys_PageUp,
		mb.Keys_PageDown, mb.Keys_Left, mb.Keys_Right, mb.Keys_Up, mb.Keys_Down:
		return true
	default:
		return false
	}
}

func toWinCursor(cursor mb.CursorType) int {
	switch cursor {
	case mb.CursorType_ARROW:
		return win32.IDC_ARROW
	case mb.CursorType_SIZE:
		return win32.IDC_SIZE
	case mb.CursorType_ICON:
		return win32.IDC_ICON
	case mb.CursorType_HELP:
		return win32.IDC_HELP
	case mb.CursorType_APPSTARTING:
		return win32.IDC_APPSTARTING
	case mb.CursorType_HAND:
		return win32.IDC_HAND
	case mb.CursorType_NO:
		return win32.IDC_NO
	case mb.CursorType_SIZEALL:
		return win32.IDC_SIZEALL
	case mb.CursorType_SIZENS:
		return win32.IDC_SIZENS
	case mb.CursorType_SIZEWE:
		return win32.IDC_SIZEWE
	case mb.CursorType_SIZENWSE:
		return win32.IDC_SIZENWSE
	case mb.CursorType_SIZENESW:
		return win32.IDC_SIZENESW
	case mb.CursorType_UPARROW:
		return win32.IDC_UPARROW
	case mb.CursorType_CROSS:
		return win32.IDC_CROSS
	case mb.CursorType_WAIT:
		return win32.IDC_WAIT
	case mb.CursorType_IBEAM:
		return win32.IDC_IBEAM
	default:
		return win32.IDC_ARROW
	}
}

func winCursorTo(cursor int) mb.CursorType {
	switch cursor {
	case 32512:
		return mb.CursorType_ARROW
	case 32513:
		return mb.CursorType_IBEAM
	case 32514:
		return mb.CursorType_WAIT
	case 32515:
		return mb.CursorType_CROSS
	case 32516:
		return mb.CursorType_UPARROW
	case 32642:
		return mb.CursorType_SIZENWSE
	case 32643:
		return mb.CursorType_SIZENESW
	case 32644:
		return mb.CursorType_SIZEWE
	case 32645:
		return mb.CursorType_SIZENS
	case 32646:
		return mb.CursorType_SIZEALL
	case 32648:
		return mb.CursorType_NO
	case 32649:
		return mb.CursorType_HAND
	case 32650:
		return mb.CursorType_APPSTARTING
	case 32651:
		return mb.CursorType_HELP
	case 32641:
		return mb.CursorType_ICON
	case 32640:
		return mb.CursorType_SIZE
	default:
		return mb.CursorType_ARROW
	}
}

func vkToKey(vk int) mb.Keys {
	switch vk {
	case win32.VK_ESCAPE:
		return mb.Keys_Esc
	case win32.VK_F1:
		return mb.Keys_F1
	case win32.VK_F2:
		return mb.Keys_F2
	case win32.VK_F3:
		return mb.Keys_F3
	case win32.VK_F4:
		return mb.Keys_F4
	case win32.VK_F5:
		return mb.Keys_F5
	case win32.VK_F6:
		return mb.Keys_F6
	case win32.VK_F7:
		return mb.Keys_F7
	case win32.VK_F8:
		return mb.Keys_F8
	case win32.VK_F9:
		return mb.Keys_F9
	case win32.VK_F10:
		return mb.Keys_F10
	case win32.VK_F11:
		return mb.Keys_F11
	case win32.VK_F12:
		return mb.Keys_F12
	case win32.VK_SPACE:
		return mb.Keys_Space
	case win32.VK_TAB:
		return mb.Keys_Tab
	case win32.VK_CAPITAL:
		return mb.Keys_CapsLock
	case win32.VK_SHIFT:
		return mb.Keys_Shift
	case win32.VK_CONTROL:
		return mb.Keys_Ctrl
	case win32.VK_BACK:
		return mb.Keys_Back
	case win32.VK_MENU:
		return mb.Keys_Alt
	case win32.VK_LWIN:
		return mb.Keys_Win
	case win32.VK_RSHIFT:
		return mb.Keys_Right_Shift
	case win32.VK_RCONTROL:
		return mb.Keys_Right_Ctrl
	case win32.VK_RMENU:
		return mb.Keys_Right_Alt
	case win32.VK_RWIN:
		return mb.Keys_Right_Win
	case win32.VK_RETURN:
		return mb.Keys_Enter
	case win32.VK_APPS:
		return mb.Keys_Apps
	case 0x30:
		return mb.Keys_0
	case 0x31:
		return mb.Keys_1
	case 0x32:
		return mb.Keys_2
	case 0x33:
		return mb.Keys_3
	case 0x34:
		return mb.Keys_4
	case 0x35:
		return mb.Keys_5
	case 0x36:
		return mb.Keys_6
	case 0x37:
		return mb.Keys_7
	case 0x38:
		return mb.Keys_8
	case 0x39:
		return mb.Keys_9
	case win32.VK_OEM_3:
		return mb.Keys_OEM_3
	case win32.VK_OEM_PLUS:
		return mb.Keys_OEM_PLUS
	case win32.VK_OEM_MINUS:
		return mb.Keys_OEM_MINUS
	case win32.VK_SNAPSHOT:
		return mb.Keys_Snapshot
	case win32.VK_SCROLL:
		return mb.Keys_Scroll_Lock
	case win32.VK_PAUSE:
		return mb.Keys_Pause
	case 0x51:
		return mb.Keys_Q
	case 0x57:
		return mb.Keys_W
	case 0x45:
		return mb.Keys_E
	case 0x52:
		return mb.Keys_R
	case 0x54:
		return mb.Keys_T
	case 0x59:
		return mb.Keys_Y
	case 0x55:
		return mb.Keys_U
	case 0x49:
		return mb.Keys_I
	case 0x4F:
		return mb.Keys_O
	case 0x50:
		return mb.Keys_P
	case 0x41:
		return mb.Keys_A
	case 0x53:
		return mb.Keys_S
	case 0x44:
		return mb.Keys_D
	case 0x46:
		return mb.Keys_F
	case 0x47:
		return mb.Keys_G
	case 0x48:
		return mb.Keys_H
	case 0x4A:
		return mb.Keys_J
	case 0x4B:
		return mb.Keys_K
	case 0x4C:
		return mb.Keys_L
	case 0x5A:
		return mb.Keys_Z
	case 0x58:
		return mb.Keys_X
	case 0x43:
		return mb.Keys_C
	case 0x56:
		return mb.Keys_V
	case 0x42:
		return mb.Keys_B
	case 0x4E:
		return mb.Keys_N
	case 0x4D:
		return mb.Keys_M
	case win32.VK_OEM_4:
		return mb.Keys_OEM_4
	case win32.VK_OEM_6:
		return mb.Keys_OEM_6
	case win32.VK_OEM_1:
		return mb.Keys_OEM_1
	case win32.VK_OEM_7:
		return mb.Keys_OEM_7
	case win32.VK_OEM_5:
		return mb.Keys_OEM_5
	case win32.VK_OEM_COMMA:
		return mb.Keys_OEM_COMMA
	case win32.VK_OEM_PERIOD:
		return mb.Keys_OEM_PERIOD
	case win32.VK_OEM_2:
		return mb.Keys_OEM_2
	case win32.VK_INSERT:
		return mb.Keys_Insert
	case win32.VK_DELETE:
		return mb.Keys_Delete
	case win32.VK_HOME:
		return mb.Keys_Home
	case win32.VK_END:
		return mb.Keys_End
	case win32.VK_PRIOR:
		return mb.Keys_PageUp
	case win32.VK_NEXT:
		return mb.Keys_PageDown
	case win32.VK_LEFT:
		return mb.Keys_Left
	case win32.VK_UP:
		return mb.Keys_Up
	case win32.VK_RIGHT:
		return mb.Keys_Right
	case win32.VK_DOWN:
		return mb.Keys_Down
	case win32.VK_NUMLOCK:
		return mb.Keys_Num_Lock
	case win32.VK_ADD:
		return mb.Keys_Num_Add
	case win32.VK_SUBTRACT:
		return mb.Keys_Num_Sub
	case win32.VK_MULTIPLY:
		return mb.Keys_Num_Multiply
	case win32.VK_DIVIDE:
		return mb.Keys_Num_Divide
	case win32.VK_DECIMAL:
		return mb.Keys_Num_Decimal
	case win32.VK_SEPARATOR:
		return mb.Keys_Num_Separator
	case win32.VK_NUMPAD0:
		return mb.Keys_Num_0
	case win32.VK_NUMPAD1:
		return mb.Keys_Num_1
	case win32.VK_NUMPAD2:
		return mb.Keys_Num_2
	case win32.VK_NUMPAD3:
		return mb.Keys_Num_3
	case win32.VK_NUMPAD4:
		return mb.Keys_Num_4
	case win32.VK_NUMPAD5:
		return mb.Keys_Num_5
	case win32.VK_NUMPAD6:
		return mb.Keys_Num_6
	case win32.VK_NUMPAD7:
		return mb.Keys_Num_7
	case win32.VK_NUMPAD8:
		return mb.Keys_Num_8
	case win32.VK_NUMPAD9:
		return mb.Keys_Num_9
	default:
		return mb.Keys_Error
	}
}

func sto16(str string) *uint16 {
	ptr, _ := syscall.UTF16PtrFromString(str)
	return ptr
}
