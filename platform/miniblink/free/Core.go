package free

import (
	"image"
	"image/draw"
	mb "qq.2564874169/goMiniblink"
	"qq.2564874169/goMiniblink/platform"
	"qq.2564874169/goMiniblink/platform/miniblink"
	"unsafe"
)

type Core struct {
	owner platform.IWindow
	wke   wkeHandle

	onPaint miniblink.PaintCallback
}

func (_this *Core) Init(window platform.IWindow) *Core {
	_this.owner = window
	_this.wke = wkeCreateWebView()
	if _this.wke == 0 {
		panic("创建失败")
	}
	wkeSetHandle(_this.wke, _this.owner.GetHandle())
	//wkeOnPaintBitUpdated(_this.wke, _this.onPaintBitUpdated, 0)
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

func (_this *Core) FillImage(bmp *image.RGBA) {
	w := wkeGetWidth(_this.wke)
	h := wkeGetHeight(_this.wke)
	img := image.NewRGBA(image.Rect(0, 0, int(w), int(h)))
	wkePaint(_this.wke, img.Pix, 0)
	draw.Draw(bmp, bmp.Bounds(), img, image.Pt(0, 0), draw.Src)
}

func (_this *Core) GetImage(bound mb.Bound) *image.RGBA {
	bmp := image.NewRGBA(image.Rect(0, 0, bound.Width, bound.Height))
	//wkePaint2(_this.wke, bmp.Pix,
	//	uint32(bmp.Bounds().Dx()), uint32(bmp.Bounds().Dy()), 0, 0,
	//	uint32(bmp.Bounds().Dx()), uint32(bmp.Bounds().Dy()),
	//	uint32(bound.X), uint32(bound.Y), true)
	return bmp
}

func (_this *Core) onPaintBitUpdated(wke wkeHandle, param, bits uintptr, rect *wkeRect, width, height int32) uintptr {
	if width == 0 || height == 0 {
		return 0
	}
	w, h := int(rect.w), int(rect.h)
	size := unsafe.Sizeof(uint32(1))
	bmp := image.NewRGBA(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x, bits = x+1, bits+size {
			rgba := *((*uint32)(unsafe.Pointer(bits)))
			pix := mb.IntToRGBA(int(rgba))
			bmp.SetRGBA(x, y, pix)
		}
	}
	e := miniblink.PaintUpdateArgs{
		Wke: uintptr(wke),
		Clip: mb.Bound{
			Point: mb.Point{
				X: int(rect.x),
				Y: int(rect.y),
			},
			Rect: mb.Rect{
				Width:  int(rect.w),
				Height: int(rect.h),
			},
		},
		Size: mb.Rect{
			Width:  w,
			Height: h,
		},
		Image: bmp,
		Param: param,
	}
	_this.onPaint(e)
	return 0
}

func (_this *Core) Resize(width, height int) {
	wkeResize(_this.wke, uint32(width), uint32(height))
}

func (_this *Core) SetOnPaint(callback miniblink.PaintCallback) {
	_this.onPaint = callback
}

func (_this *Core) LoadUri(uri string) {
	wkeLoadURL(_this.wke, uri)
}
