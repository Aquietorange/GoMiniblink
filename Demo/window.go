package main

import (
	"fmt"
	fm "qq2564874169/goMiniblink/forms"
	cs "qq2564874169/goMiniblink/forms/controls"
	"qq2564874169/goMiniblink/forms/windows"
	"unsafe"
)

func main() {
	fmt.Println("x64 is", unsafe.Sizeof(uintptr(0)) == 8)
	cs.App = new(windows.Provider).Init()
	cs.App.SetIcon("app.ico")

	frm := new(cs.Form).Init()
	frm.SetTitle("this is form")
	frm.SetSize(600, 600)

	ctrl := new(cs.Control).Init()
	ctrl.SetSize(300, 300)
	ctrl.SetLocation(100, 100)
	ctrl.SetBgColor(0x2FAEE3)
	ctrl.EvMouseClick["show_pos"] = func(s cs.GUI, e *fm.MouseEvArgs) {
		fmt.Println(s.GetHandle(), e.X, e.Y)
	}
	frm.AddChild(ctrl)
	frm.EvMouseClick["show_pos"] = func(s cs.GUI, e *fm.MouseEvArgs) {
		fmt.Println(s.GetHandle(), e.X, e.Y)
	}
	cs.Run(frm)
}
