package controls

import (
	"gitee.com/aochulai/GoMiniblink/forms/bridge"
)

type MainForm interface {
	getFormImpl() bridge.Form
}

var App bridge.Provider

func Run(form MainForm) {
	App.RunMain(form.getFormImpl())
}
