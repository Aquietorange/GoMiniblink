package Windows

import (
	"GoMiniblink/Forms/CrossPlatform"
	"GoMiniblink/Forms/CrossPlatform/Windows/win32"
	"os"
	"syscall"
	"unsafe"
)

type Provider struct {
	hInstance win32.HINSTANCE
	className string
	uncreate  []IWindow
	created   map[win32.HWND]*IWindow
}

func (_this *Provider) Init() *Provider {
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

func (_this *Provider) add(window IWindow) {
	_this.uncreate = append(_this.uncreate, window)
}

func (_this *Provider) remove(hWnd win32.HWND) {

}

func (_this *Provider) defaultWndProc(hWnd win32.HWND, msg uint32, wParam uintptr, lParam uintptr) uintptr {
	if msg == win32.WM_CREATE {
		println(hWnd)
	}
	for _, v := range _this.uncreate {
		if v.hWnd() == hWnd {
			ret := v.wndProc(hWnd, msg, wParam, lParam)
			if ret != 0 {
				return ret
			}
			break
		}
	}
	return win32.DefWindowProc(hWnd, msg, wParam, lParam)
}

func (_this *Provider) RunMain(form CrossPlatform.IForm) {
	frm, ok := form.(*winForm)
	if ok == false {
		panic("类型不正确")
	}
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
