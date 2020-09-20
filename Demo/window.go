package main

import (
	"fmt"
	"qq2564874169/goMiniblink/forms/controls"
	"qq2564874169/goMiniblink/forms/platform/windows"
	"unsafe"
)

func main() {
	fmt.Println("is x64", unsafe.Sizeof(uintptr(0)) == 8)
	controls.App = new(windows.Provider).Init()
	controls.App.SetIcon("app.ico")

	var frm = new(controls.Form).Init()
	frm.SetTitle("this is form")
	frm.SetSize(600, 600)
	frm.SetLocation(100, 100)
	controls.Run(frm)
}
