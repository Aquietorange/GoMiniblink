package GoMiniblink

import (
	"fmt"
	"golang.org/x/sys/windows"
	"strconv"
	"syscall"
	"unsafe"
)

func _toInt64(low, high int32) int64 {
	var l = int64(high)<<32 + int64(low)
	return *((*int64)(unsafe.Pointer(&l)))
}

func _toJsValue(low, high uintptr) jsValue {
	return jsValue(uintptr(_toInt64(int32(low), int32(high))))
}

func _toLH(value jsValue) (low, high int32) {
	if is64 {
		return 0, 0
	}
	return int32(int64(value)), int32(int64(value) >> 32 & 0xffffffff)
}

var is64 bool

type winFreeApi struct {
	_dll *windows.LazyDLL

	_wkeInitialize               *windows.LazyProc
	_wkeCreateWebView            *windows.LazyProc
	_wkeSetHandle                *windows.LazyProc
	_wkeOnPaintBitUpdated        *windows.LazyProc
	_wkeLoadURL                  *windows.LazyProc
	_wkeResize                   *windows.LazyProc
	_wkeNetOnResponse            *windows.LazyProc
	_wkeOnLoadUrlBegin           *windows.LazyProc
	_wkePaint                    *windows.LazyProc
	_wkeGetWidth                 *windows.LazyProc
	_wkeGetHeight                *windows.LazyProc
	_wkeFireMouseEvent           *windows.LazyProc
	_wkeFireMouseWheelEvent      *windows.LazyProc
	_wkeGetCursorInfoType        *windows.LazyProc
	_wkeFireKeyUpEvent           *windows.LazyProc
	_wkeFireKeyDownEvent         *windows.LazyProc
	_wkeFireKeyPressEvent        *windows.LazyProc
	_wkeGetCaretRect             *windows.LazyProc
	_wkeSetFocus                 *windows.LazyProc
	_wkeNetGetRequestMethod      *windows.LazyProc
	_wkeNetSetData               *windows.LazyProc
	_wkeNetCancelRequest         *windows.LazyProc
	_wkeJsBindFunction           *windows.LazyProc
	_jsArgCount                  *windows.LazyProc
	_jsArg                       *windows.LazyProc
	_jsTypeOf                    *windows.LazyProc
	_jsToTempString              *windows.LazyProc
	_jsToDoubleString            *windows.LazyProc
	_jsToInt                     *windows.LazyProc
	_jsToBoolean                 *windows.LazyProc
	_jsGetLength                 *windows.LazyProc
	_jsGetAt                     *windows.LazyProc
	_jsGetKeys                   *windows.LazyProc
	_jsGet                       *windows.LazyProc
	_jsSetGlobal                 *windows.LazyProc
	_jsGetGlobal                 *windows.LazyProc
	_wkeGlobalExec               *windows.LazyProc
	_jsCall                      *windows.LazyProc
	_jsUndefined                 *windows.LazyProc
	_jsInt                       *windows.LazyProc
	_jsBoolean                   *windows.LazyProc
	_jsDoubleString              *windows.LazyProc
	_jsString                    *windows.LazyProc
	_jsEmptyArray                *windows.LazyProc
	_jsSetLength                 *windows.LazyProc
	_jsSetAt                     *windows.LazyProc
	_jsFunction                  *windows.LazyProc
	_jsEmptyObject               *windows.LazyProc
	_jsSet                       *windows.LazyProc
	_wkeDestroyWebView           *windows.LazyProc
	_jsGetWebView                *windows.LazyProc
	_wkeKillFocus                *windows.LazyProc
	_jsEval                      *windows.LazyProc
	_wkeOnDidCreateScriptContext *windows.LazyProc
	_wkeIsMainFrame              *windows.LazyProc
	_wkeGetFrameUrl              *windows.LazyProc
	_wkeIsWebRemoteFrame         *windows.LazyProc
	_wkeGetGlobalExecByFrame     *windows.LazyProc
	_wkeOnConsole                *windows.LazyProc
	_wkeGetString                *windows.LazyProc
	_wkeNetSetHTTPHeaderField    *windows.LazyProc
	_wkeNetChangeRequestUrl      *windows.LazyProc
	_wkeNetHookRequest           *windows.LazyProc
	_wkeNetHoldJobToAsynCommit   *windows.LazyProc
	_wkeNetContinueJob           *windows.LazyProc
	_wkeOnLoadUrlEnd             *windows.LazyProc
	_wkeOnLoadUrlFail            *windows.LazyProc
	_wkeNetGetUrlByJob           *windows.LazyProc
	_wkeNetGetMIMEType           *windows.LazyProc
	_wkeNetSetMIMEType           *windows.LazyProc
	_wkeNetGetRawResponseHead    *windows.LazyProc
	_wkeOnDocumentReady2         *windows.LazyProc
}

func (_this *winFreeApi) init() *winFreeApi {
	is64 = unsafe.Sizeof(uintptr(0)) == 8
	var lib *windows.LazyDLL
	if is64 {
		lib = windows.NewLazyDLL("miniblink_x64.dll")
	} else {
		lib = windows.NewLazyDLL("miniblink_x86.dll")
	}
	_this._wkeOnDocumentReady2 = lib.NewProc("wkeOnDocumentReady2")
	_this._wkeNetGetRawResponseHead = lib.NewProc("wkeNetGetRawResponseHead")
	_this._wkeNetSetMIMEType = lib.NewProc("wkeNetSetMIMEType")
	_this._wkeNetGetMIMEType = lib.NewProc("wkeNetGetMIMEType")
	_this._wkeNetGetUrlByJob = lib.NewProc("wkeNetGetUrlByJob")
	_this._wkeOnLoadUrlFail = lib.NewProc("wkeOnLoadUrlFail")
	_this._wkeOnLoadUrlEnd = lib.NewProc("wkeOnLoadUrlEnd")
	_this._wkeNetContinueJob = lib.NewProc("wkeNetContinueJob")
	_this._wkeNetHoldJobToAsynCommit = lib.NewProc("wkeNetHoldJobToAsynCommit")
	_this._wkeNetHookRequest = lib.NewProc("wkeNetHookRequest")
	_this._wkeNetChangeRequestUrl = lib.NewProc("wkeNetChangeRequestUrl")
	_this._wkeNetSetHTTPHeaderField = lib.NewProc("wkeNetSetHTTPHeaderField")
	_this._wkeGetString = lib.NewProc("wkeGetString")
	_this._wkeOnConsole = lib.NewProc("wkeOnConsole")
	_this._wkeGetGlobalExecByFrame = lib.NewProc("wkeGetGlobalExecByFrame")
	_this._wkeIsWebRemoteFrame = lib.NewProc("wkeIsWebRemoteFrame")
	_this._wkeGetFrameUrl = lib.NewProc("wkeGetFrameUrl")
	_this._wkeIsMainFrame = lib.NewProc("wkeIsMainFrame")
	_this._wkeOnDidCreateScriptContext = lib.NewProc("wkeOnDidCreateScriptContext")
	_this._jsEval = lib.NewProc("jsEval")
	_this._wkeKillFocus = lib.NewProc("wkeKillFocus")
	_this._jsToInt = lib.NewProc("jsToInt")
	_this._jsSet = lib.NewProc("jsSet")
	_this._jsEmptyObject = lib.NewProc("jsEmptyObject")
	_this._jsFunction = lib.NewProc("jsFunction")
	_this._jsSetAt = lib.NewProc("jsSetAt")
	_this._jsSetLength = lib.NewProc("jsSetLength")
	_this._jsEmptyArray = lib.NewProc("jsEmptyArray")
	_this._jsString = lib.NewProc("jsString")
	_this._jsDoubleString = lib.NewProc("jsDoubleString")
	_this._jsBoolean = lib.NewProc("jsBoolean")
	_this._jsInt = lib.NewProc("jsInt")
	_this._jsUndefined = lib.NewProc("jsUndefined")
	_this._jsCall = lib.NewProc("jsCall")
	_this._wkeGlobalExec = lib.NewProc("wkeGlobalExec")
	_this._jsGetGlobal = lib.NewProc("jsGetGlobal")
	_this._jsSetGlobal = lib.NewProc("jsSetGlobal")
	_this._jsGet = lib.NewProc("jsGet")
	_this._jsGetKeys = lib.NewProc("jsGetKeys")
	_this._jsGetAt = lib.NewProc("jsGetAt")
	_this._jsGetLength = lib.NewProc("jsGetLength")
	_this._jsToBoolean = lib.NewProc("jsToBoolean")
	_this._jsToDoubleString = lib.NewProc("jsToDoubleString")
	_this._jsToTempString = lib.NewProc("jsToTempString")
	_this._jsTypeOf = lib.NewProc("jsTypeOf")
	_this._jsArg = lib.NewProc("jsArg")
	_this._jsArgCount = lib.NewProc("jsArgCount")
	_this._wkeJsBindFunction = lib.NewProc("wkeJsBindFunction")
	_this._wkeNetCancelRequest = lib.NewProc("wkeNetCancelRequest")
	_this._wkeNetSetData = lib.NewProc("wkeNetSetData")
	_this._wkeNetGetRequestMethod = lib.NewProc("wkeNetGetRequestMethod")
	_this._wkeFireKeyPressEvent = lib.NewProc("wkeFireKeyPressEvent")
	_this._wkeFireKeyUpEvent = lib.NewProc("wkeFireKeyUpEvent")
	_this._wkeFireKeyDownEvent = lib.NewProc("wkeFireKeyDownEvent")
	_this._wkeGetCursorInfoType = lib.NewProc("wkeGetCursorInfoType")
	_this._wkeFireMouseWheelEvent = lib.NewProc("wkeFireMouseWheelEvent")
	_this._wkeFireMouseEvent = lib.NewProc("wkeFireMouseEvent")
	_this._wkeGetHeight = lib.NewProc("wkeGetHeight")
	_this._wkeGetWidth = lib.NewProc("wkeGetWidth")
	_this._wkePaint = lib.NewProc("wkePaint")
	_this._wkeOnLoadUrlBegin = lib.NewProc("wkeOnLoadUrlBegin")
	_this._wkeNetOnResponse = lib.NewProc("wkeNetOnResponse")
	_this._wkeLoadURL = lib.NewProc("wkeLoadURL")
	_this._wkeResize = lib.NewProc("wkeResize")
	_this._wkeOnPaintBitUpdated = lib.NewProc("wkeOnPaintBitUpdated")
	_this._wkeSetHandle = lib.NewProc("wkeSetHandle")
	_this._wkeCreateWebView = lib.NewProc("wkeCreateWebView")
	_this._wkeInitialize = lib.NewProc("wkeInitialize")
	_this._wkeGetCaretRect = lib.NewProc("wkeGetCaretRect2")
	_this._wkeSetFocus = lib.NewProc("wkeSetFocus")
	_this._wkeDestroyWebView = lib.NewProc("wkeDestroyWebView")
	_this._jsGetWebView = lib.NewProc("jsGetWebView")

	ret, _, err := _this._wkeInitialize.Call()
	if ret == 0 {
		fmt.Println("初始化失败", err)
	}
	return _this
}

func (_this *winFreeApi) wkeOnDocumentReady(wke wkeHandle, callback wkeDocumentReady2Callback, param uintptr) {
	_this._wkeOnDocumentReady2.Call(uintptr(wke), syscall.NewCallbackCDecl(callback), param)
}

func (_this *winFreeApi) wkeNetGetRawResponseHead(job wkeNetJob) map[string]string {
	r, _, _ := _this._wkeNetGetRawResponseHead.Call(uintptr(job))
	var list []string
	slist := *((*wkeSlist)(unsafe.Pointer(r)))
	for slist.str != 0 {
		list = append(list, ptrToUtf8(slist.str))
		if slist.next == 0 {
			break
		} else {
			slist = *((*wkeSlist)(unsafe.Pointer(slist.next)))
		}
	}
	hMap := make(map[string]string)
	for i := 0; i < len(list); i += 2 {
		hMap[list[i]] = list[i+1]
	}
	return hMap
}

func (_this *winFreeApi) wkeNetSetMIMEType(job wkeNetJob, mime string) {
	p := toCallStr(mime)
	_this._wkeNetSetMIMEType.Call(uintptr(job), uintptr(unsafe.Pointer(&p[0])))
}

func (_this *winFreeApi) wkeNetGetMIMEType(job wkeNetJob) string {
	r, _, _ := _this._wkeNetGetMIMEType.Call(uintptr(job))
	return ptrToUtf8(r)
}

func (_this *winFreeApi) wkeNetGetUrlByJob(job wkeNetJob) string {
	r, _, _ := _this._wkeNetGetUrlByJob.Call(uintptr(job))
	return ptrToUtf8(r)
}

func (_this *winFreeApi) wkeOnLoadUrlFail(wke wkeHandle, callback wkeLoadUrlFailCallback, param uintptr) {
	_this._wkeOnLoadUrlFail.Call(uintptr(wke), syscall.NewCallbackCDecl(callback), param)
}

func (_this *winFreeApi) wkeOnLoadUrlEnd(wke wkeHandle, callback wkeLoadUrlEndCallback, param uintptr) {
	_this._wkeOnLoadUrlEnd.Call(uintptr(wke), syscall.NewCallbackCDecl(callback), param)
}

func (_this *winFreeApi) wkeNetContinueJob(job wkeNetJob) {
	_this._wkeNetContinueJob.Call(uintptr(job))
}

func (_this *winFreeApi) wkeNetHoldJobToAsynCommit(job wkeNetJob) {
	_this._wkeNetHoldJobToAsynCommit.Call(uintptr(job))
}

func (_this *winFreeApi) wkeNetHookRequest(job wkeNetJob) {
	_this._wkeNetHookRequest.Call(uintptr(job))
}

func (_this *winFreeApi) wkeNetChangeRequestUrl(job wkeNetJob, url string) {
	p := toCallStr(url)
	_this._wkeNetChangeRequestUrl.Call(uintptr(job), uintptr(unsafe.Pointer(&p[0])))
}

func (_this *winFreeApi) wkeNetSetHTTPHeaderField(job wkeNetJob, name, value string) {
	np := toCallStr(name)
	vp := toCallStr(value)
	_this._wkeNetSetHTTPHeaderField.Call(uintptr(job), uintptr(unsafe.Pointer(&np[0])), uintptr(unsafe.Pointer(&vp[0])))
}

func (_this *winFreeApi) wkeGetString(str wkeString) string {
	r, _, _ := _this._wkeGetString.Call(uintptr(str))
	return ptrToUtf8(r)
}

func (_this *winFreeApi) wkeOnConsole(wke wkeHandle, callback wkeConsoleCallback, param uintptr) {
	_this._wkeOnConsole.Call(uintptr(wke), syscall.NewCallbackCDecl(callback), param)
}

func (_this *winFreeApi) wkeGetGlobalExecByFrame(wke wkeHandle, frame wkeFrame) jsExecState {
	r, _, _ := _this._wkeGetGlobalExecByFrame.Call(uintptr(wke), uintptr(frame))
	return jsExecState(r)
}

func (_this *winFreeApi) wkeIsWebRemoteFrame(wke wkeHandle, frame wkeFrame) bool {
	r, _, _ := _this._wkeIsWebRemoteFrame.Call(uintptr(wke), uintptr(frame))
	return r != 0
}

func (_this *winFreeApi) wkeGetFrameUrl(wke wkeHandle, frame wkeFrame) string {
	r, _, _ := _this._wkeGetFrameUrl.Call(uintptr(wke), uintptr(frame))
	return ptrToUtf8(r)
}

func (_this *winFreeApi) wkeIsMainFrame(wke wkeHandle, frame wkeFrame) bool {
	r, _, _ := _this._wkeIsMainFrame.Call(uintptr(wke), uintptr(frame))
	return r != 0
}

func (_this *winFreeApi) wkeOnDidCreateScriptContext(wke wkeHandle, callback wkeDidCreateScriptContextCallback, param uintptr) {
	_this._wkeOnDidCreateScriptContext.Call(uintptr(wke), syscall.NewCallbackCDecl(callback), param)
}

func (_this *winFreeApi) wkeKillFocus(wke wkeHandle) {
	_this._wkeKillFocus.Call(uintptr(wke))
}

func (_this *winFreeApi) jsGetWebView(es jsExecState) wkeHandle {
	r, _, _ := _this._jsGetWebView.Call(uintptr(es))
	return wkeHandle(r)
}

func (_this *winFreeApi) wkeDestroyWebView(wke wkeHandle) {
	_this._wkeDestroyWebView.Call(uintptr(wke))
}

func (_this *winFreeApi) jsSet(es jsExecState, obj jsValue, name string, value jsValue) {
	ptr := []byte(name)
	if is64 {
		_this._jsSet.Call(uintptr(es), uintptr(obj), uintptr(unsafe.Pointer(&ptr[0])), uintptr(value))
	} else {
		ol, oh := _toLH(obj)
		vl, vh := _toLH(value)
		_this._jsSet.Call(uintptr(es), uintptr(ol), uintptr(oh), uintptr(unsafe.Pointer(&ptr[0])), uintptr(vl), uintptr(vh))
	}
}

func (_this *winFreeApi) jsEmptyObject(es jsExecState) jsValue {
	if is64 {
		r, _, _ := _this._jsEmptyObject.Call(uintptr(es))
		return jsValue(r)
	}
	l, h, _ := _this._jsEmptyObject.Call(uintptr(es))
	return _toJsValue(l, h)
}

func (_this *winFreeApi) jsFunction(es jsExecState, data *jsData) jsValue {
	if is64 {
		r, _, _ := _this._jsFunction.Call(uintptr(es), uintptr(unsafe.Pointer(data)))
		return jsValue(r)
	}
	l, h, _ := _this._jsFunction.Call(uintptr(es), uintptr(unsafe.Pointer(data)))
	return _toJsValue(l, h)
}

func (_this *winFreeApi) jsSetAt(es jsExecState, array jsValue, index uint32, value jsValue) {
	if is64 {
		_this._jsSetAt.Call(uintptr(es), uintptr(array), uintptr(index), uintptr(value))
	} else {
		al, ah := _toLH(array)
		vl, vh := _toLH(value)
		_this._jsSetAt.Call(uintptr(es), uintptr(al), uintptr(ah), uintptr(index), uintptr(vl), uintptr(vh))
	}
}

func (_this *winFreeApi) jsSetLength(es jsExecState, array jsValue, length uint32) {
	if is64 {
		_this._jsSetLength.Call(uintptr(es), uintptr(array), uintptr(length))
	} else {
		l, h := _toLH(array)
		_this._jsSetLength.Call(uintptr(es), uintptr(l), uintptr(h), uintptr(length))
	}
}

func (_this *winFreeApi) jsEmptyArray(es jsExecState) jsValue {
	if is64 {
		r, _, _ := _this._jsEmptyArray.Call(uintptr(es))
		return jsValue(r)
	}
	l, h, _ := _this._jsEmptyArray.Call(uintptr(es))
	return _toJsValue(l, h)
}

func (_this *winFreeApi) jsString(es jsExecState, value string) jsValue {
	ptr := toCallStr(value)
	if is64 {
		r, _, _ := _this._jsString.Call(uintptr(es), uintptr(unsafe.Pointer(&ptr[0])))
		return jsValue(r)
	}
	l, h, _ := _this._jsString.Call(uintptr(es), uintptr(unsafe.Pointer(&ptr[0])))
	return _toJsValue(l, h)
}

func (_this *winFreeApi) jsDouble(value float64) jsValue {
	ptr := toCallStr(strconv.FormatFloat(value, 'f', 9, 64))
	if is64 {
		r, _, _ := _this._jsDoubleString.Call(uintptr(unsafe.Pointer(&ptr[0])))
		return jsValue(r)
	}
	l, h, _ := _this._jsDoubleString.Call(uintptr(unsafe.Pointer(&ptr[0])))
	return _toJsValue(l, h)
}

func (_this *winFreeApi) jsBoolean(value bool) jsValue {
	if is64 {
		r, _, _ := _this._jsBoolean.Call(uintptr(toBool(value)))
		return jsValue(r)
	}
	l, h, _ := _this._jsBoolean.Call(uintptr(toBool(value)))
	return _toJsValue(l, h)
}

func (_this *winFreeApi) jsInt(value int32) jsValue {
	if is64 {
		r, _, _ := _this._jsInt.Call(uintptr(value))
		return jsValue(r)
	}
	l, h, _ := _this._jsInt.Call(uintptr(value))
	return _toJsValue(l, h)
}

func (_this *winFreeApi) jsCall(es jsExecState, fn, thisObject jsValue, args []jsValue) jsValue {
	var ptr = uintptr(0)
	if len(args) > 0 {
		ptr = uintptr(unsafe.Pointer(&args[0]))
	}
	if is64 {
		r, _, _ := _this._jsCall.Call(uintptr(es), uintptr(fn), uintptr(thisObject), ptr, uintptr(len(args)))
		return jsValue(r)
	}
	fl, fh := _toLH(fn)
	ol, oh := _toLH(thisObject)
	l, h, _ := _this._jsCall.Call(uintptr(es), uintptr(fl), uintptr(fh), uintptr(ol), uintptr(oh), ptr, uintptr(len(args)))
	return _toJsValue(l, h)
}

func (_this *winFreeApi) wkeGlobalExec(wke wkeHandle) jsExecState {
	r, _, _ := _this._wkeGlobalExec.Call(uintptr(wke))
	return jsExecState(r)
}

func (_this *winFreeApi) jsGetGlobal(es jsExecState, name string) jsValue {
	ptr := toCallStr(name)
	if is64 {
		r, _, _ := _this._jsGetGlobal.Call(uintptr(es), uintptr(unsafe.Pointer(&ptr[0])))
		return jsValue(r)
	}
	l, h, _ := _this._jsGetGlobal.Call(uintptr(es), uintptr(unsafe.Pointer(&ptr[0])))
	return _toJsValue(l, h)
}

func (_this *winFreeApi) jsSetGlobal(es jsExecState, name string, value jsValue) {
	ptr := toCallStr(name)
	if is64 {
		_this._jsSetGlobal.Call(uintptr(es), uintptr(unsafe.Pointer(&ptr[0])), uintptr(value))
	} else {
		l, h := _toLH(value)
		_this._jsSetGlobal.Call(uintptr(es), uintptr(unsafe.Pointer(&ptr[0])), uintptr(l), uintptr(h))
	}
}

func (_this *winFreeApi) jsEval(es jsExecState, js string) jsValue {
	ptr := toCallStr(js)
	if is64 {
		rs, _, _ := _this._jsEval.Call(uintptr(es), uintptr(unsafe.Pointer(&ptr[0])))
		return jsValue(rs)
	}
	l, h, _ := _this._jsEval.Call(uintptr(es), uintptr(unsafe.Pointer(&ptr[0])))
	return _toJsValue(l, h)
}

func (_this *winFreeApi) jsGetKeys(es jsExecState, value jsValue) []string {
	var rs uintptr
	if is64 {
		r, _, _ := _this._jsGetKeys.Call(uintptr(es), uintptr(value))
		rs = r
	} else {
		l, h := _toLH(value)
		r, _, _ := _this._jsGetKeys.Call(uintptr(es), uintptr(l), uintptr(h))
		rs = r
	}
	keys := *((*jsKeys)(unsafe.Pointer(rs)))
	items := make([]string, keys.length)
	for i := 0; i < len(items); i++ {
		items[i] = ptrToUtf8(*((*uintptr)(unsafe.Pointer(keys.first))))
		keys.first += unsafe.Sizeof(uintptr(0))
	}
	return items

	//_this._jsGetKeys.Call(uintptr(es), uintptr(value))
	//return []string{"n1", "n2"}

	//json := _this.jsGetGlobal(es, "Object")
	//stringify := _this.jsGet(es, json, "keys")
	//rs := _this.jsCall(es, stringify, _this.jsUndefined(), []jsValue{value})
	//alen := _this.jsGetLength(es, rs)
	//items := make([]string, alen)
	//for i := 0; i < len(items); i++ {
	//	v := _this.jsGetAt(es, rs, uint32(i))
	//	str := _this.jsToTempString(es, v)
	//	items[i] = str
	//}
	//return items
}

func (_this *winFreeApi) jsGet(es jsExecState, value jsValue, name string) jsValue {
	ptr := toCallStr(name)
	if is64 {
		r, _, _ := _this._jsGet.Call(uintptr(es), uintptr(value), uintptr(unsafe.Pointer(&ptr[0])))
		return jsValue(r)
	}
	pl, ph := _toLH(value)
	l, h, _ := _this._jsGet.Call(uintptr(es), uintptr(pl), uintptr(ph), uintptr(unsafe.Pointer(&ptr[0])))
	return _toJsValue(l, h)
}

func (_this *winFreeApi) jsGetAt(es jsExecState, value jsValue, index uint32) jsValue {
	if is64 {
		r, _, _ := _this._jsGetAt.Call(uintptr(es), uintptr(value), uintptr(index))
		return jsValue(r)
	}
	pl, ph := _toLH(value)
	rl, rh, _ := _this._jsGetAt.Call(uintptr(es), uintptr(pl), uintptr(ph), uintptr(index))
	return _toJsValue(rl, rh)
}

func (_this *winFreeApi) jsGetLength(es jsExecState, value jsValue) int {
	if is64 {
		r, _, _ := _this._jsGetLength.Call(uintptr(es), uintptr(value))
		return int(r)
	}
	l, h := _toLH(value)
	r, _, _ := _this._jsGetLength.Call(uintptr(es), uintptr(l), uintptr(h))
	return int(r)
}

func (_this *winFreeApi) jsUndefined() jsValue {
	if is64 {
		r, _, _ := _this._jsUndefined.Call()
		return jsValue(r)
	}
	l, h, _ := _this._jsUndefined.Call()
	return _toJsValue(l, h)
}

func (_this *winFreeApi) jsToBoolean(es jsExecState, value jsValue) bool {
	if is64 {
		r, _, _ := _this._jsToBoolean.Call(uintptr(es), uintptr(value))
		return byte(r) != 0
	}
	pl, ph := _toLH(value)
	r, _, _ := _this._jsToBoolean.Call(uintptr(es), uintptr(pl), uintptr(ph))
	return byte(r) != 0
}

func (_this *winFreeApi) jsToDouble(es jsExecState, value jsValue) float64 {
	var rs uintptr
	if is64 {
		r, _, _ := _this._jsToDoubleString.Call(uintptr(es), uintptr(value))
		rs = r
	} else {
		l, h := _toLH(value)
		r, _, _ := _this._jsToDoubleString.Call(uintptr(es), uintptr(l), uintptr(h))
		rs = r
	}
	str := ptrToUtf8(rs)
	n, _ := strconv.ParseFloat(str, 10)
	return n
}

func (_this *winFreeApi) jsToTempString(es jsExecState, value jsValue) string {
	if is64 {
		r, _, _ := _this._jsToTempString.Call(uintptr(es), uintptr(value))
		return ptrToUtf8(r)
	}
	l, h := _toLH(value)
	r, _, _ := _this._jsToTempString.Call(uintptr(es), uintptr(l), uintptr(h))
	return ptrToUtf8(r)
}

func (_this *winFreeApi) jsTypeOf(value jsValue) jsType {
	if is64 {
		r, _, _ := _this._jsTypeOf.Call(uintptr(value))
		return jsType(r)
	}
	l, h := _toLH(value)
	r, _, _ := _this._jsTypeOf.Call(uintptr(l), uintptr(h))
	return jsType(r)
}

func (_this *winFreeApi) jsArg(es jsExecState, index uint32) jsValue {
	if is64 {
		r, _, _ := _this._jsArg.Call(uintptr(es), uintptr(index))
		return jsValue(r)
	}
	l, h, _ := _this._jsArg.Call(uintptr(es), uintptr(index))
	return jsValue(uintptr(_toInt64(int32(l), int32(h))))
}

func (_this *winFreeApi) jsArgCount(es jsExecState) uint32 {
	r, _, _ := _this._jsArgCount.Call(uintptr(es))
	return uint32(r)
}

func (_this *winFreeApi) wkeJsBindFunction(name string, fn wkeJsNativeFunction, param uintptr, argCount uint32) {
	ptr := toCallStr(name)
	_this._wkeJsBindFunction.Call(uintptr(unsafe.Pointer(&ptr[0])), syscall.NewCallbackCDecl(fn), param, uintptr(argCount))
}

func (_this *winFreeApi) wkeNetCancelRequest(job wkeNetJob) {
	_this._wkeNetCancelRequest.Call(uintptr(job))
}

func (_this *winFreeApi) wkeNetOnResponse(wke wkeHandle, callback wkeNetResponseCallback, param uintptr) {
	_this._wkeNetOnResponse.Call(uintptr(wke), syscall.NewCallbackCDecl(callback), param)
}

func (_this *winFreeApi) wkeOnLoadUrlBegin(wke wkeHandle, callback wkeLoadUrlBeginCallback, param uintptr) {
	_this._wkeOnLoadUrlBegin.Call(uintptr(wke), syscall.NewCallbackCDecl(callback), param)
}

func (_this *winFreeApi) wkeNetGetRequestMethod(job wkeNetJob) wkeRequestType {
	r, _, _ := _this._wkeNetGetRequestMethod.Call(uintptr(job))
	return wkeRequestType(r)
}

func (_this *winFreeApi) wkeNetSetData(job wkeNetJob, buf []byte) {
	if len(buf) == 0 {
		buf = []byte{0}
	}
	_this._wkeNetSetData.Call(uintptr(job), uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)))
}

func (_this *winFreeApi) wkeGetCaretRect(wke wkeHandle) wkeRect {
	r, _, _ := _this._wkeGetCaretRect.Call(uintptr(wke))
	return *((*wkeRect)(unsafe.Pointer(r)))
}

func (_this *winFreeApi) wkeSetFocus(wke wkeHandle) {
	_this._wkeSetFocus.Call(uintptr(wke))
}

func (_this *winFreeApi) wkeFireKeyPressEvent(wke wkeHandle, code int, flags uint32, isSysKey bool) bool {
	ret, _, _ := _this._wkeFireKeyPressEvent.Call(
		uintptr(wke),
		uintptr(code),
		uintptr(flags),
		uintptr(toBool(isSysKey)))
	return byte(ret) != 0
}

func (_this *winFreeApi) wkeFireKeyDownEvent(wke wkeHandle, code, flags uint32, isSysKey bool) bool {
	ret, _, _ := _this._wkeFireKeyDownEvent.Call(
		uintptr(wke),
		uintptr(code),
		uintptr(flags),
		uintptr(toBool(isSysKey)))
	return byte(ret) != 0
}

func (_this *winFreeApi) wkeFireKeyUpEvent(wke wkeHandle, code, flags uint32, isSysKey bool) bool {
	ret, _, _ := _this._wkeFireKeyUpEvent.Call(
		uintptr(wke),
		uintptr(code),
		uintptr(flags),
		uintptr(toBool(isSysKey)))
	return byte(ret) != 0
}

func (_this *winFreeApi) wkeGetCursorInfoType(wke wkeHandle) wkeCursorType {
	r, _, _ := _this._wkeGetCursorInfoType.Call(uintptr(wke))
	return wkeCursorType(r)
}

func (_this *winFreeApi) wkeFireMouseWheelEvent(wke wkeHandle, x, y, delta, flags int32) bool {
	r, _, _ := _this._wkeFireMouseWheelEvent.Call(
		uintptr(wke),
		uintptr(x),
		uintptr(y),
		uintptr(delta),
		uintptr(flags))
	return byte(r) != 0
}

func (_this *winFreeApi) wkeFireMouseEvent(wke wkeHandle, message, x, y, flags int32) bool {
	r, _, _ := _this._wkeFireMouseEvent.Call(
		uintptr(wke),
		uintptr(message),
		uintptr(x),
		uintptr(y),
		uintptr(flags))
	return byte(r) != 0
}

func (_this *winFreeApi) wkePaint(wke wkeHandle, bits []byte, pitch uint32) {
	_this._wkePaint.Call(uintptr(wke), uintptr(unsafe.Pointer(&bits[0])), uintptr(pitch))
}

func (_this *winFreeApi) wkeGetHeight(wke wkeHandle) uint32 {
	r, _, _ := _this._wkeGetHeight.Call(uintptr(wke))
	return uint32(r)
}

func (_this *winFreeApi) wkeGetWidth(wke wkeHandle) uint32 {
	r, _, _ := _this._wkeGetWidth.Call(uintptr(wke))
	return uint32(r)
}

func (_this *winFreeApi) wkeResize(wke wkeHandle, w, h uint32) {
	_this._wkeResize.Call(uintptr(wke), uintptr(w), uintptr(h))
}

func (_this *winFreeApi) wkeLoadURL(wke wkeHandle, url string) {
	ptr := toCallStr(url)
	_this._wkeLoadURL.Call(uintptr(wke), uintptr(unsafe.Pointer(&ptr[0])))
}

func (_this *winFreeApi) wkeOnPaintBitUpdated(wke wkeHandle, callback wkePaintBitUpdatedCallback, param uintptr) {
	_this._wkeOnPaintBitUpdated.Call(uintptr(wke), syscall.NewCallbackCDecl(callback), param)
}

func (_this *winFreeApi) wkeSetHandle(wke wkeHandle, handle uintptr) {
	_this._wkeSetHandle.Call(uintptr(wke), handle)
}

func (_this *winFreeApi) wkeCreateWebView() wkeHandle {
	r, _, _ := _this._wkeCreateWebView.Call()
	return wkeHandle(r)
}
