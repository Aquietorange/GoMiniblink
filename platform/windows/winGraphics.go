package windows

import (
	"image"
	"qq.2564874169/goMiniblink/platform/windows/win32"
	"unsafe"
)

type winGraphics struct {
	onClose func(g *winGraphics)
	dc      win32.HDC
}

func (_this *winGraphics) init(hdc win32.HDC, width, height int) *winGraphics {
	_this.dc = hdc
	return _this
}

func (_this *winGraphics) Close() {
	if _this.onClose != nil {
		_this.onClose(_this)
	}
}

func (_this *winGraphics) DrawImage(src *image.RGBA, xSrc, ySrc, width, height, xDst, yDst int) {
	var head win32.BITMAPV5HEADER
	head.BiSize = uint32(unsafe.Sizeof(head))
	head.BiWidth = int32(src.Bounds().Dx())
	head.BiHeight = int32(src.Bounds().Dy() * -1)
	head.BiBitCount = 32
	head.BiPlanes = 1
	head.BiCompression = win32.BI_RGB

	var lpBits unsafe.Pointer
	bmp := win32.CreateDIBSection(_this.dc, &head.BITMAPINFOHEADER, win32.DIB_RGB_COLORS, &lpBits, 0, 0)
	bits := (*[1 << 30]byte)(lpBits)
	for i := range src.Pix {
		bits[i] = src.Pix[i]
	}
	memDc := win32.CreateCompatibleDC(_this.dc)
	win32.SelectObject(memDc, win32.HGDIOBJ(bmp))
	win32.BitBlt(_this.dc, int32(xDst), int32(yDst), int32(width), int32(height), memDc, int32(xSrc), int32(ySrc), win32.SRCCOPY)
	win32.DeleteDC(memDc)
	win32.DeleteObject(win32.HGDIOBJ(bmp))
}
