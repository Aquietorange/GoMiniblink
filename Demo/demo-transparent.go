package main

import (
	gm "GoMiniblink"
	fm "GoMiniblink/forms"
	cs "GoMiniblink/forms/controls"
	gw "GoMiniblink/forms/windows"
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("x64 is", unsafe.Sizeof(uintptr(0)) == 8)
	cs.App = new(gw.Provider).Init()

	frm := new(gm.MiniblinkForm).Init()
	frm.TransparentMode()
	frm.SetLocation(100, 100)
	frm.SetSize(300, 300)
	frm.SetBorderStyle(fm.FormBorder_None)
	frm.NoneBorderResize()
	frm.View.ResourceLoader = append(frm.View.ResourceLoader, new(gm.FileLoader).Init("Res", "local"))
	frm.EvLoad["show"] = func(s cs.GUI) {
		frm.View.LoadUri("http://local/transparent.html")
	}
	cs.Run(&frm.Form)
}
