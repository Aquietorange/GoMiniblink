package windows

import (
	"image"
	mb "qq.2564874169/miniblink"
	"qq.2564874169/miniblink/platform/windows/win32"
	"unsafe"
)

type winGraphics struct {
	hdc     win32.HDC
	memBmp  win32.HBITMAP
	onClose func()
}

func (_this *winGraphics) init(hdc win32.HDC) *winGraphics {
	_this.hdc = hdc
	return _this
}

func (_this *winGraphics) Close() {
	if _this.onClose != nil {
		_this.onClose()
	}
}

func (_this *winGraphics) DrawImage(src *image.RGBA, srcXY mb.Point, rect mb.Rect, toXY mb.Point) {
	bHead := win32.BITMAPINFOHEADER{
		BiWidth:       int32(src.Bounds().Dx()),
		BiHeight:      int32(src.Bounds().Dy() * -1),
		BiPlanes:      1,
		BiBitCount:    32,
		BiCompression: win32.BI_RGB,
	}
	bHead.BiSize = uint32(unsafe.Sizeof(bHead))
	bInfo := win32.BITMAPINFO{
		BmiHeader: bHead,
		BmiColors: nil,
	}
	isDel := _this.memBmp == 0
	if isDel {
		_this.memBmp = win32.CreateCompatibleBitmap(_this.hdc, int32(rect.Wdith), int32(rect.Height))
	}
	win32.SetDIBits(0, _this.memBmp, 0, uint32(rect.Height), &src.Pix[0], &bInfo, 0)
	memDc := win32.CreateCompatibleDC(_this.hdc)
	old := win32.SelectObject(memDc, win32.HGDIOBJ(_this.memBmp))
	win32.BitBlt(_this.hdc, int32(toXY.X), int32(toXY.Y), int32(rect.Wdith), int32(rect.Height), memDc, int32(srcXY.X), int32(srcXY.Y), win32.SRCCOPY|win32.CAPTUREBLT)
	win32.SelectObject(memDc, old)
	win32.DeleteDC(memDc)
	if isDel {
		win32.DeleteObject(win32.HGDIOBJ(_this.memBmp))
	}
}
