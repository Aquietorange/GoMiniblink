package bridge

import (
	fm "GoMiniblink/forms"
)

type FormStateProc func(state fm.FormState)
type FormActiveProc func()

type Form interface {
	Controls

	Close()
	ShowDialog()
	SetTitle(title string)
	SetBorderStyle(style fm.FormBorder)
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
