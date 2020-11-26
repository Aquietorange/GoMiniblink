package windows

import (
	fm "GoMiniblink/forms"
	br "GoMiniblink/forms/bridge"
	win "GoMiniblink/forms/windows/win32"
	"golang.org/x/sys/windows"
	"os"
	"path/filepath"
	"reflect"
	"syscall"
	"unsafe"
)

type windowsMsgProc func(hWnd win.HWND, msg uint32, wParam, lParam uintptr) uintptr

type Provider struct {
	hInstance    win.HINSTANCE
	className    string
	wndClass     win.WNDCLASSEX
	mainId       win.HWND
	tmpWnd       map[uintptr]baseWindow
	handleWnds   map[win.HWND]baseWindow
	defOwner     win.HWND
	defIcon      win.HICON
	mainThreadId uint32
	watchAll     map[win.HWND][]windowsMsgProc
	forms        map[win.HWND]br.Form
}

func (_this *Provider) Init() *Provider {
	_this.forms = make(map[win.HWND]br.Form)
	_this.watchAll = make(map[win.HWND][]windowsMsgProc)
	_this.tmpWnd = make(map[uintptr]baseWindow)
	_this.handleWnds = make(map[win.HWND]baseWindow)
	_this.className = "goMiniblinkClass"
	_this.hInstance = win.GetModuleHandle(nil)
	_this.mainThreadId = windows.GetCurrentThreadId()
	_this.registerWndClass()
	return _this
}

func (_this *Provider) watch(wnd baseWindow, proc windowsMsgProc) {
	_this.watchAll[wnd.hWnd()] = append(_this.watchAll[wnd.hWnd()], proc)
}

func (_this *Provider) MouseLocation() fm.Point {
	pos := win.POINT{}
	win.GetCursorPos(&pos)
	return fm.Point{
		X: int(pos.X),
		Y: int(pos.Y),
	}
}

func (_this *Provider) AppDir() string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return dir
}

func (_this *Provider) ModifierKeys() map[fm.Keys]bool {
	keys := make(map[fm.Keys]bool)
	cs := win.GetKeyState(int32(win.VK_CONTROL))
	ss := win.GetKeyState(int32(win.VK_SHIFT))
	as := win.GetKeyState(int32(win.VK_MENU))
	keys[fm.Keys_Ctrl] = cs < 0
	keys[fm.Keys_Shift] = ss < 0
	keys[fm.Keys_Alt] = as < 0
	return keys
}

func (_this *Provider) MouseIsDown() map[fm.MouseButtons]bool {
	keys := make(map[fm.MouseButtons]bool)
	ls := win.GetKeyState(int32(win.VK_LBUTTON))
	rs := win.GetKeyState(int32(win.VK_RBUTTON))
	ms := win.GetKeyState(int32(win.VK_MBUTTON))
	keys[fm.MouseButtons_Left] = ls < 0
	keys[fm.MouseButtons_Right] = rs < 0
	keys[fm.MouseButtons_Middle] = ms < 0
	return keys
}

func (_this *Provider) GetScreen() fm.Screen {
	var s = fm.Screen{
		Full: fm.Rect{
			Width:  int(win.GetSystemMetrics(win.SM_CXSCREEN)),
			Height: int(win.GetSystemMetrics(win.SM_CYSCREEN)),
		},
		WorkArea: fm.Rect{
			Width:  int(win.GetSystemMetrics(win.SM_CXFULLSCREEN)),
			Height: int(win.GetSystemMetrics(win.SM_CYFULLSCREEN)),
		},
	}
	return s
}

func (_this *Provider) SetIcon(file string) {
	h := win.LoadImage(_this.hInstance, sto16(file), win.IMAGE_ICON, 0, 0, win.LR_LOADFROMFILE)
	_this.defIcon = win.HICON(h)
}

func (_this *Provider) registerWndClass() {
	_this.wndClass = win.WNDCLASSEX{
		Style:         win.CS_HREDRAW | win.CS_VREDRAW | win.CS_DBLCLKS,
		LpfnWndProc:   syscall.NewCallbackCDecl(_this.classMsgProc),
		HInstance:     _this.hInstance,
		LpszClassName: sto16(_this.className),
		HCursor:       win.LoadCursor(0, win.MAKEINTRESOURCE(win.IDC_ARROW)),
		HbrBackground: win.GetSysColorBrush(win.COLOR_WINDOW),
	}
	_this.wndClass.CbSize = uint32(unsafe.Sizeof(_this.wndClass))
	win.RegisterClassEx(&_this.wndClass)
	_this.defOwner = win.CreateWindowEx(0,
		sto16(_this.className), sto16(""),
		win.WS_OVERLAPPED, 0, 0, 0, 0,
		0, 0, _this.hInstance, unsafe.Pointer(nil))
}

func (_this *Provider) add(wnd baseWindow) {
	ref := reflect.ValueOf(wnd).Pointer()
	_this.tmpWnd[ref] = wnd
}

func (_this *Provider) classMsgProc(hWnd win.HWND, msg uint32, wParam uintptr, lParam uintptr) uintptr {
	switch msg {
	case win.WM_CREATE:
		cs := *((*win.CREATESTRUCT)(unsafe.Pointer(lParam)))
		if wnd, es := _this.tmpWnd[cs.CreateParams]; es {
			delete(_this.tmpWnd, cs.CreateParams)
			_this.handleWnds[hWnd] = wnd
			if _this.mainId == 0 {
				_this.mainId = hWnd
			}
		}
	}
	for _, list := range _this.watchAll {
		for _, proc := range list {
			if rs := proc(hWnd, msg, wParam, lParam); rs != 0 {
				return rs
			}
		}
	}
	if wnd, ok := _this.handleWnds[hWnd]; ok {
		if rs := wnd.onWndMsg(hWnd, msg, wParam, lParam); rs != 0 {
			return rs
		}
	}
	switch msg {
	case win.WM_DESTROY:
		delete(_this.handleWnds, hWnd)
		delete(_this.forms, hWnd)
		delete(_this.watchAll, hWnd)
		if hWnd == _this.mainId {
			_this.Exit(0)
		}
	}
	return win.DefWindowProc(hWnd, msg, wParam, lParam)
}

func (_this *Provider) Exit(code int) {
	win.PostQuitMessage(int32(code))
}

func (_this *Provider) RunMain(form br.Form) {
	form.Show()
	var message win.MSG
	for {
		if win.GetMessage(&message, 0, 0, 0) {
			if !win.IsDialogMessage(message.HWnd, &message) {
				win.TranslateMessage(&message)
				win.DispatchMessage(&message)
			}
		} else {
			break
		}
	}
	os.Exit(0)
}
