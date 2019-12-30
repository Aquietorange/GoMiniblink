package Forms

import MB "GoMiniblink"

func (_this *Form) defOnMouseClick(e MB.MouseEvArgs) {
	for _, v := range _this.EvMouseClick {
		v(_this, e)
	}
}

func (_this *Form) defOnMouseWheel(e MB.MouseEvArgs) {
	for _, v := range _this.EvMouseWheel {
		v(_this, e)
	}
}

func (_this *Form) defOnMouseUp(e MB.MouseEvArgs) {
	for _, v := range _this.EvMouseUp {
		v(_this, e)
	}
}

func (_this *Form) defOnMouseDown(e MB.MouseEvArgs) {
	for _, v := range _this.EvMouseDown {
		v(_this, e)
	}
}

func (_this *Form) defOnMouseMove(e MB.MouseEvArgs) {
	for _, v := range _this.EvMouseMove {
		v(_this, e)
	}
}

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

func (_this *Form) defOnState(state MB.FormState) {
	for _, v := range _this.EvState {
		v(_this, state)
	}
}
