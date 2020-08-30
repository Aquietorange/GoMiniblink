package GoMiniblink

import (
	url2 "net/url"
	c "qq2564874169/goMiniblink/forms/controls"
	"strings"
)

type MiniblinkBrowser struct {
	c.Control
	_mb      Miniblink
	_initUri string

	EvRequestBefore []func(sender *MiniblinkBrowser, e RequestEvArgs)
	OnRequestBefore func(e RequestEvArgs)

	EvJsReady []func(sender *MiniblinkBrowser, e JsReadyEvArgs)
	OnJsReady func(e JsReadyEvArgs)

	EvConsole []func(sender *MiniblinkBrowser, e ConsoleEvArgs)
	OnConsole func(e ConsoleEvArgs)

	ResourceLoader []LoadResource
}

func (_this *MiniblinkBrowser) Init() *MiniblinkBrowser {
	_this.Control.Init()
	_this.OnRequestBefore = _this.defOnRequestBefore
	_this.OnJsReady = _this.defOnJsReady
	_this.OnConsole = _this.defOnConsole

	bakLoad := _this.Control.OnLoad
	_this.Control.OnLoad = func() {
		_this._mb = new(freeMiniblink).init(&_this.Control)
		_this.mbInit()
		if bakLoad != nil {
			bakLoad()
		}
	}
	_this.EvRequestBefore = append(_this.EvRequestBefore, _this.loadRes)
	return _this
}

func (_this *MiniblinkBrowser) loadRes(_ *MiniblinkBrowser, e RequestEvArgs) {
	if len(_this.ResourceLoader) == 0 {
		return
	}
	url, err := url2.Parse(e.Url())
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
	_this._mb.SetOnJsReady(func(args JsReadyEvArgs) {
		if _this.OnJsReady != nil {
			_this.OnJsReady(args)
		}
	})
	_this._mb.SetOnConsole(func(args ConsoleEvArgs) {
		if _this.OnConsole != nil {
			_this.OnConsole(args)
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

func (_this *MiniblinkBrowser) BindJsFunc(name string, fn GoFn, state interface{}) {
	_this._mb.JsFunc(name, fn, state)
}
