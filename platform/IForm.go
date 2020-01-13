package platform

import mb "qq.2564874169/miniblink"

type IForm interface {
	IWindow
	IControls

	ShowDialog()
	SetTitle(title string)
	SetBorderStyle(style mb.FormBorder)
	ShowInTaskbar(isShow bool)
	ShowToMax()
	ShowToMin()

	SetMaximizeBox(isShow bool)
	SetMinimizeBox(isShow bool)
	SetIcon(iconFile string)
	SetIconVisable(isShow bool)

	SetOnState(func(mb.FormState))
}
