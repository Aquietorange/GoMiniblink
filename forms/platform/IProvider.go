package platform

import (
	f "qq2564874169/goMiniblink/forms"
)

type IProvider interface {
	RunMain(form IForm)
	Exit(code int)
	SetIcon(file string)
	SetBgColor(color int)
	GetScreen() f.Screen
	ModifierKeys() map[f.Keys]bool
	MouseIsDown() map[f.MouseButtons]bool
	AppDir() string

	NewForm() IForm
	NewControl() IControl
}
