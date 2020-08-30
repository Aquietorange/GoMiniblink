package controls

import f "qq2564874169/goMiniblink/forms"

func (_this *FormBaseUI) defOnLostFocus() {
	for _, v := range _this.EvLostFocus {
		v(_this.instance)
	}
}

func (_this *FormBaseUI) defOnFocus() {
	for _, v := range _this.EvFocus {
		v(_this.instance)
	}
}

func (_this *FormBaseUI) defOnKeyPress(e *f.KeyPressEvArgs) {
	for _, v := range _this.EvKeyPress {
		v(_this.instance, e)
	}
}

func (_this *FormBaseUI) defOnKeyUp(e *f.KeyEvArgs) {
	for _, v := range _this.EvKeyUp {
		v(_this.instance, e)
	}
}

func (_this *FormBaseUI) defOnKeyDown(e *f.KeyEvArgs) {
	for _, v := range _this.EvKeyDown {
		v(_this.instance, e)
	}
}

func (_this *FormBaseUI) defOnPaint(e f.PaintEvArgs) {
	for _, v := range _this.EvPaint {
		v(_this.instance, e)
	}
}

func (_this *FormBaseUI) defOnMouseClick(e *f.MouseEvArgs) {
	for _, v := range _this.EvMouseClick {
		v(_this.instance, e)
	}
}

func (_this *FormBaseUI) defOnMouseWheel(e *f.MouseEvArgs) {
	for _, v := range _this.EvMouseWheel {
		v(_this.instance, e)
	}
}

func (_this *FormBaseUI) defOnMouseUp(e *f.MouseEvArgs) {
	for _, v := range _this.EvMouseUp {
		v(_this.instance, e)
	}
}

func (_this *FormBaseUI) defOnMouseDown(e *f.MouseEvArgs) {
	for _, v := range _this.EvMouseDown {
		v(_this.instance, e)
	}
}

func (_this *FormBaseUI) defOnMouseMove(e *f.MouseEvArgs) {
	for _, v := range _this.EvMouseMove {
		v(_this.instance, e)
	}
}

func (_this *FormBaseUI) defOnLoad() {
	for _, v := range _this.EvLoad {
		v(_this.instance)
	}
}

func (_this *FormBaseUI) defOnResize(e f.Rect) {
	if _this.GetSize().IsEmpty() || _this.GetSize().IsEqual(e) {
		return
	}
	for _, v := range _this.EvResize {
		v(_this.instance, e)
	}
}

func (_this *FormBaseUI) defOnMove(e f.Point) {
	if _this.GetLocation().IsEqual(e) {
		return
	}
	for _, v := range _this.EvMove {
		v(_this.instance, e)
	}
}
