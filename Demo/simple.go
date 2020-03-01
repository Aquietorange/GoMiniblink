package main

import (
	"golang.org/x/sys/windows"
	"image"
	"image/draw"
	"math"
	"syscall"
	"unsafe"
)

var (
	gdi32Lib    *windows.LazyDLL
	user32Lib   *windows.LazyDLL
	imm32Lib    *windows.LazyDLL
	kernel32Lib *windows.LazyDLL
	wkeLib      *windows.LazyDLL

	GetModuleHandle    *windows.LazyProc
	RegisterClassEx    *windows.LazyProc
	CreateWindowEx     *windows.LazyProc
	ShowWindow         *windows.LazyProc
	UpdateWindow       *windows.LazyProc
	GetMessage         *windows.LazyProc
	TranslateMessage   *windows.LazyProc
	DispatchMessage    *windows.LazyProc
	DefWindowProc      *windows.LazyProc
	CallWindowProc     *windows.LazyProc
	DestroyWindow      *windows.LazyProc
	PostQuitMessage    *windows.LazyProc
	CreateDIBSection   *windows.LazyProc
	CreateCompatibleDC *windows.LazyProc
	SelectObject       *windows.LazyProc
	BitBlt             *windows.LazyProc
	DeleteDC           *windows.LazyProc
	DeleteObject       *windows.LazyProc
	GetDC              *windows.LazyProc
	ReleaseDC          *windows.LazyProc
	GetClientRect      *windows.LazyProc
	BeginPaint         *windows.LazyProc
	EndPaint           *windows.LazyProc
	GetKeyState        *windows.LazyProc
	LoadCursor         *windows.LazyProc
	SetCursor          *windows.LazyProc

	wkeInitialize          *windows.LazyProc
	wkeCreateWebView       *windows.LazyProc
	wkeSetHandle           *windows.LazyProc
	wkeSetFocus            *windows.LazyProc
	wkeResize              *windows.LazyProc
	wkeLoadURL             *windows.LazyProc
	wkeOnPaintBitUpdated   *windows.LazyProc
	wkePaint               *windows.LazyProc
	wkeFireKeyPressEvent   *windows.LazyProc
	wkeFireKeyUpEvent      *windows.LazyProc
	wkeFireKeyDownEvent    *windows.LazyProc
	wkeGetCursorInfoType   *windows.LazyProc
	wkeFireMouseWheelEvent *windows.LazyProc
	wkeFireMouseEvent      *windows.LazyProc
	wkeGetHeight           *windows.LazyProc
	wkeGetWidth            *windows.LazyProc
	wkeOnLoadUrlBegin      *windows.LazyProc
	wkeFireWindowsMessage  *windows.LazyProc

	appInstance          uintptr
	className            string
	userdata             map[uintptr]uintptr
	refOnPaintBitUpdated uintptr
)

const (
	CS_VREDRAW          = 1
	CS_HREDRAW          = 2
	WM_CREATE           = 1
	WM_DESTROY          = 2
	WM_SIZE             = 5
	WM_PAINT            = 15
	WM_CLOSE            = 16
	WM_MOUSEMOVE        = 512
	WM_LBUTTONDOWN      = 513
	WM_LBUTTONUP        = 514
	WM_LBUTTONDBLCLK    = 515
	WM_RBUTTONDOWN      = 516
	WM_RBUTTONUP        = 517
	WM_RBUTTONDBLCLK    = 518
	WM_MBUTTONDOWN      = 519
	WM_MBUTTONUP        = 520
	WM_MBUTTONDBLCLK    = 521
	WM_MOUSEWHEEL       = 522
	WM_SETCURSOR        = 32
	WS_OVERLAPPEDWINDOW = 0x00000000 | 0x00C00000 | 0x00080000 | 0x00040000 | 0x00020000 | 0x00010000
	SW_SHOW             = 5
	BI_RGB              = 0
	DIB_RGB_COLORS      = 0
	SRCCOPY             = 0x00CC0020
	MK_LBUTTON          = 0x0001
	MK_MBUTTON          = 0x0010
	MK_RBUTTON          = 0x0002
	WKE_LBUTTON         = 0x01
	WKE_RBUTTON         = 0x02
	WKE_SHIFT           = 0x04
	WKE_CONTROL         = 0x08
	WKE_MBUTTON         = 0x10
	VK_SHIFT            = 16
	VK_CONTROL          = 17
	IDC_IBEAM           = 32513
	IDC_ARROW           = 32512
	IDC_HAND            = 32649
	IDC_SIZEWE          = 32644
	IDC_SIZENS          = 32645
)

func init() {
	//is64 := unsafe.Sizeof(uintptr(0)) == 8
	gdi32Lib = windows.NewLazySystemDLL("gdi32.dll")
	user32Lib = windows.NewLazySystemDLL("user32.dll")
	imm32Lib = windows.NewLazySystemDLL("imm32.dll")
	kernel32Lib = windows.NewLazySystemDLL("kernel32.dll")
	wkeLib = windows.NewLazyDLL("miniblink_x64.dll")

	GetModuleHandle = kernel32Lib.NewProc("GetModuleHandleW")
	RegisterClassEx = user32Lib.NewProc("RegisterClassExW")
	CreateWindowEx = user32Lib.NewProc("CreateWindowExW")
	ShowWindow = user32Lib.NewProc("ShowWindow")
	UpdateWindow = user32Lib.NewProc("UpdateWindow")
	GetMessage = user32Lib.NewProc("GetMessageW")
	TranslateMessage = user32Lib.NewProc("TranslateMessage")
	DispatchMessage = user32Lib.NewProc("DispatchMessageW")
	DefWindowProc = user32Lib.NewProc("DefWindowProcW")
	CallWindowProc = user32Lib.NewProc("CallWindowProcW")
	DestroyWindow = user32Lib.NewProc("DestroyWindow")
	PostQuitMessage = user32Lib.NewProc("PostQuitMessage")
	CreateDIBSection = gdi32Lib.NewProc("CreateDIBSection")
	CreateCompatibleDC = gdi32Lib.NewProc("CreateCompatibleDC")
	SelectObject = gdi32Lib.NewProc("SelectObject")
	BitBlt = gdi32Lib.NewProc("BitBlt")
	DeleteDC = gdi32Lib.NewProc("DeleteDC")
	DeleteObject = gdi32Lib.NewProc("DeleteObject")
	GetDC = user32Lib.NewProc("GetDC")
	ReleaseDC = user32Lib.NewProc("ReleaseDC")
	GetClientRect = user32Lib.NewProc("GetClientRect")
	BeginPaint = user32Lib.NewProc("BeginPaint")
	EndPaint = user32Lib.NewProc("EndPaint")
	GetKeyState = user32Lib.NewProc("GetKeyState")
	LoadCursor = user32Lib.NewProc("LoadCursorW")
	SetCursor = user32Lib.NewProc("SetCursor")

	wkeInitialize = wkeLib.NewProc("wkeInitialize")
	wkeFireKeyPressEvent = wkeLib.NewProc("wkeFireKeyPressEvent")
	wkeFireKeyUpEvent = wkeLib.NewProc("wkeFireKeyUpEvent")
	wkeFireKeyDownEvent = wkeLib.NewProc("wkeFireKeyDownEvent")
	wkeGetCursorInfoType = wkeLib.NewProc("wkeGetCursorInfoType")
	wkeFireMouseWheelEvent = wkeLib.NewProc("wkeFireMouseWheelEvent")
	wkeFireMouseEvent = wkeLib.NewProc("wkeFireMouseEvent")
	wkeGetHeight = wkeLib.NewProc("wkeGetHeight")
	wkeGetWidth = wkeLib.NewProc("wkeGetWidth")
	wkePaint = wkeLib.NewProc("wkePaint")
	wkeLoadURL = wkeLib.NewProc("wkeLoadURL")
	wkeResize = wkeLib.NewProc("wkeResize")
	wkeOnPaintBitUpdated = wkeLib.NewProc("wkeOnPaintBitUpdated")
	wkeSetHandle = wkeLib.NewProc("wkeSetHandle")
	wkeCreateWebView = wkeLib.NewProc("wkeCreateWebView")
	wkeInitialize = wkeLib.NewProc("wkeInitialize")
	wkeSetFocus = wkeLib.NewProc("wkeSetFocus")
	wkeOnLoadUrlBegin = wkeLib.NewProc("wkeOnLoadUrlBegin")
	wkeFireWindowsMessage = wkeLib.NewProc("wkeFireWindowsMessage")

	code, _, err := wkeInitialize.Call()
	if code == 0 {
		panic(err)
	}
	userdata = make(map[uintptr]uintptr)
	className = "miniblink"
}

func utf16PtrFromString(str string) *uint16 {
	p, _ := syscall.UTF16PtrFromString(str)
	return p
}

func drawToDc(dc uintptr, src *image.RGBA, width, height, xDst, yDst int) {
	var head struct {
		BiSize          uint32
		BiWidth         int32
		BiHeight        int32
		BiPlanes        uint16
		BiBitCount      uint16
		BiCompression   uint32
		BiSizeImage     uint32
		BiXPelsPerMeter int32
		BiYPelsPerMeter int32
		BiClrUsed       uint32
		BiClrImportant  uint32
	}
	head.BiSize = uint32(unsafe.Sizeof(head))
	head.BiWidth = int32(width)
	head.BiHeight = int32(height * -1)
	head.BiBitCount = 32
	head.BiPlanes = 1
	head.BiCompression = BI_RGB

	var lpBits unsafe.Pointer
	bmp, _, _ := CreateDIBSection.Call(dc, uintptr(unsafe.Pointer(&head)), DIB_RGB_COLORS, uintptr(unsafe.Pointer(&lpBits)), 0, 0)
	bits := (*[1 << 30]byte)(lpBits)
	for i := 0; i < len(src.Pix); i++ {
		bits[i] = src.Pix[i]
	}
	memDc, _, _ := CreateCompatibleDC.Call(dc)
	SelectObject.Call(memDc, bmp)
	BitBlt.Call(dc, uintptr(xDst), uintptr(yDst), uintptr(width), uintptr(height), memDc, 0, 0, SRCCOPY)
	DeleteDC.Call(memDc)
	DeleteObject.Call(bmp)
}

func onPaintBitUpdated(wke uintptr, param, bits uintptr, rect uintptr, width, height uint32) uintptr {
	wh, _, _ := wkeGetHeight.Call(wke)
	rx, ry := int(*(*uint32)(unsafe.Pointer(rect))), int(*(*uint32)(unsafe.Pointer(rect + 4)))
	w, h := int(*(*uint32)(unsafe.Pointer(rect + 8))), int(*(*uint32)(unsafe.Pointer(rect + 12)))
	w = int(math.Min(float64(w), float64(width)))
	h = int(math.Min(float64(h), float64(wh)))
	bmp := image.NewRGBA(image.Rect(0, 0, w, h))
	pixs := (*[1 << 30]byte)(unsafe.Pointer(bits))
	stride := int(width) * 4
	for y := 0; y < h; y++ {
		for x := 0; x < w*4; x++ {
			sp := bmp.Stride*y + x
			dp := stride*(ry+y) + rx*4 + x
			bmp.Pix[sp] = pixs[dp]
		}
	}
	dc, _, _ := GetDC.Call(param)
	drawToDc(dc, bmp, w, h, rx, ry)
	ReleaseDC.Call(param, dc)
	return 0
}

func getSize(hWnd uintptr) (w, h int) {
	rect := struct{ l, t, r, b int32 }{}
	GetClientRect.Call(hWnd, uintptr(unsafe.Pointer(&rect)))
	return int(rect.r - rect.l), int(rect.b - rect.t)
}

func utf8To(str string) unsafe.Pointer {
	ptr := []byte(str)
	return unsafe.Pointer(&ptr[0])
}

func toUtf8(ptr uintptr) string {
	var seq []byte
	for {
		b := *((*byte)(unsafe.Pointer(ptr)))
		if b != 0 {
			seq = append(seq, b)
			ptr++
		} else {
			break
		}
	}
	return string(seq)
}

func initWke(hWnd uintptr) {
	wke, _, _ := wkeCreateWebView.Call()
	userdata[hWnd] = wke
	wkeSetHandle.Call(wke, hWnd)
	refOnPaintBitUpdated = syscall.NewCallback(onPaintBitUpdated)
	wkeOnPaintBitUpdated.Call(wke, refOnPaintBitUpdated, hWnd)
	wkeOnLoadUrlBegin.Call(wke, syscall.NewCallback(func(wke, param, utf8Url, job uintptr) uintptr {
		return 0
	}), hWnd)
	w, h := getSize(hWnd)
	wkeResize.Call(wke, uintptr(int32(w)), uintptr(int32(h)))
	wkeLoadURL.Call(wke, uintptr(utf8To("https://www.baidu.com")))
}

func toLp(value uintptr) int32 {
	return int32(int16(int32(value)))
}

func toHp(value uintptr) int32 {
	return int32(int16(int32(value) >> 16 & 0xffff))
}

func windowMsgProc(hWnd uintptr, msg uint32, wParam uintptr, lParam uintptr) uintptr {
	switch msg {
	case WM_CREATE:
		initWke(hWnd)
	case WM_DESTROY:
		PostQuitMessage.Call(0)
	case WM_SIZE:
		w := toLp(lParam)
		h := toHp(lParam)
		wke := userdata[hWnd]
		wkeResize.Call(wke, uintptr(w), uintptr(h))
	case WM_PAINT:
		pt := struct {
			Hdc         uintptr
			FErase      byte
			RcPaint     struct{ l, t, r, b int32 }
			FRestore    byte
			FIncUpdate  byte
			RgbReserved [32]byte
		}{}
		hdc, _, _ := BeginPaint.Call(hWnd, uintptr(unsafe.Pointer(&pt)))
		wke := userdata[hWnd]
		w, _, _ := wkeGetWidth.Call(wke)
		h, _, _ := wkeGetHeight.Call(wke)
		if w > 0 && h > 0 {
			view := image.NewRGBA(image.Rect(0, 0, int(uint32(w)), int(uint32(h))))
			wkePaint.Call(wke, uintptr(unsafe.Pointer(&view.Pix[0])), 0)
			bmp := image.NewRGBA(image.Rect(0, 0, int(pt.RcPaint.r-pt.RcPaint.l), int(pt.RcPaint.b-pt.RcPaint.t)))
			draw.Draw(bmp, bmp.Rect, view, image.Pt(int(pt.RcPaint.l), int(pt.RcPaint.t)), draw.Src)
			drawToDc(hdc, bmp, bmp.Rect.Dx(), bmp.Rect.Dy(), int(uint32(pt.RcPaint.l)), int(uint32(pt.RcPaint.t)))
		}
		EndPaint.Call(hWnd, uintptr(unsafe.Pointer(&pt)))
	case WM_CLOSE:
		DestroyWindow.Call(hWnd)
	case WM_SETCURSOR:
		wke := userdata[hWnd]
		res := (*uint16)(unsafe.Pointer(uintptr(IDC_ARROW)))
		h, _, _ := LoadCursor.Call(0, uintptr(unsafe.Pointer(res)))
		r, _, _ := wkeFireWindowsMessage.Call(wke, hWnd, uintptr(msg), 0, 0, 0)
		if r == 0 {
			SetCursor.Call(h)
		}
	case WM_MOUSEMOVE,
		WM_LBUTTONUP, WM_LBUTTONDOWN, WM_LBUTTONDBLCLK,
		WM_RBUTTONUP, WM_RBUTTONDOWN, WM_RBUTTONDBLCLK,
		WM_MBUTTONUP, WM_MBUTTONDOWN, WM_MBUTTONDBLCLK,
		WM_MOUSEWHEEL:
		x, y, delta := toLp(lParam), toHp(lParam), 0
		wp := int(wParam)
		if msg == WM_MOUSEWHEEL {
			wp, delta = int(toLp(wParam)), int(toHp(wParam))
		}
		flags := 0
		if wp&MK_LBUTTON != 0 {
			flags |= WKE_LBUTTON
		}
		if wp&MK_MBUTTON != 0 {
			flags |= WKE_MBUTTON
		}
		if wp&MK_RBUTTON != 0 {
			flags |= WKE_RBUTTON
		}
		shift, _, _ := GetKeyState.Call(uintptr(VK_SHIFT))
		if int32(shift) < 0 {
			flags |= WKE_SHIFT
		}
		ctrl, _, _ := GetKeyState.Call(uintptr(VK_CONTROL))
		if int32(ctrl) < 0 {
			flags |= WKE_CONTROL
		}
		wke := userdata[hWnd]
		if msg == WM_MOUSEWHEEL {
			wkeFireMouseWheelEvent.Call(wke, uintptr(x), uintptr(y), uintptr(delta), uintptr(flags))
		} else {
			wkeFireMouseEvent.Call(wke, uintptr(msg), uintptr(x), uintptr(y), uintptr(flags))
		}
		return 1
	}
	code, _, _ := DefWindowProc.Call(hWnd, uintptr(msg), wParam, lParam)
	return code
}

func RegisterWindowClass() {
	class := struct {
		CbSize        uint32
		Style         uint32
		LpfnWndProc   uintptr
		CbClsExtra    int32
		CbWndExtra    int32
		HInstance     uintptr
		HIcon         uintptr
		HCursor       uintptr
		HbrBackground uintptr
		LpszMenuName  *uint16
		LpszClassName *uint16
		HIconSm       uintptr
	}{}
	class.CbSize = uint32(unsafe.Sizeof(class))
	class.Style = CS_VREDRAW | CS_HREDRAW
	class.HInstance, _, _ = GetModuleHandle.Call(0)
	class.LpszClassName = utf16PtrFromString(className)
	class.LpfnWndProc = syscall.NewCallback(windowMsgProc)
	code, _, err := RegisterClassEx.Call(uintptr(unsafe.Pointer(&class)))
	if code == 0 {
		panic(err)
	}
}

func main() {
	RegisterWindowClass()
	hWnd, _, _ := CreateWindowEx.Call(0,
		uintptr(unsafe.Pointer(utf16PtrFromString(className))),
		uintptr(unsafe.Pointer(utf16PtrFromString("simple"))),
		uintptr(WS_OVERLAPPEDWINDOW),
		100, 100, 400, 300, 0, 0, appInstance, uintptr(unsafe.Pointer(nil)))
	ShowWindow.Call(hWnd, uintptr(SW_SHOW))
	UpdateWindow.Call(hWnd)
	msg := struct {
		HWnd    uintptr
		Message uint32
		WParam  uintptr
		LParam  uintptr
		Time    uint32
		Pt      struct{ X, Y int32 }
	}{}
	for {
		code, _, _ := GetMessage.Call(uintptr(unsafe.Pointer(&msg)), 0, 0, 0)
		if code != 0 {
			TranslateMessage.Call(uintptr(unsafe.Pointer(&msg)))
			DispatchMessage.Call(uintptr(unsafe.Pointer(&msg)))
		} else {
			break
		}
	}
}
