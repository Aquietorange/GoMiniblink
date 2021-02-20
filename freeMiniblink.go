package GoMiniblink

import (
	"fmt"
	"image"
	"math"
	"strconv"
	"strings"
	"time"
	"unsafe"

	fm "gitee.com/aochulai/GoMiniblink/forms"
	cs "gitee.com/aochulai/GoMiniblink/forms/controls"
	win "gitee.com/aochulai/GoMiniblink/forms/windows/win32"
)

var fnCall = "fn" + strconv.FormatInt(time.Now().UnixNano(), 32)

func init() {
	BindJsFunc(JsFnBinding{
		Name: fnCall,
		Fn:   execGoFunc,
	})
}

func execGoFunc(ctx GoFnContext) interface{} {
	wkeName := ctx.Param[0].(string)
	fnName := ctx.Param[1].(string)
	rsName := ctx.Param[2].(string)
	if num, err := strconv.ParseUint(wkeName, 10, 64); err == nil {
		wke := wkeHandle(uintptr(num))
		mb := views[wke]
		if impl, implOk := mb.(*freeMiniblink); implOk {
			fn := impl.fnMap[fnName]
			rs := fn.Call(mb, ctx.Param[3:])
			es := mbApi.wkeGlobalExec(mb.GetHandle())
			v := toJsValue(mb, es, rs)
			mbApi.jsSetGlobal(es, rsName, v)
		}
	}
	return nil
}

type freeMiniblink struct {
	view        *cs.Control
	wke         wkeHandle
	fnMap       map[string]JsFnBinding
	jsIsReady   bool
	frames      []FrameContext
	reqMap      map[wkeNetJob]*freeRequestBeforeEvArgs
	isLockMouse bool
	isBmpPaint  bool

	onRequest     RequestBeforeCallback
	onJsReady     JsReadyCallback
	onConsole     ConsoleCallback
	documentReady DocumentReadyCallback
	paintUpdated  PaintUpdatedCallback
}

func (_this *freeMiniblink) init(control *cs.Control) *freeMiniblink {
	_this.view = control
	_this.fnMap = make(map[string]JsFnBinding)
	_this.reqMap = make(map[wkeNetJob]*freeRequestBeforeEvArgs)
	_this.setView()
	_this.mbInit()
	return _this
}

func (_this *freeMiniblink) SetProxy(info ProxyInfo) {
	mbApi.wkeSetViewProxy(_this.wke, info)
}

func (_this *freeMiniblink) MouseIsEnable() bool {
	return _this.isLockMouse == false
}

func (_this *freeMiniblink) MouseEnable(b bool) {
	_this.isLockMouse = b == false
}

func (_this *freeMiniblink) CallJsFunc(name string, param []interface{}) interface{} {
	es := mbApi.wkeGlobalExec(_this.wke)
	var ps []jsValue
	for i := 0; i < len(param); i++ {
		v := toJsValue(_this, es, param[i])
		ps = append(ps, v)
	}
	fn := mbApi.jsGetGlobal(es, name)
	rs := mbApi.jsCall(es, fn, mbApi.jsUndefined(), ps)
	return toGoValue(_this, es, rs)
}

func (_this *freeMiniblink) JsFunc(name string, fn GoFn, state interface{}) {
	_this.fnMap[name] = JsFnBinding{
		Name:  name,
		State: state,
		Fn:    fn,
	}
	if _this.jsIsReady {
		for _, f := range _this.frames {
			f.RunJs(_this.getJsBindingScript(f.IsMain()))
		}
	}
}

func (_this *freeMiniblink) getJsBindingScript(isMain bool) string {
	rsName := "rs" + strconv.FormatUint(uint64(_this.wke), 32)
	call := fnCall
	if isMain == false {
		call = fmt.Sprintf("window.top[%q]", call)
	}
	var list []string
	for k := range _this.fnMap {
		js := fmt.Sprintf(`window[%q]=function(){
							   var rs=%q;
							   var arr = Array.prototype.slice.call(arguments);
							   var args = [%q,%q,rs].concat(arr);
							   %s.apply(null,args);
							   return window.top[rs];
						   };`,
			k, rsName, strconv.FormatUint(uint64(_this.wke), 10), k, call)
		list = append(list, js)
	}
	return strings.Join(list, ";")
}

func (_this *freeMiniblink) RunJs(script string) interface{} {
	es := mbApi.wkeGlobalExec(_this.wke)
	rs := mbApi.jsEval(es, script)
	return toGoValue(_this, es, rs)
}

func (_this *freeMiniblink) SetOnDocumentReady(callback DocumentReadyCallback) {
	_this.documentReady = callback
}

func (_this *freeMiniblink) SetOnConsole(callback ConsoleCallback) {
	_this.onConsole = callback
}

func (_this *freeMiniblink) SetOnJsReady(callback JsReadyCallback) {
	_this.onJsReady = callback
}

func (_this *freeMiniblink) SetOnRequestBefore(callback RequestBeforeCallback) {
	_this.onRequest = callback
}

func (_this *freeMiniblink) SetOnPaintUpdated(callback PaintUpdatedCallback) {
	_this.paintUpdated = callback
}

func (_this *freeMiniblink) mbInit() {
	_this.wke = createWebView(_this)
	_this.viewResize(_this.view.GetBound().Rect)
	mbApi.wkeSetHandle(_this.wke, _this.view.GetHandle())
	mbApi.wkeOnPaintBitUpdated(_this.wke, _this.onPaintBitUpdated, 0)
	mbApi.wkeOnLoadUrlBegin(_this.wke, _this.onUrlBegin, 0)
	mbApi.wkeOnLoadUrlEnd(_this.wke, _this.onUrlEnd, 0)
	mbApi.wkeOnLoadUrlFail(_this.wke, _this.onUrlFail, 0)
	mbApi.wkeOnDidCreateScriptContext(_this.wke, _this.onDidCreateScriptContext, 0)
	mbApi.wkeOnConsole(_this.wke, _this.jsConsole, 0)
	mbApi.wkeOnDocumentReady(_this.wke, _this.onDocumentReady, 0)
}

func (_this *freeMiniblink) onDocumentReady(_ wkeHandle, _ uintptr, frame wkeFrame) uintptr {
	args := new(freeDocumentReadyEvArgs).init(_this, frame)
	if _this.documentReady != nil {
		_this.documentReady(args)
	}
	return 0
}

func (_this *freeMiniblink) jsConsole(_ wkeHandle, _ uintptr, level int32, msg, name wkeString, line uint32, stack wkeString) uintptr {
	if _this.onConsole == nil {
		return 0
	}
	args := new(freeConsoleMessageEvArgs).init()
	args.line = int(line)
	args.message = mbApi.wkeGetString(msg)
	args.name = mbApi.wkeGetString(name)
	args.stack = mbApi.wkeGetString(stack)
	lv := wkeConsoleLevel(level)
	switch lv {
	case wkeConsoleLevel_Log:
		args.level = "log"
	case wkeConsoleLevel_Warning:
		args.level = "warning"
	case wkeConsoleLevel_Error:
		args.level = "error"
	case wkeConsoleLevel_Debug:
		args.level = "debug"
	case wkeConsoleLevel_Info:
		args.level = "info"
	case wkeConsoleLevel_RevokedError:
		args.level = "revoke"
	default:
		panic("无法识别的类型：" + strconv.Itoa(int(lv)))
	}
	_this.onConsole(args)
	return 0
}

func (_this *freeMiniblink) onDidCreateScriptContext(_ wkeHandle, _ uintptr, frame wkeFrame, _ uintptr, _, _ int) uintptr {
	_this.jsIsReady = true
	args := new(wkeJsReadyEvArgs).init(_this, frame)
	_this.frames = append(_this.frames, args)
	args.RunJs(_this.getJsBindingScript(args.IsMain()))
	if _this.onJsReady == nil {
		return 0
	}
	_this.onJsReady(args)
	return 0
}

func (_this *freeMiniblink) onUrlBegin(_ wkeHandle, _, _ uintptr, job wkeNetJob) uintptr {
	if _this.onRequest == nil {
		return 0
	}
	e := new(freeRequestBeforeEvArgs).init(_this, job)
	e.EvFinish().AddEx(func() {
		delete(_this.reqMap, job)
	})
	_this.onRequest(e)
	e.onBegin()
	_this.reqMap[job] = e
	return 0
}

func (_this *freeMiniblink) onUrlEnd(_ wkeHandle, _, _ uintptr, job wkeNetJob, buf uintptr, len int32) uintptr {
	if req, ok := _this.reqMap[job]; ok {
		data := (*[1 << 30]byte)(unsafe.Pointer(buf))
		rs := make([]byte, int(len))
		for i := int32(0); i < len; i++ {
			rs[i] = data[i]
		}
		req.onResponse(rs)
	}
	return 0
}

func (_this *freeMiniblink) onUrlFail(_ wkeHandle, _, _ uintptr, job wkeNetJob) uintptr {
	if req, ok := _this.reqMap[job]; ok {
		req.onFail()
	}
	return 0
}

func (_this *freeMiniblink) setView() {
	bakFocus := _this.view.OnFocus
	_this.view.OnFocus = func() {
		_this.viewFocus()
		if bakFocus != nil {
			bakFocus()
		}
	}
	bakLostFocus := _this.view.OnLostFocus
	_this.view.OnLostFocus = func() {
		_this.viewLostFocus()
		if bakLostFocus != nil {
			bakLostFocus()
		}
	}
	bakResize := _this.view.OnResize
	_this.view.OnResize = func(e fm.Rect) {
		_this.viewResize(e)
		if bakResize != nil {
			bakResize(e)
		}
	}
	bakPaint := _this.view.OnPaint
	_this.view.OnPaint = func(e fm.PaintEvArgs) {
		_this.viewPaint(e)
		if bakPaint != nil {
			bakPaint(e)
		}
	}
	bakMouseMove := _this.view.OnMouseMove
	_this.view.OnMouseMove = func(e *fm.MouseEvArgs) {
		if _this.isLockMouse == false {
			_this.viewMouseMove(e)
		}
		if bakMouseMove != nil {
			bakMouseMove(e)
		}
	}
	bakMouseDown := _this.view.OnMouseDown
	_this.view.OnMouseDown = func(e *fm.MouseEvArgs) {
		if _this.isLockMouse == false {
			_this.viewMouseDown(e)
		}
		if bakMouseDown != nil {
			bakMouseDown(e)
		}
	}
	bakMouseUp := _this.view.OnMouseUp
	_this.view.OnMouseUp = func(e *fm.MouseEvArgs) {
		if _this.isLockMouse == false {
			_this.viewMouseUp(e)
		}
		if bakMouseUp != nil {
			bakMouseUp(e)
		}
	}
	bakMouseWheel := _this.view.OnMouseWheel
	_this.view.OnMouseWheel = func(e *fm.MouseEvArgs) {
		if _this.isLockMouse == false {
			_this.viewMouseWheel(e)
		}
		if bakMouseWheel != nil {
			bakMouseWheel(e)
		}
	}
	bakSetCursor := _this.view.OnSetCursor
	_this.view.OnSetCursor = func() bool {
		b := _this.viewSetCursor()
		if !b && bakSetCursor != nil {
			b = bakSetCursor()
		}
		return b
	}
	bakKeyDown := _this.view.OnKeyDown
	_this.view.OnKeyDown = func(e *fm.KeyEvArgs) {
		_this.viewKeyDown(e)
		if bakKeyDown != nil {
			bakKeyDown(e)
		}
	}
	bakKeyUp := _this.view.OnKeyUp
	_this.view.OnKeyUp = func(e *fm.KeyEvArgs) {
		_this.viewKeyUp(e)
		if bakKeyUp != nil {
			bakKeyUp(e)
		}
	}
	bakKeyPress := _this.view.OnKeyPress
	_this.view.OnKeyPress = func(e *fm.KeyPressEvArgs) {
		_this.viewKeyPress(e)
		if bakKeyPress != nil {
			bakKeyPress(e)
		}
	}
	bakImeStart := _this.view.OnImeStartComposition
	_this.view.OnImeStartComposition = func() bool {
		b := _this.viewImeStart()
		if !b && bakImeStart != nil {
			b = bakImeStart()
		}
		return b
	}
}

func (_this *freeMiniblink) viewImeStart() bool {
	rect := mbApi.wkeGetCaretRect(_this.wke)
	comp := win.COMPOSITIONFORM{
		DwStyle: win.CFS_POINT | win.CFS_FORCE_POSITION,
		Pos: win.POINT{
			X: rect.x,
			Y: rect.y,
		},
	}
	h := win.HWND(_this.view.GetHandle())
	imc := win.ImmGetContext(h)
	win.ImmSetCompositionWindow(imc, &comp)
	win.ImmReleaseContext(h, imc)
	return true
}

func (_this *freeMiniblink) viewKeyPress(e *fm.KeyPressEvArgs) {
	if mbApi.wkeFireKeyPressEvent(_this.wke, int([]rune(e.KeyChar)[0]), uint32(wkeKeyFlags_Repeat), e.IsSys) {
		e.IsHandle = true
	}
}

func (_this *freeMiniblink) viewKeyUp(e *fm.KeyEvArgs) {
	if _this.viewKeyEvent(e, false) {
		e.IsHandle = true
	}
}

func (_this *freeMiniblink) viewKeyDown(e *fm.KeyEvArgs) {
	if _this.viewKeyEvent(e, true) {
		e.IsHandle = true
	}
}

func (_this *freeMiniblink) viewKeyEvent(e *fm.KeyEvArgs, isDown bool) bool {
	flags := int(wkeKeyFlags_Repeat)
	switch e.Key {
	case fm.Keys_Insert, fm.Keys_Delete, fm.Keys_Home, fm.Keys_End, fm.Keys_PageUp,
		fm.Keys_PageDown, fm.Keys_Left, fm.Keys_Right, fm.Keys_Up, fm.Keys_Down:
		flags |= int(wkeKeyFlags_Extend)
	}
	if isDown {
		return mbApi.wkeFireKeyDownEvent(_this.wke, uint32(e.Value), uint32(flags), e.IsSys)
	} else {
		return mbApi.wkeFireKeyUpEvent(_this.wke, uint32(e.Value), uint32(flags), e.IsSys)
	}
}

func (_this *freeMiniblink) LoadUri(uri string) {
	mbApi.wkeLoadURL(_this.wke, uri)
}

func (_this *freeMiniblink) SetDebugConfig(debugString string, param string) {
	mbApi.wkeSetDebugConfig(_this.wke, debugString, param)
}

func (_this *freeMiniblink) viewSetCursor() bool {
	cur := mbApi.wkeGetCursorInfoType(_this.wke)
	newCur := fm.CursorType_Default
	switch cur {
	case wkeCursorType_Pointer:
		newCur = fm.CursorType_ARROW
	case wkeCursorType_Hand:
		newCur = fm.CursorType_HAND
	case wkeCursorType_IBeam:
		newCur = fm.CursorType_IBEAM
	case wkeCursorType_ColumnResize:
		newCur = fm.CursorType_SIZEWE
	case wkeCursorType_RowResize:
		newCur = fm.CursorType_SIZENS
	case wkeCursorType_Cross:
		newCur = fm.CursorType_CROSS
	default:
		fmt.Println("未实现的鼠标指针类型：" + strconv.Itoa(int(cur)))
	}
	_this.view.SetCursor(newCur)
	return true
}

func (_this *freeMiniblink) viewMouseWheel(e *fm.MouseEvArgs) {
	flags := wkeMouseFlags_None
	keys := cs.App.ModifierKeys()
	if s, ok := keys[fm.Keys_Ctrl]; ok && s {
		flags |= wkeMouseFlags_CONTROL
	}
	if s, ok := keys[fm.Keys_Shift]; ok && s {
		flags |= wkeMouseFlags_SHIFT
	}
	if e.Button&fm.MouseButtons_Left != 0 {
		flags |= wkeMouseFlags_LBUTTON
	}
	if e.Button&fm.MouseButtons_Right != 0 {
		flags |= wkeMouseFlags_RBUTTON
	}
	if e.Button&fm.MouseButtons_Middle != 0 {
		flags |= wkeMouseFlags_MBUTTON
	}
	if mbApi.wkeFireMouseWheelEvent(_this.wke, int32(e.X), int32(e.Y), int32(e.Delta), int32(flags)) {
		e.IsHandle = true
	}
}

func (_this *freeMiniblink) viewMouseUp(e *fm.MouseEvArgs) {
	_this.viewMouseEvent(e, false)
}

func (_this *freeMiniblink) viewMouseDown(e *fm.MouseEvArgs) {
	_this.viewMouseEvent(e, true)
}

func (_this *freeMiniblink) viewMouseEvent(e *fm.MouseEvArgs, isDown bool) {
	flags := wkeMouseFlags_None
	keys := cs.App.ModifierKeys()
	if s, ok := keys[fm.Keys_Ctrl]; ok && s {
		flags |= wkeMouseFlags_CONTROL
	}
	if s, ok := keys[fm.Keys_Shift]; ok && s {
		flags |= wkeMouseFlags_SHIFT
	}
	msg := 0
	if e.Button&fm.MouseButtons_Left != 0 {
		flags |= wkeMouseFlags_LBUTTON
		if e.IsDouble {
			msg = win.WM_LBUTTONDBLCLK
		} else if isDown {
			msg = win.WM_LBUTTONDOWN
		} else {
			msg = win.WM_LBUTTONUP
		}
	}
	if e.Button&fm.MouseButtons_Right != 0 {
		flags |= wkeMouseFlags_RBUTTON
		if e.IsDouble {
			msg = win.WM_RBUTTONDBLCLK
		} else if isDown {
			msg = win.WM_RBUTTONDOWN
		} else {
			msg = win.WM_RBUTTONUP
		}
	}
	if e.Button&fm.MouseButtons_Middle != 0 {
		flags |= wkeMouseFlags_MBUTTON
		if e.IsDouble {
			msg = win.WM_MBUTTONDBLCLK
		} else if isDown {
			msg = win.WM_MBUTTONDOWN
		} else {
			msg = win.WM_MBUTTONUP
		}
	}
	if msg != 0 && mbApi.wkeFireMouseEvent(_this.wke, int32(msg), int32(e.X), int32(e.Y), int32(flags)) {
		e.IsHandle = true
	}
}

func (_this *freeMiniblink) viewMouseMove(e *fm.MouseEvArgs) {
	flags := wkeMouseFlags_None
	if e.Button&fm.MouseButtons_Left != 0 {
		flags |= wkeMouseFlags_LBUTTON
	}
	if e.Button&fm.MouseButtons_Right != 0 {
		flags |= wkeMouseFlags_RBUTTON
	}
	if mbApi.wkeFireMouseEvent(_this.wke, int32(win.WM_MOUSEMOVE), int32(e.X), int32(e.Y), int32(flags)) {
		e.IsHandle = true
	}
}

func (_this *freeMiniblink) ToBitmap() *image.RGBA {
	w := mbApi.wkeGetWidth(_this.wke)
	h := mbApi.wkeGetHeight(_this.wke)
	w = uint32(math.Max(float64(w), 0))
	h = uint32(math.Max(float64(h), 0))
	view := image.NewRGBA(image.Rect(0, 0, int(w), int(h)))
	mbApi.wkePaint(_this.wke, view.Pix, 0)
	return view
}

func (_this *freeMiniblink) viewPaint(e fm.PaintEvArgs) {
	if _this.isBmpPaint {
		img := _this.ToBitmap()
		if img.Bounds().Empty() == false {
			e.Graphics.DrawImage(img, e.Clip.X, e.Clip.Y, e.Clip.Width, e.Clip.Height, e.Clip.X, e.Clip.Y)
		}
	} else {
		hdc := win.HDC(e.Graphics.GetHandle())
		mdc := win.HDC(mbApi.wkeGetViewDC(_this.wke))
		win.BitBlt(hdc, int32(e.Clip.X), int32(e.Clip.Y), int32(e.Clip.Width), int32(e.Clip.Height), mdc, int32(e.Clip.X), int32(e.Clip.Y), win.SRCCOPY)
	}
}

func (_this *freeMiniblink) onPaintBitUpdated(wke wkeHandle, _, bits uintptr, rect *wkeRect, width, _ int32) uintptr {
	bx, by := int(rect.x), int(rect.y)
	bw, bh := int(math.Min(float64(rect.w), float64(width))), int(math.Min(float64(rect.h), float64(mbApi.wkeGetHeight(wke))))
	var bmp *image.RGBA
	if _this.isBmpPaint {
		bmp = image.NewRGBA(image.Rect(0, 0, bw, bh))
		stride := int(width) * 4
		pixs := (*[1 << 30]byte)(unsafe.Pointer(bits))
		for y := 0; y < bh; y++ {
			for x := 0; x < bw*4; x++ {
				sp := bmp.Stride*y + x
				dp := stride*(by+y) + bx*4 + x
				bmp.Pix[sp] = pixs[dp]
			}
		}
	}
	args := new(freePaintUpdatedEvArgs).init(bmp, fm.Bound{
		Point: fm.Point{
			X: bx,
			Y: by,
		},
		Rect: fm.Rect{
			Width:  bw,
			Height: bh,
		},
	})
	if _this.paintUpdated != nil {
		_this.paintUpdated(args)
	}
	if args.IsCancel() {
		return 0
	} else if _this.isBmpPaint {
		_this.view.CreateGraphics().DrawImage(bmp, 0, 0, bw, bh, bx, by).Close()
	} else {
		r := win.RECT{
			Left:   rect.x,
			Top:    rect.y,
			Right:  rect.x + rect.w,
			Bottom: rect.y + rect.h,
		}
		win.InvalidateRect(win.HWND(_this.view.GetHandle()), &r, false)
	}
	return 0
}

func (_this *freeMiniblink) SetBmpPaintMode(b bool) {
	_this.isBmpPaint = b
}

func (_this *freeMiniblink) viewResize(e fm.Rect) {
	mbApi.wkeResize(_this.wke, uint32(e.Width), uint32(e.Height))
}

func (_this *freeMiniblink) viewLostFocus() {
	mbApi.wkeKillFocus(_this.wke)
}

func (_this *freeMiniblink) viewFocus() {
	mbApi.wkeSetFocus(_this.wke)
}

func (_this *freeMiniblink) GetHandle() wkeHandle {
	return _this.wke
}
