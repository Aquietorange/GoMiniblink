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
	cs.App.SetIcon("app.ico")

	var frm = new(gm.MiniblinkForm).Init()
	frm.SetTitle("miniblink窗口")
	frm.SetLocation(100, 100)
	frm.SetSize(800, 500)
	frm.SetBorderStyle(fm.FormBorder_None)
	frm.NoneBorderResize()
	frm.View.ResourceLoader = append(frm.View.ResourceLoader, new(gm.FileLoader).Init("Res", "local"))
	frm.View.EvConsole["show"] = func(_ *gm.MiniblinkBrowser, e gm.ConsoleEvArgs) {
		fmt.Println(e.Message())
	}
	frm.EvShow["init"] = func(s cs.GUI) {
		frm.View.LoadUri("http://local/window.html")
	}
	cs.Run(&frm.Form)
}
