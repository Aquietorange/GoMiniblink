package free

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"qq.2564874169/miniblink/platform"
	"qq.2564874169/miniblink/platform/miniblinkc-core"
	"unsafe"
)

type Core struct {
	view platform.IWindow
	wke  wkeHandle

	onPaint miniblinkc_core.PaintCallback
}

func (_this *Core) Init(window platform.IWindow) *Core {
	if wkeInitialize() == false {
		panic("初始化失败")
	}
	_this.wke = wkeCreateWebView()
	if _this.wke == 0 {
		panic("创建失败")
	}
	_this.view = window
	wkeSetHandle(_this.wke, _this.view.GetHandle())
	wkeOnPaintBitUpdated(_this.wke, _this.onPaintBitUpdated, 0)
	wkeResize(_this.wke, 500, 500)
	return _this
}

func (_this *Core) onPaintBitUpdated(wke wkeHandle, param, buf uintptr, rect *wkeRect, width, height int32) uintptr {

	if width == 0 || height == 0 {
		return 0
	}
	w, h := int(width), int(height)
	size := unsafe.Sizeof(uint32(1))
	dataLen := w * h
	data := make([]uint32, dataLen)
	bmp := image.NewRGBA(image.Rect(0, 0, int(width), int(height)))
	for i := 0; i < dataLen; i++ {
		rgba := *((*uint32)(unsafe.Pointer(buf)))
		data[i] = rgba
		buf += size
	}
	for y := 0; y < int(height); y++ {
		for x := 0; x < int(width); x++ {
			rgba := data[y*h+x]
			buf += size
			bmp.SetRGBA(x, y, color.RGBA{
				R: uint8(rgba << 24),
				G: uint8(rgba << 16),
				B: uint8(rgba << 8),
				A: uint8(rgba),
			})
		}
	}
	//file := make([]uint8, dataLen*int(unsafe.Sizeof(uint8(1))))
	//buffer := bytes.NewBuffer(file)
	//if err := binary.Write(buffer, binary.LittleEndian, data); err != nil {
	//	fmt.Println(err)
	//}
	//bmp.Pix = file
	img := bmp.SubImage(bmp.Rect)
	//if err != nil {
	//	fmt.Println(err)
	//	return 0
	//}
	if f, err := os.OpenFile("show.png", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600); err == nil {
		defer f.Close()
		_ = png.Encode(f, img)
	} else {
		fmt.Println(err)
	}

	//args := miniblinkc-core.PaintArgs{
	//	Wke: uintptr(wke),
	//	Clip: GoMiniblink.Bound{
	//		Point: GoMiniblink.Point{
	//			X: rect.x,
	//			Y: rect.y,
	//		},
	//		Rect: GoMiniblink.Rect{
	//			Wdith:  rect.w,
	//			Height: rect.h,
	//		},
	//	},
	//	Size: GoMiniblink.Rect{
	//		Wdith:  width,
	//		Height: height,
	//	},
	//	Bits:  data,
	//	Param: param,
	//}
	//_this.onPaint(args)
	return 0
}

func (_this *Core) Resize(width, height int) {
	wkeResize(_this.wke, int32(width), int32(height))
}

func (_this *Core) SetOnPaint(callback miniblinkc_core.PaintCallback) {
	_this.onPaint = callback
}

func (_this *Core) LoadUri(uri string) {
	wkeLoadURL(_this.wke, uri)
}
