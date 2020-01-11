package CrossPlatform

import (
	MB "GoMiniblink"
)

type IProvider interface {
	RunMain(form IForm, show func())
	Exit(code int)
	SetIcon(file string)
	GetScreen() MB.Screen

	NewForm() IForm
	NewControl() IControl
	NewMiniblink() IMiniblink
}
