package windows

import (
	"fmt"
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
		_this.wke.FireMouseClickEvent(_this.app, e.Button, true, e.IsDouble, e.X, e.Y)
		return false
	})
	_this.SetOnMouseUp(func(e mb.MouseEvArgs) bool {
		_this.wke.FireMouseClickEvent(_this.app, e.Button, false, e.IsDouble, e.X, e.Y)
		return false
	})
	_this.SetOnMouseMove(func(e mb.MouseEvArgs) bool {
		_this.wke.FireMouseMoveEvent(_this.app, e.Button, e.X, e.Y)
		return false
	})
	_this.SetOnMouseWheel(func(e mb.MouseEvArgs) bool {
		_this.wke.FireMouseWheelEvent(_this.app, e.Button, e.Delta, e.X, e.Y)
		return false
	})
	_this.SetOnKeyDown(func(e *mb.KeyEvArgs) bool {
		_this.wke.FireKeyEvent(e.Value, true, isExtendKey(e.Key), true, e.IsSys)
		return false
	})
	_this.SetOnKeyUp(func(e *mb.KeyEvArgs) bool {
		_this.wke.FireKeyEvent(e.Value, true, isExtendKey(e.Key), false, e.IsSys)
		return false
	})
	_this.SetOnKeyPress(func(e *mb.KeyPressEvArgs) bool {
		_this.wke.FireKeyPressEvent(int([]rune(e.KeyChar)[0]), true, false, e.IsSys)
		return false
	})
	_this.SetOnFocus(func() bool {
		_this.wke.SetFocus()
		return false
	})
	_this.SetOnImeStartComposition(func() bool {
		p := _this.wke.GetCaretPos()
		fmt.Println(p)
		comp := win32.COMPOSITIONFORM{
			DwStyle: win32.CFS_POINT | win32.CFS_FORCE_POSITION,
			Pos: win32.POINT{
				X: int32(10),
				Y: int32(10),
			},
		}
		imc := win32.ImmGetContext(_this.handle)
		win32.ImmSetCompositionWindow(imc, &comp)
		win32.ImmReleaseContext(_this.handle, imc)
		return true
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
	_this.onSetCursor = _this.defOnSetCursor
	_this.wke.SetOnPaint(_this.paintUpdate)
	if _this.initSize.Width > 0 && _this.initSize.Height > 0 {
		_this.SetSize(_this.initSize.Width, _this.initSize.Height)
	}
	if _this.initUri != "" {
		_this.LoadUri(_this.initUri)
	}
}

func (_this *winMiniblink) defOnSetCursor() bool {
	cur := _this.wke.GetCursor()
	_this.SetCursor(cur)
	return false
}

func (_this *winMiniblink) defOnPaint(e mb.PaintEvArgs) bool {
	bmp := _this.wke.GetImage(e.Clip)
	e.Graphics.DrawImage(bmp, 0, 0, e.Clip.Width, e.Clip.Height, e.Clip.X, e.Clip.Y)
	return true
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
