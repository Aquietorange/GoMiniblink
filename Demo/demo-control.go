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

	frm := new(cs.Form).Init()
	frm.SetTitle("miniblink窗口")
	frm.SetSize(800, 500)
	frm.SetStartPosition(fm.FormStart_Manual)
	frm.SetLocation(100, 100)
	//frm.SetBorderStyle(forms.FormBorder_None)
	//frm.NoneBorderResize()

	mb := new(gm.MiniblinkBrowser).Init()
	mb.SetBgColor(0x2FAEE3)
	mb.SetAnchor(fm.AnchorStyle_Left | fm.AnchorStyle_Bottom | fm.AnchorStyle_Right | fm.AnchorStyle_Top)
	mb.ResourceLoader = append(mb.ResourceLoader, new(gm.FileLoader).Init("Res", "local"))
	frm.AddChild(mb)
	frm.EvShow["init"] = func(s cs.GUI) {
		mb.LoadUri("http://local/window.html")
	}
	cs.Run(frm)
}
