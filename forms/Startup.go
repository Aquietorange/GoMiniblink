package forms

import "qq2564874169/goMiniblink/platform"

var (
	Provider platform.IProvider
)

func Run(form *Form) {
	form.runMain(Provider)
}
