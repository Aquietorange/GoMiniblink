package GoMiniblink

import (
	"fmt"
	"image"
	f "qq2564874169/goMiniblink/forms"
	c "qq2564874169/goMiniblink/forms/controls"
	"qq2564874169/goMiniblink/forms/platform/windows/win32"
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
	c.Form
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
	_this.View.SetAnchor(f.AnchorStyle_Top | f.AnchorStyle_Right | f.AnchorStyle_Bottom | f.AnchorStyle_Left)
	_this.AddChild(_this.View)

	bakOnResize := _this.OnResize
	_this.OnResize = func(e f.Rect) {
		_this.View.SetSize(e.Width, e.Height)
		bakOnResize(e)
	}
	bakOnLoad := _this.OnLoad
	_this.OnLoad = func() {
		_this._wke = wkeHandle(_this.View.GetMiniblinkHandle())
		_this.View.OnFocus()
		_this.View.JsFuncEx(_fnMax, func() {
			if _this.GetState() == f.FormState_Max {
				_this.SetState(f.FormState_Normal)
			} else {
				_this.SetState(f.FormState_Max)
			}
		})
		_this.View.JsFuncEx(_fnMin, func() {
			_this.SetState(f.FormState_Min)
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
	hWnd := win32.HWND(_this.GetHandle())
	onMove := _this.View.OnMouseMove
	onDown := _this.View.OnMouseDown
	_this.View.OnMouseDown = func(e *f.MouseEvArgs) {
		if e.Button&f.MouseButtons_Left != f.MouseButtons_Left {
			onDown(e)
			return
		}
		switch _this._rsState {
		case 8:
			win32.SendMessage(hWnd, win32.WM_SYSCOMMAND, uintptr(win32.SC_SIZE|0xF003), 0)
		case 2:
			win32.SendMessage(hWnd, win32.WM_SYSCOMMAND, uintptr(win32.SC_SIZE|0xF006), 0)
		case 4:
			win32.SendMessage(hWnd, win32.WM_SYSCOMMAND, uintptr(win32.SC_SIZE|0xF001), 0)
		case 6:
			win32.SendMessage(hWnd, win32.WM_SYSCOMMAND, uintptr(win32.SC_SIZE|0xF002), 0)
		case 7:
			win32.SendMessage(hWnd, win32.WM_SYSCOMMAND, uintptr(win32.SC_SIZE|0xF004), 0)
		case 9:
			win32.SendMessage(hWnd, win32.WM_SYSCOMMAND, uintptr(win32.SC_SIZE|0xF005), 0)
		case 1:
			win32.SendMessage(hWnd, win32.WM_SYSCOMMAND, uintptr(win32.SC_SIZE|0xF007), 0)
		case 3:
			win32.SendMessage(hWnd, win32.WM_SYSCOMMAND, uintptr(win32.SC_SIZE|0xF008), 0)
		default:
			onDown(e)
		}
		_this._rsState = 0
	}
	_this.View.OnMouseMove = func(e *f.MouseEvArgs) {
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
			_this.SetCursor(f.CursorType_SIZENS)
		case 4, 6:
			_this.SetCursor(f.CursorType_SIZEWE)
		case 7, 3:
			_this.SetCursor(f.CursorType_SIZENWSE)
		case 9, 1:
			_this.SetCursor(f.CursorType_SIZENESW)
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
	_this.SetBorderStyle(f.FormBorder_None)
	hWnd := win32.HWND(_this.GetHandle())
	style := win32.GetWindowLong(hWnd, win32.GWL_EXSTYLE)
	if style&win32.WS_EX_LAYERED != win32.WS_EX_LAYERED {
		win32.SetWindowLong(hWnd, win32.GWL_EXSTYLE, style|win32.WS_EX_LAYERED)
	}
	mbApi.wkeSetTransparent(_this._wke, true)
	_this.View.OnPaintUpdated = func(e PaintUpdatedEvArgs) {
		_this.transparentPaint(e.Bitmap(), e.Bound().Width, e.Bound().Height)
		e.Cancel()
	}
}

func (_this *MiniblinkForm) transparentPaint(image *image.RGBA, width, height int) {
	hWnd := win32.HWND(_this.GetHandle())
	hdc := win32.GetDC(hWnd)
	memDc := win32.CreateCompatibleDC(0)
	src := win32.POINT{}
	dst := win32.POINT{
		X: int32(_this.GetLocation().X),
		Y: int32(_this.GetLocation().Y),
	}
	size := win32.SIZE{
		CX: int32(width),
		CY: int32(height),
	}
	var head win32.BITMAPV5HEADER
	head.BiSize = uint32(unsafe.Sizeof(head))
	head.BiWidth = int32(width)
	head.BiHeight = int32(height * -1)
	head.BiBitCount = 32
	head.BiPlanes = 1
	head.BiCompression = win32.BI_RGB
	var lpBits unsafe.Pointer
	bmp := win32.CreateDIBSection(hdc, &head.BITMAPINFOHEADER, win32.DIB_RGB_COLORS, &lpBits, 0, 0)
	bits := (*[1 << 30]byte)(lpBits)
	stride := width * 4
	for y := 0; y < height; y++ {
		for x := 0; x < width*4; x++ {
			sp := image.Stride*(y) + x
			dp := stride*y + x
			bits[dp] = image.Pix[sp]
		}
	}
	oldBits := win32.SelectObject(memDc, win32.HGDIOBJ(bmp))
	if bmp != 0 {
		defer func() {
			win32.SelectObject(memDc, oldBits)
			win32.DeleteObject(win32.HGDIOBJ(bmp))
		}()
	}
	blend := win32.BLENDFUNCTION{
		SourceConstantAlpha: 255,
		AlphaFormat:         win32.AC_SRC_ALPHA,
	}
	win32.UpdateLayeredWindow(hWnd, 0, &dst, &size, memDc, &src, 0, &blend, 2)
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
	if _this.GetState() == f.FormState_Max ||
		c.App.MouseIsDown()[f.MouseButtons_Left] == false {
		return
	}
	//win32.SendMessage(win32.HWND(_this.GetHandle()), win32.WM_SYSCOMMAND, win32.SC_MOVE|win32.HTCAPTION, 0)
	me := _this.View.GetMouseEnable()
	srcMs := c.App.MouseLocation()
	srcFrm := _this.GetLocation()
	if me {
		_this.View.SetMouseEnable(false)
	}
	_this._isDrop = true
	_this.watchMouseMove(func(p f.Point) {
		var nx = p.X - srcMs.X
		var ny = p.Y - srcMs.Y
		nx = srcFrm.X + nx
		ny = srcFrm.Y + ny
		_this.SetLocation(nx, ny)
	}, func() {
		_this.View.SetMouseEnable(me)
		_this.View.SetCursor(f.CursorType_Default)
		_this._isDrop = false
	})
	_this.View.SetCursor(f.CursorType_SIZEALL)
}

func (_this *MiniblinkForm) watchMouseMove(onMove func(p f.Point), onEnd func()) {
	go func(mv func(p f.Point), end func()) {
		pre := c.App.MouseLocation()
		for c.App.MouseIsDown()[f.MouseButtons_Left] {
			p := c.App.MouseLocation()
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
