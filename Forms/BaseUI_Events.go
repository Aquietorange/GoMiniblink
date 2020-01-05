package Forms

import MB "GoMiniblink"

func (_this *BaseUI) defOnPaint(e MB.PaintEvArgs) {
	for _, v := range _this.EvPaint {
		v(_this, e)
	}
}

func (_this *BaseUI) defOnMouseClick(e MB.MouseEvArgs) {
	for _, v := range _this.EvMouseClick {
		v(_this, e)
	}
}

func (_this *BaseUI) defOnMouseWheel(e MB.MouseEvArgs) {
	for _, v := range _this.EvMouseWheel {
		v(_this, e)
	}
}

func (_this *BaseUI) defOnMouseUp(e MB.MouseEvArgs) {
	for _, v := range _this.EvMouseUp {
		v(_this, e)
	}
}

func (_this *BaseUI) defOnMouseDown(e MB.MouseEvArgs) {
	for _, v := range _this.EvMouseDown {
		v(_this, e)
	}
}

func (_this *BaseUI) defOnMouseMove(e MB.MouseEvArgs) {
	for _, v := range _this.EvMouseMove {
		v(_this, e)
	}
}

func (_this *BaseUI) defOnLoad() {
	for _, v := range _this.EvLoad {
		v(_this)
	}
}

func (_this *BaseUI) defOnResize(e MB.Rect) {
	for _, v := range _this.EvResize {
		v(_this, e)
	}
}

func (_this *BaseUI) defOnMove(e MB.Point) {
	for _, v := range _this.EvMove {
		v(_this, e)
	}
}
