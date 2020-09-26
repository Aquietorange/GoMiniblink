package windows

import (
	"golang.org/x/sys/windows"
	fm "qq2564874169/goMiniblink/forms"
	br "qq2564874169/goMiniblink/forms/bridge"
	win "qq2564874169/goMiniblink/forms/windows/win32"
	"strconv"
	"sync"
	"time"
	"unsafe"
)

var dbclickTime = win.GetDoubleClickTime()

type invokeContext struct {
	fn    func(state interface{})
	state interface{}
}

const (
	cmd_invoke = 100
)

type winBase struct {
	app       *Provider
	handle    win.HWND
	invokeMap sync.Map
	onWndProc windowsMsgProc
	owner     br.Form
	parent    br.Control

	onShow                br.WindowShowProc
	onCreate              br.WindowCreateProc
	onDestroy             br.WindowDestroyProc
	onResize              br.WindowResizeProc
	onMove                br.WindowMoveProc
	onMouseMove           br.WindowMouseMoveProc
	onMouseDown           br.WindowMouseDownProc
	onMouseUp             br.WindowMouseUpProc
	onMouseWheel          br.WindowMouseWheelProc
	onMouseClick          br.WindowMouseClickProc
	onPaint               br.WindowPaintProc
	onKeyDown             br.WindowKeyDownProc
	onKeyUp               br.WindowKeyUpProc
	onKeyPress            br.WindowKeyPressProc
	onSetCursor           br.WindowSetCursorProc
	onImeStartComposition br.WindowImeStartCompositionProc
	onFocus               br.WindowFocusProc
	onLostFocus           br.WindowLostFocusProc

	bgColor   int32
	cursor    fm.CursorType
	isEnable  bool
	clickTime time.Time
}

func (_this *winBase) init(provider *Provider) *winBase {
	_this.app = provider
	_this.isEnable = true
	_this.cursor = fm.CursorType_Default
	_this.bgColor = -1
	_this.app.add(_this)
	_this.onWndProc = _this.msgProc
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
	var rect win.RECT
	win.GetClientRect(_this.handle, &rect)
	win.InvalidateRect(_this.handle, &rect, false)
}

func (_this *winBase) SetCursor(cursor fm.CursorType) {
	if cursor != fm.CursorType_Default {
		res := win.MAKEINTRESOURCE(uintptr(toWinCursor(cursor)))
		win.SetCursor(win.LoadCursor(0, res))
	}
	_this.cursor = cursor
}

func (_this *winBase) GetCursor() fm.CursorType {
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
		if _this.app.defIcon != 0 && cs.ExStyle&win.WS_EX_DLGMODALFRAME == 0 && cs.Style&win.WS_CHILD == 0 {
			win.SendMessage(hWnd, win.WM_SETICON, 1, uintptr(_this.app.defIcon))
			win.SendMessage(hWnd, win.WM_SETICON, 0, uintptr(_this.app.defIcon))
		}
		_this.handle = hWnd
		if _this.onCreate != nil {
			_this.onCreate(uintptr(hWnd))
		}
	case win.WM_SHOWWINDOW:
		if _this.onShow != nil {
			_this.onShow()
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
		if _this.cursor != fm.CursorType_Default {
			_this.SetCursor(_this.cursor)
			ret = 1
		} else if _this.onSetCursor != nil && _this.onSetCursor() {
			ret = 1
		}
	case win.WM_SIZE:
		if _this.onResize != nil {
			w, h := win.GET_X_LPARAM(lParam), win.GET_Y_LPARAM(lParam)
			rect := fm.Rect{Width: int(w), Height: int(h)}
			_this.onResize(rect)
		}
	case win.WM_MOVE:
		if _this.onMove != nil {
			x, y := win.GET_X_LPARAM(lParam), win.GET_Y_LPARAM(lParam)
			pos := fm.Point{X: int(x), Y: int(y)}
			if _this.onMove(pos) {
				ret = 1
			}
		}
	case win.WM_DESTROY:
		if _this.onDestroy != nil {
			_this.onDestroy()
		}
	case win.WM_SYSKEYDOWN, win.WM_KEYDOWN:
		key := vkToKey(int(wParam))
		if _this.onKeyDown != nil && key != fm.Keys_Error {
			e := fm.KeyEvArgs{
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
		if _this.onKeyUp != nil && key != fm.Keys_Error {
			e := fm.KeyEvArgs{
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
			e := fm.KeyPressEvArgs{
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
		e := fm.PaintEvArgs{
			Clip: fm.Bound{
				Point: fm.Point{
					X: int(ps.RcPaint.Left),
					Y: int(ps.RcPaint.Top),
				},
				Rect: fm.Rect{
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
		if _this.onMouseMove != nil {
			e := fm.MouseEvArgs{
				X:    int(win.GET_X_LPARAM(lParam)),
				Y:    int(win.GET_Y_LPARAM(lParam)),
				Time: time.Now(),
			}
			wp := int(wParam)
			if wp&win.MK_LBUTTON != 0 {
				e.Button |= fm.MouseButtons_Left
			}
			if wp&win.MK_MBUTTON != 0 {
				e.Button |= fm.MouseButtons_Middle
			}
			if wp&win.MK_RBUTTON != 0 {
				e.Button |= fm.MouseButtons_Right
			}
			if _this.onMouseMove(&e); e.IsHandle {
				ret = 1
			}
		}
	case win.WM_LBUTTONDBLCLK, win.WM_MBUTTONDBLCLK, win.WM_RBUTTONDBLCLK:
		_this.clickTime = time.Now().Add(time.Hour * -1)
		if _this.onMouseClick != nil {
			e := fm.MouseEvArgs{
				X:        int(win.GET_X_LPARAM(lParam)),
				Y:        int(win.GET_Y_LPARAM(lParam)),
				Time:     time.Now(),
				IsDouble: true,
			}
			switch msg {
			case win.WM_LBUTTONDOWN:
				e.Button |= fm.MouseButtons_Left
			case win.WM_RBUTTONDOWN:
				e.Button |= fm.MouseButtons_Right
			case win.WM_MBUTTONDOWN:
				e.Button |= fm.MouseButtons_Middle
			}
			if _this.onMouseClick(&e); e.IsHandle {
				ret = 1
			}
		}
	case win.WM_LBUTTONDOWN, win.WM_RBUTTONDOWN, win.WM_MBUTTONDOWN:
		win.SetCapture(hWnd)
		_this.clickTime = time.Now()
		if _this.onMouseDown != nil {
			e := fm.MouseEvArgs{
				X:    int(win.GET_X_LPARAM(lParam)),
				Y:    int(win.GET_Y_LPARAM(lParam)),
				Time: time.Now(),
			}
			switch msg {
			case win.WM_LBUTTONDOWN:
				e.Button |= fm.MouseButtons_Left
			case win.WM_RBUTTONDOWN:
				e.Button |= fm.MouseButtons_Right
			case win.WM_MBUTTONDOWN:
				e.Button |= fm.MouseButtons_Middle
			}
			if _this.onMouseDown(&e); e.IsHandle {
				ret = 1
			}
		}
	case win.WM_LBUTTONUP, win.WM_RBUTTONUP, win.WM_MBUTTONUP:
		win.ReleaseCapture()
		if _this.onMouseUp != nil {
			e := fm.MouseEvArgs{
				X:    int(win.GET_X_LPARAM(lParam)),
				Y:    int(win.GET_Y_LPARAM(lParam)),
				Time: time.Now(),
			}
			switch msg {
			case win.WM_LBUTTONDOWN:
				e.Button |= fm.MouseButtons_Left
			case win.WM_RBUTTONDOWN:
				e.Button |= fm.MouseButtons_Right
			case win.WM_MBUTTONDOWN:
				e.Button |= fm.MouseButtons_Middle
			}
			if _this.onMouseUp(&e); e.IsHandle {
				ret = 1
			}
		}
		if _this.onMouseClick != nil {
			var rect win.RECT
			win.GetWindowRect(hWnd, &rect)
			mp := _this.app.MouseLocation()
			if mp.X < int(rect.Right) && mp.X >= int(rect.Left) && mp.Y >= int(rect.Top) && mp.Y < int(rect.Bottom) {
				go func() {
					time.Sleep(time.Duration(time.Millisecond.Nanoseconds() * int64(dbclickTime)))
					s := time.Now().Sub(_this.clickTime).Seconds()
					if s < 1 {
						e := fm.MouseEvArgs{
							X:        int(win.GET_X_LPARAM(lParam)),
							Y:        int(win.GET_Y_LPARAM(lParam)),
							Time:     time.Now(),
							IsDouble: false,
						}
						switch msg {
						case win.WM_LBUTTONDOWN:
							e.Button |= fm.MouseButtons_Left
						case win.WM_RBUTTONDOWN:
							e.Button |= fm.MouseButtons_Right
						case win.WM_MBUTTONDOWN:
							e.Button |= fm.MouseButtons_Middle
						}
						_this.Invoke(func(state interface{}) {
							st := state.(fm.MouseEvArgs)
							_this.onMouseClick(&st)
						}, e)
					}
				}()
			}
		}
	case win.WM_MOUSEWHEEL:
		if _this.onMouseWheel != nil {
			lp, hp := win.LOWORD(int32(wParam)), win.HIWORD(int32(wParam))
			e := fm.MouseEvArgs{
				X:     int(win.GET_X_LPARAM(lParam)),
				Y:     int(win.GET_Y_LPARAM(lParam)),
				Delta: int(hp),
				Time:  time.Now(),
			}
			if lp&win.MK_LBUTTON != 0 {
				e.Button |= fm.MouseButtons_Left
			}
			if lp&win.MK_MBUTTON != 0 {
				e.Button |= fm.MouseButtons_Middle
			}
			if lp&win.MK_RBUTTON != 0 {
				e.Button |= fm.MouseButtons_Right
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

func (_this *winBase) CreateGraphics() fm.Graphics {
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

func (_this *winBase) GetProvider() br.Provider {
	return _this.app
}

func (_this *winBase) IsEnable() bool {
	return _this.isEnable
}

func (_this *winBase) Enable(b bool) {
	style := win.GetWindowLong(_this.handle, win.GWL_STYLE)
	if b {
		style &= ^win.WS_DISABLED
	} else {
		style |= win.WS_DISABLED
	}
	win.SetWindowLong(_this.handle, win.GWL_STYLE, style)
}

func (_this *winBase) SetSize(width, height int) {
	wndRect := win.RECT{}
	win.GetWindowRect(_this.handle, &wndRect)
	ww := wndRect.Right - wndRect.Left
	wh := wndRect.Bottom - wndRect.Top
	clientRect := win.RECT{}
	win.GetClientRect(_this.handle, &clientRect)
	cw := clientRect.Right - clientRect.Left
	ch := clientRect.Bottom - clientRect.Top
	if cw < ww {
		width += int(ww - cw)
	}
	if ch < wh {
		height += int(wh - ch)
	}
	win.SetWindowPos(_this.handle, 0, 0, 0, int32(width), int32(height), win.SWP_NOMOVE|win.SWP_NOZORDER|win.SWP_NOACTIVATE)
}

func (_this *winBase) SetLocation(x, y int) {
	win.SetWindowPos(_this.handle, 0, int32(x), int32(y), 0, 0, win.SWP_NOSIZE|win.SWP_NOZORDER|win.SWP_NOACTIVATE)
}

func (_this *winBase) GetBound() fm.Bound {
	rect := win.RECT{}
	win.GetWindowRect(_this.handle, &rect)
	bn := fm.Bound{
		Point: fm.Point{
			X: int(rect.Left),
			Y: int(rect.Top),
		},
	}
	win.GetClientRect(_this.handle, &rect)
	bn.Rect = fm.Rect{
		Width:  int(rect.Right - rect.Left),
		Height: int(rect.Bottom - rect.Top),
	}
	if _this.GetParent() != nil {
		p := win.POINT{
			X: int32(bn.X),
			Y: int32(bn.Y),
		}
		win.ScreenToClient(win.HWND(_this.GetParent().GetHandle()), &p)
		bn.X, bn.Y = int(p.X), int(p.Y)
	}
	return bn
}

func (_this *winBase) GetParent() br.Control {
	return _this.parent
}

func (_this *winBase) GetOwner() br.Form {
	return _this.owner
}

func (_this *winBase) ToClientPoint(p fm.Point) fm.Point {
	sp := win.POINT{
		X: int32(p.X),
		Y: int32(p.Y),
	}
	win.ScreenToClient(_this.handle, &sp)
	return fm.Point{
		X: int(sp.X),
		Y: int(sp.Y),
	}
}
