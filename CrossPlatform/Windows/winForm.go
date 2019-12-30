package Windows

import (
	MB "GoMiniblink"
	"GoMiniblink/CrossPlatform/Windows/win32"
	"GoMiniblink/Utils"
	"unsafe"
)

type winForm struct {
	winControl
	onCreate func()
	onResize func(int, int)
	onMove   func(int, int)
	onClose  func() (cancel bool)
	onState  func(state MB.FormState)

	createParams win32.CREATESTRUCT
	initState    MB.FormState
}

func (_this *winForm) hWnd() win32.HWND {
	return _this.handle
}

func (_this *winForm) class() string {
	return _this.className
}

func (_this *winForm) name() string {
	return _this.idName
}

func (_this *winForm) init(provider *Provider) *winForm {
	_this.winControl.init()
	_this.provider = provider
	_this.idName = Utils.NewUUID()
	_this.className = provider.className
	_this.createParams = win32.CREATESTRUCT{
		Style:     win32.WS_SIZEBOX | win32.WS_CAPTION | win32.WS_SYSMENU | win32.WS_MAXIMIZEBOX | win32.WS_MINIMIZEBOX,
		ClassName: uintptr(unsafe.Pointer(sto16(_this.className))),
		Name:      uintptr(unsafe.Pointer(sto16(""))),
	}
	_this.initState = MB.FormState_Normal
	_this.evWndProc["WinformWndProc"] = _this.defWndProc
	_this.evWndCreate["__onCreate"] = func(hWnd win32.HWND) {
		if _this.onCreate != nil {
			_this.onCreate()
		}
	}
	provider.add(_this)
	return _this
}

func (_this *winForm) defWndProc(hWnd win32.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {
	case win32.WM_CLOSE:
		if _this.onClose != nil && _this.onClose() {
			return 1
		}
	case win32.WM_DESTROY:
		_this.provider.remove(_this.hWnd(), true)
	case win32.WM_SIZE:
		if _this.onResize != nil {
			w, h := win32.GET_X_LPARAM(lParam), win32.GET_Y_LPARAM(lParam)
			_this.onResize(int(w), int(h))
		}
		if _this.onState != nil {
			switch int(wParam) {
			case win32.SIZE_RESTORED:
				_this.onState(MB.FormState_Normal)
			case win32.SIZE_MAXIMIZED:
				_this.onState(MB.FormState_Max)
			case win32.SIZE_MINIMIZED:
				_this.onState(MB.FormState_Min)
			}
		}
	case win32.WM_MOVE:
		if _this.onMove != nil {
			x, y := win32.GET_X_LPARAM(lParam), win32.GET_Y_LPARAM(lParam)
			_this.onMove(int(x), int(y))
		}
	}
	return 0
}

func (_this *winForm) SetOnCreate(fn func()) {
	_this.onCreate = fn
}

func (_this *winForm) Create() {
	if _this.IsCreate() {
		return
	}
	win32.CreateWindowEx(
		_this.createParams.ExStyle,
		(*uint16)(unsafe.Pointer(_this.createParams.ClassName)),
		(*uint16)(unsafe.Pointer(_this.createParams.Name)),
		_this.createParams.Style,
		_this.createParams.X,
		_this.createParams.Y,
		_this.createParams.Cx,
		_this.createParams.Cy,
		_this.createParams.Parent,
		_this.createParams.Menu,
		_this.provider.hInstance,
		unsafe.Pointer(&_this.idName))
}

func (_this *winForm) Show() {
	switch _this.initState {
	case MB.FormState_Normal:
		win32.ShowWindow(_this.hWnd(), win32.SW_SHOW)
	case MB.FormState_Max:
		win32.ShowWindow(_this.hWnd(), win32.SW_MAXIMIZE)
	case MB.FormState_Min:
		win32.ShowWindow(_this.hWnd(), win32.SW_MINIMIZE)
	}
	win32.UpdateWindow(_this.hWnd())
}

func (_this *winForm) Hide() {
	win32.ShowWindow(_this.hWnd(), win32.SW_HIDE)
}

func (_this *winForm) ShowDialog() {

}

func (_this *winForm) SetSize(w, h int) {
	if _this.hWnd() == 0 {
		_this.createParams.Cx, _this.createParams.Cy = int32(w), int32(h)
	} else {
		win32.SetWindowPos(_this.hWnd(), win32.HWND(0), 0, 0, int32(w), int32(h), win32.SWP_NOMOVE)
	}
}

func (_this *winForm) SetOnResize(fn func(w, h int)) {
	_this.onResize = fn
}

func (_this *winForm) SetLocation(x, y int) {
	if _this.hWnd() == 0 {
		_this.createParams.X, _this.createParams.Y = int32(x), int32(y)
	} else {
		win32.SetWindowPos(_this.hWnd(), win32.HWND(0), int32(x), int32(y), 0, 0, win32.SWP_NOSIZE)
	}
}

func (_this *winForm) SetOnMove(fn func(x, y int)) {
	_this.onMove = fn
}

func (_this *winForm) SetTitle(title string) {
	if _this.hWnd() == 0 {
		_this.createParams.Name = uintptr(unsafe.Pointer(sto16(title)))
	} else {
		win32.SetWindowText(_this.hWnd(), title)
	}
}

func (_this *winForm) SetBorderStyle(border MB.FormBorder) {
	var style int64
	if _this.hWnd() == 0 {
		style = _this.createParams.Style
	} else {
		style = win32.GetWindowLong(_this.hWnd(), win32.GWL_STYLE)
	}
	bak := style
	switch border {
	case MB.FormBorder_Default:
		style |= win32.WS_SIZEBOX | win32.WS_CAPTION | win32.WS_SYSMENU | win32.WS_MAXIMIZEBOX | win32.WS_MINIMIZEBOX
	case MB.FormBorder_None:
		style &= ^win32.WS_SIZEBOX & ^win32.WS_CAPTION
	case MB.FormBorder_Disable_Resize:
		style &= ^win32.WS_SIZEBOX
	}
	if _this.hWnd() == 0 {
		_this.createParams.Style = style
	} else if bak != style {
		win32.SetWindowLong(_this.hWnd(), win32.GWL_STYLE, style)
	}
}

func (_this *winForm) ShowInTaskbar(isShow bool) {
	if _this.hWnd() == 0 {
		if isShow {
			_this.createParams.Style &= ^win32.WS_POPUP
			_this.createParams.Parent = 0
		} else {
			_this.createParams.Style |= win32.WS_POPUP
			_this.createParams.Parent = _this.provider.defOwner
		}
		return
	}
	style := win32.GetWindowLong(_this.hWnd(), win32.GWL_STYLE)
	bak := style
	if isShow {
		style &= ^win32.WS_POPUP
	} else {
		style |= win32.WS_POPUP
	}
	if bak != style {
		preHwnd := _this.hWnd()
		visible := win32.IsWindowVisible(preHwnd)
		var rect win32.RECT
		win32.GetWindowRect(preHwnd, &rect)
		title := sto16(win32.GetWindowText(preHwnd))
		_this.createParams.Style = style
		_this.createParams.Name = uintptr(unsafe.Pointer(title))
		_this.createParams.X = rect.Left
		_this.createParams.Y = rect.Top
		_this.createParams.Cx = rect.Right - rect.Left
		_this.createParams.Cy = rect.Bottom - rect.Top
		if isShow {
			_this.createParams.Parent = 0
		} else {
			_this.createParams.Parent = _this.provider.defOwner
		}
		_this.isCreated = false
		_this.Create()
		_this.provider.remove(preHwnd, false)
		win32.DestroyWindow(preHwnd)
		if visible {
			win32.ShowWindow(_this.hWnd(), win32.SW_SHOW)
		}
	}
}

func (_this *winForm) SetState(state MB.FormState) {
	if _this.isCreated {
		isMax := win32.IsZoomed(_this.hWnd())
		isMin := win32.IsIconic(_this.hWnd())
		switch state {
		case MB.FormState_Normal:
			if isMax || isMin {
				win32.ShowWindow(_this.hWnd(), win32.SW_RESTORE)
			}
		case MB.FormState_Max:
			if isMax == false {
				win32.ShowWindow(_this.hWnd(), win32.SW_MAXIMIZE)
			}
		case MB.FormState_Min:
			if isMin == false {
				win32.ShowWindow(_this.hWnd(), win32.SW_MINIMIZE)
			}
		}
	} else {
		_this.initState = state
	}
}

func (_this *winForm) SetOnState(fn func(state MB.FormState)) {
	_this.onState = fn
}

func (_this *winForm) SetMaximizeBox(isShow bool) {
	var style int64
	if _this.hWnd() == 0 {
		style = _this.createParams.Style
	} else {
		style = win32.GetWindowLong(_this.hWnd(), win32.GWL_STYLE)
	}
	bak := style
	if isShow {
		style |= win32.WS_MAXIMIZEBOX
	} else {
		style &= ^win32.WS_MAXIMIZEBOX
	}
	if _this.hWnd() == 0 {
		_this.createParams.Style = style
	} else if bak != style {
		win32.SetWindowLong(_this.hWnd(), win32.GWL_STYLE, style)
	}
}

func (_this *winForm) SetMinimizeBox(isShow bool) {
	var style int64
	if _this.hWnd() == 0 {
		style = _this.createParams.Style
	} else {
		style = win32.GetWindowLong(_this.hWnd(), win32.GWL_STYLE)
	}
	bak := style
	if isShow {
		style |= win32.WS_MINIMIZEBOX
	} else {
		style &= ^win32.WS_MINIMIZEBOX
	}
	if _this.hWnd() == 0 {
		_this.createParams.Style = style
	} else if bak != style {
		win32.SetWindowLong(_this.hWnd(), win32.GWL_STYLE, style)
	}
}
