package forms

import mb "qq2564874169/goMiniblink"

func (_this *MiniblinkBrowser) defOnRequest(e mb.RequestBeforeEvArgs) {
	for i := 0; i < len(_this.EvRequestBefore); i++ {
		_this.EvRequestBefore[i](e)
	}
}

func (_this *MiniblinkBrowser) defOnJsReady(e mb.JsReadyEvArgs) {
	for i := 0; i < len(_this.EvJsReady); i++ {
		_this.EvJsReady[i](e)
	}
}
