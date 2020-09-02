package GoMiniblink

import (
	f "qq2564874169/goMiniblink/forms"
	c "qq2564874169/goMiniblink/forms/controls"
)

type MiniblinkForm struct {
	c.Form
	View *MiniblinkBrowser
}

func (_this *MiniblinkForm) Init(isTransparent bool) *MiniblinkForm {
	_this.Form.Init()
	_this.View = new(MiniblinkBrowser).Init()
	_this.View.SetAnchor(f.AnchorStyle_Top | f.AnchorStyle_Right | f.AnchorStyle_Bottom | f.AnchorStyle_Left)
	_this.AddChild(_this.View)

	_this.EvResize["set_view_size"] = func(_ interface{}, e f.Rect) {
		_this.View.SetSize(e.Width, e.Height)
	}
	if isTransparent {
		_this.EvLoad["SetTransparent"] = _this.setTransparent
	}
	return _this
}

func (_this *MiniblinkForm) setTransparent(_ interface{}) {

}
