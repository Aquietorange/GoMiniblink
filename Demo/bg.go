package main

import (
	"qq.2564874169/goMiniblink/platform/windows/win32"
	"syscall"
	"unsafe"
)

func main() {
	println(1 << 30)
	className := "DemoClass"
	var class = win32.WNDCLASSEX{
		Style:         win32.CS_HREDRAW | win32.CS_VREDRAW,
		LpfnWndProc:   syscall.NewCallback(defaultMsgProc),
		HInstance:     win32.GetModuleHandle(nil),
		LpszClassName: sto16(className),
	}
	class.CbSize = uint32(unsafe.Sizeof(class))
	win32.RegisterClassEx(&class)
	hWnd := win32.CreateWindowEx(
		0,
		sto16(className),
		sto16(className),
		win32.WS_OVERLAPPEDWINDOW,
		100, 100, 400, 300, 0, 0, class.HInstance, unsafe.Pointer(nil))
	win32.ShowWindow(hWnd, win32.SW_SHOW)
	win32.UpdateWindow(hWnd)
	var msg win32.MSG
	for {
		if win32.GetMessage(&msg, 0, 0, 0) {
			win32.TranslateMessage(&msg)
			win32.DispatchMessage(&msg)
		} else {
			break
		}
	}
}

var memBmp win32.HBITMAP
var rgba = []uint8{255, 0, 0, 0}

func Paint(hWnd win32.HWND, msg uint32, wParam uintptr, lParam uintptr) {
	pt := new(win32.PAINTSTRUCT)
	hdc := win32.BeginPaint(hWnd, pt)
	w := pt.RcPaint.Right - pt.RcPaint.Left
	h := pt.RcPaint.Bottom - pt.RcPaint.Top
	if memBmp == 0 {
		pix := make([]uint8, w*h*4)
		for i := 0; i < len(pix); i++ {
			pix[i] = rgba[i%4]
		}
		var head win32.BITMAPV5HEADER
		head.BiSize = uint32(unsafe.Sizeof(head))
		head.BiWidth = w
		head.BiHeight = h
		head.BiBitCount = 32
		head.BiPlanes = 1
		head.BiCompression = win32.BI_RGB

		var lpBits unsafe.Pointer
		memBmp = win32.CreateDIBSection(hdc, &head.BITMAPINFOHEADER, win32.DIB_RGB_COLORS, &lpBits, 0, 0)
		bits := (*[1 << 30]byte)(lpBits)
		for i := range pix {
			bits[i] = pix[i]
		}
	}
	memDc := win32.CreateCompatibleDC(hdc)
	old := win32.SelectObject(memDc, win32.HGDIOBJ(memBmp))
	win32.BitBlt(hdc, pt.RcPaint.Left, pt.RcPaint.Top, w, h, memDc, 0, 0, win32.SRCCOPY)
	win32.SelectObject(memDc, old)
	win32.DeleteDC(memDc)
	win32.DeleteObject(win32.HGDIOBJ(memBmp))
	memBmp = 0
	win32.EndPaint(hWnd, pt)
}

func defaultMsgProc(hWnd win32.HWND, msg uint32, wParam uintptr, lParam uintptr) uintptr {
	switch msg {
	case win32.WM_CREATE:
		break
	case win32.WM_PAINT:
		Paint(hWnd, msg, wParam, lParam)
		break
	case win32.WM_CLOSE:
		win32.DestroyWindow(hWnd)
		break
	case win32.WM_DESTROY:
		win32.PostQuitMessage(0)
		break
	}
	return win32.DefWindowProc(hWnd, msg, wParam, lParam)
}

func sto16(str string) *uint16 {
	ptr, _ := syscall.UTF16PtrFromString(str)
	return ptr
}
