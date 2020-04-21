package GoMiniblink

import (
	url2 "net/url"
	c "qq2564874169/goMiniblink/forms/controls"
	"strings"
)

type MiniblinkBrowser struct {
	c.Control
	_mb      IMiniblink
	_initUri string

	EvRequestBefore []func(target *MiniblinkBrowser, e RequestEvArgs)
	OnRequestBefore func(e RequestEvArgs)

	ResourceLoader []ILoadResource
}

func (_this *MiniblinkBrowser) Init() *MiniblinkBrowser {
	_this.Control.Init()
	_this.OnRequestBefore = _this.defOnRequestBefore

	bakLoad := _this.Control.OnLoad
	_this.Control.OnLoad = func() {
		if bakLoad != nil {
			bakLoad()
		}
		_this._mb = new(free4x64).init(&_this.Control)
		_this.mbInit()
	}
	_this.EvRequestBefore = append(_this.EvRequestBefore, _this.loadRes)
	return _this
}

func (_this *MiniblinkBrowser) loadRes(_ *MiniblinkBrowser, e RequestEvArgs) {
	if len(_this.ResourceLoader) == 0 {
		return
	}
	url, err := url2.Parse(e.GetUrl())
	if err != nil {
		return
	}
	host := strings.ToLower(url.Host)
	for i := range _this.ResourceLoader {
		loader := _this.ResourceLoader[i]
		if strings.HasPrefix(strings.ToLower(loader.Domain()), host) == false {
			continue
		}
		data := loader.ByUri(url)
		if data != nil {
			e.SetData(data)
			break
		}
	}
}

func (_this *MiniblinkBrowser) mbInit() {
	_this._mb.SetOnRequest(func(args RequestEvArgs) {
		if _this.OnRequestBefore != nil {
			_this.OnRequestBefore(args)
		}
	})

	if _this._initUri != "" {
		_this.LoadUri(_this._initUri)
	}
}

func (_this *MiniblinkBrowser) LoadUri(uri string) {
	if _this._mb != nil {
		_this._mb.LoadUri(uri)
	} else {
		_this._initUri = uri
	}
}
