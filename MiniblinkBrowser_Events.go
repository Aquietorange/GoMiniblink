package GoMiniblink

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
	Frame() FrameContext
}

type wkeJsReadyEvArgs struct {
	ctx *freeFrameContext
}

func (_this *wkeJsReadyEvArgs) init() *wkeJsReadyEvArgs {
	return _this
}

func (_this *wkeJsReadyEvArgs) Frame() FrameContext {
	return _this.ctx
}

type RequestEvArgs interface {
	Url() string
	Method() string
	SetData([]byte)
	Data() []byte
	SetCancel(b bool)
}

type freeRequestEvArgs struct {
	_wke    Miniblink
	_job    wkeNetJob
	_url    string
	_cancel bool
	_data   []byte
	//1=before,2=async,3=post,4=net,5=valid
	_state int
}

func (_this *freeRequestEvArgs) init(wke Miniblink, url string, job wkeNetJob) *freeRequestEvArgs {
	_this._wke = wke
	_this._url = url
	_this._job = job
	_this._state = 1
	return _this
}

func (_this *freeRequestEvArgs) onBegin() bool {
	if _this._state == 1 && _this._data != nil {
		_this.wkeSetData(_this._data)
		_this._cancel = true
		_this._state = 5
	}
	if _this._cancel {
		mbApi.wkeNetCancelRequest(_this._job)
	}
	return _this._cancel
}

func (_this *freeRequestEvArgs) SetData(data []byte) {
	_this._data = data
}

func (_this *freeRequestEvArgs) Data() []byte {
	return _this._data
}

func (_this *freeRequestEvArgs) Method() string {
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

func (_this *freeRequestEvArgs) wkeSetData(data []byte) {
	mbApi.wkeNetSetData(_this._job, data)
}

func (_this *freeRequestEvArgs) Url() string {
	return _this._url
}

func (_this *freeRequestEvArgs) SetCancel(b bool) {
	_this._cancel = b
}

func (_this *MiniblinkBrowser) defOnRequestBefore(e RequestEvArgs) {
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
