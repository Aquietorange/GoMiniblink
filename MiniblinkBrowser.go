package GoMiniblink

import (
	"image"
	url2 "net/url"
	"reflect"
	"strings"

	cs "github.com/hujun528/GoMiniblink/forms/controls"
)

type MiniblinkBrowser struct {
	cs.Control
	core   Miniblink
	fnlist map[string]reflect.Value

	EvRequestBefore map[string]func(sender *MiniblinkBrowser, e RequestBeforeEvArgs)
	OnRequestBefore func(e RequestBeforeEvArgs)

	EvJsReady map[string]func(sender *MiniblinkBrowser, e JsReadyEvArgs)
	OnJsReady func(e JsReadyEvArgs)

	EvConsole map[string]func(sender *MiniblinkBrowser, e ConsoleEvArgs)
	OnConsole func(e ConsoleEvArgs)

	EvDocumentReady map[string]func(sender *MiniblinkBrowser, e DocumentReadyEvArgs)
	OnDocumentReady func(e DocumentReadyEvArgs)

	EvPaintUpdated map[string]func(sender *MiniblinkBrowser, e PaintUpdatedEvArgs)
	OnPaintUpdated func(e PaintUpdatedEvArgs)

	ResourceLoader []LoadResource
}

func (_this *MiniblinkBrowser) Init() *MiniblinkBrowser {
	_this.Control.Init()
	_this.EvRequestBefore = make(map[string]func(*MiniblinkBrowser, RequestBeforeEvArgs))
	_this.EvJsReady = make(map[string]func(*MiniblinkBrowser, JsReadyEvArgs))
	_this.EvConsole = make(map[string]func(*MiniblinkBrowser, ConsoleEvArgs))
	_this.EvDocumentReady = make(map[string]func(*MiniblinkBrowser, DocumentReadyEvArgs))

	_this.OnRequestBefore = _this.defOnRequestBefore
	_this.OnJsReady = _this.defOnJsReady
	_this.OnConsole = _this.defOnConsole
	_this.OnDocumentReady = _this.defOnDocumentReady
	_this.OnPaintUpdated = _this.defOnPaintUpdated

	_this.EvRequestBefore["__goMiniblink"] = _this.loadRes
	_this.EvDestroy["__goMiniblink"] = _this.onClosed
	_this.core = new(freeMiniblink).init(&_this.Control)
	_this.mbInit()
	return _this
}

func (_this *MiniblinkBrowser) onClosed(_ cs.GUI) {
	destroyWebView(_this.core.GetHandle())
}

func (_this *MiniblinkBrowser) loadRes(_ *MiniblinkBrowser, e RequestBeforeEvArgs) {
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
	_this.core.SetOnRequestBefore(func(args RequestBeforeEvArgs) {
		if _this.OnRequestBefore != nil {
			_this.OnRequestBefore(args)
		}
	})
	_this.core.SetOnJsReady(func(args JsReadyEvArgs) {
		if _this.OnJsReady != nil {
			_this.OnJsReady(args)
		}
	})
	_this.core.SetOnConsole(func(args ConsoleEvArgs) {
		if _this.OnConsole != nil {
			_this.OnConsole(args)
		}
	})
	_this.core.SetOnDocumentReady(func(args DocumentReadyEvArgs) {
		if _this.OnDocumentReady != nil {
			_this.OnDocumentReady(args)
		}
	})
	_this.core.SetOnPaintUpdated(func(args PaintUpdatedEvArgs) {
		if _this.OnPaintUpdated != nil {
			_this.OnPaintUpdated(args)
		}
	})
}

func (_this *MiniblinkBrowser) SetProxy(info ProxyInfo) {
	_this.core.SetProxy(info)
}

func (_this *MiniblinkBrowser) LoadUri(uri string) {
	_this.core.LoadUri(uri)
}

func (_this *MiniblinkBrowser) WkeRunMessageLoop() {
	_this.core.WkeRunMessageLoop()
}

func (_this *MiniblinkBrowser) SetDebugConfig(debugString string, param string) {
	_this.core.SetDebugConfig(debugString, param)
}

func (_this *MiniblinkBrowser) JsFunc(name string, fn GoFn, state interface{}) {
	_this.core.JsFunc(name, fn, state)
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

func (_this *MiniblinkBrowser) CallJsFunc(name string, param ...interface{}) interface{} {
	return _this.core.CallJsFunc(name, param)
}

func (_this *MiniblinkBrowser) ToBitmap() *image.RGBA {
	return _this.core.ToBitmap()
}

func (_this *MiniblinkBrowser) GetMiniblinkHandle() uintptr {
	return uintptr(_this.core.GetHandle())
}

func (_this *MiniblinkBrowser) MouseIsEnable() bool {
	return _this.core.MouseIsEnable()
}

func (_this *MiniblinkBrowser) MouseEnable(b bool) {
	_this.core.MouseEnable(b)
}

func (_this *MiniblinkBrowser) SetBmpPaintMode(b bool) {
	_this.core.SetBmpPaintMode(b)
}
