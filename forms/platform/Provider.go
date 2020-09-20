package platform

import (
	f "qq2564874169/goMiniblink/forms"
)

type Provider interface {
	RunMain(form Form)
	Exit(code int)
	SetIcon(file string)
	GetScreen() f.Screen
	ModifierKeys() map[f.Keys]bool
	MouseIsDown() map[f.MouseButtons]bool
	MouseLocation() f.Point
	AppDir() string

	NewForm() Form
	NewControl() Control
}
