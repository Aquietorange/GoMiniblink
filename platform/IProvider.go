package platform

import (
	mb "qq.2564874169/goMiniblink"
)

type IProvider interface {
	RunMain(form IForm, show func())
	Exit(code int)
	SetIcon(file string)
	SetBgColor(color int)
	GetScreen() mb.Screen

	NewForm() IForm
	NewControl() IControl
	NewMiniblink() IMiniblink
}
