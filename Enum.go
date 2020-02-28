package goMiniblink

//鼠标样式
type CursorType int

//鼠标样式
const (
	CursorType_ARROW CursorType = iota + 1
	CursorType_IBEAM
	CursorType_WAIT
	CursorType_CROSS
	CursorType_UPARROW
	CursorType_SIZENWSE
	CursorType_SIZENESW
	CursorType_SIZEWE
	CursorType_SIZENS
	CursorType_SIZEALL
	CursorType_NO
	CursorType_HAND
	CursorType_APPSTARTING
	CursorType_HELP
	CursorType_ICON
	CursorType_SIZE
)

//窗体的边框类型
type FormBorder int

//窗体的边框类型
const (
	FormBorder_Default FormBorder = iota
	FormBorder_None
	FormBorder_Disable_Resize
)

//窗体的状态类型
type FormState int

//窗体的状态类型
const (
	FormState_Normal FormState = iota
	FormState_Max
	FormState_Min
)

//窗体的首次显示位置
type FormStartPosition int

//窗体的首次显示位置
const (
	FormStartPosition_Manual FormStartPosition = iota
	FormStartPosition_Screen_Center
)

type MouseButtons int

//鼠标按键
const (
	MouseButtons_None   MouseButtons = 0
	MouseButtons_Left   MouseButtons = 1
	MouseButtons_Right  MouseButtons = 2
	MouseButtons_Middle MouseButtons = 4
)

//控件停靠方式
type AnchorStyle int

const (
	AnchorStyle_Left   AnchorStyle = 1
	AnchorStyle_Right  AnchorStyle = 2
	AnchorStyle_Top    AnchorStyle = 4
	AnchorStyle_Bottom AnchorStyle = 8
)

type Keys int

//键盘按键
const (
	Keys_Error Keys = iota
	Keys_Esc
	Keys_F1
	Keys_F2
	Keys_F3
	Keys_F4
	Keys_F5
	Keys_F6
	Keys_F7
	Keys_F8
	Keys_F9
	Keys_F10
	Keys_F11
	Keys_F12
	Keys_Space
	Keys_Tab
	Keys_CapsLock
	Keys_Shift
	Keys_Ctrl
	Keys_Alt
	Keys_Win
	Keys_Enter
	Keys_Right_Shift
	Keys_Right_Ctrl
	Keys_Right_Alt
	Keys_Right_Win
	Keys_Apps
	Keys_0
	Keys_1
	Keys_2
	Keys_3
	Keys_4
	Keys_5
	Keys_6
	Keys_7
	Keys_8
	Keys_9
	Keys_OEM_3
	Keys_OEM_PLUS
	Keys_OEM_MINUS
	Keys_Back
	Keys_Snapshot
	Keys_Scroll_Lock
	Keys_Pause
	Keys_Q
	Keys_W
	Keys_E
	Keys_R
	Keys_T
	Keys_Y
	Keys_U
	Keys_I
	Keys_O
	Keys_P
	Keys_A
	Keys_S
	Keys_D
	Keys_F
	Keys_G
	Keys_H
	Keys_J
	Keys_K
	Keys_L
	Keys_Z
	Keys_X
	Keys_C
	Keys_V
	Keys_B
	Keys_N
	Keys_M
	Keys_OEM_4
	Keys_OEM_6
	Keys_OEM_1
	Keys_OEM_7
	Keys_OEM_5
	Keys_OEM_COMMA
	Keys_OEM_PERIOD
	Keys_OEM_2
	Keys_Insert
	Keys_Delete
	Keys_Home
	Keys_End
	Keys_PageUp
	Keys_PageDown
	Keys_Left
	Keys_Up
	Keys_Right
	Keys_Down
	Keys_Num_Lock
	Keys_Num_Add
	Keys_Num_Sub
	Keys_Num_Multiply
	Keys_Num_Divide
	Keys_Num_Decimal
	Keys_Num_Separator
	Keys_Num_0
	Keys_Num_1
	Keys_Num_2
	Keys_Num_3
	Keys_Num_4
	Keys_Num_5
	Keys_Num_6
	Keys_Num_7
	Keys_Num_8
	Keys_Num_9
	KeysCount
)
