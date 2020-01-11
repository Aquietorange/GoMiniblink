package free

import (
	"GoMiniblink"
	"GoMiniblink/CrossPlatform"
	"GoMiniblink/CrossPlatform/miniblink"
	"unsafe"
)

type Core struct {
	view CrossPlatform.IWindow
	wke  wkeHandle

	paintCallback miniblink.PaintCallback
}

func (_this *Core) Init(window CrossPlatform.IWindow) *Core {
	if wkeInitialize() == false {
		panic("初始化失败")
	}
	_this.wke = wkeCreateWebView()
	if _this.wke == 0 {
		panic("创建失败")
	}
	_this.view = window
	wkeSetHandle(_this.wke, _this.view.GetHandle())
	wkeOnPaintBitUpdated(_this.wke, _this.onPaint, 0)
	return _this
}

func (_this *Core) onPaint(wke wkeHandle, param, bufPtr uintptr, rect wkeRect, width, height int) {
	stride := width*4 + width*4%4
	len := stride * height
	buf := (*[]byte)(unsafe.Pointer(bufPtr))
	bits := (*buf)[:len]
	args := miniblink.PaintArgs{
		Wke: uintptr(wke),
		Update: GoMiniblink.Bound{
			Point: GoMiniblink.Point{
				X: rect.x,
				Y: rect.y,
			},
			Rect: GoMiniblink.Rect{
				Wdith:  rect.w,
				Height: rect.h,
			},
		},
		Size: GoMiniblink.Rect{
			Wdith:  width,
			Height: height,
		},
		Bits:  bits,
		Param: param,
	}
	_this.paintCallback(args)
}

func (_this *Core) SetOnPaint(callback miniblink.PaintCallback) {
	_this.paintCallback = callback
}

func (_this *Core) LoadUri(uri string) {
	wkeLoadURL(_this.wke, uri)
}
