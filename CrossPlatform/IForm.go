package CrossPlatform

import MB "GoMiniblink"

type IForm interface {
	IWindow
	ShowDialog()
	SetTitle(title string)
	SetBorderStyle(style MB.FormBorder)
	ShowInTaskbar(isShow bool)
	ShowToMax()
	ShowToMin()

	SetSize(w int, h int)
	SetLocation(x int, y int)
	SetMaximizeBox(isShow bool)
	SetMinimizeBox(isShow bool)
	SetIcon(iconFile string)
	SetIconVisable(isShow bool)

	SetOnResize(func(w int, h int))
	SetOnMove(func(x int, y int))
	SetOnState(func(MB.FormState))
}
