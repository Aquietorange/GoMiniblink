package free

import (
	"image"
	"image/draw"
	mb "qq.2564874169/goMiniblink"
	plat "qq.2564874169/goMiniblink/platform"
	core "qq.2564874169/goMiniblink/platform/miniblink"
	"unsafe"
)

type Core struct {
	owner plat.IWindow
	wke   wkeHandle

	onPaint core.PaintCallback
}

func (_this *Core) Init(window plat.IWindow) *Core {
	_this.owner = window
	_this.wke = wkeCreateWebView()
	if _this.wke == 0 {
		panic("创建失败")
	}
	wkeSetHandle(_this.wke, _this.owner.GetHandle())
	wkeOnPaintBitUpdated(_this.wke, _this.onPaintBitUpdated, 0)
	//wkeNetOnResponse(_this.wke, _this.onNetResponse, 0)
	//wkeOnLoadUrlBegin(_this.wke, _this.onLoadUrlBegin, 0)
	return _this
}

func (_this *Core) onLoadUrlBegin(wke wkeHandle, param, utf8Url uintptr, job wkeNetJob) uintptr {
	url := utf8PtrToString(utf8Url)
	println("begin", url)
	//println("begin")
	return 0
}

func (_this *Core) onNetResponse(wke wkeHandle, param, utf8Url uintptr, job wkeNetJob) uintptr {
	//url := *(*[]rune)(unsafe.Pointer(utf8Url))
	//println("resp", url)
	println("resp")
	return 0
}

func (_this *Core) GetImage(bound mb.Bound) *image.RGBA {
	bmp := image.NewRGBA(image.Rect(0, 0, bound.Width, bound.Height))
	wkePaint(_this.wke, bmp.Pix, 0)
	sub := image.NewRGBA(image.Rect(0, 0, bound.Width, bound.Height))
	draw.Draw(sub, image.Rect(0, 0, bound.Width, bound.Height), bmp, image.Pt(bound.X, bound.Y), draw.Src)
	return sub
}

func (_this *Core) onPaintBitUpdated(wke wkeHandle, param, bits uintptr, rect *wkeRect, width, height int32) uintptr {
	if width == 0 || height == 0 {
		return 0
	}
	w, h := int(rect.w), int(rect.h)
	e := core.PaintUpdateArgs{
		Wke: uintptr(wke),
		Clip: mb.Bound{
			Point: mb.Point{
				X: int(rect.x),
				Y: int(rect.y),
			},
			Rect: mb.Rect{
				Width:  w,
				Height: h,
			},
		},
		Size: mb.Rect{
			Width:  int(width),
			Height: int(height),
		},
		Param: param,
	}
	bmp := image.NewRGBA(image.Rect(0, 0, w, h))
	stride := e.Size.Width * 4
	pixs := (*[1 << 30]byte)(unsafe.Pointer(bits))
	for y := 0; y < h; y++ {
		for x := 0; x < w*4; x++ {
			sp := bmp.Stride*y + x
			dp := stride*(e.Clip.Y+y) + e.Clip.X*4 + x
			bmp.Pix[sp] = pixs[dp]
		}
	}
	e.Image = bmp
	_this.onPaint(e)
	return 0
}

func (_this *Core) Resize(width, height int) {
	wkeResize(_this.wke, uint32(width), uint32(height))
}

func (_this *Core) SetOnPaint(callback core.PaintCallback) {
	_this.onPaint = callback
}

func (_this *Core) LoadUri(uri string) {
	wkeLoadURL(_this.wke, uri)
}
