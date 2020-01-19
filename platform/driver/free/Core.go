package free

import (
	"image"
	"qq.2564874169/miniblink"
	"qq.2564874169/miniblink/platform"
	"qq.2564874169/miniblink/platform/driver"
	"unsafe"
)

type Core struct {
	owner platform.IWindow
	wke   wkeHandle

	onPaint driver.PaintCallback
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

func (_this *Core) GetView(bound miniblink.Bound) image.Image {
	w := int(wkeGetWidth(_this.wke))
	h := int(wkeGetHeight(_this.wke))
	bmp := image.NewRGBA(image.Rect(0, 0, w, h))
	wkePaint(_this.wke, unsafe.Pointer(&bmp.Pix[0]), 0)
	return bmp.SubImage(image.Rect(bound.X, bound.Y, bound.Wdith, bound.Height))
}

func (_this *Core) onPaintBitUpdated(wke wkeHandle, param, bits uintptr, rect *wkeRect, width, height int32) uintptr {
	if width == 0 || height == 0 {
		return 0
	}
	w, h := int(rect.w), int(rect.h)
	size := unsafe.Sizeof(uint32(1))
	bmp := image.NewRGBA(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			rgba := *((*uint32)(unsafe.Pointer(bits)))
			bits += size
			pix := miniblink.IntToRGBA(int(rgba))
			bmp.SetRGBA(x, y, pix)
		}
	}
	args := driver.PaintArgs{
		Wke: uintptr(wke),
		Clip: miniblink.Bound{
			Point: miniblink.Point{
				X: int(rect.x),
				Y: int(rect.y),
			},
			Rect: miniblink.Rect{
				Wdith:  int(rect.w),
				Height: int(rect.h),
			},
		},
		Size: miniblink.Rect{
			Wdith:  w,
			Height: h,
		},
		Image: bmp,
		Param: param,
	}
	_this.onPaint(args)
	return 0
}

func (_this *Core) Resize(width, height int) {
	wkeResize(_this.wke, uint32(width), uint32(height))
}

func (_this *Core) SetOnPaint(callback driver.PaintCallback) {
	_this.onPaint = callback
}

func (_this *Core) LoadUri(uri string) {
	wkeLoadURL(_this.wke, uri)
	println(uri)
}
