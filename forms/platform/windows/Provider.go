package windows

import (
	"golang.org/x/sys/windows"
	"os"
	"path/filepath"
	f "qq2564874169/goMiniblink/forms"
	p "qq2564874169/goMiniblink/forms/platform"
	w "qq2564874169/goMiniblink/forms/platform/windows/win32"
	"reflect"
	"syscall"
	"unsafe"
)

type windowsMsgProc func(hWnd w.HWND, msg uint32, wParam, lParam uintptr) uintptr

type Provider struct {
	hInstance    w.HINSTANCE
	className    string
	wndClass     w.WNDCLASSEX
	mainId       w.HWND
	tmpWnd       map[uintptr]baseWindow
	handleWnds   map[w.HWND]baseWindow
	nameWnds     map[string]baseWindow
	dlgOwner     w.HWND
	defIcon      w.HICON
	mainThreadId uint32
	watchAll     map[w.HWND]windowsMsgProc
}

func (_this *Provider) Init() *Provider {
	_this.watchAll = make(map[w.HWND]windowsMsgProc)
	_this.tmpWnd = make(map[uintptr]baseWindow)
	_this.handleWnds = make(map[w.HWND]baseWindow)
	_this.nameWnds = make(map[string]baseWindow)
	_this.className = "goMiniblinkClass"
	_this.hInstance = w.GetModuleHandle(nil)
	_this.mainThreadId = windows.GetCurrentThreadId()
	_this.registerWndClass()
	return _this
}

func (_this *Provider) watch(wnd baseWindow, proc windowsMsgProc) {
	_this.watchAll[wnd.hWnd()] = proc
}

func (_this *Provider) unWatch(wnd baseWindow) {
	delete(_this.watchAll, wnd.hWnd())
}

func (_this *Provider) MouseLocation() f.Point {
	pos := w.POINT{}
	w.GetCursorPos(&pos)
	return f.Point{
		X: int(pos.X),
		Y: int(pos.Y),
	}
}

func (_this *Provider) AppDir() string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return dir
}

func (_this *Provider) ModifierKeys() map[f.Keys]bool {
	keys := make(map[f.Keys]bool)
	cs := w.GetKeyState(int32(w.VK_CONTROL))
	ss := w.GetKeyState(int32(w.VK_SHIFT))
	as := w.GetKeyState(int32(w.VK_MENU))
	keys[f.Keys_Ctrl] = cs < 0
	keys[f.Keys_Shift] = ss < 0
	keys[f.Keys_Alt] = as < 0
	return keys
}

func (_this *Provider) MouseIsDown() map[f.MouseButtons]bool {
	keys := make(map[f.MouseButtons]bool)
	ls := w.GetKeyState(int32(w.VK_LBUTTON))
	rs := w.GetKeyState(int32(w.VK_RBUTTON))
	ms := w.GetKeyState(int32(w.VK_MBUTTON))
	keys[f.MouseButtons_Left] = ls < 0
	keys[f.MouseButtons_Right] = rs < 0
	keys[f.MouseButtons_Middle] = ms < 0
	return keys
}

func (_this *Provider) GetScreen() f.Screen {
	var s = f.Screen{
		Full: f.Rect{
			Width:  int(w.GetSystemMetrics(w.SM_CXSCREEN)),
			Height: int(w.GetSystemMetrics(w.SM_CYSCREEN)),
		},
		WorkArea: f.Rect{
			Width:  int(w.GetSystemMetrics(w.SM_CXFULLSCREEN)),
			Height: int(w.GetSystemMetrics(w.SM_CYFULLSCREEN)),
		},
	}
	return s
}

func (_this *Provider) SetIcon(file string) {
	h := w.LoadImage(_this.hInstance, sto16(file), w.IMAGE_ICON, 0, 0, w.LR_LOADFROMFILE)
	_this.defIcon = w.HICON(h)
}

func (_this *Provider) registerWndClass() {
	_this.wndClass = w.WNDCLASSEX{
		Style:         w.CS_HREDRAW | w.CS_VREDRAW | w.CS_DBLCLKS,
		LpfnWndProc:   syscall.NewCallbackCDecl(_this.classMsgProc),
		HInstance:     _this.hInstance,
		LpszClassName: sto16(_this.className),
		HIcon:         _this.defIcon,
		HIconSm:       _this.defIcon,
		HCursor:       w.LoadCursor(0, w.MAKEINTRESOURCE(w.IDC_ARROW)),
		HbrBackground: w.GetSysColorBrush(w.COLOR_WINDOW),
	}
	_this.wndClass.CbSize = uint32(unsafe.Sizeof(_this.wndClass))
	w.RegisterClassEx(&_this.wndClass)
	_this.dlgOwner = w.CreateWindowEx(0,
		sto16(_this.className), sto16(""),
		w.WS_OVERLAPPED, 0, 0, 0, 0,
		0, 0, _this.hInstance, unsafe.Pointer(nil))
}

func (_this *Provider) add(wnd baseWindow) {
	ref := reflect.ValueOf(wnd).Pointer()
	_this.tmpWnd[ref] = wnd
}

func (_this *Provider) remove(hWnd w.HWND, isExit bool) {
	if wnd, ok := _this.handleWnds[hWnd]; ok {
		delete(_this.handleWnds, hWnd)
		if isExit && wnd.hWnd() == _this.mainId {
			_this.Exit(0)
		}
	}
}

func (_this *Provider) classMsgProc(hWnd w.HWND, msg uint32, wParam uintptr, lParam uintptr) uintptr {
	if msg == w.WM_CREATE {
		cs := *((*w.CREATESTRUCT)(unsafe.Pointer(lParam)))
		if wnd, ok := _this.tmpWnd[cs.CreateParams]; ok {
			if _this.mainId == 0 {
				_this.mainId = hWnd
			}
			_this.handleWnds[hWnd] = wnd
			delete(_this.tmpWnd, cs.CreateParams)
		}
	}
	for _, proc := range _this.watchAll {
		if code := proc(hWnd, msg, wParam, lParam); code != 0 {
			return code
		}
	}
	if wnd, ok := _this.handleWnds[hWnd]; ok {
		if code := wnd.onWndMsg(hWnd, msg, wParam, lParam); code != 0 {
			return code
		}
	}
	return w.DefWindowProc(hWnd, msg, wParam, lParam)
}

func (_this *Provider) Exit(code int) {
	w.PostQuitMessage(int32(code))
}

func (_this *Provider) RunMain(form p.Form) {
	if fm, ok := form.(*winForm); ok {
		fm.Show()
	}
	var message w.MSG
	for {
		if w.GetMessage(&message, 0, 0, 0) {
			w.TranslateMessage(&message)
			w.DispatchMessage(&message)
		} else {
			break
		}
	}
	os.Exit(0)
}
