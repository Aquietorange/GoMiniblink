package GoMiniblink

/*
窗体的边框类型
*/
type FormBorder int

const (
	FormBorder_Default FormBorder = iota
	FormBorder_None
	FormBorder_Disable_Resize
)

/*
窗体的状态类型
*/
type FormState int

const (
	FormState_Normal FormState = iota
	FormState_Max
	FormState_Min
)

/*
窗体的首次显示位置
*/
type FormStartPosition int

const (
	FormStartPosition_Manual FormStartPosition = iota
	FormStartPosition_Screen_Center
)

type MouseButtons int

const (
	MouseButtons_Left       MouseButtons = 2
	MouseButtons_Right      MouseButtons = 4
	MouseButtons_Middle     MouseButtons = 8
	MouseButtons_Left_Right MouseButtons = 16
)
