package platform

import mb "qq2564874169/goMiniblink"

type IForm interface {
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
