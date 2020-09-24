package main

import (
	"fmt"
	"qq2564874169/goMiniblink/forms"
	cs "qq2564874169/goMiniblink/forms/controls"
	"qq2564874169/goMiniblink/forms/windows"
	"unsafe"
)

func main() {
	fmt.Println("is x64", unsafe.Sizeof(uintptr(0)) == 8)
	cs.App = new(windows.Provider).Init()
	cs.App.SetIcon("app.ico")

	frm := new(cs.Form).Init()
	frm.SetTitle("this is form")
	frm.SetSize(600, 600)

	ctrl := new(cs.Control).Init()
	ctrl.SetSize(200, 300)
	ctrl.SetLocation(50, 50)
	ctrl.SetBgColor(0x2FAEE3)
	ctrl.EvMouseClick["show"] = func(s cs.GUI, e *forms.MouseEvArgs) {
		fmt.Println("click", e.IsDouble, e.X, e.Y)
	}
	ctrl.EvMouseUp["showForm"] = func(s cs.GUI, e *forms.MouseEvArgs) {

	}
	frm.AddChild(ctrl)

	cs.Run(frm)
}
