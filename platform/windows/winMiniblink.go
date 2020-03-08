package windows

import (
	mb "qq2564874169/goMiniblink"
	plat "qq2564874169/goMiniblink/platform"
	core "qq2564874169/goMiniblink/platform/miniblink"
	"qq2564874169/goMiniblink/platform/miniblink/free"
	"qq2564874169/goMiniblink/platform/miniblink/vip"
	"qq2564874169/goMiniblink/platform/windows/win32"
)

type winMiniblink struct {
	winControl

	wke      core.ICore
	initUri  string
	initSize mb.Rect
}

func (_this *winMiniblink) init(provider *Provider) *winMiniblink {
	_this.winControl.init(provider)
	var baseCreateProc plat.WindowCreateProc
	baseCreateProc = _this.SetOnCreate(func(handle uintptr) bool {
		if baseCreateProc != nil {
			baseCreateProc(handle)
		}
		_this.initWke()
		return false
	})
	_this.SetOnMouseDown(func(e mb.MouseEvArgs) bool {
		return _this.wke.FireMouseClickEvent(e.Button, true, e.IsDouble, e.X, e.Y)
	})
	_this.SetOnMouseUp(func(e mb.MouseEvArgs) bool {
		return _this.wke.FireMouseClickEvent(e.Button, false, e.IsDouble, e.X, e.Y)
	})
	_this.SetOnMouseMove(func(e mb.MouseEvArgs) bool {
		return _this.wke.FireMouseMoveEvent(e.Button, e.X, e.Y)
	})
	_this.SetOnMouseWheel(func(e mb.MouseEvArgs) bool {
		return _this.wke.FireMouseWheelEvent(e.Button, e.Delta, e.X, e.Y)
	})
	_this.SetOnKeyDown(func(e *mb.KeyEvArgs) bool {
		return _this.wke.FireKeyEvent(*e, true, e.IsSys)
	})
	_this.SetOnKeyUp(func(e *mb.KeyEvArgs) bool {
		return _this.wke.FireKeyEvent(*e, false, e.IsSys)
	})
	_this.SetOnKeyPress(func(e *mb.KeyPressEvArgs) bool {
		return _this.wke.FireKeyPressEvent(int([]rune(e.KeyChar)[0]), e.IsSys)
	})
	_this.SetOnFocus(func() bool {
		_this.wke.SetFocus()
		return false
	})
	_this.SetOnImeStartComposition(func() bool {
		p := _this.wke.GetCaretPos()
		comp := win32.COMPOSITIONFORM{
			DwStyle: win32.CFS_POINT | win32.CFS_FORCE_POSITION,
			Pos: win32.POINT{
				X: int32(p.X),
				Y: int32(p.Y),
			},
		}
		imc := win32.ImmGetContext(_this.handle)
		win32.ImmSetCompositionWindow(imc, &comp)
		win32.ImmReleaseContext(_this.handle, imc)
		return true
	})
	_this.SetOnPaint(func(e mb.PaintEvArgs) bool {
		bmp := _this.wke.GetImage(e.Clip)
		e.Graphics.DrawImage(bmp, 0, 0, e.Clip.Width, e.Clip.Height, e.Clip.X, e.Clip.Y)
		return true
	})
	_this.SetOnCursor(func() bool {
		t := _this.wke.GetCursor()
		if t != mb.CursorType_Default {
			_this.SetCursor(t)
			return true
		}
		return false
	})
	return _this
}

func (_this *winMiniblink) initWke() {
	if vip.Exists() {
		//todo
	} else {
		_this.wke = new(free.Core).Init(_this)
	}
	_this.wke.SetOnPaint(_this.paintUpdate)
	if _this.initSize.Width > 0 && _this.initSize.Height > 0 {
		_this.SetSize(_this.initSize.Width, _this.initSize.Height)
	}
	if _this.initUri != "" {
		_this.LoadUri(_this.initUri)
	}
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
