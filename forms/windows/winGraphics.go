package windows

import (
	"image"
	"unsafe"

	fm "github.com/hujun528/GoMiniblink/forms"
	win "github.com/hujun528/GoMiniblink/forms/windows/win32"
)

type winGraphics struct {
	onClose func(g *winGraphics)
	dc      win.HDC
}

func (_this *winGraphics) init(hdc win.HDC) *winGraphics {
	_this.dc = hdc
	return _this
}

func (_this *winGraphics) GetHandle() uintptr {
	return uintptr(_this.dc)
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
	_this.DrawByDc(memDc, 0, 0, width, height, xDst, yDst)
	win.DeleteDC(memDc)
	win.DeleteObject(win.HGDIOBJ(bmp))
	return _this
}

func (_this *winGraphics) DrawByDc(dc win.HDC, xSrc, ySrc, width, height, xDst, yDst int) *winGraphics {
	win.BitBlt(_this.dc, int32(xDst), int32(yDst), int32(width), int32(height), dc, int32(xSrc), int32(ySrc), win.SRCCOPY)
	return _this
}
