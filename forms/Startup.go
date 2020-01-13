package forms

import "qq.2564874169/miniblink/platform"

var (
	Provider platform.IProvider
)

func Run(form *Form) {
	form.runMain(Provider)
}
