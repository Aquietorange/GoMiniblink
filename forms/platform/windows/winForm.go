package windows

import (
	f "qq2564874169/goMiniblink/forms"
	plat "qq2564874169/goMiniblink/forms/platform"
	win "qq2564874169/goMiniblink/forms/platform/windows/win32"
	"unsafe"
)

type winForm struct {
	winBase
	_onClose func() (cancel bool)
	_onState plat.FormStateProc
	_ctrls   []plat.Control
	_border  f.FormBorder
}

func (_this *winForm) init(provider *Provider, param plat.FormParam) *winForm {
	_this.winBase.init(provider)
	_this.onWndProc = _this.msgProc
	parent := win.HWND(0)
	exStyle := win.WS_EX_APPWINDOW | win.WS_EX_CONTROLPARENT
	if param.HideInTaskbar {
		exStyle &= ^win.WS_EX_APPWINDOW
		parent = _this.app.defOwner
	}
	if param.HideIcon {
		exStyle |= win.WS_EX_DLGMODALFRAME
	}
	win.CreateWindowEx(
		uint64(exStyle),
		sto16(_this.app.className),
		sto16(""),
		win.WS_OVERLAPPEDWINDOW,
		0, 0, 0, 0, parent, 0, _this.app.hInstance, unsafe.Pointer(_this))
	return _this
}

func (_this *winForm) GetChilds() []plat.Control {
	return _this._ctrls
}

func (_this *winForm) AddControl(control plat.Control) {
	control.SetParent(_this)
	_this._ctrls = append(_this._ctrls, control)
}

func (_this *winForm) RemoveControl(control plat.Control) {
	//if ctrl, ok := control.(*winControl); ok {
	//	if ctrl.IsCreate() {
	//
	//	}
	//	delete(_this._ctrls, ctrl.id())
	//}
}

func (_this *winForm) msgProc(hWnd win.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	rs := _this.winBase.msgProc(hWnd, msg, wParam, lParam)
	if rs != 0 {
		return rs
	}
	switch msg {
	case win.WM_CREATE:
		_this.SetSize(200, 300)
	case win.WM_SIZE:
		if _this._onState != nil {
			switch int(wParam) {
			case win.SIZE_RESTORED:
				_this._onState(f.FormState_Normal)
			case win.SIZE_MAXIMIZED:
				_this._onState(f.FormState_Max)
			case win.SIZE_MINIMIZED:
				_this._onState(f.FormState_Min)
			}
		}
	case win.WM_CLOSE:
		if _this._onClose != nil && _this._onClose() {
			rs = 1
		}
	}
	return rs
}

func (_this *winForm) NoneBorderResize() {
	padd := 5
	rsState := new(int)
	_this.app.watch(_this, func(hWnd win.HWND, msg uint32, wParam, lParam uintptr) uintptr {
		if _this._border != f.FormBorder_None {
			return 0
		}
		switch msg {
		case win.WM_MOUSEMOVE:
			if hWnd != _this.handle {
				if wnd, ok := _this.app.handleWnds[hWnd].(plat.Control); ok {
					if wnd.GetOwner().GetHandle() != uintptr(hWnd) {
						return 0
					}
				} else {
					return 0
				}
			}
			w, h := _this.GetSize()
			p := _this.MousePosition()
			if p.X <= padd {
				if p.Y <= padd {
					*rsState = 7
				} else if p.Y+padd >= h {
					*rsState = 1
				} else {
					*rsState = 4
				}
			} else if p.Y <= padd {
				if p.X <= padd {
					*rsState = 7
				} else if p.X+padd >= w {
					*rsState = 9
				} else {
					*rsState = 8
				}
			} else if p.X+padd >= w {
				if p.Y <= padd {
					*rsState = 9
				} else if p.Y+padd >= h {
					*rsState = 3
				} else {
					*rsState = 6
				}
			} else if p.Y+padd >= h {
				if p.X <= padd {
					*rsState = 1
				} else if p.X+padd >= w {
					*rsState = 3
				} else {
					*rsState = 2
				}
			} else {
				*rsState = 0
			}
		case win.WM_SETCURSOR:
			cur := f.CursorType_Default
			switch *rsState {
			case 8, 2:
				cur = f.CursorType_SIZENS
			case 4, 6:
				cur = f.CursorType_SIZEWE
			case 7, 3:
				cur = f.CursorType_SIZENWSE
			case 9, 1:
				cur = f.CursorType_SIZENESW
			}
			if cur != f.CursorType_Default {
				res := win.MAKEINTRESOURCE(uintptr(toWinCursor(cur)))
				win.SetCursor(win.LoadCursor(0, res))
				return 1
			} else {
				return 0
			}
		case win.WM_LBUTTONDOWN:
			switch *rsState {
			case 4:
				win.SendMessage(hWnd, win.WM_SYSCOMMAND, uintptr(win.SC_SIZE|0xF001), lParam)
			case 6:
				win.SendMessage(hWnd, win.WM_SYSCOMMAND, uintptr(win.SC_SIZE|0xF002), lParam)
			case 8:
				win.SendMessage(hWnd, win.WM_SYSCOMMAND, uintptr(win.SC_SIZE|0xF003), lParam)
			case 7:
				win.SendMessage(hWnd, win.WM_SYSCOMMAND, uintptr(win.SC_SIZE|0xF004), lParam)
			case 9:
				win.SendMessage(hWnd, win.WM_SYSCOMMAND, uintptr(win.SC_SIZE|0xF005), lParam)
			case 2:
				win.SendMessage(hWnd, win.WM_SYSCOMMAND, uintptr(win.SC_SIZE|0xF006), lParam)
			case 1:
				win.SendMessage(hWnd, win.WM_SYSCOMMAND, uintptr(win.SC_SIZE|0xF007), lParam)
			case 3:
				win.SendMessage(hWnd, win.WM_SYSCOMMAND, uintptr(win.SC_SIZE|0xF008), lParam)
			default:
				return 0
			}
			return 1
		}
		return 0
	})
}

func (_this *winForm) Show() {
	isMax := win.IsZoomed(_this.handle)
	isMin := win.IsIconic(_this.handle)
	if isMax || isMin {
		win.ShowWindow(_this.handle, win.SW_RESTORE)
	} else {
		win.ShowWindow(_this.handle, win.SW_SHOW)
	}
	win.UpdateWindow(_this.handle)
}

func (_this *winForm) Close() {
	win.SendMessage(_this.handle, win.WM_CLOSE, 0, 0)
}

func (_this *winForm) ShowDialog() {
	//todo miss
}

func (_this *winForm) ShowToMax() {
	win.ShowWindow(_this.handle, win.SW_MAXIMIZE)
	win.UpdateWindow(_this.handle)
}

func (_this *winForm) ShowToMin() {
	win.ShowWindow(_this.handle, win.SW_MINIMIZE)
}

func (_this *winForm) SetTitle(title string) {
	win.SetWindowText(_this.handle, title)
}

func (_this *winForm) SetBorderStyle(border f.FormBorder) {
	style := win.GetWindowLong(_this.handle, win.GWL_STYLE)
	switch border {
	case f.FormBorder_Default:
		style |= win.WS_OVERLAPPEDWINDOW
	case f.FormBorder_None:
		style &= ^win.WS_SIZEBOX & ^win.WS_CAPTION
	case f.FormBorder_Disable_Resize:
		style &= ^win.WS_SIZEBOX
	}
	win.SetWindowLong(_this.handle, win.GWL_STYLE, style)
	_this._border = border
}

func (_this *winForm) SetOnState(proc plat.FormStateProc) plat.FormStateProc {
	pre := _this._onState
	_this._onState = proc
	return pre
}

func (_this *winForm) SetMaximizeBox(isShow bool) {
	style := win.GetWindowLong(_this.handle, win.GWL_STYLE)
	if isShow {
		style |= win.WS_MAXIMIZEBOX
	} else {
		style &= ^win.WS_MAXIMIZEBOX
	}
	win.SetWindowLong(_this.handle, win.GWL_STYLE, style)
}

func (_this *winForm) SetMinimizeBox(isShow bool) {
	style := win.GetWindowLong(_this.handle, win.GWL_STYLE)
	if isShow {
		style |= win.WS_MINIMIZEBOX
	} else {
		style &= ^win.WS_MINIMIZEBOX
	}
	win.SetWindowLong(_this.handle, win.GWL_STYLE, style)
}

func (_this *winForm) SetIcon(iconFile string) {
	style := win.GetWindowLong(_this.handle, win.GWL_EXSTYLE)
	if style&win.WS_EX_DLGMODALFRAME != 0 {
		return
	}
	h := win.LoadImage(_this.app.hInstance, sto16(iconFile), win.IMAGE_ICON, 0, 0, win.LR_LOADFROMFILE)
	if h != 0 {
		win.SendMessage(_this.handle, win.WM_SETICON, 1, uintptr(h))
		win.SendMessage(_this.handle, win.WM_SETICON, 0, uintptr(h))
	}
}
