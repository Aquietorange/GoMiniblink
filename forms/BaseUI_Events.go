package forms

import mb "qq.2564874169/goMiniblink"

func (_this *BaseUI) defOnKeyPress(e *mb.KeyPressEvArgs) {
	for _, v := range _this.EvKeyPress {
		v(_this.instance, e)
	}
}

func (_this *BaseUI) defOnKeyUp(e *mb.KeyEvArgs) {
	for _, v := range _this.EvKeyUp {
		v(_this.instance, e)
	}
}

func (_this *BaseUI) defOnKeyDown(e *mb.KeyEvArgs) {
	for _, v := range _this.EvKeyDown {
		v(_this.instance, e)
	}
}

func (_this *BaseUI) defOnPaint(e mb.PaintEvArgs) {
	for _, v := range _this.EvPaint {
		v(_this.instance, e)
	}
}

func (_this *BaseUI) defOnMouseClick(e mb.MouseEvArgs) {
	for _, v := range _this.EvMouseClick {
		v(_this.instance, e)
	}
}

func (_this *BaseUI) defOnMouseWheel(e mb.MouseEvArgs) {
	for _, v := range _this.EvMouseWheel {
		v(_this.instance, e)
	}
}

func (_this *BaseUI) defOnMouseUp(e mb.MouseEvArgs) {
	for _, v := range _this.EvMouseUp {
		v(_this.instance, e)
	}
}

func (_this *BaseUI) defOnMouseDown(e mb.MouseEvArgs) {
	for _, v := range _this.EvMouseDown {
		v(_this.instance, e)
	}
}

func (_this *BaseUI) defOnMouseMove(e mb.MouseEvArgs) {
	for _, v := range _this.EvMouseMove {
		v(_this.instance, e)
	}
}

func (_this *BaseUI) defOnLoad() {
	for _, v := range _this.EvLoad {
		v(_this.instance)
	}
}

func (_this *BaseUI) defOnResize(e mb.Rect) {
	if _this.GetSize().IsEmpty() || _this.GetSize().IsEqual(e) {
		return
	}
	for _, v := range _this.EvResize {
		v(_this.instance, e)
	}
}

func (_this *BaseUI) defOnMove(e mb.Point) {
	if _this.GetLocation().IsEqual(e) {
		return
	}
	for _, v := range _this.EvMove {
		v(_this.instance, e)
	}
}
