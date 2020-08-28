package forms

import (
	"fmt"
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
	broName := ctx.Param[0].(string)
	fnName := ctx.Param[1].(string)
	rsName := ctx.Param[2].(string)
	bro := broMap[broName]
	fn := bro.jsfns[fnName]
	rs := fn.OnExecute(ctx.Param[3:])
	bro.impl.SetWindowProp(rsName, rs)
	return nil
}

type MiniblinkBrowser struct {
	BaseControl
	impl      p.IMiniblink
	name      string
	jsIsReady bool
	frames    []g.FrameContext
	jsfns     map[string]g.JsFuncBinding

	ResourceLoader []ILoadResource

	EvRequestBefore []func(e g.RequestBeforeEvArgs)
	OnRequestBefore func(e g.RequestBeforeEvArgs)
	EvJsReady       []func(e g.JsReadyEvArgs)
	OnJsReady       func(e g.JsReadyEvArgs)
}

func (_this *MiniblinkBrowser) Init() *MiniblinkBrowser {
	_this.jsfns = make(map[string]g.JsFuncBinding)
	_this.name = strconv.FormatInt(int64(uintptr(unsafe.Pointer(_this))), 10)
	_this.impl = Provider.NewMiniblink()
	_this.BaseControl.Init(_this.impl)
	_this.BaseControl.SetBgColor(-1)
	_this.setCallback()
	_this.EvRequestBefore = append(_this.EvRequestBefore, _this.loadRes)
	return _this
}

func (_this *MiniblinkBrowser) BindJsFunc(name string, fn g.GoFn, state interface{}) {
	_this.jsfns[name] = g.JsFuncBinding{
		Name:  name,
		State: state,
		Fn:    fn,
	}
	if _this.jsIsReady {
		for _, f := range _this.frames {
			f.RunJs(_this.getJsBindingScript(false))
		}
		_this.impl.RunJs(_this.getJsBindingScript(true))
	}
}

func (_this *MiniblinkBrowser) getJsBindingScript(isMain bool) string {
	rsName := "rs" + strconv.FormatInt(g.NewId(), 10)
	call := callFnName
	if isMain == false {
		call = "window.top['" + call + "']"
	}
	var list []string
	for k, _ := range _this.jsfns {
		js := `window.%s=function(){
               var rs=%q;
               var arr = Array.prototype.slice.call(arguments);
               var args = [%q,%q,rs].concat(arr);
               %s.apply(null,args);
               return window.top[rs];
           };`
		js = fmt.Sprintf(js, k, rsName, _this.name, k, call)
		list = append(list, js)
	}
	return strings.Join(list, ";")
}

func (_this *MiniblinkBrowser) loadRes(e g.RequestBeforeEvArgs) {
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
	_this.OnRequestBefore = _this.defOnRequest
	_this.OnJsReady = _this.defOnJsReady

	_this.impl.SetOnJsReady(func(e g.JsReadyEvArgs) {
		_this.jsIsReady = true
		if e.Frame().IsMain() == false {
			_this.frames = append(_this.frames, e.Frame())
		}
		if _this.OnJsReady != nil {
			_this.OnJsReady(e)
		}
	})
	_this.impl.SetOnRequest(func(e g.RequestBeforeEvArgs) {
		if _this.OnRequestBefore != nil {
			_this.OnRequestBefore(e)
		}
	})
}

func (_this *MiniblinkBrowser) LoadUri(uri string) {
	_this.impl.LoadUri(uri)
}
