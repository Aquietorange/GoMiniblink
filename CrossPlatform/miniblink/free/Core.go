package free

import (
	"GoMiniblink/CrossPlatform"
	"GoMiniblink/CrossPlatform/miniblink"
	"bytes"
	"encoding/binary"
	"os"
	"unsafe"
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
	wkeOnPaintBitUpdated(_this.wke, _this.paintBitUpdate, 0)
	wkeResize(_this.wke, 500, 500)
	return _this
}

func (_this *Core) paintBitUpdate(wke wkeHandle, param, buf uintptr, rect *wkeRect, width, height int32) uintptr {
	dataLen := width * height
	if dataLen == 0 {
		return 0
	}
	data := make([]uint32, dataLen)
	size := unsafe.Sizeof(uint32(1))
	for i := 0; i < int(dataLen); i++ {
		b := *((*uint32)(unsafe.Pointer(buf)))
		buf = buf + size
		data = append(data, b)
	}
	file := make([]byte, int(dataLen)*int(size))
	buffer := bytes.NewBuffer(file)
	binary.Write(buffer, binary.LittleEndian, data)
	f, _ := os.Create("111.png")
	defer f.Close()
	f.Write(file)
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
