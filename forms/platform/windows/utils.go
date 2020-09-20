package windows

import (
	"fmt"
	"image/color"
	"qq2564874169/goMiniblink/forms"
	plat "qq2564874169/goMiniblink/forms/platform"
	win "qq2564874169/goMiniblink/forms/platform/windows/win32"
	"syscall"
)

func findChild(container plat.Controls, hWnd win.HWND) plat.Window {
	if container.GetHandle() == uintptr(hWnd) {
		return container
	}
	for _, i := range container.GetChilds() {
		if c, ok := i.(plat.Controls); ok {
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

func intToRGBA(rgba int) color.RGBA {
	return color.RGBA{
		R: uint8(rgba),
		G: uint8(rgba >> 8),
		B: uint8(rgba >> 16),
		A: uint8(rgba >> 24),
	}
}

func isExtendKey(key forms.Keys) bool {
	switch key {
	case forms.Keys_Insert, forms.Keys_Delete, forms.Keys_Home, forms.Keys_End, forms.Keys_PageUp,
		forms.Keys_PageDown, forms.Keys_Left, forms.Keys_Right, forms.Keys_Up, forms.Keys_Down:
		return true
	default:
		return false
	}
}

func toWinCursor(cursor forms.CursorType) int {
	switch cursor {
	case forms.CursorType_ARROW:
		return win.IDC_ARROW
	case forms.CursorType_SIZE:
		return win.IDC_SIZE
	case forms.CursorType_ICON:
		return win.IDC_ICON
	case forms.CursorType_HELP:
		return win.IDC_HELP
	case forms.CursorType_APPSTARTING:
		return win.IDC_APPSTARTING
	case forms.CursorType_HAND:
		return win.IDC_HAND
	case forms.CursorType_NO:
		return win.IDC_NO
	case forms.CursorType_SIZEALL:
		return win.IDC_SIZEALL
	case forms.CursorType_SIZENS:
		return win.IDC_SIZENS
	case forms.CursorType_SIZEWE:
		return win.IDC_SIZEWE
	case forms.CursorType_SIZENWSE:
		return win.IDC_SIZENWSE
	case forms.CursorType_SIZENESW:
		return win.IDC_SIZENESW
	case forms.CursorType_UPARROW:
		return win.IDC_UPARROW
	case forms.CursorType_CROSS:
		return win.IDC_CROSS
	case forms.CursorType_WAIT:
		return win.IDC_WAIT
	case forms.CursorType_IBEAM:
		return win.IDC_IBEAM
	default:
		return win.IDC_ARROW
	}
}

func winCursorTo(cursor int) forms.CursorType {
	switch cursor {
	case win.IDC_ARROW:
		return forms.CursorType_ARROW
	case win.IDC_IBEAM:
		return forms.CursorType_IBEAM
	case win.IDC_WAIT:
		return forms.CursorType_WAIT
	case win.IDC_CROSS:
		return forms.CursorType_CROSS
	case win.IDC_UPARROW:
		return forms.CursorType_UPARROW
	case win.IDC_SIZENWSE:
		return forms.CursorType_SIZENWSE
	case win.IDC_SIZENESW:
		return forms.CursorType_SIZENESW
	case win.IDC_SIZEWE:
		return forms.CursorType_SIZEWE
	case win.IDC_SIZENS:
		return forms.CursorType_SIZENS
	case win.IDC_SIZEALL:
		return forms.CursorType_SIZEALL
	case win.IDC_NO:
		return forms.CursorType_NO
	case win.IDC_HAND:
		return forms.CursorType_HAND
	case win.IDC_APPSTARTING:
		return forms.CursorType_APPSTARTING
	case win.IDC_HELP:
		return forms.CursorType_HELP
	case win.IDC_ICON:
		return forms.CursorType_ICON
	case win.IDC_SIZE:
		return forms.CursorType_SIZE
	default:
		return forms.CursorType_ARROW
	}
}

func vkToKey(vk int) forms.Keys {
	switch vk {
	case win.VK_ESCAPE:
		return forms.Keys_Esc
	case win.VK_F1:
		return forms.Keys_F1
	case win.VK_F2:
		return forms.Keys_F2
	case win.VK_F3:
		return forms.Keys_F3
	case win.VK_F4:
		return forms.Keys_F4
	case win.VK_F5:
		return forms.Keys_F5
	case win.VK_F6:
		return forms.Keys_F6
	case win.VK_F7:
		return forms.Keys_F7
	case win.VK_F8:
		return forms.Keys_F8
	case win.VK_F9:
		return forms.Keys_F9
	case win.VK_F10:
		return forms.Keys_F10
	case win.VK_F11:
		return forms.Keys_F11
	case win.VK_F12:
		return forms.Keys_F12
	case win.VK_SPACE:
		return forms.Keys_Space
	case win.VK_TAB:
		return forms.Keys_Tab
	case win.VK_CAPITAL:
		return forms.Keys_CapsLock
	case win.VK_SHIFT:
		return forms.Keys_Shift
	case win.VK_CONTROL:
		return forms.Keys_Ctrl
	case win.VK_BACK:
		return forms.Keys_Back
	case win.VK_MENU:
		return forms.Keys_Alt
	case win.VK_LWIN:
		return forms.Keys_Win
	case win.VK_RSHIFT:
		return forms.Keys_Right_Shift
	case win.VK_RCONTROL:
		return forms.Keys_Right_Ctrl
	case win.VK_RMENU:
		return forms.Keys_Right_Alt
	case win.VK_RWIN:
		return forms.Keys_Right_Win
	case win.VK_RETURN:
		return forms.Keys_Enter
	case win.VK_APPS:
		return forms.Keys_Apps
	case 0x30:
		return forms.Keys_0
	case 0x31:
		return forms.Keys_1
	case 0x32:
		return forms.Keys_2
	case 0x33:
		return forms.Keys_3
	case 0x34:
		return forms.Keys_4
	case 0x35:
		return forms.Keys_5
	case 0x36:
		return forms.Keys_6
	case 0x37:
		return forms.Keys_7
	case 0x38:
		return forms.Keys_8
	case 0x39:
		return forms.Keys_9
	case win.VK_OEM_3:
		return forms.Keys_OEM_3
	case win.VK_OEM_PLUS:
		return forms.Keys_OEM_PLUS
	case win.VK_OEM_MINUS:
		return forms.Keys_OEM_MINUS
	case win.VK_SNAPSHOT:
		return forms.Keys_Snapshot
	case win.VK_SCROLL:
		return forms.Keys_Scroll_Lock
	case win.VK_PAUSE:
		return forms.Keys_Pause
	case 0x51:
		return forms.Keys_Q
	case 0x57:
		return forms.Keys_W
	case 0x45:
		return forms.Keys_E
	case 0x52:
		return forms.Keys_R
	case 0x54:
		return forms.Keys_T
	case 0x59:
		return forms.Keys_Y
	case 0x55:
		return forms.Keys_U
	case 0x49:
		return forms.Keys_I
	case 0x4F:
		return forms.Keys_O
	case 0x50:
		return forms.Keys_P
	case 0x41:
		return forms.Keys_A
	case 0x53:
		return forms.Keys_S
	case 0x44:
		return forms.Keys_D
	case 0x46:
		return forms.Keys_F
	case 0x47:
		return forms.Keys_G
	case 0x48:
		return forms.Keys_H
	case 0x4A:
		return forms.Keys_J
	case 0x4B:
		return forms.Keys_K
	case 0x4C:
		return forms.Keys_L
	case 0x5A:
		return forms.Keys_Z
	case 0x58:
		return forms.Keys_X
	case 0x43:
		return forms.Keys_C
	case 0x56:
		return forms.Keys_V
	case 0x42:
		return forms.Keys_B
	case 0x4E:
		return forms.Keys_N
	case 0x4D:
		return forms.Keys_M
	case win.VK_OEM_4:
		return forms.Keys_OEM_4
	case win.VK_OEM_6:
		return forms.Keys_OEM_6
	case win.VK_OEM_1:
		return forms.Keys_OEM_1
	case win.VK_OEM_7:
		return forms.Keys_OEM_7
	case win.VK_OEM_5:
		return forms.Keys_OEM_5
	case win.VK_OEM_COMMA:
		return forms.Keys_OEM_COMMA
	case win.VK_OEM_PERIOD:
		return forms.Keys_OEM_PERIOD
	case win.VK_OEM_2:
		return forms.Keys_OEM_2
	case win.VK_INSERT:
		return forms.Keys_Insert
	case win.VK_DELETE:
		return forms.Keys_Delete
	case win.VK_HOME:
		return forms.Keys_Home
	case win.VK_END:
		return forms.Keys_End
	case win.VK_PRIOR:
		return forms.Keys_PageUp
	case win.VK_NEXT:
		return forms.Keys_PageDown
	case win.VK_LEFT:
		return forms.Keys_Left
	case win.VK_UP:
		return forms.Keys_Up
	case win.VK_RIGHT:
		return forms.Keys_Right
	case win.VK_DOWN:
		return forms.Keys_Down
	case win.VK_NUMLOCK:
		return forms.Keys_Num_Lock
	case win.VK_ADD:
		return forms.Keys_Num_Add
	case win.VK_SUBTRACT:
		return forms.Keys_Num_Sub
	case win.VK_MULTIPLY:
		return forms.Keys_Num_Multiply
	case win.VK_DIVIDE:
		return forms.Keys_Num_Divide
	case win.VK_DECIMAL:
		return forms.Keys_Num_Decimal
	case win.VK_SEPARATOR:
		return forms.Keys_Num_Separator
	case win.VK_NUMPAD0:
		return forms.Keys_Num_0
	case win.VK_NUMPAD1:
		return forms.Keys_Num_1
	case win.VK_NUMPAD2:
		return forms.Keys_Num_2
	case win.VK_NUMPAD3:
		return forms.Keys_Num_3
	case win.VK_NUMPAD4:
		return forms.Keys_Num_4
	case win.VK_NUMPAD5:
		return forms.Keys_Num_5
	case win.VK_NUMPAD6:
		return forms.Keys_Num_6
	case win.VK_NUMPAD7:
		return forms.Keys_Num_7
	case win.VK_NUMPAD8:
		return forms.Keys_Num_8
	case win.VK_NUMPAD9:
		return forms.Keys_Num_9
	default:
		return forms.Keys_Error
	}
}

func keyToVk(key forms.Keys) int {
	switch key {
	case forms.Keys_Esc:
		return win.VK_ESCAPE
	case forms.Keys_F1:
		return win.VK_F1
	case forms.Keys_F2:
		return win.VK_F2
	case forms.Keys_F3:
		return win.VK_F3
	case forms.Keys_F4:
		return win.VK_F4
	case forms.Keys_F5:
		return win.VK_F5
	case forms.Keys_F6:
		return win.VK_F6
	case forms.Keys_F7:
		return win.VK_F7
	case forms.Keys_F8:
		return win.VK_F8
	case forms.Keys_F9:
		return win.VK_F9
	case forms.Keys_F10:
		return win.VK_F10
	case forms.Keys_F11:
		return win.VK_F11
	case forms.Keys_F12:
		return win.VK_F12
	case forms.Keys_Space:
		return win.VK_SPACE
	case forms.Keys_Tab:
		return win.VK_TAB
	case forms.Keys_CapsLock:
		return win.VK_CAPITAL
	case forms.Keys_Shift:
		return win.VK_SHIFT
	case forms.Keys_Ctrl:
		return win.VK_CONTROL
	case forms.Keys_Back:
		return win.VK_BACK
	case forms.Keys_Alt:
		return win.VK_MENU
	case forms.Keys_Win:
		return win.VK_LWIN
	case forms.Keys_Right_Shift:
		return win.VK_RSHIFT
	case forms.Keys_Right_Ctrl:
		return win.VK_RCONTROL
	case forms.Keys_Right_Alt:
		return win.VK_RMENU
	case forms.Keys_Right_Win:
		return win.VK_RWIN
	case forms.Keys_Enter:
		return win.VK_RETURN
	case forms.Keys_Apps:
		return win.VK_APPS
	case forms.Keys_0:
		return 0x30
	case forms.Keys_1:
		return 0x31
	case forms.Keys_2:
		return 0x32
	case forms.Keys_3:
		return 0x33
	case forms.Keys_4:
		return 0x34
	case forms.Keys_5:
		return 0x35
	case forms.Keys_6:
		return 0x36
	case forms.Keys_7:
		return 0x37
	case forms.Keys_8:
		return 0x38
	case forms.Keys_9:
		return 0x39
	case forms.Keys_OEM_3:
		return win.VK_OEM_3
	case forms.Keys_OEM_PLUS:
		return win.VK_OEM_PLUS
	case forms.Keys_OEM_MINUS:
		return win.VK_OEM_MINUS
	case forms.Keys_Snapshot:
		return win.VK_SNAPSHOT
	case forms.Keys_Scroll_Lock:
		return win.VK_SCROLL
	case forms.Keys_Pause:
		return win.VK_PAUSE
	case forms.Keys_Q:
		return 0x51
	case forms.Keys_W:
		return 0x57
	case forms.Keys_E:
		return 0x45
	case forms.Keys_R:
		return 0x52
	case forms.Keys_T:
		return 0x54
	case forms.Keys_Y:
		return 0x59
	case forms.Keys_U:
		return 0x55
	case forms.Keys_I:
		return 0x49
	case forms.Keys_O:
		return 0x4F
	case forms.Keys_P:
		return 0x50
	case forms.Keys_A:
		return 0x41
	case forms.Keys_S:
		return 0x53
	case forms.Keys_D:
		return 0x44
	case forms.Keys_F:
		return 0x46
	case forms.Keys_G:
		return 0x47
	case forms.Keys_H:
		return 0x48
	case forms.Keys_J:
		return 0x4A
	case forms.Keys_K:
		return 0x4B
	case forms.Keys_L:
		return 0x4C
	case forms.Keys_Z:
		return 0x5A
	case forms.Keys_X:
		return 0x58
	case forms.Keys_C:
		return 0x43
	case forms.Keys_V:
		return 0x56
	case forms.Keys_B:
		return 0x42
	case forms.Keys_N:
		return 0x4E
	case forms.Keys_M:
		return 0x4D
	case forms.Keys_OEM_4:
		return win.VK_OEM_4
	case forms.Keys_OEM_6:
		return win.VK_OEM_6
	case forms.Keys_OEM_1:
		return win.VK_OEM_1
	case forms.Keys_OEM_7:
		return win.VK_OEM_7
	case forms.Keys_OEM_5:
		return win.VK_OEM_5
	case forms.Keys_OEM_COMMA:
		return win.VK_OEM_COMMA
	case forms.Keys_OEM_PERIOD:
		return win.VK_OEM_PERIOD
	case forms.Keys_OEM_2:
		return win.VK_OEM_2
	case forms.Keys_Insert:
		return win.VK_INSERT
	case forms.Keys_Delete:
		return win.VK_DELETE
	case forms.Keys_Home:
		return win.VK_HOME
	case forms.Keys_End:
		return win.VK_END
	case forms.Keys_PageUp:
		return win.VK_PRIOR
	case forms.Keys_PageDown:
		return win.VK_NEXT
	case forms.Keys_Left:
		return win.VK_LEFT
	case forms.Keys_Up:
		return win.VK_UP
	case forms.Keys_Right:
		return win.VK_RIGHT
	case forms.Keys_Down:
		return win.VK_DOWN
	case forms.Keys_Num_Lock:
		return win.VK_NUMLOCK
	case forms.Keys_Num_Add:
		return win.VK_ADD
	case forms.Keys_Num_Sub:
		return win.VK_SUBTRACT
	case forms.Keys_Num_Multiply:
		return win.VK_MULTIPLY
	case forms.Keys_Num_Divide:
		return win.VK_DIVIDE
	case forms.Keys_Num_Decimal:
		return win.VK_DECIMAL
	case forms.Keys_Num_Separator:
		return win.VK_SEPARATOR
	case forms.Keys_Num_0:
		return win.VK_NUMPAD0
	case forms.Keys_Num_1:
		return win.VK_NUMPAD1
	case forms.Keys_Num_2:
		return win.VK_NUMPAD2
	case forms.Keys_Num_3:
		return win.VK_NUMPAD3
	case forms.Keys_Num_4:
		return win.VK_NUMPAD4
	case forms.Keys_Num_5:
		return win.VK_NUMPAD5
	case forms.Keys_Num_6:
		return win.VK_NUMPAD6
	case forms.Keys_Num_7:
		return win.VK_NUMPAD7
	case forms.Keys_Num_8:
		return win.VK_NUMPAD8
	case forms.Keys_Num_9:
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
