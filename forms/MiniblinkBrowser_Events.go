package forms

import mb "qq2564874169/goMiniblink"

func (_this *MiniblinkBrowser) defOnRequest(e mb.RequestEvArgs) {
	for i := 0; i < len(_this.EvRequest); i++ {
		_this.EvRequest[i](e)
	}
}
