package GoMiniblink

import (
	"fmt"
	f "qq2564874169/goMiniblink/forms"
	c "qq2564874169/goMiniblink/forms/controls"
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
		_this.View.JsFuncEx(_fnClose, func() {
			_this.Close()
		})
		//bakEvJsReady := _this.View.OnJsReady
		//_this.View.OnJsReady = func(e JsReadyEvArgs) {
		//	_this.setFormFn(e)
		//	bakEvJsReady(e)
		//}
		_this.View.EvDocumentReady["_"] = func(_ interface{}, e DocumentReadyEvArgs) {
			_this.setFormFn(e)
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
			window.mbFormMax(document);
			window.mbFormMin(document);
			window.mbFormClose(document);
	`
	js = fmt.Sprintf(js, _fnMax, _fnMin, _fnClose)
	frame.RunJs(js)
}
