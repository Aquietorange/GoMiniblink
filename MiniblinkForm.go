package GoMiniblink

import (
	"fmt"
	fm "gitee.com/aochulai/GoMiniblink/forms"
	cs "gitee.com/aochulai/GoMiniblink/forms/controls"
	gww "gitee.com/aochulai/GoMiniblink/forms/windows/win32"
	"image"
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

	wke           wkeHandle
	isTransparent bool
	resizeState   int
}

func (_this *MiniblinkForm) Init() *MiniblinkForm {
	_this.Form.Init()
	_this.View = new(MiniblinkBrowser).Init()
	_this.View.SetAnchor(fm.AnchorStyle_Fill)
	_this.AddChild(_this.View)
	_this.wke = wkeHandle(_this.View.GetMiniblinkHandle())
	_this.setOn()
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
	_this.View.EvDocumentReady["__goMiniblink_init_js"] = func(_ *MiniblinkBrowser, e DocumentReadyEvArgs) {
		e.RunJs("window.setFormButton();window.mbFormDrop();")
	}
	_this.setDrop()
	return _this
}

func (_this *MiniblinkForm) setDrop() {
	isDrop := false
	var anchor fm.Point
	var base fm.Point
	_this.View.JsFuncEx(_fnDrop, func() {
		isDrop = true
	})
	_this.View.EvMouseDown["__goMiniblink_drop"] = func(s cs.GUI, e *fm.MouseEvArgs) {
		if isDrop {
			base = _this.GetBound().Point
			anchor = fm.Point{
				X: e.ScreenX,
				Y: e.ScreenY,
			}
			_this.View.SetCursor(fm.CursorType_SIZEALL)
			_this.View.MouseEnable(false)
		}
	}
	_this.View.EvMouseUp["__goMiniblink_drop"] = func(s cs.GUI, e *fm.MouseEvArgs) {
		isDrop = false
		_this.View.SetCursor(fm.CursorType_Default)
		_this.View.MouseEnable(true)
	}
	_this.View.EvMouseMove["__goMiniblink_drop"] = func(s cs.GUI, e *fm.MouseEvArgs) {
		if isDrop {
			var nx = e.ScreenX - anchor.X
			var ny = e.ScreenY - anchor.Y
			nx = base.X + nx
			ny = base.Y + ny
			_this.SetLocation(nx, ny)
			e.IsHandle = true
		}
	}
}

func (_this *MiniblinkForm) setOn() {
	bakOnResize := _this.OnResize
	_this.OnResize = func(e fm.Rect) {
		_this.View.SetSize(e.Width, e.Height)
		bakOnResize(e)
	}
	bakOnLoad := _this.OnLoad
	_this.OnLoad = func() {
		if _this.isTransparent {
			hWnd := gww.HWND(_this.GetHandle())
			style := gww.GetWindowLong(hWnd, gww.GWL_EXSTYLE)
			if style&gww.WS_EX_LAYERED != gww.WS_EX_LAYERED {
				gww.SetWindowLong(hWnd, gww.GWL_EXSTYLE, style|gww.WS_EX_LAYERED)
			}
			mbApi.wkeSetTransparent(_this.wke, true)
			_this.View.OnPaintUpdated = func(e PaintUpdatedEvArgs) {
				_this.transparentPaint(e.Bitmap(), e.Bound().Width, e.Bound().Height)
				e.Cancel()
			}
			img := _this.View.ToBitmap()
			_this.transparentPaint(img, img.Bounds().Dx(), img.Bounds().Dy())
		}
		bakOnLoad()
	}
	bakOnJsReady := _this.View.OnJsReady
	_this.View.OnJsReady = func(e JsReadyEvArgs) {
		bakOnJsReady(e)
		_this.setFormFn(e)
	}
}

func (_this *MiniblinkForm) TransparentMode() {
	_this.isTransparent = true
	_this.SetBorderStyle(fm.FormBorder_None)
}

func (_this *MiniblinkForm) transparentPaint(image *image.RGBA, width, height int) {
	bn := _this.GetBound()
	hWnd := gww.HWND(_this.GetHandle())
	hdc := gww.GetDC(hWnd)
	memDc := gww.CreateCompatibleDC(0)
	src := gww.POINT{}
	dst := gww.POINT{
		X: int32(bn.X),
		Y: int32(bn.Y),
	}
	size := gww.SIZE{
		CX: int32(width),
		CY: int32(height),
	}
	var head gww.BITMAPV5HEADER
	head.BiSize = uint32(unsafe.Sizeof(head))
	head.BiWidth = int32(width)
	head.BiHeight = int32(height * -1)
	head.BiBitCount = 32
	head.BiPlanes = 1
	head.BiCompression = gww.BI_RGB
	var lpBits unsafe.Pointer
	bmp := gww.CreateDIBSection(hdc, &head.BITMAPINFOHEADER, gww.DIB_RGB_COLORS, &lpBits, 0, 0)
	bits := (*[1 << 30]byte)(lpBits)
	stride := width * 4
	for y := 0; y < height; y++ {
		for x := 0; x < width*4; x++ {
			sp := image.Stride*(y) + x
			dp := stride*y + x
			bits[dp] = image.Pix[sp]
		}
	}
	oldBits := gww.SelectObject(memDc, gww.HGDIOBJ(bmp))
	if bmp != 0 {
		defer func() {
			gww.SelectObject(memDc, oldBits)
			gww.DeleteObject(gww.HGDIOBJ(bmp))
		}()
	}
	blend := gww.BLENDFUNCTION{
		SourceConstantAlpha: 255,
		AlphaFormat:         gww.AC_SRC_ALPHA,
	}
	gww.UpdateLayeredWindow(hWnd, 0, &dst, &size, memDc, &src, 0, &blend, 2)
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
