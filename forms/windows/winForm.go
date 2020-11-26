package windows

import (
	fm "gitee.com/aochulai/GoMiniblink/forms"
	br "gitee.com/aochulai/GoMiniblink/forms/bridge"
	win "gitee.com/aochulai/GoMiniblink/forms/windows/win32"
	"unsafe"
)

type winForm struct {
	winContainer
	_onClose  func() (cancel bool)
	_onState  br.FormStateProc
	_onActive br.FormActiveProc
	border    fm.FormBorder
	_isModal  bool
}

func (_this *winForm) init(provider *Provider, param br.FormParam) *winForm {
	_this.winContainer.init(provider, _this)
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
	x := 100 + len(_this.app.forms)*25
	y := 100 + len(_this.app.forms)*25
	win.CreateWindowEx(
		uint64(exStyle),
		sto16(_this.app.className),
		sto16(""),
		win.WS_OVERLAPPEDWINDOW,
		int32(x), int32(y), 200, 300, parent, 0, _this.app.hInstance, unsafe.Pointer(_this))
	return _this
}

func (_this *winForm) msgProc(hWnd win.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	rs := _this.winBase.msgProc(hWnd, msg, wParam, lParam)
	if rs != 0 {
		return rs
	}
	switch msg {
	case win.WM_CREATE:
		_this.app.forms[hWnd] = _this
	case win.WM_DESTROY:
		if _this._isModal {
			win.PostQuitMessage(0)
			return 1
		}
	case win.WM_SIZE:
		if _this._onState != nil {
			switch int(wParam) {
			case win.SIZE_RESTORED:
				_this._onState(fm.FormState_Normal)
			case win.SIZE_MAXIMIZED:
				_this._onState(fm.FormState_Max)
			case win.SIZE_MINIMIZED:
				_this._onState(fm.FormState_Min)
			}
		}
	case win.WM_ACTIVATE:
		if _this._onActive != nil {
			_this._onActive()
		}
	case win.WM_CLOSE:
		if _this._onClose != nil && _this._onClose() {
			rs = 1
		}
		if rs != 0 && _this._isModal {
			_this._isModal = false
		}
	}
	return rs
}

func (_this *winForm) SetOnActive(proc br.FormActiveProc) br.FormActiveProc {
	pre := _this._onActive
	_this._onActive = proc
	return pre
}

func (_this *winForm) NoneBorderResize() {
	padd := 5
	rsState := new(int)
	_this.app.watch(_this, func(hWnd win.HWND, msg uint32, wParam, lParam uintptr) uintptr {
		if _this.border != fm.FormBorder_None {
			return 0
		}
		switch msg {
		case win.WM_MOUSEMOVE:
			if hWnd != _this.handle {
				if wnd, ok := _this.app.handleWnds[hWnd].(br.Control); ok {
					if wnd.GetOwner().GetHandle() != uintptr(_this.handle) {
						return 0
					}
				} else {
					return 0
				}
			}
			sz := _this.GetBound().Rect
			p := _this.ToClientPoint(_this.app.MouseLocation())
			if p.X <= padd {
				if p.Y <= padd {
					*rsState = 7
				} else if p.Y+padd >= sz.Height {
					*rsState = 1
				} else {
					*rsState = 4
				}
			} else if p.Y <= padd {
				if p.X <= padd {
					*rsState = 7
				} else if p.X+padd >= sz.Width {
					*rsState = 9
				} else {
					*rsState = 8
				}
			} else if p.X+padd >= sz.Width {
				if p.Y <= padd {
					*rsState = 9
				} else if p.Y+padd >= sz.Height {
					*rsState = 3
				} else {
					*rsState = 6
				}
			} else if p.Y+padd >= sz.Height {
				if p.X <= padd {
					*rsState = 1
				} else if p.X+padd >= sz.Width {
					*rsState = 3
				} else {
					*rsState = 2
				}
			} else {
				*rsState = 0
			}
		case win.WM_SETCURSOR:
			cur := fm.CursorType_Default
			switch *rsState {
			case 8, 2:
				cur = fm.CursorType_SIZENS
			case 4, 6:
				cur = fm.CursorType_SIZEWE
			case 7, 3:
				cur = fm.CursorType_SIZENWSE
			case 9, 1:
				cur = fm.CursorType_SIZENESW
			}
			if cur != fm.CursorType_Default {
				res := win.MAKEINTRESOURCE(uintptr(toWinCursor(cur)))
				win.SetCursor(win.LoadCursor(0, res))
				return 1
			} else {
				return 0
			}
		case win.WM_LBUTTONDOWN:
			switch *rsState {
			case 4:
				win.SendMessage(_this.handle, win.WM_SYSCOMMAND, uintptr(win.SC_SIZE|0xF001), lParam)
			case 6:
				win.SendMessage(_this.handle, win.WM_SYSCOMMAND, uintptr(win.SC_SIZE|0xF002), lParam)
			case 8:
				win.SendMessage(_this.handle, win.WM_SYSCOMMAND, uintptr(win.SC_SIZE|0xF003), lParam)
			case 7:
				win.SendMessage(_this.handle, win.WM_SYSCOMMAND, uintptr(win.SC_SIZE|0xF004), lParam)
			case 9:
				win.SendMessage(_this.handle, win.WM_SYSCOMMAND, uintptr(win.SC_SIZE|0xF005), lParam)
			case 2:
				win.SendMessage(_this.handle, win.WM_SYSCOMMAND, uintptr(win.SC_SIZE|0xF006), lParam)
			case 1:
				win.SendMessage(_this.handle, win.WM_SYSCOMMAND, uintptr(win.SC_SIZE|0xF007), lParam)
			case 3:
				win.SendMessage(_this.handle, win.WM_SYSCOMMAND, uintptr(win.SC_SIZE|0xF008), lParam)
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
	acHwnd := win.GetActiveWindow()
	if acfm, ok := _this.app.forms[acHwnd].(*winForm); ok {
		acfm.Enable(false)
		_this._isModal = true
		_this.Show()
		var msg win.MSG
		for {
			if win.GetMessage(&msg, 0, 0, 0) && _this._isModal {
				win.TranslateMessage(&msg)
				win.DispatchMessage(&msg)
			} else {
				break
			}
		}
		acfm.Enable(true)
		acfm.Active()
	}
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

func (_this *winForm) SetBorderStyle(border fm.FormBorder) {
	style := win.GetWindowLong(_this.handle, win.GWL_STYLE)
	switch border {
	case fm.FormBorder_Default:
		style |= win.WS_OVERLAPPEDWINDOW
	case fm.FormBorder_None:
		style &= ^win.WS_SIZEBOX & ^win.WS_CAPTION
	case fm.FormBorder_Disable_Resize:
		style &= ^win.WS_SIZEBOX
	}
	win.SetWindowLong(_this.handle, win.GWL_STYLE, style)
	bn := _this.GetBound()
	_this.SetSize(bn.Width, bn.Height-1)
	_this.SetSize(bn.Width, bn.Height)
	_this.border = border
}

func (_this *winForm) SetOnState(proc br.FormStateProc) br.FormStateProc {
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

func (_this *winForm) Enable(b bool) {
	_this.isEnable = b
	win.EnableWindow(_this.handle, b)
}

func (_this *winForm) Active() {
	win.SetActiveWindow(_this.handle)
}
