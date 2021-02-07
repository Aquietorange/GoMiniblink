package GoMiniblink

type (
	wkeString   uintptr
	wkeFrame    uintptr
	wkeHandle   uintptr
	jsExecState uintptr
	jsValue     int64
)

type wkeSlist struct {
	str  uintptr
	next uintptr
}

type jsType uint32

const (
	jsType_NUMBER jsType = iota
	jsType_STRING
	jsType_BOOLEAN
	jsType_OBJECT
	jsType_FUNCTION
	jsType_UNDEFINED
	jsType_ARRAY
	jsType_NULL
)

type jsData struct {
	name [100]byte
	propertyGet,
	propertySet,
	finalize,
	callAsFunction uintptr
}

type jsKeys struct {
	length uint32
	first  uintptr
}

type wkeRequestType int

const (
	wkeRequestType_Unknow = 1
	wkeRequestType_Get    = 2
	wkeRequestType_Post   = 3
	wkeRequestType_Put    = 4
)

type wkeKeyFlags int

const (
	wkeKeyFlags_Extend wkeKeyFlags = 0x0100
	wkeKeyFlags_Repeat wkeKeyFlags = 0x4000
)

type wkeRect struct {
	x, y, w, h int32
}

type wkeNetJob uintptr

type wkeMouseFlags int

const (
	wkeMouseFlags_None    wkeMouseFlags = 0
	wkeMouseFlags_LBUTTON wkeMouseFlags = 0x01
	wkeMouseFlags_RBUTTON wkeMouseFlags = 0x02
	wkeMouseFlags_SHIFT   wkeMouseFlags = 0x04
	wkeMouseFlags_CONTROL wkeMouseFlags = 0x08
	wkeMouseFlags_MBUTTON wkeMouseFlags = 0x10
)

type wkeConsoleLevel int

const (
	wkeConsoleLevel_Log          wkeConsoleLevel = 1
	wkeConsoleLevel_Warning      wkeConsoleLevel = 2
	wkeConsoleLevel_Error        wkeConsoleLevel = 3
	wkeConsoleLevel_Debug        wkeConsoleLevel = 4
	wkeConsoleLevel_Info         wkeConsoleLevel = 5
	wkeConsoleLevel_RevokedError wkeConsoleLevel = 6
)

type wkePaintBitUpdatedCallback func(wke wkeHandle, param, buf uintptr, rect *wkeRect, width, height int32) uintptr
type wkeNetResponseCallback func(wke wkeHandle, param, utf8Url uintptr, job wkeNetJob) uintptr
type wkeLoadUrlBeginCallback func(wke wkeHandle, param, utf8Url uintptr, job wkeNetJob) uintptr
type wkeJsNativeFunction func(es jsExecState, param uintptr) uintptr
type wkeDidCreateScriptContextCallback func(wke wkeHandle, param uintptr, frame wkeFrame, context uintptr, exGroup, worldId int) uintptr
type wkeConsoleCallback func(wke wkeHandle, param uintptr, level int32, msg, name wkeString, line uint32, stack wkeString) uintptr
type wkeLoadUrlEndCallback func(wke wkeHandle, param, url uintptr, job wkeNetJob, buf uintptr, len int32) uintptr
type wkeLoadUrlFailCallback func(wke wkeHandle, param, url uintptr, job wkeNetJob) uintptr
type wkeDocumentReady2Callback func(wke wkeHandle, param uintptr, frame wkeFrame) uintptr

var mbApi freeApi

type freeApi interface {
	wkeCreateWebView() wkeHandle
	wkeDestroyWebView(wke wkeHandle)
	jsGetWebView(es jsExecState) wkeHandle
	jsSet(es jsExecState, obj jsValue, name string, value jsValue)
	jsEmptyObject(es jsExecState) jsValue
	jsFunction(es jsExecState, data *jsData) jsValue
	jsSetAt(es jsExecState, array jsValue, index uint32, value jsValue)
	jsSetLength(es jsExecState, array jsValue, length uint32)
	jsEmptyArray(es jsExecState) jsValue
	jsString(es jsExecState, value string) jsValue
	jsDouble(value float64) jsValue
	jsBoolean(value bool) jsValue
	jsInt(value int32) jsValue
	jsCall(es jsExecState, fn, thisObject jsValue, args []jsValue) jsValue
	wkeGlobalExec(wke wkeHandle) jsExecState
	jsGetGlobal(es jsExecState, name string) jsValue
	jsSetGlobal(es jsExecState, name string, value jsValue)
	jsGetKeys(es jsExecState, value jsValue) []string
	jsGet(es jsExecState, value jsValue, name string) jsValue
	jsGetAt(es jsExecState, value jsValue, index uint32) jsValue
	jsGetLength(es jsExecState, value jsValue) int
	jsUndefined() jsValue
	jsToBoolean(es jsExecState, value jsValue) bool
	jsToDouble(es jsExecState, value jsValue) float64
	jsToTempString(es jsExecState, value jsValue) string
	jsTypeOf(value jsValue) jsType
	jsArg(es jsExecState, index uint32) jsValue
	jsArgCount(es jsExecState) uint32
	wkeJsBindFunction(name string, fn wkeJsNativeFunction, param uintptr, argCount uint32)
	wkeNetCancelRequest(job wkeNetJob)
	wkeNetOnResponse(wke wkeHandle, callback wkeNetResponseCallback, param uintptr)
	wkeOnLoadUrlBegin(wke wkeHandle, callback wkeLoadUrlBeginCallback, param uintptr)
	wkeNetGetRequestMethod(job wkeNetJob) wkeRequestType
	wkeNetSetData(job wkeNetJob, buf []byte)
	wkeGetCaretRect(wke wkeHandle) wkeRect
	wkeSetFocus(wke wkeHandle)
	wkeKillFocus(wke wkeHandle)
	wkeFireKeyPressEvent(wke wkeHandle, code int, flags uint32, isSysKey bool) bool
	wkeFireKeyDownEvent(wke wkeHandle, code, flags uint32, isSysKey bool) bool
	wkeFireKeyUpEvent(wke wkeHandle, code, flags uint32, isSysKey bool) bool
	wkeGetCursorInfoType(wke wkeHandle) wkeCursorType
	wkeFireMouseWheelEvent(wke wkeHandle, x, y, delta, flags int32) bool
	wkeFireMouseEvent(wke wkeHandle, message, x, y, flags int32) bool
	wkePaint(wke wkeHandle, bits []byte, pitch uint32)
	wkeGetHeight(wke wkeHandle) uint32
	wkeGetWidth(wke wkeHandle) uint32
	wkeResize(wke wkeHandle, w, h uint32)
	wkeLoadURL(wke wkeHandle, url string)
	wkeOnPaintBitUpdated(wke wkeHandle, callback wkePaintBitUpdatedCallback, param uintptr)
	wkeSetHandle(wke wkeHandle, handle uintptr)
	jsEval(es jsExecState, js string) jsValue
	wkeGetGlobalExecByFrame(wke wkeHandle, frame wkeFrame) jsExecState
	wkeIsWebRemoteFrame(wke wkeHandle, frame wkeFrame) bool
	wkeGetFrameUrl(wke wkeHandle, frame wkeFrame) string
	wkeIsMainFrame(wke wkeHandle, frame wkeFrame) bool
	wkeOnDidCreateScriptContext(wke wkeHandle, callback wkeDidCreateScriptContextCallback, param uintptr)
	wkeOnConsole(wke wkeHandle, callback wkeConsoleCallback, param uintptr)
	wkeGetString(str wkeString) string
	wkeNetSetHTTPHeaderField(job wkeNetJob, name, value string)
	wkeNetChangeRequestUrl(job wkeNetJob, url string)
	wkeNetHookRequest(job wkeNetJob)
	wkeNetHoldJobToAsynCommit(job wkeNetJob)
	wkeNetContinueJob(job wkeNetJob)
	wkeOnLoadUrlEnd(wke wkeHandle, callback wkeLoadUrlEndCallback, param uintptr)
	wkeOnLoadUrlFail(wke wkeHandle, callback wkeLoadUrlFailCallback, param uintptr)
	wkeNetGetUrlByJob(job wkeNetJob) string
	wkeNetGetMIMEType(job wkeNetJob) string
	wkeNetSetMIMEType(job wkeNetJob, mime string)
	wkeNetGetRawResponseHead(job wkeNetJob) map[string]string
	wkeOnDocumentReady(wke wkeHandle, callback wkeDocumentReady2Callback, param uintptr)
	wkeSetTransparent(wke wkeHandle, enable bool)
	wkeSetViewProxy(wke wkeHandle, proxy ProxyInfo)
	wkeGetViewDC(wke wkeHandle) uintptr
}
