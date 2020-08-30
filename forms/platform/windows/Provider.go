package windows

import (
	"golang.org/x/sys/windows"
	"os"
	"path/filepath"
	"qq2564874169/goMiniblink/forms"
	"qq2564874169/goMiniblink/forms/platform"
	"qq2564874169/goMiniblink/forms/platform/windows/win32"
	"syscall"
	"time"
	"unsafe"
)

type windowsMsgProc func(hWnd win32.HWND, msg uint32, wParam, lParam uintptr) uintptr

type Provider struct {
	hInstance    win32.HINSTANCE
	className    string
	mainId       win32.HWND
	handleWnds   map[win32.HWND]baseWindow
	nameWnds     map[string]baseWindow
	defOwner     win32.HWND
	defIcon      win32.HICON
	msClick      *mouseClickWorker
	defBgColor   int
	mainThreadId uint32
}

func (_this *Provider) Init() *Provider {
	_this.handleWnds = make(map[win32.HWND]baseWindow)
	_this.nameWnds = make(map[string]baseWindow)
	_this.className = "goMiniblinkClass"
	_this.hInstance = win32.GetModuleHandle(nil)
	_this.msClick = new(mouseClickWorker).init()
	_this.mainThreadId = windows.GetCurrentThreadId()
	return _this
}

func (_this *Provider) AppDir() string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return dir
}

func (_this *Provider) ModifierKeys() map[forms.Keys]bool {
	keys := make(map[forms.Keys]bool)
	cs := win32.GetKeyState(int32(win32.VK_CONTROL))
	ss := win32.GetKeyState(int32(win32.VK_SHIFT))
	as := win32.GetKeyState(int32(win32.VK_MENU))
	keys[forms.Keys_Ctrl] = cs < 0
	keys[forms.Keys_Shift] = ss < 0
	keys[forms.Keys_Alt] = as < 0
	return keys
}

func (_this *Provider) MouseIsDown() map[forms.MouseButtons]bool {
	keys := make(map[forms.MouseButtons]bool)
	ls := win32.GetKeyState(int32(win32.VK_LBUTTON))
	rs := win32.GetKeyState(int32(win32.VK_RBUTTON))
	ms := win32.GetKeyState(int32(win32.VK_MBUTTON))
	keys[forms.MouseButtons_Left] = ls < 0
	keys[forms.MouseButtons_Right] = rs < 0
	keys[forms.MouseButtons_Middle] = ms < 0
	return keys
}

func (_this *Provider) SetBgColor(color int) {
	_this.defBgColor = color
}

func (_this *Provider) GetScreen() forms.Screen {
	var s = forms.Screen{
		Full: forms.Rect{
			Width:  int(win32.GetSystemMetrics(win32.SM_CXSCREEN)),
			Height: int(win32.GetSystemMetrics(win32.SM_CYSCREEN)),
		},
		WorkArea: forms.Rect{
			Width:  int(win32.GetSystemMetrics(win32.SM_CXFULLSCREEN)),
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
		LpfnWndProc:   syscall.NewCallback(_this.classMsgProc),
		HInstance:     _this.hInstance,
		LpszClassName: sto16(_this.className),
		HIcon:         _this.defIcon,
		HIconSm:       _this.defIcon,
	}
	class.CbSize = uint32(unsafe.Sizeof(class))
	win32.RegisterClassEx(&class)
	_this.defOwner = win32.CreateWindowEx(0,
		sto16(_this.className), sto16(""),
		win32.WS_OVERLAPPED, 0, 0, 0, 0,
		0, 0, _this.hInstance, unsafe.Pointer(nil))
}

func (_this *Provider) add(wnd baseWindow) {
	if hWnd := wnd.hWnd(); hWnd != 0 {
		if _this.mainId == 0 {
			_this.mainId = hWnd
		}
		_this.handleWnds[hWnd] = wnd
	}
}

func (_this *Provider) remove(hWnd win32.HWND, isExit bool) {
	if w, ok := _this.handleWnds[hWnd]; ok {
		delete(_this.handleWnds, hWnd)
		if isExit && w.hWnd() == _this.mainId {
			_this.Exit(0)
		}
	}
}

func (_this *Provider) classMsgProc(hWnd win32.HWND, msg uint32, wParam uintptr, lParam uintptr) uintptr {
	if w, ok := _this.handleWnds[hWnd]; ok {
		if code := w.wndMsgProc(hWnd, msg, wParam, lParam); code != 0 {
			return code
		}
	}
	return win32.DefWindowProc(hWnd, msg, wParam, lParam)
}

func (_this *Provider) sendMouseClick(hWnd win32.HWND, msg uint32, lParam uintptr) uintptr {
	switch msg {
	case win32.WM_LBUTTONDOWN:
		_this.msClick.mouseDown(hWnd, forms.MouseButtons_Left, int(lParam))
	case win32.WM_LBUTTONUP:
		if _this.msClick.mouseUp(hWnd, forms.MouseButtons_Left, int(lParam)) {
			return 1
		}
	case win32.WM_RBUTTONDOWN:
		_this.msClick.mouseDown(hWnd, forms.MouseButtons_Right, int(lParam))
	case win32.WM_RBUTTONUP:
		if _this.msClick.mouseUp(hWnd, forms.MouseButtons_Right, int(lParam)) {
			return 1
		}
	case win32.WM_MBUTTONDOWN:
		_this.msClick.mouseDown(hWnd, forms.MouseButtons_Middle, int(lParam))
	case win32.WM_MBUTTONUP:
		if _this.msClick.mouseUp(hWnd, forms.MouseButtons_Middle, int(lParam)) {
			return 1
		}
	}
	return 0
}

func (_this *Provider) Exit(code int) {
	win32.PostQuitMessage(int32(code))
}

func (_this *Provider) RunMain(form platform.Form) {
	_, ok := form.(*winForm)
	if ok == false {
		panic("类型不正确")
	}
	_this.registerWndClass()
	form.Create()
	form.Show()
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

type mouseClickWorker struct {
	down_hWnd  win32.HWND
	down_key   forms.MouseButtons
	down_pos   int
	click_hWnd win32.HWND
	click_key  forms.MouseButtons
	click_pos  int
	time       int64
	isDouble   bool
}

func (_this *mouseClickWorker) init() *mouseClickWorker {
	return _this
}

func (_this *mouseClickWorker) fire() {
	defer func() {
		recover()
	}()
	for {
		time.Sleep(time.Millisecond)
		if _this.click_hWnd != 0 && _this.time <= time.Now().UnixNano() {
			x, y := int(win32.LOWORD(int32(_this.click_pos))), int(win32.HIWORD(int32(_this.click_pos)))
			e := forms.MouseEvArgs{
				X:        x,
				Y:        y,
				Delta:    0,
				IsDouble: _this.isDouble,
				Time:     time.Now(),
				Button:   _this.click_key,
			}
			win32.PostMessage(_this.click_hWnd, uint32(win32.WM_COMMAND), uintptr(cmd_mouse_click), uintptr(unsafe.Pointer(&e)))
			_this.click_hWnd = 0
			_this.click_key = 0
			_this.click_pos = 0
		}
	}
}

func (_this *mouseClickWorker) mouseDown(hWnd win32.HWND, key forms.MouseButtons, pos int) {
	_this.down_hWnd = hWnd
	_this.down_key = key
	_this.down_pos = pos
}

func (_this *mouseClickWorker) mouseUp(hWnd win32.HWND, key forms.MouseButtons, pos int) bool {
	if _this.down_hWnd == hWnd {
		if _this.down_key == key && _this.down_pos == pos {
			_this.click_key = key
			_this.click_pos = pos
			if _this.click_hWnd != hWnd {
				_this.click_hWnd = hWnd
				_this.isDouble = false
				_this.time = time.Now().Add(time.Millisecond * 200).UnixNano()
				return true
			} else if _this.time > time.Now().UnixNano() {
				_this.isDouble = true
				return true
			}
		}
	}
	_this.click_hWnd = 0
	return false
}
