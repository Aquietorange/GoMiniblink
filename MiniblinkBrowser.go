package GoMiniblink

import (
	c "qq2564874169/goMiniblink/forms/controls"
)

type MiniblinkBrowser struct {
	c.Control
	mb IMiniblink
}

func (_this *MiniblinkBrowser) Init() *MiniblinkBrowser {
	_this.Control.Init()
	_this.mb = new(free4x64).init(&_this.Control)
	return _this
}
