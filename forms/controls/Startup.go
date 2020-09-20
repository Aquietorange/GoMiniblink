package controls

import p "qq2564874169/goMiniblink/forms/platform"

type MainForm interface {
	getFormImpl() p.Form
}

var App p.Provider

func Run(form MainForm) {
	App.RunMain(form.getFormImpl())
}
