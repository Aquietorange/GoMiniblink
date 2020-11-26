package windows

import (
	fm "GoMiniblink/forms"
	win "GoMiniblink/forms/windows/win32"
	"image"
	"unsafe"
)

type winGraphics struct {
	onClose func(g *winGraphics)
	dc      win.HDC
}

func (_this *winGraphics) init(hdc win.HDC) *winGraphics {
	_this.dc = hdc
	return _this
}

func (_this *winGraphics) Close() {
	if _this.onClose != nil {
		_this.onClose(_this)
		_this.onClose = nil
	}
}

func (_this *winGraphics) DrawImage(src *image.RGBA, xSrc, ySrc, width, height, xDst, yDst int) fm.Graphics {
	var head win.BITMAPV5HEADER
	head.BiSize = uint32(unsafe.Sizeof(head))
	head.BiWidth = int32(width)
	head.BiHeight = int32(height * -1)
	head.BiBitCount = 32
	head.BiPlanes = 1
	head.BiCompression = win.BI_RGB

	var lpBits unsafe.Pointer
	bmp := win.CreateDIBSection(_this.dc, &head.BITMAPINFOHEADER, win.DIB_RGB_COLORS, &lpBits, 0, 0)
	bits := (*[1 << 30]byte)(lpBits)
	stride := width * 4
	for y := 0; y < height; y++ {
		for x := 0; x < width*4; x++ {
			sp := src.Stride*(ySrc+y) + xSrc*4 + x
			dp := stride*y + x
			bits[dp] = src.Pix[sp]
		}
	}
	memDc := win.CreateCompatibleDC(_this.dc)
	win.SelectObject(memDc, win.HGDIOBJ(bmp))
	win.BitBlt(_this.dc, int32(xDst), int32(yDst), int32(width), int32(height), memDc, 0, 0, win.SRCCOPY)
	win.DeleteDC(memDc)
	win.DeleteObject(win.HGDIOBJ(bmp))
	return _this
}
