package Forms

import "GoMiniblink/Forms/CrossPlatform"

var (
	Provider CrossPlatform.IProvider
)

func Run(form *Form) {
	form.RunMain(Provider)
}
