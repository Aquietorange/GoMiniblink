package windows

import (
	f "qq2564874169/goMiniblink/forms"
	p "qq2564874169/goMiniblink/forms/platform"
	w "qq2564874169/goMiniblink/forms/platform/windows/win32"
	"syscall"
)

type winForm struct {
	winBase
	onClose func() (cancel bool)
	onState p.FormStateProc

	createParams *w.DLGTEMPLATEEX
	initTitle    string
	initIcon     string
	ctrls        []p.Control
}

func (_this *winForm) init(provider *Provider) *winForm {
	_this.winBase.init(provider)
	_this.createParams = &w.DLGTEMPLATEEX{
		Ver:         1,
		Sign:        0xFFFF,
		WindowClass: sto16(provider.className),
		ExStyle:     w.WS_EX_APPWINDOW,
		Style:       w.WS_SIZEBOX | w.WS_CAPTION | w.WS_SYSMENU | w.WS_MAXIMIZEBOX | w.WS_MINIMIZEBOX | w.DS_ABSALIGN | w.WS_CLIPCHILDREN,
	}
	_this.initTitle = ""
	return _this
}

func (_this *winForm) GetHandle() uintptr {
	return uintptr(_this.handle)
}

func (_this *winForm) AddControl(control p.Control) {
	_this.ctrls = append(_this.ctrls, control)
	if _this.IsCreate() {
		control.SetParent(_this)
		control.Create()
		control.Show()
	}
}

func (_this *winForm) RemoveControl(control p.Control) {
	//if ctrl, ok := control.(*winControl); ok {
	//	if ctrl.IsCreate() {
	//
	//	}
	//	delete(_this.ctrls, ctrl.id())
	//}
}

func (_this *winForm) formWndProc(hWnd w.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	var b bool
	switch msg {
	case w.WM_INITDIALOG:
		_this.isCreated = true
		_this.handle = hWnd
		_this.app.add(_this)
		w.SetWindowText(hWnd, _this.initTitle)
		w.SetWindowPos(hWnd, 0,
			int32(_this.createParams.X),
			int32(_this.createParams.Y),
			int32(_this.createParams.CX),
			int32(_this.createParams.CY),
			w.SWP_NOZORDER)
		if _this.createParams.Style&w.DS_MODALFRAME == 0 {
			if _this.initIcon != "" {
				_this.SetIcon(_this.initIcon)
			} else if _this.app.defIcon != 0 {
				w.SendMessage(hWnd, w.WM_SETICON, 1, uintptr(_this.app.defIcon))
				w.SendMessage(hWnd, w.WM_SETICON, 0, uintptr(_this.app.defIcon))
			}
		}
		for _, v := range _this.ctrls {
			if v.GetHandle() == 0 {
				v.SetParent(_this)
				v.Create()
				v.Show()
			}
		}
		if _this.onCreate != nil {
			_this.onCreate(uintptr(hWnd))
		}
	case w.WM_CLOSE:
		if _this.onClose != nil && _this.onClose() {
			b = true
		} else {
			return w.DefWindowProc(hWnd, msg, wParam, lParam)
		}
	case w.WM_SIZE:
		if _this.onState != nil {
			switch int(wParam) {
			case w.SIZE_RESTORED:
				_this.onState(f.FormState_Normal)
			case w.SIZE_MAXIMIZED:
				_this.onState(f.FormState_Max)
			case w.SIZE_MINIMIZED:
				_this.onState(f.FormState_Min)
			}
		}
	}
	if b == false {
		if r := _this.winBase.msgProc(hWnd, msg, wParam, lParam); r != 0 {
			b = true
		}
	}
	if b {
		return uintptr(byte(1))
	}
	return uintptr(byte(0))
}

func (_this *winForm) Create() {
	if _this.IsCreate() == false {
		hWnd := w.CreateDialogIndirectParam(
			_this.app.hInstance,
			_this.createParams,
			_this.app.defOwner,
			syscall.NewCallback(_this.formWndProc),
			nil)
		if hWnd == 0 {
			panic("创建失败")
		}
	}
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
	if _this.IsCreate() {
		w.SetWindowPos(_this.hWnd(), 0, 0, 0, int32(width), int32(height), w.SWP_NOMOVE)
	} else {
		_this.createParams.CX, _this.createParams.CY = int16(width), int16(height)
	}
}

func (_this *winForm) SetLocation(x, y int) {
	if _this.IsCreate() {
		w.SetWindowPos(_this.hWnd(), 0, int32(x), int32(y), 0, 0, w.SWP_NOSIZE)
	} else {
		_this.createParams.X, _this.createParams.Y = int16(x), int16(y)
	}
}

func (_this *winForm) SetTitle(title string) {
	if _this.IsCreate() {
		w.SetWindowText(_this.hWnd(), title)
	} else {
		_this.initTitle = title
	}
}

func (_this *winForm) SetBorderStyle(border f.FormBorder) {
	var style int64
	if _this.IsCreate() {
		style = w.GetWindowLong(_this.hWnd(), w.GWL_STYLE)
	} else {
		style = int64(_this.createParams.Style)
	}
	bak := style
	switch border {
	case f.FormBorder_Default:
		style |= w.WS_SIZEBOX | w.WS_CAPTION | w.WS_SYSMENU | w.WS_MAXIMIZEBOX | w.WS_MINIMIZEBOX | w.DS_ABSALIGN
	case f.FormBorder_None:
		style &= ^w.WS_SIZEBOX & ^w.WS_CAPTION
	case f.FormBorder_Disable_Resize:
		style &= ^w.WS_SIZEBOX
	}
	if _this.IsCreate() {
		w.SetWindowLong(_this.hWnd(), w.GWL_STYLE, style)
	} else if bak != style {
		_this.createParams.Style = uint32(style)
	}
}

func (_this *winForm) ShowInTaskbar(isShow bool) {
	if _this.IsCreate() == false {
		if isShow {
			_this.createParams.ExStyle |= w.WS_EX_APPWINDOW
		} else {
			_this.createParams.ExStyle &= ^uint32(w.WS_EX_APPWINDOW)
		}
		return
	}
	//exStyle := uint32(w.GetWindowLong(_this.hWnd(), w.GWL_EXSTYLE))
	//bak := exStyle
	//if isShow {
	//	exStyle |= w.WS_EX_APPWINDOW
	//} else {
	//	exStyle &= ^uint32(w.WS_EX_APPWINDOW)
	//}
	//if bak != exStyle {
	//	_this.createParams.Style = uint32(w.GetWindowLong(_this.hWnd(), w.GWL_STYLE))
	//	_this.createParams.ExStyle = exStyle
	//	_this.reCreate()
	//}
}

func (_this *winForm) SetOnState(proc p.FormStateProc) p.FormStateProc {
	pre := _this.onState
	_this.onState = proc
	return pre
}

func (_this *winForm) SetMaximizeBox(isShow bool) {
	var style uint32
	if _this.IsCreate() {
		style = uint32(w.GetWindowLong(_this.hWnd(), w.GWL_STYLE))
	} else {
		style = _this.createParams.Style
	}
	bak := style
	if isShow {
		style |= w.WS_MAXIMIZEBOX
	} else {
		style &= ^uint32(w.WS_MAXIMIZEBOX)
	}
	if bak != style {
		if _this.IsCreate() {
			w.SetWindowLong(_this.hWnd(), w.GWL_STYLE, int64(style))
		} else if bak != style {
			_this.createParams.Style = style
		}
	}
}

func (_this *winForm) SetMinimizeBox(isShow bool) {
	var style uint32
	if _this.IsCreate() {
		style = uint32(w.GetWindowLong(_this.hWnd(), w.GWL_STYLE))
	} else {
		style = _this.createParams.Style
	}
	bak := style
	if isShow {
		style |= w.WS_MINIMIZEBOX
	} else {
		style &= ^uint32(w.WS_MINIMIZEBOX)
	}
	if bak != style {
		if _this.IsCreate() {
			w.SetWindowLong(_this.hWnd(), w.GWL_STYLE, int64(style))
		} else if bak != style {
			_this.createParams.Style = style
		}
	}
}

func (_this *winForm) SetIcon(iconFile string) {
	_this.initIcon = iconFile
	if _this.IsCreate() {
		h := w.LoadImage(_this.app.hInstance, sto16(_this.initIcon), w.IMAGE_ICON, 0, 0, w.LR_LOADFROMFILE)
		if h != 0 {
			w.SendMessage(_this.hWnd(), w.WM_SETICON, 1, uintptr(h))
			w.SendMessage(_this.hWnd(), w.WM_SETICON, 0, uintptr(h))
		}
	}
}

func (_this *winForm) SetIconVisable(isShow bool) {
	var style uint32
	if _this.IsCreate() {
		//style = uint32(w.GetWindowLong(_this.hWnd(), w.GWL_STYLE))
	} else {
		style = _this.createParams.Style
	}
	bak := style
	if isShow {
		style &= ^uint32(w.DS_MODALFRAME)
	} else {
		style |= w.DS_MODALFRAME
	}
	if bak != style {
		_this.createParams.Style = style
		//if _this.IsCreate() {
		//	_this.reCreate()
		//}
	}
}
