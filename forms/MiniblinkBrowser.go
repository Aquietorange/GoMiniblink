package forms

import (
	u "net/url"
	g "qq2564874169/goMiniblink"
	p "qq2564874169/goMiniblink/platform"
	m "qq2564874169/goMiniblink/platform/miniblink"
	"strconv"
	"strings"
	"unsafe"
)

var callFnName = "fn" + strconv.FormatInt(g.NewId(), 10)
var broMap = make(map[string]MiniblinkBrowser)

func init() {
	m.BindJsFunc(&g.JsFuncBinding{
		Name: callFnName,
		Fn:   execGoFunc,
	})
}

func execGoFunc(ctx g.GoFnContext) interface{} {
	//todo
}

type MiniblinkBrowser struct {
	BaseControl
	impl p.IMiniblink
	name string

	ResourceLoader []ILoadResource

	EvRequest []func(e g.RequestEvArgs)
	OnRequest func(e g.RequestEvArgs)
}

func (_this *MiniblinkBrowser) Init() *MiniblinkBrowser {
	_this.name = strconv.FormatInt(int64(uintptr(unsafe.Pointer(_this))), 10)
	_this.impl = Provider.NewMiniblink()
	_this.BaseControl.Init(_this.impl)
	_this.BaseControl.SetBgColor(-1)
	_this.setCallback()
	_this.EvRequest = append(_this.EvRequest, _this.loadRes)
	return _this
}

func (_this *MiniblinkBrowser) BindJsFunc(name string, fn g.GoFn, state interface{}) {
	_this.impl.BindJsFunc(g.JsFuncBinding{
		Name:  name,
		State: state,
		Fn:    fn,
	})
}

func (_this *MiniblinkBrowser) loadRes(e g.RequestEvArgs) {
	if len(_this.ResourceLoader) == 0 {
		return
	}
	url, err := u.Parse(e.GetUrl())
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
	_this.impl.SetOnRequest(func(e g.RequestEvArgs) {
		if _this.OnRequest != nil {
			_this.OnRequest(e)
		}
	})
}

func (_this *MiniblinkBrowser) LoadUri(uri string) {
	_this.impl.LoadUri(uri)
}
