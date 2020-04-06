package GoMiniblink

import "unsafe"

var (
	views   map[wkeHandle]IMiniblink
	keepRef map[uintptr]interface{}
)

func init() {
	keepRef = make(map[uintptr]interface{})
	views = make(map[wkeHandle]IMiniblink)
	mbApi = new(freeApiForWindows).init()
}

func createWebView(miniblink IMiniblink) wkeHandle {
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

func BindGoFunc(fn GoFunc) {
	fn.jsFunc = func(es jsExecState, param uintptr) jsValue {
		handle := mbApi.jsGetWebView(es)
		if mb, ok := views[handle]; ok {
			arglen := mbApi.jsArgCount(es)
			args := make([]interface{}, arglen)
			for i := uint32(0); i < arglen; i++ {
				value := mbApi.jsArg(es, i)
				args[i] = toGoValue(mb, es, value)
			}
			g := *((*GoFunc)(unsafe.Pointer(param)))
			rs := g.Call(mb, args)
			return toJsValue(mb, es, rs)
		}
		return mbApi.jsUndefined()
	}
	pm := unsafe.Pointer(&fn)
	mbApi.wkeJsBindFunction(fn.Name, fn.jsFunc, pm, 0)
	keepRef[uintptr(seq())] = fn
}
