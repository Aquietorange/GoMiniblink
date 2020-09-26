package main

import (
	"fmt"
	gm "qq2564874169/goMiniblink"
	cs "qq2564874169/goMiniblink/forms/controls"
	"qq2564874169/goMiniblink/forms/windows"
	"unsafe"
)

func main() {
	fmt.Println("is x64", unsafe.Sizeof(uintptr(0)) == 8)
	cs.App = new(windows.Provider).Init()
	cs.App.SetIcon("app.ico")

	frm := new(cs.Form).Init()
	frm.SetTitle("miniblink窗口")
	frm.SetSize(800, 500)
	frm.SetLocation(100, 100)
	frm.SetBgColor(0x2FAEE3)

	mb := new(gm.MiniblinkBrowser).Init()
	mb.SetBgColor(0x2FAEE3)
	mb.SetSize(700, 400)
	mb.SetLocation(50, 50)
	mb.ResourceLoader = append(mb.ResourceLoader, new(gm.FileLoader).Init("Res", "local"))

	frm.AddChild(mb)
	frm.EvShow["init"] = func(s cs.GUI) {
		mb.LoadUri("http://local/control.html")
	}
	cs.Run(frm)
}
