package forms

import MB "qq.2564874169/goMiniblink"

func (_this *BaseUI) defOnKeyPress(e *MB.KeyPressEvArgs) {
	for _, v := range _this.EvKeyPress {
		v(_this.instance, e)
	}
}

func (_this *BaseUI) defOnKeyUp(e *MB.KeyEvArgs) {
	for _, v := range _this.EvKeyUp {
		v(_this.instance, e)
	}
}

func (_this *BaseUI) defOnKeyDown(e *MB.KeyEvArgs) {
	for _, v := range _this.EvKeyDown {
		v(_this.instance, e)
	}
}

func (_this *BaseUI) defOnPaint(e MB.PaintEvArgs) {
	for _, v := range _this.EvPaint {
		v(_this.instance, e)
	}
}

func (_this *BaseUI) defOnMouseClick(e MB.MouseEvArgs) {
	for _, v := range _this.EvMouseClick {
		v(_this.instance, e)
	}
}

func (_this *BaseUI) defOnMouseWheel(e MB.MouseEvArgs) {
	for _, v := range _this.EvMouseWheel {
		v(_this.instance, e)
	}
}

func (_this *BaseUI) defOnMouseUp(e MB.MouseEvArgs) {
	for _, v := range _this.EvMouseUp {
		v(_this.instance, e)
	}
}

func (_this *BaseUI) defOnMouseDown(e MB.MouseEvArgs) {
	for _, v := range _this.EvMouseDown {
		v(_this.instance, e)
	}
}

func (_this *BaseUI) defOnMouseMove(e MB.MouseEvArgs) {
	for _, v := range _this.EvMouseMove {
		v(_this.instance, e)
	}
}

func (_this *BaseUI) defOnLoad() {
	for _, v := range _this.EvLoad {
		v(_this.instance)
	}
}

func (_this *BaseUI) defOnResize(e MB.Rect) {
	if _this.GetSize().IsEmpty() || _this.GetSize().IsEqual(e) {
		return
	}
	for _, v := range _this.EvResize {
		v(_this.instance, e)
	}
}

func (_this *BaseUI) defOnMove(e MB.Point) {
	if _this.GetLocation().IsEqual(e) {
		return
	}
	for _, v := range _this.EvMove {
		v(_this.instance, e)
	}
}
