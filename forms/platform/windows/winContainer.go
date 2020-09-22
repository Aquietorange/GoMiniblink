package windows

import (
	plat "qq2564874169/goMiniblink/forms/platform"
	win "qq2564874169/goMiniblink/forms/platform/windows/win32"
)

type winContainer struct {
	winBase
	_self  plat.Control
	_ctrls []plat.Control
}

func (_this *winContainer) init(provider *Provider, self plat.Control) *winContainer {
	_this.winBase.init(provider)
	_this._self = self
	return _this
}

func (_this *winContainer) GetChilds() []plat.Control {
	return _this._ctrls
}

func (_this *winContainer) AddControl(control plat.Control) {
	if win.SetParent(win.HWND(control.GetHandle()), win.HWND(_this.GetHandle())) != 0 {
		if wc, ok := control.(*winControl); ok {
			wc.parent = _this
		}
		if ow, ok := _this._self.(plat.Form); ok {
			_this.owner = ow
		} else {
			_this.owner = _this.GetOwner()
		}
	}
	_this._ctrls = append(_this._ctrls, control)
}

func (_this *winContainer) RemoveControl(control plat.Control) {
	for i, n := range _this._ctrls {
		if n.GetHandle() == control.GetHandle() {
			hWnd := win.HWND(control.GetHandle())
			if win.SetParent(hWnd, _this.app.defOwner) != 0 {
				win.SendMessage(hWnd, win.WM_DESTROY, 0, 0)
				if wc, ok := control.(*winControl); ok {
					wc.parent = nil
					wc.owner = nil
				}
			}
			_this._ctrls = append(_this._ctrls[0:i], _this._ctrls[i+1:]...)
			break
		}
	}
}
