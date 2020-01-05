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

	SetMaximizeBox(isShow bool)
	SetMinimizeBox(isShow bool)
	SetIcon(iconFile string)
	SetIconVisable(isShow bool)

	SetOnState(func(MB.FormState))
}
