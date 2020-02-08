package windows

import (
	"image"
	"qq.2564874169/goMiniblink/platform/windows/win32"
	"unsafe"
)

type winGraphics struct {
	dc      win32.HDC
	onClose func()
}

func (_this *winGraphics) init(hdc win32.HDC) *winGraphics {
	_this.dc = hdc
	return _this
}

func (_this *winGraphics) Close() {
	if _this.onClose != nil {
		_this.onClose()
	}
}

func (_this *winGraphics) DrawImage(src *image.RGBA, xSrc, ySrc, width, height, xDst, yDst int) {
	//bHead := win32.BITMAPINFOHEADER{
	//	BiWidth:       int32(src.Bounds().Dx()),
	//	BiHeight:      int32(src.Bounds().Dy() * -1),
	//	BiPlanes:      1,
	//	BiBitCount:    32,
	//	BiCompression: win32.BI_RGB,
	//}
	//bHead.BiSize = uint32(unsafe.Sizeof(bHead))
	//bInfo := win32.BITMAPINFO{
	//	BmiHeader: bHead,
	//}
	//bmp := win32.CreateCompatibleBitmap(_this.dc, int32(src.Bounds().Dx()), int32(src.Bounds().Dy()))
	//ret := win32.SetDIBits(0, bmp, 0, uint32(src.Bounds().Dy()), &src.Pix[0], &bInfo, win32.DIB_RGB_COLORS)
	//memDc := win32.CreateCompatibleDC(_this.dc)
	//old := win32.SelectObject(memDc, win32.HGDIOBJ(bmp))
	//if ret != 0 {
	//	win32.BitBlt(_this.dc, int32(toXY.X), int32(toXY.Y), int32(rect.Width), int32(rect.Height), memDc, int32(srcXY.X), int32(srcXY.Y), win32.SRCCOPY)
	//} else {
	//	println(src.Bounds().Dx(), src.Bounds().Dy(), len(src.Pix), cap(src.Pix), src.Stride, src.Stride%4)
	//}
	//win32.SelectObject(memDc, old)
	//win32.DeleteDC(memDc)
	//win32.DeleteObject(win32.HGDIOBJ(bmp))

	if src.Bounds().Dx()%4 != 0 {
		panic("图像宽度必须是4的倍数")
	}
	bHead := win32.BITMAPV5HEADER{
		BITMAPV4HEADER: win32.BITMAPV4HEADER{
			BITMAPINFOHEADER: win32.BITMAPINFOHEADER{
				BiWidth:         int32(src.Bounds().Dx()),
				BiHeight:        int32(src.Bounds().Dy() * -1),
				BiPlanes:        1,
				BiBitCount:      32,
				BiXPelsPerMeter: 3780,
				BiYPelsPerMeter: 3780,
				BiCompression:   win32.BI_RGB,
			},
		},
	}
	bHead.BiSize = uint32(unsafe.Sizeof(bHead))
	bInfo := win32.BITMAPINFO{
		BmiHeader: bHead.BITMAPINFOHEADER,
	}
	bmp := win32.CreateDIBitmap(_this.dc, &bHead, win32.CBM_INIT, &src.Pix[0], &bInfo, win32.DIB_RGB_COLORS)
	if bmp != 0 {
		memDc := win32.CreateCompatibleDC(_this.dc)
		old := win32.SelectObject(memDc, win32.HGDIOBJ(bmp))
		win32.BitBlt(_this.dc, int32(xDst), int32(yDst), int32(width), int32(height), memDc, int32(xSrc), int32(ySrc), win32.SRCCOPY)
		win32.SelectObject(memDc, old)
		win32.DeleteDC(memDc)
		win32.DeleteObject(win32.HGDIOBJ(bmp))
	} else {
		len1 := bHead.BiWidth * bHead.BiHeight * -1 * 4
		len2 := len(src.Pix)
		if int(len1) != len2 {
			println(len1, len2)
		}
	}
}
