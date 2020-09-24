package main

import (
	"fmt"
	g "qq2564874169/goMiniblink"
	"qq2564874169/goMiniblink/forms/controls"
	"qq2564874169/goMiniblink/forms/windows"
	"unsafe"
)

func main() {
	fmt.Println("is x64", unsafe.Sizeof(uintptr(0)) == 8)
	controls.App = new(windows.Provider).Init()

	var frm = new(g.MiniblinkForm).Init()
	frm.EvShow["set_form"] = func(target interface{}) {
		frm.TransparentMode()
		frm.SetLocation(100, 100)
		frm.SetSize(300, 300)
		frm.NoneBorderResize()
	}
	frm.EvShow["set_view"] = func(target interface{}) {
		frm.View.ResourceLoader = append(frm.View.ResourceLoader, new(g.FileLoader).Init("Res", "local"))
		frm.View.LoadUri("http://local/transparent.html")
	}
	controls.Run(&frm.Form)
}
