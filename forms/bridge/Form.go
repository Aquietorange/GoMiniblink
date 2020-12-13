package bridge

import (
	fm "gitee.com/aochulai/GoMiniblink/forms"
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
	/*
		允许在无边框模式下调整窗体大小
	*/
	NoneBorderResize()
	Active()

	SetMaximizeBox(isShow bool)
	SetMinimizeBox(isShow bool)
	SetIcon(iconFile string)

	SetOnState(proc FormStateProc) FormStateProc
	SetOnActive(proc FormActiveProc) FormActiveProc
}
