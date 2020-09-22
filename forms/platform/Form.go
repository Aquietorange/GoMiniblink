package platform

import (
	f "qq2564874169/goMiniblink/forms"
)

type FormStateProc func(state f.FormState)
type FormActiveProc func()

type Form interface {
	Controls

	Close()
	ShowDialog()
	SetTitle(title string)
	SetBorderStyle(style f.FormBorder)
	ShowToMax()
	ShowToMin()
	NoneBorderResize()
	Active()

	SetMaximizeBox(isShow bool)
	SetMinimizeBox(isShow bool)
	SetIcon(iconFile string)

	SetOnState(proc FormStateProc) FormStateProc
	SetOnActive(proc FormActiveProc) FormActiveProc
}
