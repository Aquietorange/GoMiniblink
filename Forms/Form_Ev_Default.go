package Forms

func (_this *Form) defOnLoad() {
	for _, v := range _this.EvLoad {
		v(_this)
	}
}

func (_this *Form) defOnResize(w, h int) {
	for _, v := range _this.EvResize[:] {
		v(_this, w, h)
	}
}

func (_this *Form) defOnMove(x, y int) {
	for _, v := range _this.EvMove[:] {
		v(_this, x, y)
	}
}
