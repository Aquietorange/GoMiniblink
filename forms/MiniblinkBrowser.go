package forms

import (
	mb "qq2564874169/goMiniblink"
	"qq2564874169/goMiniblink/platform"
)

type MiniblinkBrowser struct {
	BaseControl

	impl platform.IMiniblink
}

func (_this *MiniblinkBrowser) Init() *MiniblinkBrowser {
	_this.impl = Provider.NewMiniblink()
	_this.BaseControl.Init(_this.impl)
	_this.BaseControl.SetBgColor(-1)
	_this.EvResize["__syncResize"] = _this.mbOnResize
	return _this
}

func (_this *MiniblinkBrowser) mbOnResize(target interface{}, e mb.Rect) {
	if _this.GetSize().IsEmpty() || _this.GetSize().IsEqual(e) {
		return
	}
	_this.impl.SetSize(e.Width, e.Height)
}

func (_this *MiniblinkBrowser) LoadUri(uri string) {
	_this.impl.LoadUri(uri)
}
