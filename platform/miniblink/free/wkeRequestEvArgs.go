package free

import core "qq2564874169/goMiniblink/platform/miniblink"

type wkeRequestEvArgs struct {
	wke    core.ICore
	job    wkeNetJob
	url    string
	cancel bool
	data   []byte
	state  int //1=before,2=async,3=post,4=net,5=valid
}

func (_this *wkeRequestEvArgs) init(wke core.ICore, url string, job wkeNetJob) *wkeRequestEvArgs {
	_this.wke = wke
	_this.url = url
	_this.job = job
	_this.state = 1
	return _this
}

func (_this *wkeRequestEvArgs) OnBegin() bool {
	if _this.state == 1 && _this.data != nil {
		_this.wkeSetData(_this.data)
		_this.cancel = true
		_this.state = 5
	}
	if _this.cancel {
		wkeNetCancelRequest(_this.job)
	}
	return _this.cancel
}

func (_this *wkeRequestEvArgs) SetData(data []byte) {
	_this.data = data
}

func (_this *wkeRequestEvArgs) GetData() []byte {
	return _this.data
}

func (_this *wkeRequestEvArgs) Method() string {
	t := wkeNetGetRequestMethod(_this.job)
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

func (_this *wkeRequestEvArgs) Url() string {
	return _this.url
}

func (_this *wkeRequestEvArgs) IsCancel() bool {
	return _this.cancel
}

func (_this *wkeRequestEvArgs) SetCancel(b bool) {
	_this.cancel = b
}

func (_this *wkeRequestEvArgs) wkeSetData(data []byte) {
	if data == nil || len(data) == 0 {
		var buf = [1]byte{0}
		wkeNetSetData(_this.job, &buf[0], 1)
	} else {
		wkeNetSetData(_this.job, &data[0], uint32(len(data)))
	}
}
