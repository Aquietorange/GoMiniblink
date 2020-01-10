package free

type Miniblink struct {
	view uintptr
	wke  WKE
}

func (_this *Miniblink) Init(view uintptr) *Miniblink {
	if wkeIsInitialize() == false {
		panic("初始化失败")
	}
	_this.wke = wkeCreateWebView()
	if _this.wke == 0 {
		panic("创建失败")
	}
	_this.view = view
	wkeSetHandle(_this.wke, _this.view)

	return _this
}

func (_this *Miniblink) LoadUri(uri string) {
	wkeLoadURL(_this.wke, uri)
}
