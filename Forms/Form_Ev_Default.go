package Forms

import "GoMiniblink"

func (_this *Form) defOnLoad() {
	for _, v := range _this.EvLoad {
		v(_this)
	}
}

func (_this *Form) defOnResize(w, h int) {
	for _, v := range _this.EvResize {
		v(_this, w, h)
	}
}

func (_this *Form) defOnMove(x, y int) {
	for _, v := range _this.EvMove {
		v(_this, x, y)
	}
}

func (_this *Form) defOnState(state GoMiniblink.FormState) {
	for _, v := range _this.EvState {
		v(_this, state)
	}
}
