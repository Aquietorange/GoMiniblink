package Windows

import (
	"GoMiniblink/Forms/CrossPlatform"
	"GoMiniblink/Forms/CrossPlatform/Windows/win32"
	"os"
	"syscall"
	"unsafe"
)

type Provider struct {
	hInstance    win32.HINSTANCE
	className    string
	main         win32.HWND
	wnds         map[win32.HWND]IWindow
	unCreateWnds map[string]IWindow
}

func (_this *Provider) Init() *Provider {
	_this.wnds = make(map[win32.HWND]IWindow)
	_this.unCreateWnds = make(map[string]IWindow)
	_this.className = "GoMiniblinkForms"
	_this.hInstance = win32.GetModuleHandle(nil)
	_this.registerWndClass()
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
	_this.unCreateWnds[wnd.name()] = wnd
}

func (_this *Provider) remove(hWnd win32.HWND) {
	delete(_this.wnds, hWnd)
}

func (_this *Provider) defaultWndProc(hWnd win32.HWND, msg uint32, wParam uintptr, lParam uintptr) uintptr {
	if msg == win32.WM_DESTROY {
		if _this.main == hWnd {
			_this.Exit(0)
		}
	} else if w, ok := _this.wnds[hWnd]; ok {
		ret := w.fireWndProc(hWnd, msg, wParam, lParam)
		if ret != 0 {
			return ret
		}
	} else if msg == win32.WM_CREATE {
		cp := *((*win32.CREATESTRUCT)(unsafe.Pointer(lParam)))
		id := *((*string)(unsafe.Pointer(cp.CreateParams)))
		if w, ok := _this.unCreateWnds[id]; ok {
			w.onWndCreate(hWnd)
			delete(_this.unCreateWnds, id)
			_this.wnds[hWnd] = w
		}
	}
	return win32.DefWindowProc(hWnd, msg, wParam, lParam)
}

func (_this *Provider) Exit(code int) {
	win32.PostQuitMessage(int32(code))
}

func (_this *Provider) RunMain(form CrossPlatform.IForm) {
	frm, ok := form.(*winForm)
	if ok == false {
		panic("类型不正确")
	}
	frm.Create()
	_this.main = frm.hWnd()
	frm.Show()
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
