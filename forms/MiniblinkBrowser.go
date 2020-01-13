package forms

import (
	mb "qq.2564874169/miniblink"
	"qq.2564874169/miniblink/platform"
)

type MiniblinkBrowser struct {
	BaseControl

	impl platform.IMiniblink
}

func (_this *MiniblinkBrowser) Init() *MiniblinkBrowser {
	_this.impl = Provider.NewMiniblink()
	_this.BaseControl.Init(_this.impl)
	_this.EvResize["__miniblink"] = _this.mbOnResize
	return _this
}

func (_this *MiniblinkBrowser) mbOnResize(target interface{}, e mb.Rect) {
	_this.impl.Resize(e.Wdith, e.Height)
}

func (_this *MiniblinkBrowser) LoadUri(uri string) {
	_this.impl.LoadUri(uri)
}
