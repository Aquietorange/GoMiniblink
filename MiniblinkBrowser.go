package GoMiniblink

import (
	url2 "net/url"
	c "qq2564874169/goMiniblink/forms/controls"
	"reflect"
	"strings"
)

type MiniblinkBrowser struct {
	c.Control
	_mb     Miniblink
	_fnlist map[string]reflect.Value

	EvRequestBefore map[string]func(sender interface{}, e RequestEvArgs)
	OnRequestBefore func(e RequestEvArgs)

	EvJsReady map[string]func(sender interface{}, e JsReadyEvArgs)
	OnJsReady func(e JsReadyEvArgs)

	EvConsole map[string]func(sender interface{}, e ConsoleEvArgs)
	OnConsole func(e ConsoleEvArgs)

	ResourceLoader []LoadResource
}

func (_this *MiniblinkBrowser) Init() *MiniblinkBrowser {
	_this.Control.Init()
	_this.EvRequestBefore = make(map[string]func(interface{}, RequestEvArgs))
	_this.EvJsReady = make(map[string]func(interface{}, JsReadyEvArgs))
	_this.EvConsole = make(map[string]func(interface{}, ConsoleEvArgs))
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
	_this.EvRequestBefore["load_resource"] = _this.loadRes
	return _this
}

func (_this *MiniblinkBrowser) loadRes(_ interface{}, e RequestEvArgs) {
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
}

func (_this *MiniblinkBrowser) LoadUri(uri string) {
	_this._mb.LoadUri(uri)
}

func (_this *MiniblinkBrowser) JsFunc(name string, fn GoFn, state interface{}) {
	_this._mb.JsFunc(name, fn, state)
}

func (_this *MiniblinkBrowser) CallJsFunc(name string, param ...interface{}) interface{} {
	return _this._mb.CallJsFunc(name, param)
}

func (_this *MiniblinkBrowser) JsFuncEx(name string, fn interface{}) {
	p := reflect.TypeOf(fn)
	if p.Kind() != reflect.Func {
		return
	}
	_this.JsFunc(name, func(ctx GoFnContext) interface{} {
		rt := reflect.TypeOf(ctx.State)
		rv := reflect.ValueOf(ctx.State)
		var args []reflect.Value
		for i := 0; i < rt.NumIn() && i < len(ctx.Param); i++ {
			args = append(args, reflect.ValueOf(ctx.Param[i]))
		}
		rs := rv.Call(args)
		if rt.NumOut() > 0 {
			return rs[0].Interface()
		}
		return nil
	}, fn)
}
