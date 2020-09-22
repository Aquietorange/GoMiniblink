package main

import (
	"fmt"
	"qq2564874169/goMiniblink/forms"
	"qq2564874169/goMiniblink/forms/controls"
	"qq2564874169/goMiniblink/forms/platform/windows"
	"unsafe"
)

func main() {
	fmt.Println("is x64", unsafe.Sizeof(uintptr(0)) == 8)
	controls.App = new(windows.Provider).Init()
	controls.App.SetIcon("app.ico")

	frm := new(controls.Form).Init()
	frm.SetTitle("this is form")
	frm.SetSize(600, 600)

	ctrl := new(controls.Control).Init()
	ctrl.SetSize(200, 300)
	ctrl.SetLocation(50, 50)
	ctrl.SetBgColor(0x2FAEE3)
	ctrl.EvMouseUp["showForm"] = func(sender interface{}, e *forms.MouseEvArgs) {
		dlg := new(controls.Form).Init()
		dlg.SetTitle("this is dialog")
		dlg.SetSize(200, 200)
		dlg.SetLocation(500, 500)
		dlg.ShowDialog()
		//dlg.Show()
	}
	frm.AddChild(ctrl)

	controls.Run(frm)
}
