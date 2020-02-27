package windows

import (
	mb "qq.2564874169/goMiniblink"
	"qq.2564874169/goMiniblink/platform"
	core "qq.2564874169/goMiniblink/platform/miniblink"
	"qq.2564874169/goMiniblink/platform/miniblink/free"
	"qq.2564874169/goMiniblink/platform/miniblink/vip"
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
		_this.initWke()
	})
	return _this
}

func (_this *winMiniblink) initWke() {
	if vip.Exists() {
		//todo
	} else {
		_this.wke = new(free.Core).Init(_this)
	}
	_this.onPaint = _this.defOnPaint
	_this.wke.SetOnPaint(_this.paintUpdate)
	if _this.initSize.Width > 0 && _this.initSize.Height > 0 {
		_this.SetSize(_this.initSize.Width, _this.initSize.Height)
	}
	if _this.initUri != "" {
		_this.LoadUri(_this.initUri)
	}
}

func (_this *winMiniblink) defOnPaint(e mb.PaintEvArgs) {
	bmp := _this.wke.GetImage(e.Clip)
	e.Graphics.DrawImage(bmp, 0, 0, e.Clip.Width, e.Clip.Height, e.Clip.X, e.Clip.Y)
}

func (_this *winMiniblink) paintUpdate(args core.PaintUpdateArgs) {
	g := _this.CreateGraphics()
	g.DrawImage(args.Image, 0, 0, args.Clip.Width, args.Clip.Height, args.Clip.X, args.Clip.Y).Close()
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
			Width:  width,
			Height: height,
		}
	}
}
