package GoMiniblink

type RequestEvArgs interface {
	GetUrl() string
	GetMethod() string
	SetData([]byte)
	GetData() []byte
	SetCancel(b bool)
	IsCancel() bool
}

type freeRequestEvArgs struct {
	_wke    IMiniblink
	_job    wkeNetJob
	_url    string
	_cancel bool
	_data   []byte
	//1=before,2=async,3=post,4=net,5=valid
	_state int
}

func (_this *freeRequestEvArgs) init(wke IMiniblink, url string, job wkeNetJob) *freeRequestEvArgs {
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

func (_this *freeRequestEvArgs) GetData() []byte {
	return _this._data
}

func (_this *freeRequestEvArgs) GetMethod() string {
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

func (_this *freeRequestEvArgs) GetUrl() string {
	return _this._url
}

func (_this *freeRequestEvArgs) IsCancel() bool {
	return _this._cancel
}

func (_this *freeRequestEvArgs) SetCancel(b bool) {
	_this._cancel = b
}
