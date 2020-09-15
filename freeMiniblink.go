package GoMiniblink

import (
	"fmt"
	"image"
	"math"
	f "qq2564874169/goMiniblink/forms"
	c "qq2564874169/goMiniblink/forms/controls"
	w "qq2564874169/goMiniblink/forms/platform/windows/win32"
	"strconv"
	"strings"
	"time"
	"unsafe"
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
			fn := impl._fnMap[fnName]
			rs := fn.Call(mb, ctx.Param[3:])
			es := mbApi.wkeGlobalExec(mb.GetHandle())
			v := toJsValue(mb, es, rs)
			mbApi.jsSetGlobal(es, rsName, v)
		}
	}
	return nil
}

type freeMiniblink struct {
	_view      *c.Control
	_wke       wkeHandle
	_fnMap     map[string]JsFnBinding
	_jsIsReady bool
	_frames    []FrameContext
	_reqMap    map[wkeNetJob]*freeRequestBeforeEvArgs

	_onRequest     RequestBeforeCallback
	_onJsReady     JsReadyCallback
	_onConsole     ConsoleCallback
	_documentReady DocumentReadyCallback
	_paintUpdated  PaintUpdatedCallback
}

func (_this *freeMiniblink) init(control *c.Control) *freeMiniblink {
	_this._view = control
	_this._fnMap = make(map[string]JsFnBinding)
	_this._reqMap = make(map[wkeNetJob]*freeRequestBeforeEvArgs)
	_this.setView()
	_this.mbInit()
	return _this
}

func (_this *freeMiniblink) CallJsFunc(name string, param ...interface{}) interface{} {
	es := mbApi.wkeGlobalExec(_this.GetHandle())
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
	_this._fnMap[name] = JsFnBinding{
		Name:  name,
		State: state,
		Fn:    fn,
	}
	if _this._jsIsReady {
		for _, f := range _this._frames {
			f.RunJs(_this.getJsBindingScript(f.IsMain()))
		}
	}
}

func (_this *freeMiniblink) getJsBindingScript(isMain bool) string {
	rsName := "rs" + strconv.FormatUint(uint64(_this._wke), 32)
	call := fnCall
	if isMain == false {
		call = fmt.Sprintf("window.top[%q]", call)
	}
	var list []string
	for k, _ := range _this._fnMap {
		js := `window[%q]=function(){
               var rs=%q;
               var arr = Array.prototype.slice.call(arguments);
               var args = [%q,%q,rs].concat(arr);
               %s.apply(null,args);
               return window.top[rs];
           };`
		js = fmt.Sprintf(js, k, rsName, strconv.FormatUint(uint64(_this._wke), 10), k, call)
		list = append(list, js)
	}
	return strings.Join(list, ";")
}

func (_this *freeMiniblink) RunJs(script string) interface{} {
	es := mbApi.wkeGlobalExec(_this._wke)
	rs := mbApi.jsEval(es, script)
	return toGoValue(_this, es, rs)
}

func (_this *freeMiniblink) SetOnDocumentReady(callback DocumentReadyCallback) {
	_this._documentReady = callback
}

func (_this *freeMiniblink) SetOnConsole(callback ConsoleCallback) {
	_this._onConsole = callback
}

func (_this *freeMiniblink) SetOnJsReady(callback JsReadyCallback) {
	_this._onJsReady = callback
}

func (_this *freeMiniblink) SetOnRequestBefore(callback RequestBeforeCallback) {
	_this._onRequest = callback
}

func (_this *freeMiniblink) SetOnPaintUpdated(callback PaintUpdatedCallback) {
	_this._paintUpdated = callback
}

func (_this *freeMiniblink) mbInit() {
	_this._wke = createWebView(_this)
	_this.viewResize(_this._view.GetSize())
	mbApi.wkeSetHandle(_this._wke, _this._view.GetHandle())
	mbApi.wkeOnPaintBitUpdated(_this._wke, _this.onPaintBitUpdated, 0)
	mbApi.wkeOnLoadUrlBegin(_this._wke, _this.onUrlBegin, 0)
	mbApi.wkeOnLoadUrlEnd(_this._wke, _this.onUrlEnd, 0)
	mbApi.wkeOnLoadUrlFail(_this._wke, _this.onUrlFail, 0)
	mbApi.wkeOnDidCreateScriptContext(_this._wke, _this.onDidCreateScriptContext, 0)
	mbApi.wkeOnConsole(_this._wke, _this.onConsole, 0)
	mbApi.wkeOnDocumentReady(_this._wke, _this.onDocumentReady, 0)
}

func (_this *freeMiniblink) onDocumentReady(_ wkeHandle, _ uintptr, frame wkeFrame) uintptr {
	args := new(freeDocumentReadyEvArgs).init(_this, frame)
	if _this._documentReady != nil {
		_this._documentReady(args)
	}
	return 0
}

func (_this *freeMiniblink) onConsole(_ wkeHandle, _ uintptr, level int32, msg, name wkeString, line uint32, stack wkeString) uintptr {
	if _this._onConsole == nil {
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
	_this._onConsole(args)
	return 0
}

func (_this *freeMiniblink) onDidCreateScriptContext(_ wkeHandle, _ uintptr, frame wkeFrame, _ uintptr, _, _ int) uintptr {
	_this._jsIsReady = true
	args := new(wkeJsReadyEvArgs).init(_this, frame)
	_this._frames = append(_this._frames, args)
	args.RunJs(_this.getJsBindingScript(args.IsMain()))
	if _this._onJsReady == nil {
		return 0
	}
	_this._onJsReady(args)
	return 0
}

func (_this *freeMiniblink) onUrlBegin(_ wkeHandle, _, _ uintptr, job wkeNetJob) uintptr {
	if _this._onRequest == nil {
		return 0
	}
	e := new(freeRequestBeforeEvArgs).init(_this, job)
	e.EvFinish().AddEx(func() {
		delete(_this._reqMap, job)
	})
	_this._onRequest(e)
	e.onBegin()
	_this._reqMap[job] = e
	return 0
}

func (_this *freeMiniblink) onUrlEnd(_ wkeHandle, _, _ uintptr, job wkeNetJob, buf uintptr, len int32) uintptr {
	if req, ok := _this._reqMap[job]; ok {
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
	if req, ok := _this._reqMap[job]; ok {
		req.onFail()
	}
	return 0
}

func (_this *freeMiniblink) setView() {
	bakFocus := _this._view.OnFocus
	_this._view.OnFocus = func() {
		_this.viewFocus()
		if bakFocus != nil {
			bakFocus()
		}
	}
	bakLostFocus := _this._view.OnLostFocus
	_this._view.OnLostFocus = func() {
		_this.viewLostFocus()
		if bakLostFocus != nil {
			bakLostFocus()
		}
	}
	bakResize := _this._view.OnResize
	_this._view.OnResize = func(e f.Rect) {
		_this.viewResize(e)
		if bakResize != nil {
			bakResize(e)
		}
	}
	bakPaint := _this._view.OnPaint
	_this._view.OnPaint = func(e f.PaintEvArgs) {
		_this.viewPaint(e)
		if bakPaint != nil {
			bakPaint(e)
		}
	}
	bakMouseMove := _this._view.OnMouseMove
	_this._view.OnMouseMove = func(e *f.MouseEvArgs) {
		_this.viewMouseMove(e)
		if bakMouseMove != nil {
			bakMouseMove(e)
		}
	}
	bakMouseDown := _this._view.OnMouseDown
	_this._view.OnMouseDown = func(e *f.MouseEvArgs) {
		_this.viewMouseDown(e)
		if bakMouseDown != nil {
			bakMouseDown(e)
		}
	}
	bakMouseUp := _this._view.OnMouseUp
	_this._view.OnMouseUp = func(e *f.MouseEvArgs) {
		_this.viewMouseUp(e)
		if bakMouseUp != nil {
			bakMouseUp(e)
		}
	}
	bakMouseWheel := _this._view.OnMouseWheel
	_this._view.OnMouseWheel = func(e *f.MouseEvArgs) {
		_this.viewMouseWheel(e)
		if bakMouseWheel != nil {
			bakMouseWheel(e)
		}
	}
	bakSetCursor := _this._view.OnSetCursor
	_this._view.OnSetCursor = func() bool {
		b := _this.viewSetCursor()
		if !b && bakSetCursor != nil {
			b = bakSetCursor()
		}
		return b
	}
	bakKeyDown := _this._view.OnKeyDown
	_this._view.OnKeyDown = func(e *f.KeyEvArgs) {
		_this.viewKeyDown(e)
		if bakKeyDown != nil {
			bakKeyDown(e)
		}
	}
	bakKeyUp := _this._view.OnKeyUp
	_this._view.OnKeyUp = func(e *f.KeyEvArgs) {
		_this.viewKeyUp(e)
		if bakKeyUp != nil {
			bakKeyUp(e)
		}
	}
	bakKeyPress := _this._view.OnKeyPress
	_this._view.OnKeyPress = func(e *f.KeyPressEvArgs) {
		_this.viewKeyPress(e)
		if bakKeyPress != nil {
			bakKeyPress(e)
		}
	}
	bakImeStart := _this._view.OnImeStartComposition
	_this._view.OnImeStartComposition = func() bool {
		b := _this.viewImeStart()
		if !b && bakImeStart != nil {
			b = bakImeStart()
		}
		return b
	}
}

func (_this *freeMiniblink) viewImeStart() bool {
	rect := mbApi.wkeGetCaretRect(_this._wke)
	comp := w.COMPOSITIONFORM{
		DwStyle: w.CFS_POINT | w.CFS_FORCE_POSITION,
		Pos: w.POINT{
			X: rect.x,
			Y: rect.y,
		},
	}
	h := w.HWND(_this._view.GetHandle())
	imc := w.ImmGetContext(h)
	w.ImmSetCompositionWindow(imc, &comp)
	w.ImmReleaseContext(h, imc)
	return true
}

func (_this *freeMiniblink) viewKeyPress(e *f.KeyPressEvArgs) {
	if mbApi.wkeFireKeyPressEvent(_this._wke, int([]rune(e.KeyChar)[0]), uint32(wkeKeyFlags_Repeat), e.IsSys) {
		e.IsHandle = true
	}
}

func (_this *freeMiniblink) viewKeyUp(e *f.KeyEvArgs) {
	if _this.viewKeyEvent(e, false) {
		e.IsHandle = true
	}
}

func (_this *freeMiniblink) viewKeyDown(e *f.KeyEvArgs) {
	if _this.viewKeyEvent(e, true) {
		e.IsHandle = true
	}
}

func (_this *freeMiniblink) viewKeyEvent(e *f.KeyEvArgs, isDown bool) bool {
	flags := int(wkeKeyFlags_Repeat)
	switch e.Key {
	case f.Keys_Insert, f.Keys_Delete, f.Keys_Home, f.Keys_End, f.Keys_PageUp,
		f.Keys_PageDown, f.Keys_Left, f.Keys_Right, f.Keys_Up, f.Keys_Down:
		flags |= int(wkeKeyFlags_Extend)
	}
	if isDown {
		return mbApi.wkeFireKeyDownEvent(_this._wke, uint32(e.Value), uint32(flags), e.IsSys)
	} else {
		return mbApi.wkeFireKeyUpEvent(_this._wke, uint32(e.Value), uint32(flags), e.IsSys)
	}
}

func (_this *freeMiniblink) LoadUri(uri string) {
	mbApi.wkeLoadURL(_this._wke, uri)
}

func (_this *freeMiniblink) viewSetCursor() bool {
	cur := mbApi.wkeGetCursorInfoType(_this._wke)
	newCur := f.CursorType_Default
	switch cur {
	case wkeCursorType_Pointer:
		newCur = f.CursorType_ARROW
	case wkeCursorType_Hand:
		newCur = f.CursorType_HAND
	case wkeCursorType_IBeam:
		newCur = f.CursorType_IBEAM
	case wkeCursorType_ColumnResize:
		newCur = f.CursorType_SIZEWE
	case wkeCursorType_RowResize:
		newCur = f.CursorType_SIZENS
	case wkeCursorType_Cross:
		newCur = f.CursorType_CROSS
	default:
		fmt.Println("未实现的鼠标指针类型：" + strconv.Itoa(int(cur)))
	}
	_this._view.SetCursor(newCur)
	return true
}

func (_this *freeMiniblink) viewMouseWheel(e *f.MouseEvArgs) {
	flags := wkeMouseFlags_None
	keys := c.App.ModifierKeys()
	if s, ok := keys[f.Keys_Ctrl]; ok && s {
		flags |= wkeMouseFlags_CONTROL
	}
	if s, ok := keys[f.Keys_Shift]; ok && s {
		flags |= wkeMouseFlags_SHIFT
	}
	if e.Button&f.MouseButtons_Left != 0 {
		flags |= wkeMouseFlags_LBUTTON
	}
	if e.Button&f.MouseButtons_Right != 0 {
		flags |= wkeMouseFlags_RBUTTON
	}
	if e.Button&f.MouseButtons_Middle != 0 {
		flags |= wkeMouseFlags_MBUTTON
	}
	if mbApi.wkeFireMouseWheelEvent(_this._wke, int32(e.X), int32(e.Y), int32(e.Delta), int32(flags)) {
		e.IsHandle = true
	}
}

func (_this *freeMiniblink) viewMouseUp(e *f.MouseEvArgs) {
	_this.viewMouseEvent(e, false)
}

func (_this *freeMiniblink) viewMouseDown(e *f.MouseEvArgs) {
	_this.viewMouseEvent(e, true)
}

func (_this *freeMiniblink) viewMouseEvent(e *f.MouseEvArgs, isDown bool) {
	flags := wkeMouseFlags_None
	keys := c.App.ModifierKeys()
	if s, ok := keys[f.Keys_Ctrl]; ok && s {
		flags |= wkeMouseFlags_CONTROL
	}
	if s, ok := keys[f.Keys_Shift]; ok && s {
		flags |= wkeMouseFlags_SHIFT
	}
	msg := 0
	if e.Button&f.MouseButtons_Left != 0 {
		flags |= wkeMouseFlags_LBUTTON
		if e.IsDouble {
			msg = w.WM_LBUTTONDBLCLK
		} else if isDown {
			msg = w.WM_LBUTTONDOWN
		} else {
			msg = w.WM_LBUTTONUP
		}
	}
	if e.Button&f.MouseButtons_Right != 0 {
		flags |= wkeMouseFlags_RBUTTON
		if e.IsDouble {
			msg = w.WM_RBUTTONDBLCLK
		} else if isDown {
			msg = w.WM_RBUTTONDOWN
		} else {
			msg = w.WM_RBUTTONUP
		}
	}
	if e.Button&f.MouseButtons_Middle != 0 {
		flags |= wkeMouseFlags_MBUTTON
		if e.IsDouble {
			msg = w.WM_MBUTTONDBLCLK
		} else if isDown {
			msg = w.WM_MBUTTONDOWN
		} else {
			msg = w.WM_MBUTTONUP
		}
	}
	if msg != 0 && mbApi.wkeFireMouseEvent(_this._wke, int32(msg), int32(e.X), int32(e.Y), int32(flags)) {
		e.IsHandle = true
	}
}

func (_this *freeMiniblink) viewMouseMove(e *f.MouseEvArgs) {
	flags := wkeMouseFlags_None
	if e.Button&f.MouseButtons_Left != 0 {
		flags |= wkeMouseFlags_LBUTTON
	}
	if e.Button&f.MouseButtons_Right != 0 {
		flags |= wkeMouseFlags_RBUTTON
	}
	if mbApi.wkeFireMouseEvent(_this._wke, int32(w.WM_MOUSEMOVE), int32(e.X), int32(e.Y), int32(flags)) {
		e.IsHandle = true
	}
}

func (_this *freeMiniblink) ToBitmap() *image.RGBA {
	w := mbApi.wkeGetWidth(_this._wke)
	h := mbApi.wkeGetHeight(_this._wke)
	w = uint32(math.Max(float64(w), 0))
	h = uint32(math.Max(float64(h), 0))
	view := image.NewRGBA(image.Rect(0, 0, int(w), int(h)))
	mbApi.wkePaint(_this._wke, view.Pix, 0)
	return view
}

func (_this *freeMiniblink) viewPaint(e f.PaintEvArgs) {
	img := _this.ToBitmap()
	if img.Bounds().Empty() == false {
		e.Graphics.DrawImage(img, e.Clip.X, e.Clip.Y, e.Clip.Width, e.Clip.Height, e.Clip.X, e.Clip.Y)
	}
}

func (_this *freeMiniblink) onPaintBitUpdated(wke wkeHandle, _, bits uintptr, rect *wkeRect, width, _ int32) uintptr {
	bx, by := int(rect.x), int(rect.y)
	bw, bh := int(math.Min(float64(rect.w), float64(width))), int(math.Min(float64(rect.h), float64(mbApi.wkeGetHeight(wke))))
	bmp := image.NewRGBA(image.Rect(0, 0, bw, bh))
	stride := int(width) * 4
	pixs := (*[1 << 30]byte)(unsafe.Pointer(bits))
	for y := 0; y < bh; y++ {
		for x := 0; x < bw*4; x++ {
			sp := bmp.Stride*y + x
			dp := stride*(by+y) + bx*4 + x
			bmp.Pix[sp] = pixs[dp]
		}
	}
	args := new(freePaintUpdatedEvArgs).init(bmp, f.Bound{
		Point: f.Point{
			X: bx,
			Y: by,
		},
		Rect: f.Rect{
			Width:  bw,
			Height: bh,
		},
	})
	if _this._paintUpdated != nil {
		_this._paintUpdated(args)
	}
	if args.IsCancel() == false {
		_this._view.CreateGraphics().DrawImage(bmp, 0, 0, bw, bh, bx, by).Close()
	}
	return 0
}

func (_this *freeMiniblink) viewResize(e f.Rect) {
	mbApi.wkeResize(_this._wke, uint32(e.Width), uint32(e.Height))
}

func (_this *freeMiniblink) viewLostFocus() {
	mbApi.wkeKillFocus(_this._wke)
}

func (_this *freeMiniblink) viewFocus() {
	mbApi.wkeSetFocus(_this._wke)
}

func (_this *freeMiniblink) GetHandle() wkeHandle {
	return _this._wke
}
