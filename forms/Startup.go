package forms

import "qq.2564874169/goMiniblink/platform"

var (
	Provider platform.IProvider
)

func Run(form *Form) {
	form.runMain(Provider)
}
