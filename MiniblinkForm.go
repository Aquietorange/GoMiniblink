package GoMiniblink

import (
	"fmt"
	f "qq2564874169/goMiniblink/forms"
	c "qq2564874169/goMiniblink/forms/controls"
	"time"
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
		_this.View.JsFuncEx(_fnClose, _this.Close)
		_this.View.JsFuncEx(_fnDrop, _this.fnDrop)
		bakEvJsReady := _this.View.OnJsReady
		_this.View.OnJsReady = func(e JsReadyEvArgs) {
			bakEvJsReady(e)
			_this.setFormFn(e)
		}
		_this.View.EvDocumentReady["__goMiniblink"] = func(_ interface{}, e DocumentReadyEvArgs) {
			e.RunJs("window.setFormButton();window.mbFormDrop();")
		}
		bakOnLoad()
	}
	return _this
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
	me := _this.View.GetMouseEnable()
	srcMs := c.App.MouseLocation()
	srcFrm := _this.GetLocation()
	if me {
		_this.View.SetMouseEnable(false)
	}
	_this.watchMouseMove(func(p f.Point) {
		var nx = p.X - srcMs.X
		var ny = p.Y - srcMs.Y
		nx = srcFrm.X + nx
		ny = srcFrm.Y + ny
		_this.SetLocation(nx, ny)
	}, func() {
		_this.View.SetMouseEnable(me)
		_this.View.SetCursor(f.CursorType_Default)
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
