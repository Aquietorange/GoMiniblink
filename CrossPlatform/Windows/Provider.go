package Windows

import (
	MB "GoMiniblink"
	"GoMiniblink/CrossPlatform"
	"GoMiniblink/CrossPlatform/Windows/win32"
	"os"
	"syscall"
	"unsafe"
)

type Provider struct {
	hInstance  win32.HINSTANCE
	className  string
	main       string
	handleWnds map[win32.HWND]baseWindow
	nameWnds   map[string]baseWindow
	defOwner   win32.HWND
	defIcon    win32.HICON
}

func (_this *Provider) Init() *Provider {
	_this.handleWnds = make(map[win32.HWND]baseWindow)
	_this.nameWnds = make(map[string]baseWindow)
	_this.className = "GoMiniblinkForms"
	_this.hInstance = win32.GetModuleHandle(nil)
	return _this
}

func (_this *Provider) GetScreen() MB.Screen {
	var s = MB.Screen{
		Full: MB.Rect{
			Wdith:  int(win32.GetSystemMetrics(win32.SM_CXSCREEN)),
			Height: int(win32.GetSystemMetrics(win32.SM_CYSCREEN)),
		},
		WorkArea: MB.Rect{
			Wdith:  int(win32.GetSystemMetrics(win32.SM_CXFULLSCREEN)),
			Height: int(win32.GetSystemMetrics(win32.SM_CYFULLSCREEN)),
		},
	}
	return s
}

func (_this *Provider) SetIcon(file string) {
	h := win32.LoadImage(_this.hInstance, sto16(file), win32.IMAGE_ICON, 0, 0, win32.LR_LOADFROMFILE)
	_this.defIcon = win32.HICON(h)
}

func (_this *Provider) registerWndClass() {
	var class = win32.WNDCLASSEX{
		Style:         win32.CS_HREDRAW | win32.CS_VREDRAW,
		LpfnWndProc:   syscall.NewCallback(_this.defaultWndProc),
		HInstance:     _this.hInstance,
		LpszClassName: sto16(_this.className),
		HIcon:         _this.defIcon,
		HIconSm:       _this.defIcon,
	}
	class.CbSize = uint32(unsafe.Sizeof(class))
	win32.RegisterClassEx(&class)
	_this.defOwner = win32.CreateWindowEx(0, sto16(_this.className), sto16(""),
		win32.WS_BORDER, 0, 0, 0, 0,
		0, 0, _this.hInstance, unsafe.Pointer(nil))
}

func (_this *Provider) add(wnd baseWindow) {
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
				w.fireWndCreate(hWnd)
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
	_this.registerWndClass()
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
