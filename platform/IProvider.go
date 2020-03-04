package platform

import (
	mb "qq2564874169/goMiniblink"
)

type IProvider interface {
	RunMain(form IForm, show func())
	Exit(code int)
	SetIcon(file string)
	SetBgColor(color int)
	GetScreen() mb.Screen
	KeyIsDown(key mb.Keys) bool
	MouseIsDown(button mb.MouseButtons) bool

	NewForm() IForm
	NewControl() IControl
	NewMiniblink() IMiniblink
}
