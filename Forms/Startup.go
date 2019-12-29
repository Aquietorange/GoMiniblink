package Forms

import "GoMiniblink/CrossPlatform"

var (
	Provider CrossPlatform.IProvider
)

func Run(form *Form) {
	form.runMain(Provider)
}
