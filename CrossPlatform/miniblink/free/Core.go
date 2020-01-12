package free

import (
	"GoMiniblink/CrossPlatform"
	"GoMiniblink/CrossPlatform/miniblink"
	"fmt"
)

type Core struct {
	view CrossPlatform.IWindow
	wke  wkeHandle

	onPaint miniblink.PaintCallback
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
	wkeOnPaintBitUpdated(_this.wke, _this.firePaint, 0)
	wkeResize(_this.wke, 500, 500)
	return _this
}

func (_this *Core) firePaint(wke wkeHandle, param, buf uintptr, rect *wkeRect, width, height int32) uintptr {
	fmt.Println(rect)
	//stride := width*4 + width*4%4
	//data := make([]byte, stride*height)
	//binary.LittleEndian.PutUint32(data, uint32(buf))
	//args := miniblink.PaintArgs{
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

func (_this *Core) SetOnPaint(callback miniblink.PaintCallback) {
	_this.onPaint = callback
}

func (_this *Core) LoadUri(uri string) {
	wkeLoadURL(_this.wke, uri)
}
