package windows

import (
	"image"
	"image/draw"
	f "qq2564874169/goMiniblink/forms"
	p "qq2564874169/goMiniblink/forms/platform"
	w "qq2564874169/goMiniblink/forms/platform/windows/win32"
	"sync"
	"time"
	"unsafe"
)

type winBase struct {
	app       *Provider
	handle    w.HWND
	isCreated bool
	invokeMap sync.Map
	onWndProc windowsMsgProc
	isLoad    bool

	onLoad                p.WindowLoadProc
	onCreate              p.WindowCreateProc
	onDestroy             p.WindowDestroyProc
	onResize              p.WindowResizeProc
	onMove                p.WindowMoveProc
	onMouseMove           p.WindowMouseMoveProc
	onMouseDown           p.WindowMouseDownProc
	onMouseUp             p.WindowMouseUpProc
	onMouseWheel          p.WindowMouseWheelProc
	onMouseClick          p.WindowMouseClickProc
	onPaint               p.WindowPaintProc
	onKeyDown             p.WindowKeyDownProc
	onKeyUp               p.WindowKeyUpProc
	onKeyPress            p.WindowKeyPressProc
	onSetCursor           p.WindowSetCursorProc
	onImeStartComposition p.WindowImeStartCompositionProc
	onFocus               p.WindowFocusProc
	onLostFocus           p.WindowLostFocusProc

	bgColor     int
	fiexdCursor bool
}

func (_this *winBase) init(provider *Provider) *winBase {
	_this.app = provider
	_this.onWndProc = _this.msgProc
	_this.SetBgColor(provider.defBgColor)
	return _this
}

func (_this *winBase) SetBgColor(color int) {
	_this.bgColor = color
}

func (_this *winBase) IsCreate() bool {
	return _this.isCreated
}

func (_this *winBase) SetCursor(cursor f.CursorType) {
	res := w.MAKEINTRESOURCE(uintptr(toWinCursor(cursor)))
	cur := w.LoadCursor(0, res)
	w.SetCursor(cur)
}

func (_this *winBase) wndMsgProc(hWnd w.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	if _this.onWndProc != nil {
		return _this.onWndProc(hWnd, msg, wParam, lParam)
	}
	return 0
}

func (_this *winBase) fireCreate(hWnd w.HWND) {
	_this.isCreated = true
	_this.handle = hWnd
	_this.app.add(_this)
	if _this.onCreate != nil {
		_this.onCreate(uintptr(hWnd))
	}
}

func (_this *winBase) msgProc(hWnd w.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	var ret int
	switch msg {
	case w.WM_SHOWWINDOW:
		if _this.isLoad == false {
			if _this.onLoad != nil {
				_this.onLoad()
			}
			_this.isLoad = true
		}
	case w.WM_COMMAND:
		ret = _this.execCmd(wParam, lParam)
	case w.WM_KILLFOCUS:
		if _this.onLostFocus != nil && _this.onLostFocus() {
			ret = 1
		}
	case w.WM_SETFOCUS:
		if _this.onFocus != nil && _this.onFocus() {
			ret = 1
		}
	case w.WM_IME_STARTCOMPOSITION:
		if _this.onImeStartComposition != nil && _this.onImeStartComposition() {
			ret = 1
		}
	case w.WM_SETCURSOR:
		if _this.fiexdCursor || (_this.onSetCursor != nil && _this.onSetCursor()) {
			ret = 1
		}
	case w.WM_SIZE:
		if _this.onResize != nil {
			w, h := w.GET_X_LPARAM(lParam), w.GET_Y_LPARAM(lParam)
			rect := f.Rect{Width: int(w), Height: int(h)}
			if _this.onResize(rect) {
				ret = 1
			}
		}
	case w.WM_MOVE:
		if _this.onMove != nil {
			x, y := w.GET_X_LPARAM(lParam), w.GET_Y_LPARAM(lParam)
			pos := f.Point{X: int(x), Y: int(y)}
			if _this.onMove(pos) {
				ret = 1
			}
		}
	case w.WM_DESTROY:
		if _this.onDestroy != nil {
			_this.onDestroy()
		}
		_this.app.remove(_this.hWnd(), true)
	case w.WM_SYSKEYDOWN, w.WM_KEYDOWN:
		key := vkToKey(int(wParam))
		if _this.onKeyDown != nil && key != f.Keys_Error {
			e := f.KeyEvArgs{
				Key:   key,
				Value: wParam,
				IsSys: msg == w.WM_SYSKEYDOWN,
			}
			if _this.onKeyDown(&e); e.IsHandle {
				ret = 1
			}
		}
	case w.WM_SYSKEYUP, w.WM_KEYUP:
		key := vkToKey(int(wParam))
		if _this.onKeyUp != nil && key != f.Keys_Error {
			e := f.KeyEvArgs{
				Key:   key,
				Value: wParam,
				IsSys: msg == w.WM_SYSKEYUP,
			}
			if _this.onKeyUp(&e); e.IsHandle {
				ret = 1
			}
		}
	case w.WM_SYSCHAR, w.WM_CHAR:
		if _this.onKeyPress != nil {
			e := f.KeyPressEvArgs{
				KeyChar: string(wParam),
				Value:   wParam,
				IsSys:   msg == w.WM_SYSCHAR,
			}
			if _this.onKeyPress(&e); e.IsHandle {
				ret = 1
			}
		}
	case w.WM_ERASEBKGND:
		return 1
	case w.WM_PAINT:
		ps := new(w.PAINTSTRUCT)
		hdc := w.BeginPaint(hWnd, ps)
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
			e.Graphics = new(winGraphics).init(hdc)
			if _this.bgColor >= 0 {
				bg := image.NewRGBA(image.Rect(0, 0, e.Clip.Width, e.Clip.Height))
				draw.Draw(bg, bg.Rect, image.NewUniform(intToRGBA(_this.bgColor)), image.Pt(0, 0), draw.Src)
				e.Graphics.DrawImage(bg, 0, 0, e.Clip.Width, e.Clip.Height, e.Clip.X, e.Clip.Y)
			}
			if _this.onPaint != nil && _this.onPaint(e) {
				ret = 1
			}
			e.Graphics.Close()
		}
		w.EndPaint(hWnd, ps)
	case w.WM_MOUSEMOVE:
		if _this.onMouseMove != nil {
			e := f.MouseEvArgs{
				X:    int(w.GET_X_LPARAM(lParam)),
				Y:    int(w.GET_Y_LPARAM(lParam)),
				Time: time.Now(),
			}
			wp := int(wParam)
			if wp&w.MK_LBUTTON != 0 {
				e.Button |= f.MouseButtons_Left
			}
			if wp&w.MK_MBUTTON != 0 {
				e.Button |= f.MouseButtons_Middle
			}
			if wp&w.MK_RBUTTON != 0 {
				e.Button |= f.MouseButtons_Right
			}
			if _this.onMouseMove(&e); e.IsHandle {
				ret = 1
			}
		}
	case w.WM_LBUTTONDOWN, w.WM_RBUTTONDOWN, w.WM_MBUTTONDOWN:
		_this.fiexdCursor = true
		w.SetCapture(hWnd)
		if _this.onMouseDown != nil {
			e := f.MouseEvArgs{
				X:    int(w.GET_X_LPARAM(lParam)),
				Y:    int(w.GET_Y_LPARAM(lParam)),
				Time: time.Now(),
			}
			switch msg {
			case w.WM_LBUTTONDOWN:
				e.Button |= f.MouseButtons_Left
			case w.WM_RBUTTONDOWN:
				e.Button |= f.MouseButtons_Right
			case w.WM_MBUTTONDOWN:
				e.Button |= f.MouseButtons_Middle
			}
			if _this.onMouseDown(&e); e.IsHandle {
				ret = 1
			}
		}
	case w.WM_LBUTTONUP, w.WM_RBUTTONUP, w.WM_MBUTTONUP:
		_this.fiexdCursor = false
		w.ReleaseCapture()
		if _this.onMouseUp != nil {
			e := f.MouseEvArgs{
				X:    int(w.GET_X_LPARAM(lParam)),
				Y:    int(w.GET_Y_LPARAM(lParam)),
				Time: time.Now(),
			}
			switch msg {
			case w.WM_LBUTTONDOWN:
				e.Button |= f.MouseButtons_Left
			case w.WM_RBUTTONDOWN:
				e.Button |= f.MouseButtons_Right
			case w.WM_MBUTTONDOWN:
				e.Button |= f.MouseButtons_Middle
			}
			if _this.onMouseUp(&e); e.IsHandle {
				ret = 1
			}
		}
	case w.WM_MOUSEWHEEL:
		if _this.onMouseWheel != nil {
			lp, hp := w.LOWORD(int32(wParam)), w.HIWORD(int32(wParam))
			e := f.MouseEvArgs{
				X:     int(w.GET_X_LPARAM(lParam)),
				Y:     int(w.GET_Y_LPARAM(lParam)),
				Delta: int(hp),
				Time:  time.Now(),
			}
			if lp&w.MK_LBUTTON != 0 {
				e.Button |= f.MouseButtons_Left
			}
			if lp&w.MK_MBUTTON != 0 {
				e.Button |= f.MouseButtons_Middle
			}
			if lp&w.MK_RBUTTON != 0 {
				e.Button |= f.MouseButtons_Right
			}
			if _this.onMouseWheel(&e); e.IsHandle {
				ret = 1
			}
		}
	default:
		return 0
	}
	return uintptr(ret)
}

func (_this *winBase) CreateGraphics() f.Graphics {
	hdc := w.GetDC(_this.hWnd())
	g := new(winGraphics).init(hdc)
	g.onClose = func(e *winGraphics) {
		w.ReleaseDC(_this.hWnd(), e.dc)
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
	w.PostMessage(_this.hWnd(), uint32(w.WM_COMMAND), uintptr(cmd_invoke), ptr)
}

func (_this *winBase) execCmd(wParam, lParam uintptr) int {
	switch int(wParam) {
	case cmd_invoke:
		ctx := *((*invokeContext)(unsafe.Pointer(lParam)))
		ctx.fn(ctx.state)
		_this.invokeMap.Delete(lParam)
	case cmd_mouse_click:
		e := *((*f.MouseEvArgs)(unsafe.Pointer(lParam)))
		_this.onMouseClick(&e)
	default:
		return 0
	}
	return 1
}

func (_this *winBase) hWnd() w.HWND {
	return _this.handle
}

func (_this *winBase) GetProvider() p.IProvider {
	return _this.app
}

type invokeContext struct {
	fn    func(state interface{})
	state interface{}
}
