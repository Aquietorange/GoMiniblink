package windows

import (
	"fmt"
	f "qq2564874169/goMiniblink/forms"
	p "qq2564874169/goMiniblink/forms/platform"
	w "qq2564874169/goMiniblink/forms/platform/windows/win32"
	"unsafe"
)

type winForm struct {
	winBase
	_onClose func() (cancel bool)
	_onState p.FormStateProc

	_ctrls []p.Control
}

func (_this *winForm) init(provider *Provider) *winForm {
	_this.winBase.init(provider)
	_this.onWndProc = _this.msgProc
	rs := w.CreateWindowEx(
		w.WS_EX_APPWINDOW,
		(*uint16)(unsafe.Pointer(sto16(_this.app.className))),
		(*uint16)(unsafe.Pointer(sto16(""))),
		w.WS_SIZEBOX|w.WS_CAPTION|w.WS_SYSMENU|w.WS_MAXIMIZEBOX|w.WS_MINIMIZEBOX|w.WS_CLIPCHILDREN,
		0, 0, 0, 0, 0, 0, _this.app.hInstance, unsafe.Pointer(_this))
	if rs == 0 {
		fmt.Println("创建失败")
	}
	return _this
}

func (_this *winForm) GetHandle() uintptr {
	return uintptr(_this.handle)
}

func (_this *winForm) AddControl(control p.Control) {
	control.SetParent(_this)
	_this._ctrls = append(_this._ctrls, control)
}

func (_this *winForm) RemoveControl(control p.Control) {
	//if ctrl, ok := control.(*winControl); ok {
	//	if ctrl.IsCreate() {
	//
	//	}
	//	delete(_this._ctrls, ctrl.id())
	//}
}

func (_this *winForm) msgProc(hWnd w.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	rs := _this.winBase.msgProc(hWnd, msg, wParam, lParam)
	if rs != 0 {
		return rs
	}
	switch msg {
	case w.WM_CREATE:
		_this.SetSize(200, 300)
	case w.WM_SIZE:
		if _this._onState != nil {
			switch int(wParam) {
			case w.SIZE_RESTORED:
				_this._onState(f.FormState_Normal)
			case w.SIZE_MAXIMIZED:
				_this._onState(f.FormState_Max)
			case w.SIZE_MINIMIZED:
				_this._onState(f.FormState_Min)
			}
		}
	case w.WM_CLOSE:
		if _this._onClose != nil && _this._onClose() {
			rs = 1
		}
	}
	return rs
}

func (_this *winForm) Create() {
	//if _this.IsCreate() == false {
	//	hWnd := w.CreateDialogIndirectParam(
	//		_this.app.hInstance,
	//		_this.createParams,
	//		_this.app.defOwner,
	//		syscall.NewCallback(_this.formWndProc),
	//		nil)
	//	if hWnd == 0 {
	//		panic("创建失败")
	//	}
	//}
}

//func (_this *winForm) reCreate() {
//	preHwnd := _this.hWnd()
//	var rect w.RECT
//	w.GetWindowRect(preHwnd, &rect)
//	isVisible := w.IsWindowVisible(preHwnd)
//	isMax := w.IsZoomed(preHwnd)
//	isMin := w.IsIconic(preHwnd)
//	_this.initTitle = w.GetWindowText(preHwnd)
//	_this.createParams.X = int16(rect.Left)
//	_this.createParams.Y = int16(rect.Top)
//	_this.createParams.CX = int16(rect.Right - rect.Left)
//	_this.createParams.CY = int16(rect.Bottom - rect.Top)
//	_this.isCreated = false
//	bakEvCreate := _this.evWndCreate
//	_this.evWndCreate = nil
//	_this.Create()
//	_this.evWndCreate = bakEvCreate
//	_this.app.remove(preHwnd, false)
//	w.DestroyWindow(preHwnd)
//	if isVisible {
//		if isMax {
//			w.ShowWindow(_this.hWnd(), w.SW_MAXIMIZE)
//		} else if isMin {
//			w.ShowWindow(_this.hWnd(), w.SW_MINIMIZE)
//		} else {
//			w.ShowWindow(_this.hWnd(), w.SW_SHOW)
//		}
//	}
//}

func (_this *winForm) Show() {
	isMax := w.IsZoomed(_this.hWnd())
	isMin := w.IsIconic(_this.hWnd())
	if isMax || isMin {
		w.ShowWindow(_this.hWnd(), w.SW_RESTORE)
	} else {
		w.ShowWindow(_this.hWnd(), w.SW_SHOW)
	}
	w.UpdateWindow(_this.hWnd())
}

func (_this *winForm) Hide() {
	w.ShowWindow(_this.hWnd(), w.SW_HIDE)
}

func (_this *winForm) Close() {
	w.SendMessage(_this.hWnd(), w.WM_CLOSE, 0, 0)
}

func (_this *winForm) ShowDialog() {

}

func (_this *winForm) ShowToMax() {
	w.ShowWindow(_this.hWnd(), w.SW_MAXIMIZE)
	w.UpdateWindow(_this.hWnd())
}

func (_this *winForm) ShowToMin() {
	w.ShowWindow(_this.hWnd(), w.SW_MINIMIZE)
}

func (_this *winForm) SetSize(width, height int) {
	w.SetWindowPos(_this.hWnd(), 0, 0, 0, int32(width), int32(height), w.SWP_NOMOVE)
}

func (_this *winForm) SetLocation(x, y int) {
	w.SetWindowPos(_this.hWnd(), 0, int32(x), int32(y), 0, 0, w.SWP_NOSIZE)
}

func (_this *winForm) SetTitle(title string) {
	w.SetWindowText(_this.hWnd(), title)
}

func (_this *winForm) SetBorderStyle(border f.FormBorder) {
	style := w.GetWindowLong(_this.hWnd(), w.GWL_STYLE)
	switch border {
	case f.FormBorder_Default:
		style |= w.WS_SIZEBOX | w.WS_CAPTION | w.WS_SYSMENU | w.WS_MAXIMIZEBOX | w.WS_MINIMIZEBOX | w.DS_ABSALIGN
	case f.FormBorder_None:
		style &= ^w.WS_SIZEBOX & ^w.WS_CAPTION
	case f.FormBorder_Disable_Resize:
		style &= ^w.WS_SIZEBOX
	}
	w.SetWindowLong(_this.hWnd(), w.GWL_STYLE, style)
}

func (_this *winForm) ShowInTaskbar(isShow bool) {
	style := uint32(w.GetWindowLong(_this.hWnd(), w.GWL_STYLE))
	if isShow {
		style |= w.WS_EX_APPWINDOW
	} else {
		style &= ^uint32(w.WS_EX_APPWINDOW)
	}
}

func (_this *winForm) SetOnState(proc p.FormStateProc) p.FormStateProc {
	pre := _this._onState
	_this._onState = proc
	return pre
}

func (_this *winForm) SetMaximizeBox(isShow bool) {
	style := uint32(w.GetWindowLong(_this.hWnd(), w.GWL_STYLE))
	if isShow {
		style |= w.WS_MAXIMIZEBOX
	} else {
		style &= ^uint32(w.WS_MAXIMIZEBOX)
	}
	w.SetWindowLong(_this.hWnd(), w.GWL_STYLE, int64(style))
}

func (_this *winForm) SetMinimizeBox(isShow bool) {
	style := uint32(w.GetWindowLong(_this.hWnd(), w.GWL_STYLE))
	if isShow {
		style |= w.WS_MINIMIZEBOX
	} else {
		style &= ^uint32(w.WS_MINIMIZEBOX)
	}
	w.SetWindowLong(_this.hWnd(), w.GWL_STYLE, int64(style))
}

func (_this *winForm) SetIcon(iconFile string) {
	h := w.LoadImage(_this.app.hInstance, sto16(iconFile), w.IMAGE_ICON, 0, 0, w.LR_LOADFROMFILE)
	if h != 0 {
		w.SendMessage(_this.hWnd(), w.WM_SETICON, 1, uintptr(h))
		w.SendMessage(_this.hWnd(), w.WM_SETICON, 0, uintptr(h))
	}
}

func (_this *winForm) SetIconVisable(isShow bool) {
	style := uint32(w.GetWindowLong(_this.hWnd(), w.GWL_STYLE))
	if isShow {
		style &= ^uint32(w.DS_MODALFRAME)
	} else {
		style |= w.DS_MODALFRAME
	}
}
