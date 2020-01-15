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
	_this.EvResize["__syncResize"] = _this.mbOnResize
	return _this
}

func (_this *MiniblinkBrowser) mbOnResize(target interface{}, e mb.Rect) {
	if _this.GetSize().IsEmpty() || _this.GetSize().IsEqual(e) {
		return
	}
	_this.impl.SetSize(e.Wdith, e.Height)
}

func (_this *MiniblinkBrowser) LoadUri(uri string) {
	_this.impl.LoadUri(uri)
}
