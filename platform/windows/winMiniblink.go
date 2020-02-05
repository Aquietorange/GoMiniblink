package windows

import (
	mb "qq.2564874169/miniblink"
	"qq.2564874169/miniblink/platform"
	core "qq.2564874169/miniblink/platform/miniblink_core"
	"qq.2564874169/miniblink/platform/miniblink_core/free"
	"qq.2564874169/miniblink/platform/miniblink_core/vip"
	"qq.2564874169/miniblink/platform/windows/win32"
)

type winMiniblink struct {
	winControl

	wke      core.ICore
	initUri  string
	initSize mb.Rect
}

func (_this *winMiniblink) init(provider *Provider) *winMiniblink {
	_this.winControl.init(provider)
	var baseCreateProc platform.WindowCreateProc
	baseCreateProc = _this.SetOnCreate(func(handle uintptr) {
		if baseCreateProc != nil {
			baseCreateProc(handle)
		}
		_this.initWke(win32.HWND(handle))
	})
	return _this
}

func (_this *winMiniblink) initWke(hWnd win32.HWND) {
	if vip.Exists() {
		//todo
	} else {
		_this.wke = new(free.Core).Init(_this)
	}
	_this.onPaint = _this.defOnPaint
	_this.wke.SetOnPaint(_this.paintUpdate)
	if _this.initSize.Wdith > 0 && _this.initSize.Height > 0 {
		_this.SetSize(_this.initSize.Wdith, _this.initSize.Height)
	}
	if _this.initUri != "" {
		_this.LoadUri(_this.initUri)
	}
}

func (_this *winMiniblink) defOnPaint(e mb.PaintEvArgs) {
	bmp := _this.wke.GetView(e.Clip)
	e.Graphics.DrawImage(bmp, mb.Point{X: 0, Y: 0}, e.Clip.Rect, e.Clip.Point)
}

func (_this *winMiniblink) paintUpdate(args core.PaintUpdateArgs) {

}

func (_this *winMiniblink) LoadUri(uri string) {
	if _this.IsCreate() {
		_this.wke.LoadUri(uri)
	} else {
		_this.initUri = uri
	}
}

func (_this *winMiniblink) SetSize(width, height int) {
	if _this.IsCreate() {
		_this.wke.Resize(width, height)
		_this.winControl.SetSize(width, height)
	} else {
		_this.initSize = mb.Rect{
			Wdith:  width,
			Height: height,
		}
	}
}
