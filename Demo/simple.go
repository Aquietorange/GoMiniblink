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

	GetModuleHandle         *windows.LazyProc
	RegisterClassEx         *windows.LazyProc
	CreateWindowEx          *windows.LazyProc
	ShowWindow              *windows.LazyProc
	UpdateWindow            *windows.LazyProc
	GetMessage              *windows.LazyProc
	TranslateMessage        *windows.LazyProc
	DispatchMessage         *windows.LazyProc
	DefWindowProc           *windows.LazyProc
	DestroyWindow           *windows.LazyProc
	PostQuitMessage         *windows.LazyProc
	CreateDIBSection        *windows.LazyProc
	CreateCompatibleDC      *windows.LazyProc
	SelectObject            *windows.LazyProc
	BitBlt                  *windows.LazyProc
	DeleteDC                *windows.LazyProc
	DeleteObject            *windows.LazyProc
	GetDC                   *windows.LazyProc
	ReleaseDC               *windows.LazyProc
	GetClientRect           *windows.LazyProc
	BeginPaint              *windows.LazyProc
	EndPaint                *windows.LazyProc
	GetKeyState             *windows.LazyProc
	ImmGetContext           *windows.LazyProc
	ImmSetCompositionWindow *windows.LazyProc
	ImmReleaseContext       *windows.LazyProc
	SetCapture              *windows.LazyProc
	ReleaseCapture          *windows.LazyProc

	wkeInitialize          *windows.LazyProc
	wkeCreateWebView       *windows.LazyProc
	wkeSetHandle           *windows.LazyProc
	wkeSetFocus            *windows.LazyProc
	wkeKillFocus           *windows.LazyProc
	wkeResize              *windows.LazyProc
	wkeLoadURL             *windows.LazyProc
	wkeOnPaintBitUpdated   *windows.LazyProc
	wkePaint               *windows.LazyProc
	wkeFireKeyPressEvent   *windows.LazyProc
	wkeFireKeyUpEvent      *windows.LazyProc
	wkeFireKeyDownEvent    *windows.LazyProc
	wkeFireMouseWheelEvent *windows.LazyProc
	wkeFireMouseEvent      *windows.LazyProc
	wkeGetHeight           *windows.LazyProc
	wkeGetWidth            *windows.LazyProc
	wkeOnLoadUrlBegin      *windows.LazyProc
	wkeFireWindowsMessage  *windows.LazyProc
	wkeGetCaretRect2       *windows.LazyProc

	is64        bool
	appInstance uintptr
	className   string
	userdata    map[uintptr]uintptr
)

const (
	CS_VREDRAW              = 1
	CS_HREDRAW              = 2
	WM_CREATE               = 1
	WM_DESTROY              = 2
	WM_SIZE                 = 5
	WM_PAINT                = 15
	WM_CLOSE                = 16
	WM_MOUSEMOVE            = 512
	WM_LBUTTONDOWN          = 513
	WM_LBUTTONUP            = 514
	WM_LBUTTONDBLCLK        = 515
	WM_RBUTTONDOWN          = 516
	WM_RBUTTONUP            = 517
	WM_RBUTTONDBLCLK        = 518
	WM_MBUTTONDOWN          = 519
	WM_MBUTTONUP            = 520
	WM_MBUTTONDBLCLK        = 521
	WM_MOUSEWHEEL           = 522
	WM_SETCURSOR            = 32
	WM_SETFOCUS             = 7
	WM_KILLFOCUS            = 8
	WM_SYSKEYDOWN           = 260
	WM_KEYDOWN              = 256
	WM_SYSKEYUP             = 261
	WM_KEYUP                = 257
	WM_SYSCHAR              = 262
	WM_CHAR                 = 258
	WM_IME_STARTCOMPOSITION = 269
	WS_OVERLAPPEDWINDOW     = 0x00000000 | 0x00C00000 | 0x00080000 | 0x00040000 | 0x00020000 | 0x00010000
	SW_SHOW                 = 5
	BI_RGB                  = 0
	DIB_RGB_COLORS          = 0
	SRCCOPY                 = 0x00CC0020
	MK_LBUTTON              = 0x0001
	MK_MBUTTON              = 0x0010
	MK_RBUTTON              = 0x0002
	VK_SHIFT                = 16
	VK_CONTROL              = 17
	KF_REPEAT               = 16384
	KF_EXTENDED             = 256
	CFS_POINT               = 2
	CFS_FORCE_POSITION      = 32

	WKE_LBUTTON  = 0x01
	WKE_RBUTTON  = 0x02
	WKE_SHIFT    = 0x04
	WKE_CONTROL  = 0x08
	WKE_MBUTTON  = 0x10
	WKE_EXTENDED = 0x0100
	WKE_REPEAT   = 0x4000
)

func init() {
	is64 = unsafe.Sizeof(uintptr(0)) == 8
	gdi32Lib = windows.NewLazySystemDLL("gdi32.dll")
	user32Lib = windows.NewLazySystemDLL("user32.dll")
	imm32Lib = windows.NewLazySystemDLL("imm32.dll")
	kernel32Lib = windows.NewLazySystemDLL("kernel32.dll")
	if is64 {
		wkeLib = windows.NewLazyDLL("miniblink_x64.dll")
	} else {
		wkeLib = windows.NewLazyDLL("miniblink_x86.dll")
	}

	GetModuleHandle = kernel32Lib.NewProc("GetModuleHandleW")
	RegisterClassEx = user32Lib.NewProc("RegisterClassExW")
	CreateWindowEx = user32Lib.NewProc("CreateWindowExW")
	ShowWindow = user32Lib.NewProc("ShowWindow")
	UpdateWindow = user32Lib.NewProc("UpdateWindow")
	GetMessage = user32Lib.NewProc("GetMessageW")
	TranslateMessage = user32Lib.NewProc("TranslateMessage")
	DispatchMessage = user32Lib.NewProc("DispatchMessageW")
	DefWindowProc = user32Lib.NewProc("DefWindowProcW")
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
	ImmGetContext = imm32Lib.NewProc("ImmGetContext")
	ImmSetCompositionWindow = imm32Lib.NewProc("ImmSetCompositionWindow")
	ImmReleaseContext = imm32Lib.NewProc("ImmReleaseContext")
	SetCapture = user32Lib.NewProc("SetCapture")
	ReleaseCapture = user32Lib.NewProc("ReleaseCapture")

	wkeInitialize = wkeLib.NewProc("wkeInitialize")
	wkeFireKeyPressEvent = wkeLib.NewProc("wkeFireKeyPressEvent")
	wkeFireKeyUpEvent = wkeLib.NewProc("wkeFireKeyUpEvent")
	wkeFireKeyDownEvent = wkeLib.NewProc("wkeFireKeyDownEvent")
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
	wkeKillFocus = wkeLib.NewProc("wkeKillFocus")
	wkeOnLoadUrlBegin = wkeLib.NewProc("wkeOnLoadUrlBegin")
	wkeFireWindowsMessage = wkeLib.NewProc("wkeFireWindowsMessage")
	wkeGetCaretRect2 = wkeLib.NewProc("wkeGetCaretRect2")

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

func onPaintBitUpdated(wke uintptr, param, bits uintptr, rect uintptr, width, _ uint32) uintptr {
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
	buf := append([]byte(str), 0)
	return unsafe.Pointer(&buf[0])
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

func toLp(value uintptr) int32 {
	return int32(int16(int32(value)))
}

func toHp(value uintptr) int32 {
	return int32(int16(int32(value) >> 16 & 0xffff))
}

func initWke(hWnd uintptr) {
	wke, _, _ := wkeCreateWebView.Call()
	wkeSetHandle.Call(wke, hWnd)
	wkeOnPaintBitUpdated.Call(wke, syscall.NewCallbackCDecl(onPaintBitUpdated), hWnd)
	wkeOnLoadUrlBegin.Call(wke, syscall.NewCallbackCDecl(func(wke, param, utf8Url, job uintptr) uintptr {
		//println(toUtf8(utf8Url))
		return uintptr(byte(0))
	}), hWnd)
	w, h := getSize(hWnd)
	wkeResize.Call(wke, uintptr(int32(w)), uintptr(int32(h)))
	userdata[hWnd] = wke
}

func windowMsgProc(hWnd uintptr, msg uint32, wParam uintptr, lParam uintptr) uintptr {
	switch msg {
	case WM_CREATE:
		initWke(hWnd)
	case WM_CLOSE:
		DestroyWindow.Call(hWnd)
	case WM_DESTROY:
		PostQuitMessage.Call(0)
	case WM_SIZE:
		w := toLp(lParam)
		h := toHp(lParam)
		wkeResize.Call(userdata[hWnd], uintptr(w), uintptr(h))
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
		return 0
	case WM_SETCURSOR:
		//todo 在从窗体右边进入时不能及时更新光标
		r, _, _ := wkeFireWindowsMessage.Call(userdata[hWnd], hWnd, uintptr(msg), wParam, lParam, 0)
		c := byte(r)
		if c != 0 {
			return 0
		}
	case WM_SETFOCUS:
		wkeSetFocus.Call(userdata[hWnd])
		return 0
	case WM_KILLFOCUS:
		wkeKillFocus.Call(userdata[hWnd])
		return 0
	case WM_MOUSEMOVE,
		WM_LBUTTONUP, WM_LBUTTONDOWN, WM_LBUTTONDBLCLK,
		WM_RBUTTONUP, WM_RBUTTONDOWN, WM_RBUTTONDBLCLK,
		WM_MBUTTONUP, WM_MBUTTONDOWN, WM_MBUTTONDBLCLK,
		WM_MOUSEWHEEL:
		x, y, delta := toLp(lParam), toHp(lParam), 0
		wp := int(wParam)
		switch msg {
		case WM_MOUSEWHEEL:
			wp, delta = int(toLp(wParam)), int(toHp(wParam))
		case WM_LBUTTONDOWN, WM_MBUTTONDOWN, WM_RBUTTONDOWN:
			SetCapture.Call(hWnd)
		case WM_LBUTTONUP, WM_MBUTTONUP, WM_RBUTTONUP:
			ReleaseCapture.Call()
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
		var r uintptr
		if msg == WM_MOUSEWHEEL {
			r, _, _ = wkeFireMouseWheelEvent.Call(wke, uintptr(x), uintptr(y), uintptr(delta), uintptr(flags))
		} else {
			r, _, _ = wkeFireMouseEvent.Call(wke, uintptr(msg), uintptr(x), uintptr(y), uintptr(flags))
		}
		if byte(r) != 0 {
			return 0
		}
	case WM_SYSKEYDOWN, WM_KEYDOWN:
		flags := 0
		lp := int32(lParam)
		if lp>>16&KF_REPEAT != 0 {
			flags |= WKE_REPEAT
		}
		if lp>>16&KF_EXTENDED != 0 {
			flags |= WKE_EXTENDED
		}
		isSys := 0
		if msg == WM_SYSKEYDOWN {
			isSys = 1
		}
		r, _, _ := wkeFireKeyDownEvent.Call(userdata[hWnd], wParam, uintptr(flags), uintptr(isSys))
		if byte(r) != 0 {
			return 0
		}
	case WM_SYSKEYUP, WM_KEYUP:
		flags := 0
		lp := int32(lParam)
		if lp>>16&KF_REPEAT != 0 {
			flags |= WKE_REPEAT
		}
		if lp>>16&KF_EXTENDED != 0 {
			flags |= WKE_EXTENDED
		}
		isSys := 0
		if msg == WM_SYSKEYDOWN {
			isSys = 1
		}
		r, _, _ := wkeFireKeyUpEvent.Call(userdata[hWnd], wParam, uintptr(flags), uintptr(isSys))
		if byte(r) != 0 {
			return 0
		}
	case WM_SYSCHAR, WM_CHAR:
		flags := WKE_REPEAT
		if int32(lParam)>>16&KF_EXTENDED != 0 {
			flags |= WKE_EXTENDED
		}
		isSys := 0
		if msg == WM_SYSKEYDOWN {
			isSys = 1
		}
		r, _, _ := wkeFireKeyPressEvent.Call(userdata[hWnd], wParam, uintptr(flags), uintptr(isSys))
		if byte(r) != 0 {
			return 0
		}
	case WM_IME_STARTCOMPOSITION:
		rect, _, _ := wkeGetCaretRect2.Call(userdata[hWnd])
		rx, ry := int(*(*uint32)(unsafe.Pointer(rect))), int(*(*uint32)(unsafe.Pointer(rect + 4)))
		comp := struct {
			style, x, y, l, t, r, b int32
		}{}
		comp.style = CFS_POINT | CFS_FORCE_POSITION
		comp.x, comp.y = int32(rx), int32(ry)
		imc, _, _ := ImmGetContext.Call(hWnd)
		ImmSetCompositionWindow.Call(imc, uintptr(unsafe.Pointer(&comp)))
		ImmReleaseContext.Call(hWnd, imc)
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
	class.LpfnWndProc = syscall.NewCallbackCDecl(windowMsgProc)
	RegisterClassEx.Call(uintptr(unsafe.Pointer(&class)))
}

func main() {
	RegisterWindowClass()
	hWnd, _, _ := CreateWindowEx.Call(0,
		uintptr(unsafe.Pointer(utf16PtrFromString(className))),
		uintptr(unsafe.Pointer(utf16PtrFromString("simple"))),
		uintptr(WS_OVERLAPPEDWINDOW),
		100, 100, 1480, 800, 0, 0, appInstance, 0)
	wkeLoadURL.Call(userdata[hWnd], uintptr(utf8To("https://www.baidu.com")))
	ShowWindow.Call(hWnd, uintptr(SW_SHOW))
	UpdateWindow.Call(hWnd)
	msg := struct {
		HWnd    uintptr
		Message uint32
		WParam  uint32
		LParam  int32
		Time    uint32
		Pt      struct{ X, Y int32 }
	}{}
	for {
		if code, _, _ := GetMessage.Call(uintptr(unsafe.Pointer(&msg)), 0, 0, 0); code != 0 {
			TranslateMessage.Call(uintptr(unsafe.Pointer(&msg)))
			DispatchMessage.Call(uintptr(unsafe.Pointer(&msg)))
		} else {
			break
		}
	}
}
