package main

import (
	"fmt"
	gm "qq2564874169/goMiniblink"
	fm "qq2564874169/goMiniblink/forms"
	cs "qq2564874169/goMiniblink/forms/controls"
	"qq2564874169/goMiniblink/forms/windows"
	"unsafe"
)

func main() {
	fmt.Println("is x64", unsafe.Sizeof(uintptr(0)) == 8)
	cs.App = new(windows.Provider).Init()

	var frm = new(gm.MiniblinkForm).Init()
	frm.TransparentMode()
	frm.SetLocation(100, 100)
	frm.SetSize(300, 300)
	frm.SetBorderStyle(fm.FormBorder_None)
	frm.NoneBorderResize()
	frm.View.ResourceLoader = append(frm.View.ResourceLoader, new(gm.FileLoader).Init("Res", "local"))
	frm.EvLoad["init"] = func(s cs.GUI) {
		frm.View.LoadUri("http://local/transparent.html")
	}
	cs.Run(&frm.Form)
}
