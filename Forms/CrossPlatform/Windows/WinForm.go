package Windows

import (
	"GoMiniblink/Forms/CrossPlatform/Windows/win32"
	"GoMiniblink/Utils"
	"unsafe"
)

type winForm struct {
	winControl
	onResize func(int, int)
	onMove   func(int, int)
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
	_this.provider = provider
	_this.idName = Utils.NewUUID()
	_this.className = provider.className
	_this.defWndProc = _this.defWndProcImpl
	_this.owner = _this
	provider.add(_this)
	return _this
}

func (_this *winForm) defWndProcImpl(hWnd win32.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {
	case win32.WM_DESTROY:
		win32.PostQuitMessage(0)
		break
	case win32.WM_SIZE:
		w, h := win32.GET_X_LPARAM(lParam), win32.GET_Y_LPARAM(lParam)
		if _this.onResize != nil {
			_this.onResize(int(w), int(h))
		}
		break
	case win32.WM_MOVE:
		x, y := win32.GET_X_LPARAM(lParam), win32.GET_Y_LPARAM(lParam)
		if _this.onMove != nil {
			_this.onMove(int(x), int(y))
		}
		break
	}
	return 0
}

func (_this *winForm) fireWndProc(hWnd win32.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	return _this.defWndProc(hWnd, msg, wParam, lParam)
}

func (_this *winForm) Show() {
	if _this.isCreated == false {
		hWnd := win32.CreateWindowEx(
			win32.WS_TILED,
			sto16(_this.className),
			sto16(""),
			win32.WS_EX_CLIENTEDGE,
			int32(_this.x),
			int32(_this.y),
			int32(_this.width),
			int32(_this.height),
			0, 0,
			_this.provider.hInstance,
			unsafe.Pointer(nil))
		_this.onCreate(hWnd)
	}
	win32.ShowWindow(_this.hWnd(), win32.SW_SHOW)
	win32.UpdateWindow(_this.hWnd())
}

func (_this *winForm) Hide() {
	win32.ShowWindow(_this.hWnd(), win32.SW_HIDE)
}

func (_this *winForm) SetSize(w, h int) {
	win32.SetWindowPos(_this.hWnd(), win32.HWND(0), 0, 0, int32(w), int32(h), win32.SWP_NOMOVE)
}

func (_this *winForm) SetLocation(x, y int) {
	win32.SetWindowPos(_this.hWnd(), win32.HWND(0), int32(_this.x), int32(_this.y), 0, 0, win32.SWP_NOSIZE)
}

func (_this *winForm) SetTitle(title string) {

}

func (_this *winForm) ShowDialog() {

}
