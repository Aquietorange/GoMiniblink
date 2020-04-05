package controls

import p "qq2564874169/goMiniblink/forms/platform"

var App p.IProvider

func Run(form *Form) {
	App.RunMain(form.getImpl())
}
