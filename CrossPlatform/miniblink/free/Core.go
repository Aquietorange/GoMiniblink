package free

import (
	"GoMiniblink/CrossPlatform"
	"GoMiniblink/CrossPlatform/miniblink"
)

type Core struct {
	view CrossPlatform.IWindow
	wke  wkeHandle

	paintCallback miniblink.PaintCallback
}

func (_this *Core) Init(window CrossPlatform.IWindow) *Core {
	if wkeInitialize() == false {
		panic("初始化失败")
	}
	_this.wke = wkeCreateWebView()
	if _this.wke == 0 {
		panic("创建失败")
	}
	_this.view = window
	wkeSetHandle(_this.wke, _this.view.GetHandle())
	return _this
}

func (_this *Core) SetOnPaint(callback miniblink.PaintCallback) {
	_this.paintCallback = callback
}

func (_this *Core) LoadUri(uri string) {
	wkeLoadURL(_this.wke, uri)
}
