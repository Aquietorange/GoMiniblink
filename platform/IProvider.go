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
	ModifierKeys() map[mb.Keys]bool
	MouseIsDown() map[mb.MouseButtons]bool
	AppDir() string

	NewForm() IForm
	NewControl() IControl
	NewMiniblink() IMiniblink
}
