package platform

import (
	f "qq2564874169/goMiniblink/forms"
)

type FormStateProc func(state f.FormState)

type Form interface {
	Controls

	Close()
	ShowDialog()
	SetTitle(title string)
	SetBorderStyle(style f.FormBorder)
	ShowToMax()
	ShowToMin()
	NoneBorderResize()

	SetMaximizeBox(isShow bool)
	SetMinimizeBox(isShow bool)
	SetIcon(iconFile string)

	SetOnState(proc FormStateProc) FormStateProc
}
