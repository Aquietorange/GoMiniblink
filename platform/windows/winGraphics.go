package windows

import (
	"image"
	"qq.2564874169/goMiniblink/platform/windows/win32"
	"unsafe"
)

type winGraphics struct {
	onClose func(g *winGraphics)
	dc      win32.HDC
	bmp     win32.HBITMAP
	width   int
	height  int
	stride  int
	pix     uintptr
}

func (_this *winGraphics) init(hdc win32.HDC, width, height int) *winGraphics {
	_this.stride = width * 4
	_this.width = width
	_this.height = height
	_this.dc = hdc
	var head win32.BITMAPV5HEADER
	head.BiSize = uint32(unsafe.Sizeof(head))
	head.BiWidth = int32(width)
	head.BiHeight = int32(height * -1)
	head.BiBitCount = 32
	head.BiPlanes = 1
	head.BiCompression = win32.BI_RGB

	var lpBits unsafe.Pointer
	_this.bmp = win32.CreateDIBSection(_this.dc, &head.BITMAPINFOHEADER, win32.DIB_RGB_COLORS, &lpBits, 0, 0)
	_this.pix = uintptr(lpBits)
	if _this.bmp == 0 {
		panic("创建失败")
	}
	return _this
}

func (_this *winGraphics) Close() {
	if _this.bmp == 0 {
		return
	}
	if _this.onClose != nil {
		_this.onClose(_this)
		_this.onClose = nil
	}
	memDc := win32.CreateCompatibleDC(_this.dc)
	win32.SelectObject(memDc, win32.HGDIOBJ(_this.bmp))
	win32.BitBlt(_this.dc, 0, 0, int32(_this.width), int32(_this.height), memDc, 0, 0, win32.SRCCOPY)
	win32.DeleteDC(memDc)
	win32.DeleteObject(win32.HGDIOBJ(_this.bmp))
	_this.bmp = 0
}

func (_this *winGraphics) DrawImage(src *image.RGBA, xSrc, ySrc, width, height, xDst, yDst int) {
	bits := (*[1 << 30]byte)(unsafe.Pointer(_this.pix))
	for y := 0; y < height; y++ {
		for x := 0; x < width*4; x++ {
			sp := src.Stride*(ySrc+y) + xSrc*4 + x
			dp := _this.stride*(yDst+y) + xDst*4 + x
			bits[dp] = src.Pix[sp]
		}
	}
}
