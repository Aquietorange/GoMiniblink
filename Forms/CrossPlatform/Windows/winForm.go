package Windows

import (
	plat "GoMiniblink/Forms/CrossPlatform"
	"GoMiniblink/Forms/CrossPlatform/Windows/win32"
	"GoMiniblink/Utils"
	"unsafe"
)

type winForm struct {
	winControl
	onCreate func()
	onResize func(int, int)
	onMove   func(int, int)
	onClose  func() (cancel bool)

	hideParent win32.HWND
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
	_this.owner = _this
	_this.evWndProc["WinformWndProc"] = _this.defWndProc
	provider.add(_this)
	return _this
}

func (_this *winForm) defWndProc(hWnd win32.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {
	case win32.WM_CLOSE:
		if _this.onClose == nil || _this.onClose() == false {
			_this.provider.remove(_this.hWnd())
			return 0
		}
	case win32.WM_SIZE:
		w, h := win32.GET_X_LPARAM(lParam), win32.GET_Y_LPARAM(lParam)
		if _this.onResize != nil {
			_this.onResize(int(w), int(h))
		}
	case win32.WM_MOVE:
		x, y := win32.GET_X_LPARAM(lParam), win32.GET_Y_LPARAM(lParam)
		if _this.onMove != nil {
			_this.onMove(int(x), int(y))
		}
	default:
		return 0
	}
	return 1
}

func (_this *winForm) SetOnCreate(fn func()) {
	_this.onCreate = fn
}

func (_this *winForm) Create() {
	if _this.IsCreate() {
		return
	}
	_this.evWndCreate["__onCreate"] = func(hWnd win32.HWND) {
		if _this.onCreate != nil {
			_this.onCreate()
		}
	}
	lp := _this.name()
	win32.CreateWindowEx(
		0,
		sto16(_this.className),
		sto16(""),
		win32.WS_OVERLAPPEDWINDOW,
		0, 0, 0, 0, 0, 0,
		_this.provider.hInstance,
		unsafe.Pointer(&lp))
}

func (_this *winForm) Show() {
	win32.ShowWindow(_this.hWnd(), win32.SW_SHOW)
	win32.UpdateWindow(_this.hWnd())
}

func (_this *winForm) ShowDialog() {

}

func (_this *winForm) Hide() {
	win32.ShowWindow(_this.hWnd(), win32.SW_HIDE)
}

func (_this *winForm) SetSize(w, h int) {
	win32.SetWindowPos(_this.hWnd(), win32.HWND(0), 0, 0, int32(w), int32(h), win32.SWP_NOMOVE)
}

func (_this *winForm) SetOnResize(fn func(w, h int)) {
	_this.onResize = fn
}

func (_this *winForm) SetLocation(x, y int) {
	win32.SetWindowPos(_this.hWnd(), win32.HWND(0), int32(x), int32(y), 0, 0, win32.SWP_NOSIZE)
}

func (_this *winForm) SetOnMove(fn func(x, y int)) {
	_this.onMove = fn
}

func (_this *winForm) SetTitle(title string) {
	win32.SetWindowText(_this.hWnd(), sto16(title))
}

func (_this *winForm) SetBorderStyle(border plat.IFormBorder) {
	style := win32.GetWindowLong(_this.hWnd(), win32.GWL_STYLE)
	bak := style
	switch border {
	case plat.IFormBorder_Default:
		style |= win32.WS_OVERLAPPEDWINDOW
	case plat.IFormBorder_None:
		style &= ^win32.WS_SIZEBOX & ^win32.WS_CAPTION
	case plat.IFormBorder_Disable_Resize:
		style |= win32.WS_OVERLAPPEDWINDOW
		style &= ^win32.WS_SIZEBOX
	}
	if bak != style {
		win32.SetWindowLong(_this.hWnd(), win32.GWL_STYLE, style)
	}
}

func (_this *winForm) ShowInTaskbar(isShow bool) {
	style := win32.GetWindowLong(_this.hWnd(), win32.GWL_STYLE)
	bak := style
	if isShow {
		style |= win32.WS_EX_APPWINDOW
	} else {
		style &= ^win32.WS_EX_APPWINDOW
		//style |= win32.WS_POPUP
	}
	if bak != style {
		if _this.hideParent == 0 {
			_this.hideParent = win32.CreateWindowEx(
				0,
				sto16(_this.className),
				sto16(""),
				win32.WS_OVERLAPPEDWINDOW,
				0, 0, 0, 0, 0, 0,
				_this.provider.hInstance,
				unsafe.Pointer(nil))
			win32.SetParent(_this.hWnd(), _this.hideParent)
		}
		win32.SetWindowLong(_this.hWnd(), win32.GWL_STYLE, style)
	}
}
