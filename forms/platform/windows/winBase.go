package windows

import (
	"golang.org/x/sys/windows"
	f "qq2564874169/goMiniblink/forms"
	plat "qq2564874169/goMiniblink/forms/platform"
	win "qq2564874169/goMiniblink/forms/platform/windows/win32"
	"strconv"
	"sync"
	"time"
	"unsafe"
)

type invokeContext struct {
	fn    func(state interface{})
	state interface{}
}

type winBase struct {
	app       *Provider
	handle    win.HWND
	invokeMap sync.Map
	onWndProc windowsMsgProc
	isLoad    bool
	owner     plat.Form
	parent    plat.Control

	onLoad                plat.WindowLoadProc
	onCreate              plat.WindowCreateProc
	onDestroy             plat.WindowDestroyProc
	onResize              plat.WindowResizeProc
	onMove                plat.WindowMoveProc
	onMouseMove           plat.WindowMouseMoveProc
	onMouseDown           plat.WindowMouseDownProc
	onMouseUp             plat.WindowMouseUpProc
	onMouseWheel          plat.WindowMouseWheelProc
	onMouseClick          plat.WindowMouseClickProc
	onPaint               plat.WindowPaintProc
	onKeyDown             plat.WindowKeyDownProc
	onKeyUp               plat.WindowKeyUpProc
	onKeyPress            plat.WindowKeyPressProc
	onSetCursor           plat.WindowSetCursorProc
	onImeStartComposition plat.WindowImeStartCompositionProc
	onFocus               plat.WindowFocusProc
	onLostFocus           plat.WindowLostFocusProc

	bgColor  int32
	msEnable bool
	cursor   f.CursorType
}

func (_this *winBase) init(provider *Provider) *winBase {
	_this.app = provider
	_this.app.add(_this)
	_this.onWndProc = _this.msgProc
	_this.msEnable = true
	_this.cursor = f.CursorType_Default
	_this.bgColor = -1
	return _this
}

func (_this *winBase) Hide() {
	win.ShowWindow(_this.handle, win.SW_HIDE)
}

func (_this *winBase) Show() {
	win.ShowWindow(_this.handle, win.SW_SHOW)
	win.UpdateWindow(_this.handle)
}

func (_this *winBase) IsInvoke() bool {
	return _this.app.mainThreadId == windows.GetCurrentThreadId()
}

func (_this *winBase) SetBgColor(color int32) {
	_this.bgColor = color
}

func (_this *winBase) SetCursor(cursor f.CursorType) {
	if cursor != f.CursorType_Default {
		res := win.MAKEINTRESOURCE(uintptr(toWinCursor(cursor)))
		win.SetCursor(win.LoadCursor(0, res))
	}
	_this.cursor = cursor
}

func (_this *winBase) GetCursor() f.CursorType {
	return _this.cursor
}

func (_this *winBase) onWndMsg(hWnd win.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	if _this.onWndProc != nil {
		return _this.onWndProc(hWnd, msg, wParam, lParam)
	}
	return 0
}

func (_this *winBase) msgProc(hWnd win.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	if _this.IsInvoke() == false {
		return 0
	}
	var ret uintptr
	switch msg {
	case win.WM_CREATE:
		cs := *((*win.CREATESTRUCT)(unsafe.Pointer(lParam)))
		if _this.app.defIcon != 0 && cs.ExStyle&win.WS_EX_DLGMODALFRAME == 0 {
			win.SendMessage(hWnd, win.WM_SETICON, 1, uintptr(_this.app.defIcon))
			win.SendMessage(hWnd, win.WM_SETICON, 0, uintptr(_this.app.defIcon))
		}
		_this.handle = hWnd
		if _this.onCreate != nil {
			_this.onCreate(uintptr(hWnd))
		}
	case win.WM_SHOWWINDOW:
		if _this.isLoad == false {
			if _this.onLoad != nil {
				_this.onLoad()
			}
			_this.isLoad = true
		}
	case win.WM_COMMAND:
		ret = _this.execCmd(wParam, lParam)
	case win.WM_KILLFOCUS:
		if _this.onLostFocus != nil && _this.onLostFocus() {
			ret = 1
		}
	case win.WM_SETFOCUS:
		if _this.onFocus != nil && _this.onFocus() {
			ret = 1
		}
	case win.WM_IME_STARTCOMPOSITION:
		if _this.onImeStartComposition != nil && _this.onImeStartComposition() {
			ret = 1
		}
	case win.WM_SETCURSOR:
		if _this.cursor != f.CursorType_Default {
			_this.SetCursor(_this.cursor)
			ret = 1
		} else if _this.onSetCursor != nil && _this.onSetCursor() {
			ret = 1
		}
	case win.WM_SIZE:
		if _this.onResize != nil {
			xl, yl := win.GET_X_LPARAM(lParam), win.GET_Y_LPARAM(lParam)
			rect := f.Rect{Width: int(xl), Height: int(yl)}
			if _this.onResize(rect) {
				ret = 1
			}
		}
	case win.WM_MOVE:
		if _this.onMove != nil {
			x, y := win.GET_X_LPARAM(lParam), win.GET_Y_LPARAM(lParam)
			pos := f.Point{X: int(x), Y: int(y)}
			if _this.onMove(pos) {
				ret = 1
			}
		}
	case win.WM_DESTROY:
		if _this.onDestroy != nil {
			_this.onDestroy()
		}
		_this.app.remove(_this.handle, true)
	case win.WM_SYSKEYDOWN, win.WM_KEYDOWN:
		key := vkToKey(int(wParam))
		if _this.onKeyDown != nil && key != f.Keys_Error {
			e := f.KeyEvArgs{
				Key:   key,
				Value: wParam,
				IsSys: msg == win.WM_SYSKEYDOWN,
			}
			if _this.onKeyDown(&e); e.IsHandle {
				ret = 1
			}
		}
	case win.WM_SYSKEYUP, win.WM_KEYUP:
		key := vkToKey(int(wParam))
		if _this.onKeyUp != nil && key != f.Keys_Error {
			e := f.KeyEvArgs{
				Key:   key,
				Value: wParam,
				IsSys: msg == win.WM_SYSKEYUP,
			}
			if _this.onKeyUp(&e); e.IsHandle {
				ret = 1
			}
		}
	case win.WM_SYSCHAR, win.WM_CHAR:
		if _this.onKeyPress != nil {
			e := f.KeyPressEvArgs{
				KeyChar: strconv.Itoa(int(wParam)),
				Value:   wParam,
				IsSys:   msg == win.WM_SYSCHAR,
			}
			if _this.onKeyPress(&e); e.IsHandle {
				ret = 1
			}
		}
	case win.WM_PAINT:
		ps := new(win.PAINTSTRUCT)
		hdc := win.BeginPaint(hWnd, ps)
		e := f.PaintEvArgs{
			Clip: f.Bound{
				Point: f.Point{
					X: int(ps.RcPaint.Left),
					Y: int(ps.RcPaint.Top),
				},
				Rect: f.Rect{
					Width:  int(ps.RcPaint.Right - ps.RcPaint.Left),
					Height: int(ps.RcPaint.Bottom - ps.RcPaint.Top),
				},
			},
		}
		if e.Clip.IsEmpty() == false {
			if _this.bgColor >= 0 {
				r, g, b := uint8(_this.bgColor), uint8(_this.bgColor>>8), uint8(_this.bgColor>>16)
				bg := (int32(r) << 16) | (int32(g) << 8) | int32(b)
				sb := win.CreateSolidBrush(win.COLORREF(bg))
				win.FillRect(hdc, &ps.RcPaint, sb)
				win.DeleteObject(win.HGDIOBJ(sb))
			}
			e.Graphics = new(winGraphics).init(hdc)
			if _this.onPaint != nil && _this.onPaint(e) {
				ret = 1
			}
			e.Graphics.Close()
		}
		win.EndPaint(hWnd, ps)
	case win.WM_MOUSEMOVE:
		if _this.onMouseMove != nil && _this.msEnable {
			e := f.MouseEvArgs{
				X:    int(win.GET_X_LPARAM(lParam)),
				Y:    int(win.GET_Y_LPARAM(lParam)),
				Time: time.Now(),
			}
			wp := int(wParam)
			if wp&win.MK_LBUTTON != 0 {
				e.Button |= f.MouseButtons_Left
			}
			if wp&win.MK_MBUTTON != 0 {
				e.Button |= f.MouseButtons_Middle
			}
			if wp&win.MK_RBUTTON != 0 {
				e.Button |= f.MouseButtons_Right
			}
			if _this.onMouseMove(&e); e.IsHandle {
				ret = 1
			}
		}
	case win.WM_LBUTTONDBLCLK, win.WM_MBUTTONDBLCLK, win.WM_RBUTTONDBLCLK:
		if _this.onMouseClick != nil && _this.msEnable {
			e := f.MouseEvArgs{
				X:        int(win.GET_X_LPARAM(lParam)),
				Y:        int(win.GET_Y_LPARAM(lParam)),
				Time:     time.Now(),
				IsDouble: true,
			}
			switch msg {
			case win.WM_LBUTTONDOWN:
				e.Button |= f.MouseButtons_Left
			case win.WM_RBUTTONDOWN:
				e.Button |= f.MouseButtons_Right
			case win.WM_MBUTTONDOWN:
				e.Button |= f.MouseButtons_Middle
			}
			if _this.onMouseClick(&e); e.IsHandle {
				ret = 1
			}
		}
	case win.WM_LBUTTONDOWN, win.WM_RBUTTONDOWN, win.WM_MBUTTONDOWN:
		win.SetCapture(hWnd)
		if _this.onMouseDown != nil && _this.msEnable {
			e := f.MouseEvArgs{
				X:    int(win.GET_X_LPARAM(lParam)),
				Y:    int(win.GET_Y_LPARAM(lParam)),
				Time: time.Now(),
			}
			switch msg {
			case win.WM_LBUTTONDOWN:
				e.Button |= f.MouseButtons_Left
			case win.WM_RBUTTONDOWN:
				e.Button |= f.MouseButtons_Right
			case win.WM_MBUTTONDOWN:
				e.Button |= f.MouseButtons_Middle
			}
			if _this.onMouseDown(&e); e.IsHandle {
				ret = 1
			}
		}
	case win.WM_LBUTTONUP, win.WM_RBUTTONUP, win.WM_MBUTTONUP:
		win.ReleaseCapture()
		if _this.onMouseUp != nil && _this.msEnable {
			e := f.MouseEvArgs{
				X:    int(win.GET_X_LPARAM(lParam)),
				Y:    int(win.GET_Y_LPARAM(lParam)),
				Time: time.Now(),
			}
			switch msg {
			case win.WM_LBUTTONDOWN:
				e.Button |= f.MouseButtons_Left
			case win.WM_RBUTTONDOWN:
				e.Button |= f.MouseButtons_Right
			case win.WM_MBUTTONDOWN:
				e.Button |= f.MouseButtons_Middle
			}
			if _this.onMouseUp(&e); e.IsHandle {
				ret = 1
			}
		}
	case win.WM_MOUSEWHEEL:
		if _this.onMouseWheel != nil && _this.msEnable {
			lp, hp := win.LOWORD(int32(wParam)), win.HIWORD(int32(wParam))
			e := f.MouseEvArgs{
				X:     int(win.GET_X_LPARAM(lParam)),
				Y:     int(win.GET_Y_LPARAM(lParam)),
				Delta: int(hp),
				Time:  time.Now(),
			}
			if lp&win.MK_LBUTTON != 0 {
				e.Button |= f.MouseButtons_Left
			}
			if lp&win.MK_MBUTTON != 0 {
				e.Button |= f.MouseButtons_Middle
			}
			if lp&win.MK_RBUTTON != 0 {
				e.Button |= f.MouseButtons_Right
			}
			if _this.onMouseWheel(&e); e.IsHandle {
				ret = 1
			}
		}
	default:
		return 0
	}
	return ret
}

func (_this *winBase) CreateGraphics() f.Graphics {
	hdc := win.GetDC(_this.handle)
	g := new(winGraphics).init(hdc)
	g.onClose = func(e *winGraphics) {
		win.ReleaseDC(_this.handle, e.dc)
	}
	return g
}

func (_this *winBase) Invoke(fn func(state interface{}), state interface{}) {
	ctx := invokeContext{
		fn:    fn,
		state: state,
	}
	ptr := uintptr(unsafe.Pointer(&ctx))
	_this.invokeMap.Store(ptr, &ctx)
	win.PostMessage(_this.handle, uint32(win.WM_COMMAND), uintptr(cmd_invoke), ptr)
}

func (_this *winBase) execCmd(wParam, lParam uintptr) uintptr {
	switch int(wParam) {
	case cmd_invoke:
		ctx := *((*invokeContext)(unsafe.Pointer(lParam)))
		ctx.fn(ctx.state)
		_this.invokeMap.Delete(lParam)
		return 1
	default:
		return 0
	}
}

func (_this *winBase) hWnd() win.HWND {
	return _this.handle
}

func (_this *winBase) GetHandle() uintptr {
	return uintptr(_this.hWnd())
}

func (_this *winBase) GetProvider() plat.Provider {
	return _this.app
}

func (_this *winBase) SetMouseEnable(enable bool) {
	_this.msEnable = enable
}

func (_this *winBase) GetMouseEnable() bool {
	return _this.msEnable
}

func (_this *winBase) GetSize() (width, height int) {
	rect := win.RECT{}
	win.GetWindowRect(_this.handle, &rect)
	return int(rect.Right - rect.Left), int(rect.Bottom - rect.Top)
}

func (_this *winBase) SetSize(width, height int) {
	win.SetWindowPos(_this.handle, 0, 0, 0, int32(width), int32(height), win.SWP_NOMOVE|win.SWP_NOZORDER)
}

func (_this *winBase) SetLocation(x, y int) {
	win.SetWindowPos(_this.handle, 0, int32(x), int32(y), 0, 0, win.SWP_NOSIZE|win.SWP_NOZORDER)
}

func (_this *winBase) GetLocation() (x, y int) {
	rect := win.RECT{}
	win.GetWindowRect(_this.handle, &rect)
	return int(rect.Left), int(rect.Top)
}

func (_this *winBase) SetParent(parent plat.Control) {
	if win.SetParent(_this.handle, win.HWND(parent.GetHandle())) != 0 {
		_this.parent = parent
		if ow, ok := parent.(plat.Form); ok {
			_this.owner = ow
		} else {
			_this.owner = parent.GetOwner()
		}
	}
}

func (_this *winBase) GetParent() plat.Control {
	return _this.parent
}

func (_this *winBase) GetOwner() plat.Form {
	return _this.owner
}

func (_this *winBase) MousePosition() f.Point {
	p := _this.app.MouseLocation()
	sp := win.POINT{
		X: int32(p.X),
		Y: int32(p.Y),
	}
	win.ScreenToClient(_this.handle, &sp)
	return f.Point{
		X: int(sp.X),
		Y: int(sp.Y),
	}
}
