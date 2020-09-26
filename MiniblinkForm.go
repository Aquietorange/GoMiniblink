package GoMiniblink

import (
	"fmt"
	"image"
	fm "qq2564874169/goMiniblink/forms"
	cs "qq2564874169/goMiniblink/forms/controls"
	win "qq2564874169/goMiniblink/forms/windows/win32"
	"time"
	"unsafe"
)

var (
	_fnMax   = "__formMax"
	_fnMin   = "__formMin"
	_fnClose = "__formClose"
	_fnDrop  = "__formDrop"
)

type MiniblinkForm struct {
	cs.Form
	View *MiniblinkBrowser

	_isLoad        bool
	_isTransparent bool
	_wke           wkeHandle
	_rsState       int
	_isDrop        bool
}

func (_this *MiniblinkForm) Init() *MiniblinkForm {
	_this.Form.Init()
	_this.View = new(MiniblinkBrowser).Init()
	_this.View.SetAnchor(fm.AnchorStyle_Top | fm.AnchorStyle_Right | fm.AnchorStyle_Bottom | fm.AnchorStyle_Left)
	_this.AddChild(_this.View)

	bakOnResize := _this.OnResize
	_this.OnResize = func(e fm.Rect) {
		_this.View.SetSize(e.Width, e.Height)
		bakOnResize(e)
	}
	bakOnLoad := _this.OnShow
	_this.OnShow = func() {
		_this._wke = wkeHandle(_this.View.GetMiniblinkHandle())
		_this.View.OnFocus()
		_this.View.JsFuncEx(_fnMax, func() {
			if _this.GetState() == fm.FormState_Max {
				_this.SetState(fm.FormState_Normal)
			} else {
				_this.SetState(fm.FormState_Max)
			}
		})
		_this.View.JsFuncEx(_fnMin, func() {
			_this.SetState(fm.FormState_Min)
		})
		_this.View.JsFuncEx(_fnClose, func() {
			_this.Close()
		})
		_this.View.JsFuncEx(_fnDrop, _this.fnDrop)

		bakEvJsReady := _this.View.OnJsReady
		_this.View.OnJsReady = func(e JsReadyEvArgs) {
			bakEvJsReady(e)
			_this.setFormFn(e)
		}
		_this.View.EvDocumentReady["__goMiniblink"] = func(_ *MiniblinkBrowser, e DocumentReadyEvArgs) {
			e.RunJs("window.setFormButton();window.mbFormDrop();")
		}
		bakOnLoad()
		if _this._isTransparent {
			img := _this.View.ToBitmap()
			_this.transparentPaint(img, img.Bounds().Dx(), img.Bounds().Dy())
		}
		_this._isLoad = true
	}
	return _this
}

func (_this *MiniblinkForm) NoneBorderResize() {
	padd := 5
	hWnd := win.HWND(_this.GetHandle())
	onMove := _this.View.OnMouseMove
	onDown := _this.View.OnMouseDown
	_this.View.OnMouseDown = func(e *fm.MouseEvArgs) {
		if e.Button&fm.MouseButtons_Left != fm.MouseButtons_Left {
			onDown(e)
			return
		}
		switch _this._rsState {
		case 8:
			win.SendMessage(hWnd, win.WM_SYSCOMMAND, uintptr(win.SC_SIZE|0xF003), 0)
		case 2:
			win.SendMessage(hWnd, win.WM_SYSCOMMAND, uintptr(win.SC_SIZE|0xF006), 0)
		case 4:
			win.SendMessage(hWnd, win.WM_SYSCOMMAND, uintptr(win.SC_SIZE|0xF001), 0)
		case 6:
			win.SendMessage(hWnd, win.WM_SYSCOMMAND, uintptr(win.SC_SIZE|0xF002), 0)
		case 7:
			win.SendMessage(hWnd, win.WM_SYSCOMMAND, uintptr(win.SC_SIZE|0xF004), 0)
		case 9:
			win.SendMessage(hWnd, win.WM_SYSCOMMAND, uintptr(win.SC_SIZE|0xF005), 0)
		case 1:
			win.SendMessage(hWnd, win.WM_SYSCOMMAND, uintptr(win.SC_SIZE|0xF007), 0)
		case 3:
			win.SendMessage(hWnd, win.WM_SYSCOMMAND, uintptr(win.SC_SIZE|0xF008), 0)
		default:
			onDown(e)
		}
		_this._rsState = 0
	}
	_this.View.OnMouseMove = func(e *fm.MouseEvArgs) {
		size := _this.GetSize()
		if e.X <= padd {
			if e.Y <= padd {
				_this._rsState = 7
			} else if e.Y+padd >= size.Height {
				_this._rsState = 1
			} else {
				_this._rsState = 4
			}
		} else if e.Y <= padd {
			if e.X <= padd {
				_this._rsState = 7
			} else if e.X+padd >= size.Width {
				_this._rsState = 9
			} else {
				_this._rsState = 8
			}
		} else if e.X+padd >= size.Width {
			if e.Y <= padd {
				_this._rsState = 9
			} else if e.Y+padd >= size.Height {
				_this._rsState = 3
			} else {
				_this._rsState = 6
			}
		} else if e.Y+padd >= size.Height {
			if e.X <= padd {
				_this._rsState = 1
			} else if e.X+padd >= size.Width {
				_this._rsState = 3
			} else {
				_this._rsState = 2
			}
		} else if _this._isDrop == false {
			_this._rsState = 0
		}
		switch _this._rsState {
		case 8, 2:
			_this.SetCursor(fm.CursorType_SIZENS)
		case 4, 6:
			_this.SetCursor(fm.CursorType_SIZEWE)
		case 7, 3:
			_this.SetCursor(fm.CursorType_SIZENWSE)
		case 9, 1:
			_this.SetCursor(fm.CursorType_SIZENESW)
		}
		onMove(e)
	}
}

func (_this *MiniblinkForm) TransparentMode() {
	if _this._wke == 0 {
		panic("必须在Miniblink初始化之后设置")
	}
	if _this._isLoad {
		panic("必须在窗口加载完毕之前设置")
	}
	_this._isTransparent = true
	_this.SetBorderStyle(fm.FormBorder_None)
	hWnd := win.HWND(_this.GetHandle())
	style := win.GetWindowLong(hWnd, win.GWL_EXSTYLE)
	if style&win.WS_EX_LAYERED != win.WS_EX_LAYERED {
		win.SetWindowLong(hWnd, win.GWL_EXSTYLE, style|win.WS_EX_LAYERED)
	}
	mbApi.wkeSetTransparent(_this._wke, true)
	_this.View.OnPaintUpdated = func(e PaintUpdatedEvArgs) {
		_this.transparentPaint(e.Bitmap(), e.Bound().Width, e.Bound().Height)
		e.Cancel()
	}
}

func (_this *MiniblinkForm) transparentPaint(image *image.RGBA, width, height int) {
	hWnd := win.HWND(_this.GetHandle())
	hdc := win.GetDC(hWnd)
	memDc := win.CreateCompatibleDC(0)
	src := win.POINT{}
	dst := win.POINT{
		X: int32(_this.GetLocation().X),
		Y: int32(_this.GetLocation().Y),
	}
	size := win.SIZE{
		CX: int32(width),
		CY: int32(height),
	}
	var head win.BITMAPV5HEADER
	head.BiSize = uint32(unsafe.Sizeof(head))
	head.BiWidth = int32(width)
	head.BiHeight = int32(height * -1)
	head.BiBitCount = 32
	head.BiPlanes = 1
	head.BiCompression = win.BI_RGB
	var lpBits unsafe.Pointer
	bmp := win.CreateDIBSection(hdc, &head.BITMAPINFOHEADER, win.DIB_RGB_COLORS, &lpBits, 0, 0)
	bits := (*[1 << 30]byte)(lpBits)
	stride := width * 4
	for y := 0; y < height; y++ {
		for x := 0; x < width*4; x++ {
			sp := image.Stride*(y) + x
			dp := stride*y + x
			bits[dp] = image.Pix[sp]
		}
	}
	oldBits := win.SelectObject(memDc, win.HGDIOBJ(bmp))
	if bmp != 0 {
		defer func() {
			win.SelectObject(memDc, oldBits)
			win.DeleteObject(win.HGDIOBJ(bmp))
		}()
	}
	blend := win.BLENDFUNCTION{
		SourceConstantAlpha: 255,
		AlphaFormat:         win.AC_SRC_ALPHA,
	}
	win.UpdateLayeredWindow(hWnd, 0, &dst, &size, memDc, &src, 0, &blend, 2)
}

func (_this *MiniblinkForm) setFormFn(frame FrameContext) {
	js := `
			var fnMax=window[%q];
			var fnMin=window[%q];
			var fnClose=window[%q];
			var fnDrop=window[%q];
			window.mbFormDrop=function(){
				document.getElementsByTagName("body")[0].addEventListener("mousedown",
					function (e) {
						var obj = e.target || e.srcElement;
						if ({ "INPUT": 1, "SELECT": 1 }[obj.tagName.toUpperCase()])
							return;
					
						while (obj) {
							for (var i = 0; i < obj.classList.length; i++) {
								if (obj.classList[i] === "mbform-nodrag")
									return;
								if (obj.classList[i] === "mbform-drag") {
									fnDrop();
									return;
								}
							}
							obj = obj.parentElement;
						}
					});
			};
			window.mbFormMax=function(obj){
				if(fnMax){
					var els = obj.getElementsByClassName("mbform-max");
					for (var i = 0; i < els.length; i++) {
						els[i].removeEventListener("click");
						els[i].addEventListener("click", function(){fnMax()});
					}
				}
			};
			window.mbFormMin=function(obj){
				if(fnMin){
					var els = obj.getElementsByClassName("mbform-min");
					for (var i = 0; i < els.length; i++) {
						els[i].removeEventListener("click");
						els[i].addEventListener("click", function(){fnMin()});
					}
				}
			};
			window.mbFormClose=function(obj){
				if(fnClose){
					var els = obj.getElementsByClassName("mbform-close");
					for (var i = 0; i < els.length; i++) {
						els[i].removeEventListener("click");
						els[i].addEventListener("click", function(){fnClose()});
					}
				}
			};
			window.setFormButton=function(){
				window.mbFormMax(document);
				window.mbFormMin(document);
				window.mbFormClose(document);
			};
	`
	js = fmt.Sprintf(js, _fnMax, _fnMin, _fnClose, _fnDrop)
	frame.RunJs(js)
}

func (_this *MiniblinkForm) fnDrop() {
	if _this.GetState() != fm.FormState_Normal ||
		cs.App.MouseIsDown()[fm.MouseButtons_Left] == false {
		return
	}
	win.SendMessage(win.HWND(_this.GetHandle()), win.WM_SYSCOMMAND, win.SC_MOVE|win.HTCAPTION, 0)
	//me := _this.View.GetMouseEnable()
	//srcMs := cs.App.MouseLocation()
	//srcFrm := _this.GetLocation()
	//if me {
	//	_this.View.SetMouseEnable(false)
	//}
	//_this._isDrop = true
	//_this.watchMouseMove(func(p fm.Point) {
	//	var nx = p.X - srcMs.X
	//	var ny = p.Y - srcMs.Y
	//	nx = srcFrm.X + nx
	//	ny = srcFrm.Y + ny
	//	_this.SetLocation(nx, ny)
	//}, func() {
	//	_this.View.SetMouseEnable(me)
	//	_this.View.SetCursor(fm.CursorType_Default)
	//	_this._isDrop = false
	//})
	//_this.View.SetCursor(fm.CursorType_SIZEALL)
}

func (_this *MiniblinkForm) watchMouseMove(onMove func(p fm.Point), onEnd func()) {
	go func(mv func(p fm.Point), end func()) {
		pre := cs.App.MouseLocation()
		for cs.App.MouseIsDown()[fm.MouseButtons_Left] {
			p := cs.App.MouseLocation()
			if pre.IsEqual(p) == false {
				_this.Invoke(func(state interface{}) {
					mv(p)
				}, nil)
				pre = p
			}
			time.Sleep(time.Millisecond * 10)
		}
		_this.Invoke(func(_ interface{}) {
			onEnd()
		}, nil)
	}(onMove, onEnd)
}
