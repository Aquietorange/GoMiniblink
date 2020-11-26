package GoMiniblink

import (
	"gitee.com/aochulai/GoMiniblink/forms"
	"image"
	"strconv"
)

type PaintUpdatedEvArgs interface {
	Bitmap() *image.RGBA
	Bound() forms.Bound
	Cancel()
	IsCancel() bool
}

type freePaintUpdatedEvArgs struct {
	bitmap *image.RGBA
	bound  forms.Bound
	cancel bool
}

func (_this *freePaintUpdatedEvArgs) init(bitmap *image.RGBA, bound forms.Bound) *freePaintUpdatedEvArgs {
	_this.bitmap = bitmap
	_this.bound = bound
	return _this
}

func (_this *freePaintUpdatedEvArgs) Bitmap() *image.RGBA {
	return _this.bitmap
}

func (_this *freePaintUpdatedEvArgs) Bound() forms.Bound {
	return _this.bound
}

func (_this *freePaintUpdatedEvArgs) IsCancel() bool {
	return _this.cancel
}

func (_this *freePaintUpdatedEvArgs) Cancel() {
	_this.cancel = true
}

type DocumentReadyEvArgs interface {
	FrameContext
}

type freeDocumentReadyEvArgs struct {
	*freeFrameContext
}

func (_this *freeDocumentReadyEvArgs) init(mb Miniblink, frame wkeFrame) *freeDocumentReadyEvArgs {
	_this.freeFrameContext = new(freeFrameContext).init(mb, frame)
	return _this
}

type FrameContext interface {
	FrameId() uintptr
	IsMain() bool
	Url() string
	IsRemote() bool
	RunJs(script string) interface{}
}

type freeFrameContext struct {
	id       uintptr
	isMain   bool
	url      string
	isRemote bool
	core     Miniblink
}

func (_this *freeFrameContext) init(mb Miniblink, frame wkeFrame) *freeFrameContext {
	_this.core = mb
	_this.id = uintptr(frame)
	_this.isMain = mbApi.wkeIsMainFrame(_this.core.GetHandle(), frame)
	_this.isRemote = mbApi.wkeIsWebRemoteFrame(_this.core.GetHandle(), frame)
	_this.url = mbApi.wkeGetFrameUrl(_this.core.GetHandle(), frame)
	return _this
}

func (_this *freeFrameContext) RunJs(script string) interface{} {
	if len(script) > 0 {
		es := mbApi.wkeGetGlobalExecByFrame(_this.core.GetHandle(), wkeFrame(_this.id))
		rs := mbApi.jsEval(es, script)
		return toGoValue(_this.core, es, rs)
	}
	return nil
}

func (_this *freeFrameContext) IsRemote() bool {
	return _this.isRemote
}

func (_this *freeFrameContext) Url() string {
	return _this.url
}

func (_this *freeFrameContext) IsMain() bool {
	return _this.isMain
}

func (_this *freeFrameContext) FrameId() uintptr {
	return _this.id
}

type ConsoleEvArgs interface {
	Level() string
	Message() string
	SourceName() string
	SourceLine() int
	StackTrace() string
}

type freeConsoleMessageEvArgs struct {
	level   string
	message string
	name    string
	line    int
	stack   string
}

func (_this *freeConsoleMessageEvArgs) init() *freeConsoleMessageEvArgs {
	return _this
}

func (_this *freeConsoleMessageEvArgs) Level() string {
	return _this.level
}
func (_this *freeConsoleMessageEvArgs) Message() string {
	return _this.message
}
func (_this *freeConsoleMessageEvArgs) SourceName() string {
	return _this.name
}
func (_this *freeConsoleMessageEvArgs) SourceLine() int {
	return _this.line
}
func (_this *freeConsoleMessageEvArgs) StackTrace() string {
	return _this.stack
}

type JsReadyEvArgs interface {
	FrameContext
}

type wkeJsReadyEvArgs struct {
	*freeFrameContext
}

func (_this *wkeJsReadyEvArgs) init(mb Miniblink, frame wkeFrame) *wkeJsReadyEvArgs {
	_this.freeFrameContext = new(freeFrameContext).init(mb, frame)
	return _this
}

type ResponseEvArgs interface {
	RequestBefore() RequestBeforeEvArgs
	ContentType() string
	SetContentType(contentType string)
	Data() []byte
	SetData(data []byte)
	Headers() map[string]string
}

type freeResponseEvArgs struct {
	_req  *freeRequestBeforeEvArgs
	_data []byte
}

func (_this *freeResponseEvArgs) init(request *freeRequestBeforeEvArgs, data []byte) *freeResponseEvArgs {
	_this._req = request
	_this._data = data
	return _this
}

func (_this *freeResponseEvArgs) Headers() map[string]string {
	return mbApi.wkeNetGetRawResponseHead(_this._req._job)
}

func (_this *freeResponseEvArgs) SetData(data []byte) {
	_this._data = data
	mbApi.wkeNetSetData(_this._req._job, _this._data)
}

func (_this *freeResponseEvArgs) Data() []byte {
	return _this._data
}

func (_this *freeResponseEvArgs) RequestBefore() RequestBeforeEvArgs {
	return _this._req
}

func (_this *freeResponseEvArgs) ContentType() string {
	return mbApi.wkeNetGetMIMEType(_this._req._job)
}

func (_this *freeResponseEvArgs) SetContentType(contentType string) {
	mbApi.wkeNetSetMIMEType(_this._req._job, contentType)
}

type LoadFailEvArgs interface {
	RequestBefore() RequestBeforeEvArgs
}

type freeLoadFailEvArgs struct {
	_req *freeRequestBeforeEvArgs
}

func (_this *freeLoadFailEvArgs) init(request *freeRequestBeforeEvArgs) *freeLoadFailEvArgs {
	_this._req = request
	return _this
}

func (_this *freeLoadFailEvArgs) RequestBefore() RequestBeforeEvArgs {
	return _this._req
}

type RequestBeforeEvArgs interface {
	Url() string
	Method() string
	SetData([]byte)
	Data() []byte
	SetCancel(b bool)
	ResetUrl(url string)
	SetHeader(name, value string)
	/**
	内容最终呈现时触发
	args:intf, ResponseEvArgs
	*/
	EvResponse() *EventDispatcher
	/**
	加载失败时触发
	args:intf, LoadFailEvArgs
	*/
	EvLoadFail() *EventDispatcher
	/**
	请求流程全部完成时触发
	args:intf, RequestBeforeEvArgs
	*/
	EvFinish() *EventDispatcher
}

type freeRequestBeforeEvArgs struct {
	_wke    Miniblink
	_job    wkeNetJob
	_url    string
	_cancel bool
	_data   []byte
	//1=发送之前,2=异步处理,3=已发送,4=收到真实数据,5=完成
	_state         int
	_evResponseKey string
	_evResponse    *EventDispatcher
	_evLoadFailKey string
	_evLoadFail    *EventDispatcher
	_evFinishKey   string
	_evFinish      *EventDispatcher
}

func (_this *freeRequestBeforeEvArgs) init(wke Miniblink, job wkeNetJob) *freeRequestBeforeEvArgs {
	_this._wke = wke
	_this._url = mbApi.wkeNetGetUrlByJob(job)
	_this._job = job
	_this._state = 1
	_this._evResponseKey = "evResp" + strconv.FormatUint(uint64(job), 10)
	_this._evResponse = new(EventDispatcher).Init(_this._evResponseKey)
	_this._evLoadFailKey = "evFail" + strconv.FormatUint(uint64(job), 10)
	_this._evLoadFail = new(EventDispatcher).Init(_this._evLoadFailKey)
	_this._evFinishKey = "evFsh" + strconv.FormatUint(uint64(job), 10)
	_this._evFinish = new(EventDispatcher).Init(_this._evFinishKey)
	return _this
}

func (_this *freeRequestBeforeEvArgs) EvFinish() *EventDispatcher {
	return _this._evFinish
}

func (_this *freeRequestBeforeEvArgs) EvLoadFail() *EventDispatcher {
	return _this._evLoadFail
}

func (_this *freeRequestBeforeEvArgs) EvResponse() *EventDispatcher {
	return _this._evResponse
}

func (_this *freeRequestBeforeEvArgs) ResetUrl(url string) {
	mbApi.wkeNetChangeRequestUrl(_this._job, url)
	_this._url = url
}

func (_this *freeRequestBeforeEvArgs) SetHeader(name, value string) {
	mbApi.wkeNetSetHTTPHeaderField(_this._job, name, value)
}

func (_this *freeRequestBeforeEvArgs) SetData(data []byte) {
	_this._data = data
}

func (_this *freeRequestBeforeEvArgs) Data() []byte {
	return _this._data
}

func (_this *freeRequestBeforeEvArgs) Method() string {
	t := mbApi.wkeNetGetRequestMethod(_this._job)
	switch t {
	case wkeRequestType_Get:
		return "GET"
	case wkeRequestType_Post:
		return "POST"
	case wkeRequestType_Put:
		return "PUT"
	default:
		return "UNKNOW"
	}
}

func (_this *freeRequestBeforeEvArgs) Url() string {
	return _this._url
}

func (_this *freeRequestBeforeEvArgs) SetCancel(b bool) {
	_this._cancel = b
}

func (_this *freeRequestBeforeEvArgs) onBegin() {
	if _this._state == 2 {
		return
	}
	if _this._state == 1 && _this._data != nil {
		mbApi.wkeNetSetData(_this._job, _this._data)
		_this._cancel = true
		_this._state = 5
	} else if _this._evResponse.IsEmtpy() == false {
		mbApi.wkeNetHookRequest(_this._job)
		_this._cancel = false
	}
	if _this._cancel {
		mbApi.wkeNetCancelRequest(_this._job)
		if _this._data != nil {
			_this.onResponse(_this._data)
		} else {
			_this._evFinish.Fire(_this._evFinishKey, _this, _this)
		}
	} else {
		_this._state = 3
	}
}

func (_this *freeRequestBeforeEvArgs) onResponse(data []byte) {
	_this._state = 5
	args := new(freeResponseEvArgs).init(_this, data)
	_this._evResponse.Fire(_this._evResponseKey, _this, args)
	_this._evFinish.Fire(_this._evFinishKey, _this, _this)
}

func (_this *freeRequestBeforeEvArgs) onFail() {
	_this._state = 4
	args := new(freeLoadFailEvArgs).init(_this)
	_this._evLoadFail.Fire(_this._evLoadFailKey, _this, args)
	if _this._evResponse.IsEmtpy() {
		_this._evFinish.Fire(_this._evFinishKey, _this, _this)
	}
}

func (_this *MiniblinkBrowser) defOnRequestBefore(e RequestBeforeEvArgs) {
	for _, v := range _this.EvRequestBefore {
		v(_this, e)
	}
}

func (_this *MiniblinkBrowser) defOnJsReady(e JsReadyEvArgs) {
	for _, v := range _this.EvJsReady {
		v(_this, e)
	}
}

func (_this *MiniblinkBrowser) defOnConsole(e ConsoleEvArgs) {
	for _, v := range _this.EvConsole {
		v(_this, e)
	}
}

func (_this *MiniblinkBrowser) defOnDocumentReady(e DocumentReadyEvArgs) {
	for _, v := range _this.EvDocumentReady {
		v(_this, e)
	}
}

func (_this *MiniblinkBrowser) defOnPaintUpdated(e PaintUpdatedEvArgs) {
	for _, v := range _this.EvPaintUpdated {
		v(_this, e)
	}
}
