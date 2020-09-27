package windows

import (
	"fmt"
	"image/color"
	fm "qq2564874169/goMiniblink/forms"
	br "qq2564874169/goMiniblink/forms/bridge"
	win "qq2564874169/goMiniblink/forms/windows/win32"
	"syscall"
)

func findChild(container br.Controls, hWnd win.HWND) br.Window {
	if container.GetHandle() == uintptr(hWnd) {
		return container
	}
	for _, i := range container.GetChilds() {
		if c, ok := i.(br.Controls); ok {
			w := findChild(c, hWnd)
			if w != nil {
				return w
			}
		} else if i.GetHandle() == uintptr(hWnd) {
			return i
		}
	}
	return nil
}

func intToRGBA(rgba int32) color.RGBA {
	return color.RGBA{
		R: uint8(rgba),
		G: uint8(rgba >> 8),
		B: uint8(rgba >> 16),
		A: uint8(rgba >> 24),
	}
}

func isExtendKey(key fm.Keys) bool {
	switch key {
	case fm.Keys_Insert, fm.Keys_Delete, fm.Keys_Home, fm.Keys_End, fm.Keys_PageUp,
		fm.Keys_PageDown, fm.Keys_Left, fm.Keys_Right, fm.Keys_Up, fm.Keys_Down:
		return true
	default:
		return false
	}
}

func toWinCursor(cursor fm.CursorType) int {
	switch cursor {
	case fm.CursorType_ARROW:
		return win.IDC_ARROW
	case fm.CursorType_SIZE:
		return win.IDC_SIZE
	case fm.CursorType_ICON:
		return win.IDC_ICON
	case fm.CursorType_HELP:
		return win.IDC_HELP
	case fm.CursorType_APPSTARTING:
		return win.IDC_APPSTARTING
	case fm.CursorType_HAND:
		return win.IDC_HAND
	case fm.CursorType_NO:
		return win.IDC_NO
	case fm.CursorType_SIZEALL:
		return win.IDC_SIZEALL
	case fm.CursorType_SIZENS:
		return win.IDC_SIZENS
	case fm.CursorType_SIZEWE:
		return win.IDC_SIZEWE
	case fm.CursorType_SIZENWSE:
		return win.IDC_SIZENWSE
	case fm.CursorType_SIZENESW:
		return win.IDC_SIZENESW
	case fm.CursorType_UPARROW:
		return win.IDC_UPARROW
	case fm.CursorType_CROSS:
		return win.IDC_CROSS
	case fm.CursorType_WAIT:
		return win.IDC_WAIT
	case fm.CursorType_IBEAM:
		return win.IDC_IBEAM
	default:
		return win.IDC_ARROW
	}
}

func winCursorTo(cursor int) fm.CursorType {
	switch cursor {
	case win.IDC_ARROW:
		return fm.CursorType_ARROW
	case win.IDC_IBEAM:
		return fm.CursorType_IBEAM
	case win.IDC_WAIT:
		return fm.CursorType_WAIT
	case win.IDC_CROSS:
		return fm.CursorType_CROSS
	case win.IDC_UPARROW:
		return fm.CursorType_UPARROW
	case win.IDC_SIZENWSE:
		return fm.CursorType_SIZENWSE
	case win.IDC_SIZENESW:
		return fm.CursorType_SIZENESW
	case win.IDC_SIZEWE:
		return fm.CursorType_SIZEWE
	case win.IDC_SIZENS:
		return fm.CursorType_SIZENS
	case win.IDC_SIZEALL:
		return fm.CursorType_SIZEALL
	case win.IDC_NO:
		return fm.CursorType_NO
	case win.IDC_HAND:
		return fm.CursorType_HAND
	case win.IDC_APPSTARTING:
		return fm.CursorType_APPSTARTING
	case win.IDC_HELP:
		return fm.CursorType_HELP
	case win.IDC_ICON:
		return fm.CursorType_ICON
	case win.IDC_SIZE:
		return fm.CursorType_SIZE
	default:
		return fm.CursorType_ARROW
	}
}

func vkToKey(vk int) fm.Keys {
	switch vk {
	case win.VK_ESCAPE:
		return fm.Keys_Esc
	case win.VK_F1:
		return fm.Keys_F1
	case win.VK_F2:
		return fm.Keys_F2
	case win.VK_F3:
		return fm.Keys_F3
	case win.VK_F4:
		return fm.Keys_F4
	case win.VK_F5:
		return fm.Keys_F5
	case win.VK_F6:
		return fm.Keys_F6
	case win.VK_F7:
		return fm.Keys_F7
	case win.VK_F8:
		return fm.Keys_F8
	case win.VK_F9:
		return fm.Keys_F9
	case win.VK_F10:
		return fm.Keys_F10
	case win.VK_F11:
		return fm.Keys_F11
	case win.VK_F12:
		return fm.Keys_F12
	case win.VK_SPACE:
		return fm.Keys_Space
	case win.VK_TAB:
		return fm.Keys_Tab
	case win.VK_CAPITAL:
		return fm.Keys_CapsLock
	case win.VK_SHIFT:
		return fm.Keys_Shift
	case win.VK_CONTROL:
		return fm.Keys_Ctrl
	case win.VK_BACK:
		return fm.Keys_Back
	case win.VK_MENU:
		return fm.Keys_Alt
	case win.VK_LWIN:
		return fm.Keys_Win
	case win.VK_RSHIFT:
		return fm.Keys_Right_Shift
	case win.VK_RCONTROL:
		return fm.Keys_Right_Ctrl
	case win.VK_RMENU:
		return fm.Keys_Right_Alt
	case win.VK_RWIN:
		return fm.Keys_Right_Win
	case win.VK_RETURN:
		return fm.Keys_Enter
	case win.VK_APPS:
		return fm.Keys_Apps
	case 0x30:
		return fm.Keys_0
	case 0x31:
		return fm.Keys_1
	case 0x32:
		return fm.Keys_2
	case 0x33:
		return fm.Keys_3
	case 0x34:
		return fm.Keys_4
	case 0x35:
		return fm.Keys_5
	case 0x36:
		return fm.Keys_6
	case 0x37:
		return fm.Keys_7
	case 0x38:
		return fm.Keys_8
	case 0x39:
		return fm.Keys_9
	case win.VK_OEM_3:
		return fm.Keys_OEM_3
	case win.VK_OEM_PLUS:
		return fm.Keys_OEM_PLUS
	case win.VK_OEM_MINUS:
		return fm.Keys_OEM_MINUS
	case win.VK_SNAPSHOT:
		return fm.Keys_Snapshot
	case win.VK_SCROLL:
		return fm.Keys_Scroll_Lock
	case win.VK_PAUSE:
		return fm.Keys_Pause
	case 0x51:
		return fm.Keys_Q
	case 0x57:
		return fm.Keys_W
	case 0x45:
		return fm.Keys_E
	case 0x52:
		return fm.Keys_R
	case 0x54:
		return fm.Keys_T
	case 0x59:
		return fm.Keys_Y
	case 0x55:
		return fm.Keys_U
	case 0x49:
		return fm.Keys_I
	case 0x4F:
		return fm.Keys_O
	case 0x50:
		return fm.Keys_P
	case 0x41:
		return fm.Keys_A
	case 0x53:
		return fm.Keys_S
	case 0x44:
		return fm.Keys_D
	case 0x46:
		return fm.Keys_F
	case 0x47:
		return fm.Keys_G
	case 0x48:
		return fm.Keys_H
	case 0x4A:
		return fm.Keys_J
	case 0x4B:
		return fm.Keys_K
	case 0x4C:
		return fm.Keys_L
	case 0x5A:
		return fm.Keys_Z
	case 0x58:
		return fm.Keys_X
	case 0x43:
		return fm.Keys_C
	case 0x56:
		return fm.Keys_V
	case 0x42:
		return fm.Keys_B
	case 0x4E:
		return fm.Keys_N
	case 0x4D:
		return fm.Keys_M
	case win.VK_OEM_4:
		return fm.Keys_OEM_4
	case win.VK_OEM_6:
		return fm.Keys_OEM_6
	case win.VK_OEM_1:
		return fm.Keys_OEM_1
	case win.VK_OEM_7:
		return fm.Keys_OEM_7
	case win.VK_OEM_5:
		return fm.Keys_OEM_5
	case win.VK_OEM_COMMA:
		return fm.Keys_OEM_COMMA
	case win.VK_OEM_PERIOD:
		return fm.Keys_OEM_PERIOD
	case win.VK_OEM_2:
		return fm.Keys_OEM_2
	case win.VK_INSERT:
		return fm.Keys_Insert
	case win.VK_DELETE:
		return fm.Keys_Delete
	case win.VK_HOME:
		return fm.Keys_Home
	case win.VK_END:
		return fm.Keys_End
	case win.VK_PRIOR:
		return fm.Keys_PageUp
	case win.VK_NEXT:
		return fm.Keys_PageDown
	case win.VK_LEFT:
		return fm.Keys_Left
	case win.VK_UP:
		return fm.Keys_Up
	case win.VK_RIGHT:
		return fm.Keys_Right
	case win.VK_DOWN:
		return fm.Keys_Down
	case win.VK_NUMLOCK:
		return fm.Keys_Num_Lock
	case win.VK_ADD:
		return fm.Keys_Num_Add
	case win.VK_SUBTRACT:
		return fm.Keys_Num_Sub
	case win.VK_MULTIPLY:
		return fm.Keys_Num_Multiply
	case win.VK_DIVIDE:
		return fm.Keys_Num_Divide
	case win.VK_DECIMAL:
		return fm.Keys_Num_Decimal
	case win.VK_SEPARATOR:
		return fm.Keys_Num_Separator
	case win.VK_NUMPAD0:
		return fm.Keys_Num_0
	case win.VK_NUMPAD1:
		return fm.Keys_Num_1
	case win.VK_NUMPAD2:
		return fm.Keys_Num_2
	case win.VK_NUMPAD3:
		return fm.Keys_Num_3
	case win.VK_NUMPAD4:
		return fm.Keys_Num_4
	case win.VK_NUMPAD5:
		return fm.Keys_Num_5
	case win.VK_NUMPAD6:
		return fm.Keys_Num_6
	case win.VK_NUMPAD7:
		return fm.Keys_Num_7
	case win.VK_NUMPAD8:
		return fm.Keys_Num_8
	case win.VK_NUMPAD9:
		return fm.Keys_Num_9
	default:
		return fm.Keys_Error
	}
}

func keyToVk(key fm.Keys) int {
	switch key {
	case fm.Keys_Esc:
		return win.VK_ESCAPE
	case fm.Keys_F1:
		return win.VK_F1
	case fm.Keys_F2:
		return win.VK_F2
	case fm.Keys_F3:
		return win.VK_F3
	case fm.Keys_F4:
		return win.VK_F4
	case fm.Keys_F5:
		return win.VK_F5
	case fm.Keys_F6:
		return win.VK_F6
	case fm.Keys_F7:
		return win.VK_F7
	case fm.Keys_F8:
		return win.VK_F8
	case fm.Keys_F9:
		return win.VK_F9
	case fm.Keys_F10:
		return win.VK_F10
	case fm.Keys_F11:
		return win.VK_F11
	case fm.Keys_F12:
		return win.VK_F12
	case fm.Keys_Space:
		return win.VK_SPACE
	case fm.Keys_Tab:
		return win.VK_TAB
	case fm.Keys_CapsLock:
		return win.VK_CAPITAL
	case fm.Keys_Shift:
		return win.VK_SHIFT
	case fm.Keys_Ctrl:
		return win.VK_CONTROL
	case fm.Keys_Back:
		return win.VK_BACK
	case fm.Keys_Alt:
		return win.VK_MENU
	case fm.Keys_Win:
		return win.VK_LWIN
	case fm.Keys_Right_Shift:
		return win.VK_RSHIFT
	case fm.Keys_Right_Ctrl:
		return win.VK_RCONTROL
	case fm.Keys_Right_Alt:
		return win.VK_RMENU
	case fm.Keys_Right_Win:
		return win.VK_RWIN
	case fm.Keys_Enter:
		return win.VK_RETURN
	case fm.Keys_Apps:
		return win.VK_APPS
	case fm.Keys_0:
		return 0x30
	case fm.Keys_1:
		return 0x31
	case fm.Keys_2:
		return 0x32
	case fm.Keys_3:
		return 0x33
	case fm.Keys_4:
		return 0x34
	case fm.Keys_5:
		return 0x35
	case fm.Keys_6:
		return 0x36
	case fm.Keys_7:
		return 0x37
	case fm.Keys_8:
		return 0x38
	case fm.Keys_9:
		return 0x39
	case fm.Keys_OEM_3:
		return win.VK_OEM_3
	case fm.Keys_OEM_PLUS:
		return win.VK_OEM_PLUS
	case fm.Keys_OEM_MINUS:
		return win.VK_OEM_MINUS
	case fm.Keys_Snapshot:
		return win.VK_SNAPSHOT
	case fm.Keys_Scroll_Lock:
		return win.VK_SCROLL
	case fm.Keys_Pause:
		return win.VK_PAUSE
	case fm.Keys_Q:
		return 0x51
	case fm.Keys_W:
		return 0x57
	case fm.Keys_E:
		return 0x45
	case fm.Keys_R:
		return 0x52
	case fm.Keys_T:
		return 0x54
	case fm.Keys_Y:
		return 0x59
	case fm.Keys_U:
		return 0x55
	case fm.Keys_I:
		return 0x49
	case fm.Keys_O:
		return 0x4F
	case fm.Keys_P:
		return 0x50
	case fm.Keys_A:
		return 0x41
	case fm.Keys_S:
		return 0x53
	case fm.Keys_D:
		return 0x44
	case fm.Keys_F:
		return 0x46
	case fm.Keys_G:
		return 0x47
	case fm.Keys_H:
		return 0x48
	case fm.Keys_J:
		return 0x4A
	case fm.Keys_K:
		return 0x4B
	case fm.Keys_L:
		return 0x4C
	case fm.Keys_Z:
		return 0x5A
	case fm.Keys_X:
		return 0x58
	case fm.Keys_C:
		return 0x43
	case fm.Keys_V:
		return 0x56
	case fm.Keys_B:
		return 0x42
	case fm.Keys_N:
		return 0x4E
	case fm.Keys_M:
		return 0x4D
	case fm.Keys_OEM_4:
		return win.VK_OEM_4
	case fm.Keys_OEM_6:
		return win.VK_OEM_6
	case fm.Keys_OEM_1:
		return win.VK_OEM_1
	case fm.Keys_OEM_7:
		return win.VK_OEM_7
	case fm.Keys_OEM_5:
		return win.VK_OEM_5
	case fm.Keys_OEM_COMMA:
		return win.VK_OEM_COMMA
	case fm.Keys_OEM_PERIOD:
		return win.VK_OEM_PERIOD
	case fm.Keys_OEM_2:
		return win.VK_OEM_2
	case fm.Keys_Insert:
		return win.VK_INSERT
	case fm.Keys_Delete:
		return win.VK_DELETE
	case fm.Keys_Home:
		return win.VK_HOME
	case fm.Keys_End:
		return win.VK_END
	case fm.Keys_PageUp:
		return win.VK_PRIOR
	case fm.Keys_PageDown:
		return win.VK_NEXT
	case fm.Keys_Left:
		return win.VK_LEFT
	case fm.Keys_Up:
		return win.VK_UP
	case fm.Keys_Right:
		return win.VK_RIGHT
	case fm.Keys_Down:
		return win.VK_DOWN
	case fm.Keys_Num_Lock:
		return win.VK_NUMLOCK
	case fm.Keys_Num_Add:
		return win.VK_ADD
	case fm.Keys_Num_Sub:
		return win.VK_SUBTRACT
	case fm.Keys_Num_Multiply:
		return win.VK_MULTIPLY
	case fm.Keys_Num_Divide:
		return win.VK_DIVIDE
	case fm.Keys_Num_Decimal:
		return win.VK_DECIMAL
	case fm.Keys_Num_Separator:
		return win.VK_SEPARATOR
	case fm.Keys_Num_0:
		return win.VK_NUMPAD0
	case fm.Keys_Num_1:
		return win.VK_NUMPAD1
	case fm.Keys_Num_2:
		return win.VK_NUMPAD2
	case fm.Keys_Num_3:
		return win.VK_NUMPAD3
	case fm.Keys_Num_4:
		return win.VK_NUMPAD4
	case fm.Keys_Num_5:
		return win.VK_NUMPAD5
	case fm.Keys_Num_6:
		return win.VK_NUMPAD6
	case fm.Keys_Num_7:
		return win.VK_NUMPAD7
	case fm.Keys_Num_8:
		return win.VK_NUMPAD8
	case fm.Keys_Num_9:
		return win.VK_NUMPAD9
	default:
		fmt.Println("未定义的按键", key)
		return 0
	}
}

func sto16(str string) *uint16 {
	ptr, _ := syscall.UTF16PtrFromString(str)
	return ptr
}
