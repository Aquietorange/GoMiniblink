package forms

//鼠标样式
type CursorType int

//鼠标样式
const (
	CursorType_Default CursorType = iota + 1
	CursorType_ARROW
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
type FormStart int

//窗体的首次显示位置
const (
	FormStart_Default FormStart = iota
	FormStart_Manual
	FormStart_Screen_Center
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
	AnchorStyle_Fill   AnchorStyle = 16
)
