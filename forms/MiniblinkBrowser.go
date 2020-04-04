package forms

import (
	url2 "net/url"
	mb "qq2564874169/goMiniblink"
	plat "qq2564874169/goMiniblink/platform"
	"strings"
)

type MiniblinkBrowser struct {
	BaseControl
	impl plat.IMiniblink

	ResourceLoader []ILoadResource

	EvRequest []func(e mb.RequestEvArgs)
	OnRequest func(e mb.RequestEvArgs)
}

func (_this *MiniblinkBrowser) Init() *MiniblinkBrowser {
	_this.impl = Provider.NewMiniblink()
	_this.BaseControl.Init(_this.impl)
	_this.BaseControl.SetBgColor(-1)
	_this.setCallback()
	_this.EvRequest = append(_this.EvRequest, _this.loadRes)
	return _this
}

func (_this *MiniblinkBrowser) BindFunc(name string, fn mb.GoFuncFn, state interface{}) {
	_this.impl.BindGoFunc(mb.GoFunc{
		Name:  name,
		State: state,
		Fn:    fn,
	})
}

func (_this *MiniblinkBrowser) loadRes(e mb.RequestEvArgs) {
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

func (_this *MiniblinkBrowser) setCallback() {
	_this.OnRequest = _this.defOnRequest
	_this.impl.SetOnRequest(func(e mb.RequestEvArgs) {
		if _this.OnRequest != nil {
			_this.OnRequest(e)
		}
	})
}

func (_this *MiniblinkBrowser) LoadUri(uri string) {
	_this.impl.LoadUri(uri)
}
