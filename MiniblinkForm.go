package GoMiniblink

import (
	"fmt"
	fm "gitee.com/aochulai/GoMiniblink/forms"
	br "gitee.com/aochulai/GoMiniblink/forms/bridge"
	cs "gitee.com/aochulai/GoMiniblink/forms/controls"
	gw "gitee.com/aochulai/GoMiniblink/forms/windows/win32"
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
	return _this.InitEx(br.FormParam{})
}

func (_this *MiniblinkForm) InitEx(param br.FormParam) *MiniblinkForm {
	_this.Form.InitEx(param)
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
			hWnd := gw.HWND(_this.GetHandle())
			style := gw.GetWindowLong(hWnd, gw.GWL_EXSTYLE)
			if style&gw.WS_EX_LAYERED != gw.WS_EX_LAYERED {
				gw.SetWindowLong(hWnd, gw.GWL_EXSTYLE, style|gw.WS_EX_LAYERED)
			}
			b := _this.GetBound()
			_this.transparentPaint(b.Width, b.Height)
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
	_this.View.OnPaintUpdated = func(e PaintUpdatedEvArgs) {
		_this.transparentPaint(e.Bound().Width, e.Bound().Height)
		e.Cancel()
	}
	mbApi.wkeSetTransparent(_this.wke, true)
}

func (_this *MiniblinkForm) transparentPaint(width, height int) {
	bn := _this.GetBound()
	hWnd := gw.HWND(_this.GetHandle())
	mdc := gw.HDC(mbApi.wkeGetViewDC(_this.View.core.GetHandle()))
	src := gw.POINT{}
	dst := gw.POINT{
		X: int32(bn.X),
		Y: int32(bn.Y),
	}
	size := gw.SIZE{
		CX: int32(width),
		CY: int32(height),
	}
	blend := gw.BLENDFUNCTION{
		SourceConstantAlpha: 255,
		AlphaFormat:         gw.AC_SRC_ALPHA,
	}
	gw.UpdateLayeredWindow(hWnd, 0, &dst, &size, mdc, &src, 0, &blend, 2)
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
