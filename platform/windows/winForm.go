package windows

import (
	mb "qq.2564874169/goMiniblink"
	"qq.2564874169/goMiniblink/platform"
	"qq.2564874169/goMiniblink/platform/windows/win32"
	"syscall"
	"unsafe"
)

type winForm struct {
	winBase
	onClose func() (cancel bool)
	onState func(state mb.FormState)

	createParams *win32.DLGTEMPLATEEX
	initTitle    string
	initIcon     string
	ctrls        map[string]platform.IControl
}

func (_this *winForm) hWnd() win32.HWND {
	return _this.handle
}

func (_this *winForm) class() string {
	return win32.UTF16PtrToString(_this.createParams.WindowClass)
}

func (_this *winForm) id() string {
	return _this.idName
}

func (_this *winForm) init(provider *Provider) *winForm {
	_this.winBase.init(provider, mb.NewUUID())
	_this.ctrls = make(map[string]platform.IControl)
	_this.thisIsDialog = true
	_this.createParams = &win32.DLGTEMPLATEEX{
		Ver:         1,
		Sign:        0xFFFF,
		WindowClass: sto16(provider.className),
		ExStyle:     win32.WS_EX_APPWINDOW,
		Style:       win32.WS_SIZEBOX | win32.WS_CAPTION | win32.WS_SYSMENU | win32.WS_MAXIMIZEBOX | win32.WS_MINIMIZEBOX | win32.DS_ABSALIGN | win32.WS_CLIPCHILDREN,
	}
	_this.initTitle = ""
	bakWndProc := _this.onWndProc
	_this.onWndProc = func(hWnd win32.HWND, msg uint32, wParam, lParam uintptr) uintptr {
		ret := _this.formWndProc(hWnd, msg, wParam, lParam)
		if ret == 0 {
			ret = bakWndProc(hWnd, msg, wParam, lParam)
		}
		return ret
	}
	provider.add(_this)
	return _this
}

func (_this *winForm) GetHandle() uintptr {
	return uintptr(_this.handle)
}

func (_this *winForm) AddControl(control platform.IControl) {
	_this.ctrls[control.Id()] = control
	if _this.IsCreate() {
		control.SetParent(_this)
		control.Create()
		control.Show()
	}
}

func (_this *winForm) RemoveControl(control platform.IControl) {
	//if ctrl, ok := control.(*winControl); ok {
	//	if ctrl.IsCreate() {
	//
	//	}
	//	delete(_this.ctrls, ctrl.id())
	//}
}

func (_this *winForm) formWndProc(hWnd win32.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {
	case win32.WM_CLOSE:
		if _this.onClose != nil && _this.onClose() {
			return 1
		}
	case win32.WM_SIZE:
		if _this.onState != nil {
			switch int(wParam) {
			case win32.SIZE_RESTORED:
				_this.onState(mb.FormState_Normal)
			case win32.SIZE_MAXIMIZED:
				_this.onState(mb.FormState_Max)
			case win32.SIZE_MINIMIZED:
				_this.onState(mb.FormState_Min)
			}
		}
	}
	return 0
}

func (_this *winForm) Create() {
	if _this.IsCreate() == false {
		win32.CreateDialogIndirectParam(
			_this.app.hInstance,
			_this.createParams,
			_this.app.defOwner,
			syscall.NewCallback(_this.app.defaultMsgProc),
			unsafe.Pointer(&_this.idName))
		win32.SetWindowText(_this.hWnd(), _this.initTitle)
		win32.SetWindowPos(_this.hWnd(), 0,
			int32(_this.createParams.X),
			int32(_this.createParams.Y),
			int32(_this.createParams.CX),
			int32(_this.createParams.CY),
			win32.SWP_NOZORDER)
		if _this.createParams.Style&win32.DS_MODALFRAME == 0 {
			if _this.initIcon != "" {
				_this.SetIcon(_this.initIcon)
			} else if _this.app.defIcon != 0 {
				win32.SendMessage(_this.hWnd(), win32.WM_SETICON, 1, uintptr(_this.app.defIcon))
				win32.SendMessage(_this.hWnd(), win32.WM_SETICON, 0, uintptr(_this.app.defIcon))
			}
		}
		for _, v := range _this.ctrls {
			if v.GetHandle() == 0 {
				v.SetParent(_this)
				v.Create()
				v.Show()
			}
		}
	}
}

//func (_this *winForm) reCreate() {
//	preHwnd := _this.hWnd()
//	var rect win32.RECT
//	win32.GetWindowRect(preHwnd, &rect)
//	isVisible := win32.IsWindowVisible(preHwnd)
//	isMax := win32.IsZoomed(preHwnd)
//	isMin := win32.IsIconic(preHwnd)
//	_this.initTitle = win32.GetWindowText(preHwnd)
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
//	win32.DestroyWindow(preHwnd)
//	if isVisible {
//		if isMax {
//			win32.ShowWindow(_this.hWnd(), win32.SW_MAXIMIZE)
//		} else if isMin {
//			win32.ShowWindow(_this.hWnd(), win32.SW_MINIMIZE)
//		} else {
//			win32.ShowWindow(_this.hWnd(), win32.SW_SHOW)
//		}
//	}
//}

func (_this *winForm) Show() {
	isMax := win32.IsZoomed(_this.hWnd())
	isMin := win32.IsIconic(_this.hWnd())
	if isMax || isMin {
		win32.ShowWindow(_this.hWnd(), win32.SW_RESTORE)
	} else {
		win32.ShowWindow(_this.hWnd(), win32.SW_SHOW)
	}
	win32.UpdateWindow(_this.hWnd())
}

func (_this *winForm) ShowToMax() {
	win32.ShowWindow(_this.hWnd(), win32.SW_MAXIMIZE)
	win32.UpdateWindow(_this.hWnd())
}

func (_this *winForm) ShowToMin() {
	win32.ShowWindow(_this.hWnd(), win32.SW_MINIMIZE)
}

func (_this *winForm) Hide() {
	win32.ShowWindow(_this.hWnd(), win32.SW_HIDE)
}

func (_this *winForm) ShowDialog() {

}

func (_this *winForm) SetSize(w, h int) {
	if _this.IsCreate() {
		win32.SetWindowPos(_this.hWnd(), 0, 0, 0, int32(w), int32(h), win32.SWP_NOMOVE)
	} else {
		_this.createParams.CX, _this.createParams.CY = int16(w), int16(h)
	}
}

func (_this *winForm) SetLocation(x, y int) {
	if _this.IsCreate() {
		win32.SetWindowPos(_this.hWnd(), 0, int32(x), int32(y), 0, 0, win32.SWP_NOSIZE)
	} else {
		_this.createParams.X, _this.createParams.Y = int16(x), int16(y)
	}
}

func (_this *winForm) SetTitle(title string) {
	if _this.IsCreate() {
		win32.SetWindowText(_this.hWnd(), title)
	} else {
		_this.initTitle = title
	}
}

func (_this *winForm) SetBorderStyle(border mb.FormBorder) {
	var style int64
	if _this.IsCreate() {
		style = win32.GetWindowLong(_this.hWnd(), win32.GWL_STYLE)
	} else {
		style = int64(_this.createParams.Style)
	}
	bak := style
	switch border {
	case mb.FormBorder_Default:
		style |= win32.WS_SIZEBOX | win32.WS_CAPTION | win32.WS_SYSMENU | win32.WS_MAXIMIZEBOX | win32.WS_MINIMIZEBOX | win32.DS_ABSALIGN
	case mb.FormBorder_None:
		style &= ^win32.WS_SIZEBOX & ^win32.WS_CAPTION
	case mb.FormBorder_Disable_Resize:
		style &= ^win32.WS_SIZEBOX
	}
	if _this.IsCreate() {
		win32.SetWindowLong(_this.hWnd(), win32.GWL_STYLE, style)
	} else if bak != style {
		_this.createParams.Style = uint32(style)
	}
}

func (_this *winForm) ShowInTaskbar(isShow bool) {
	if _this.IsCreate() == false {
		if isShow {
			_this.createParams.ExStyle |= win32.WS_EX_APPWINDOW
		} else {
			_this.createParams.ExStyle &= ^uint32(win32.WS_EX_APPWINDOW)
		}
		return
	}
	//exStyle := uint32(win32.GetWindowLong(_this.hWnd(), win32.GWL_EXSTYLE))
	//bak := exStyle
	//if isShow {
	//	exStyle |= win32.WS_EX_APPWINDOW
	//} else {
	//	exStyle &= ^uint32(win32.WS_EX_APPWINDOW)
	//}
	//if bak != exStyle {
	//	_this.createParams.Style = uint32(win32.GetWindowLong(_this.hWnd(), win32.GWL_STYLE))
	//	_this.createParams.ExStyle = exStyle
	//	_this.reCreate()
	//}
}

func (_this *winForm) SetOnState(fn func(state mb.FormState)) {
	_this.onState = fn
}

func (_this *winForm) SetMaximizeBox(isShow bool) {
	var style uint32
	if _this.IsCreate() {
		style = uint32(win32.GetWindowLong(_this.hWnd(), win32.GWL_STYLE))
	} else {
		style = _this.createParams.Style
	}
	bak := style
	if isShow {
		style |= win32.WS_MAXIMIZEBOX
	} else {
		style &= ^uint32(win32.WS_MAXIMIZEBOX)
	}
	if bak != style {
		if _this.IsCreate() {
			win32.SetWindowLong(_this.hWnd(), win32.GWL_STYLE, int64(style))
		} else if bak != style {
			_this.createParams.Style = style
		}
	}
}

func (_this *winForm) SetMinimizeBox(isShow bool) {
	var style uint32
	if _this.IsCreate() {
		style = uint32(win32.GetWindowLong(_this.hWnd(), win32.GWL_STYLE))
	} else {
		style = _this.createParams.Style
	}
	bak := style
	if isShow {
		style |= win32.WS_MINIMIZEBOX
	} else {
		style &= ^uint32(win32.WS_MINIMIZEBOX)
	}
	if bak != style {
		if _this.IsCreate() {
			win32.SetWindowLong(_this.hWnd(), win32.GWL_STYLE, int64(style))
		} else if bak != style {
			_this.createParams.Style = style
		}
	}
}

func (_this *winForm) SetIcon(iconFile string) {
	_this.initIcon = iconFile
	if _this.IsCreate() {
		h := win32.LoadImage(_this.app.hInstance, sto16(_this.initIcon), win32.IMAGE_ICON, 0, 0, win32.LR_LOADFROMFILE)
		if h != 0 {
			win32.SendMessage(_this.hWnd(), win32.WM_SETICON, 1, uintptr(h))
			win32.SendMessage(_this.hWnd(), win32.WM_SETICON, 0, uintptr(h))
		}
	}
}

func (_this *winForm) SetIconVisable(isShow bool) {
	var style uint32
	if _this.IsCreate() {
		//style = uint32(win32.GetWindowLong(_this.hWnd(), win32.GWL_STYLE))
	} else {
		style = _this.createParams.Style
	}
	bak := style
	if isShow {
		style &= ^uint32(win32.DS_MODALFRAME)
	} else {
		style |= win32.DS_MODALFRAME
	}
	if bak != style {
		_this.createParams.Style = style
		//if _this.IsCreate() {
		//	_this.reCreate()
		//}
	}
}
