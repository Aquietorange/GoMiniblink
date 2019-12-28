package Windows

import (
	"GoMiniblink/Forms/CrossPlatform"
	"GoMiniblink/Forms/CrossPlatform/Windows/win32"
	"os"
	"syscall"
	"unsafe"
)

type Provider struct {
	hInstance  win32.HINSTANCE
	className  string
	main       string
	handleWnds map[win32.HWND]IWindow
	nameWnds   map[string]IWindow
	defOwner   win32.HWND
}

func (_this *Provider) Init() *Provider {
	_this.handleWnds = make(map[win32.HWND]IWindow)
	_this.nameWnds = make(map[string]IWindow)
	_this.className = "GoMiniblinkForms"
	_this.hInstance = win32.GetModuleHandle(nil)
	_this.registerWndClass()
	_this.defOwner = win32.CreateWindowEx(0, sto16(_this.className), sto16(""),
		win32.WS_BORDER, 0, 0, 0, 0,
		0, 0, _this.hInstance, unsafe.Pointer(nil))
	return _this
}

func (_this *Provider) registerWndClass() {
	var class = win32.WNDCLASSEX{
		Style:         win32.CS_HREDRAW | win32.CS_VREDRAW,
		LpfnWndProc:   syscall.NewCallback(_this.defaultWndProc),
		HInstance:     _this.hInstance,
		LpszClassName: sto16(_this.className),
	}
	class.CbSize = uint32(unsafe.Sizeof(class))
	win32.RegisterClassEx(&class)
}

func (_this *Provider) add(wnd IWindow) {
	_this.nameWnds[wnd.name()] = wnd
}

func (_this *Provider) remove(hWnd win32.HWND, isExit bool) {
	if w, ok := _this.handleWnds[hWnd]; ok {
		delete(_this.nameWnds, w.name())
		delete(_this.handleWnds, hWnd)
		if isExit && w.name() == _this.main {
			_this.Exit(0)
		}
	}
}

func (_this *Provider) defaultWndProc(hWnd win32.HWND, msg uint32, wParam uintptr, lParam uintptr) uintptr {
	if msg == win32.WM_CREATE {
		cp := *((*win32.CREATESTRUCT)(unsafe.Pointer(lParam)))
		if cp.CreateParams != 0 {
			id := *((*string)(unsafe.Pointer(cp.CreateParams)))
			if w, ok := _this.nameWnds[id]; ok {
				_this.handleWnds[hWnd] = w
				w.onWndCreate(hWnd)
			}
		}
	} else if w, ok := _this.handleWnds[hWnd]; ok {
		ret := w.fireWndProc(hWnd, msg, wParam, lParam)
		if ret != 0 {
			return ret
		}
	}
	return win32.DefWindowProc(hWnd, msg, wParam, lParam)
}

func (_this *Provider) Exit(code int) {
	win32.PostQuitMessage(int32(code))
}

func (_this *Provider) RunMain(form CrossPlatform.IForm, show func()) {
	frm, ok := form.(*winForm)
	if ok == false {
		panic("类型不正确")
	}
	_this.main = frm.name()
	show()
	var message win32.MSG
	for {
		if win32.GetMessage(&message, 0, 0, 0) {
			win32.TranslateMessage(&message)
			win32.DispatchMessage(&message)
		} else {
			break
		}
	}
	os.Exit(0)
}
