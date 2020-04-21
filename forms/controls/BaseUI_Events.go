package controls

import f "qq2564874169/goMiniblink/forms"

func (_this *BaseUI) defOnLostFocus() {
	for _, v := range _this.EvLostFocus {
		v(_this.instance)
	}
}

func (_this *BaseUI) defOnFocus() {
	for _, v := range _this.EvFocus {
		v(_this.instance)
	}
}

func (_this *BaseUI) defOnKeyPress(e *f.KeyPressEvArgs) {
	for _, v := range _this.EvKeyPress {
		v(_this.instance, e)
	}
}

func (_this *BaseUI) defOnKeyUp(e *f.KeyEvArgs) {
	for _, v := range _this.EvKeyUp {
		v(_this.instance, e)
	}
}

func (_this *BaseUI) defOnKeyDown(e *f.KeyEvArgs) {
	for _, v := range _this.EvKeyDown {
		v(_this.instance, e)
	}
}

func (_this *BaseUI) defOnPaint(e f.PaintEvArgs) {
	for _, v := range _this.EvPaint {
		v(_this.instance, e)
	}
}

func (_this *BaseUI) defOnMouseClick(e *f.MouseEvArgs) {
	for _, v := range _this.EvMouseClick {
		v(_this.instance, e)
	}
}

func (_this *BaseUI) defOnMouseWheel(e *f.MouseEvArgs) {
	for _, v := range _this.EvMouseWheel {
		v(_this.instance, e)
	}
}

func (_this *BaseUI) defOnMouseUp(e *f.MouseEvArgs) {
	for _, v := range _this.EvMouseUp {
		v(_this.instance, e)
	}
}

func (_this *BaseUI) defOnMouseDown(e *f.MouseEvArgs) {
	for _, v := range _this.EvMouseDown {
		v(_this.instance, e)
	}
}

func (_this *BaseUI) defOnMouseMove(e *f.MouseEvArgs) {
	for _, v := range _this.EvMouseMove {
		v(_this.instance, e)
	}
}

func (_this *BaseUI) defOnLoad() {
	for _, v := range _this.EvLoad {
		v(_this.instance)
	}
}

func (_this *BaseUI) defOnResize(e f.Rect) {
	if _this.GetSize().IsEmpty() || _this.GetSize().IsEqual(e) {
		return
	}
	for _, v := range _this.EvResize {
		v(_this.instance, e)
	}
}

func (_this *BaseUI) defOnMove(e f.Point) {
	if _this.GetLocation().IsEqual(e) {
		return
	}
	for _, v := range _this.EvMove {
		v(_this.instance, e)
	}
}
