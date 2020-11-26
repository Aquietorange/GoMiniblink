package GoMiniblink

import (
	"strconv"
)

var (
	views   map[wkeHandle]Miniblink
	keepRef map[string]interface{}
)

func init() {
	keepRef = make(map[string]interface{})
	views = make(map[wkeHandle]Miniblink)
	mbApi = new(winFreeApi).init()
}

func createWebView(miniblink Miniblink) wkeHandle {
	wke := mbApi.wkeCreateWebView()
	views[wke] = miniblink
	return wke
}

func destroyWebView(handle wkeHandle) {
	if _, ok := views[handle]; ok {
		mbApi.wkeDestroyWebView(handle)
		delete(views, handle)
	}
}

func BindJsFunc(fn JsFnBinding) {
	fn.core = func(es jsExecState, param uintptr) uintptr {
		handle := mbApi.jsGetWebView(es)
		if mb, ok := views[handle]; ok {
			arglen := mbApi.jsArgCount(es)
			args := make([]interface{}, arglen)
			for i := uint32(0); i < arglen; i++ {
				value := mbApi.jsArg(es, i)
				args[i] = toGoValue(mb, es, value)
			}
			g := keepRef["__mbJsFn_"+strconv.FormatUint(uint64(param), 10)].(JsFnBinding)
			rs := g.Call(mb, args)
			if rs != nil {
				return uintptr(toJsValue(mb, es, rs))
			}
		}
		return 0
	}
	pm := seq()
	mbApi.wkeJsBindFunction(fn.Name, fn.core, uintptr(pm), 0)
	keepRef["__mbJsFn_"+strconv.FormatUint(pm, 10)] = fn
}
