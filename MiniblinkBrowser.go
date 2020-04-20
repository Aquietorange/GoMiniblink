package GoMiniblink

import (
	c "qq2564874169/goMiniblink/forms/controls"
)

type MiniblinkBrowser struct {
	c.Control
	mb IMiniblink

	_initUri string
}

func (_this *MiniblinkBrowser) Init() *MiniblinkBrowser {
	_this.Control.Init()
	bakOnLoad := _this.Control.OnLoad
	_this.Control.OnLoad = func() {
		if bakOnLoad != nil {
			bakOnLoad()
		}
		_this.mb = new(free4x64).init(&_this.Control)
		_this.mbInit()
	}
	return _this
}

func (_this *MiniblinkBrowser) mbInit() {
	if _this._initUri != "" {
		_this.LoadUri(_this._initUri)
	}
}

func (_this *MiniblinkBrowser) LoadUri(uri string) {
	if _this.mb != nil {
		_this.mb.LoadUri(uri)
	} else {
		_this._initUri = uri
	}
}
