package main

import (
	"fmt"
	g "qq2564874169/goMiniblink"
	"qq2564874169/goMiniblink/forms"
	"qq2564874169/goMiniblink/forms/controls"
	"qq2564874169/goMiniblink/forms/windows"
	"unsafe"
)

func main() {
	fmt.Println("is x64", unsafe.Sizeof(uintptr(0)) == 8)
	controls.App = new(windows.Provider).Init()
	controls.App.SetIcon("app.ico")

	var frm = new(g.MiniblinkForm).Init()
	frm.View.EvShow["init"] = func(target interface{}) {
		frm.SetTitle("miniblink窗口")
		frm.SetStartPosition(forms.FormStartPosition_Manual)
		frm.SetLocation(100, 100)
		frm.SetBorderStyle(forms.FormBorder_None)
		frm.NoneBorderResize()
		frm.SetSize(800, 500)

		frm.View.EvConsole["show"] = func(_ *g.MiniblinkBrowser, e g.ConsoleEvArgs) {
			fmt.Println(e.Message())
		}
		frm.View.ResourceLoader = append(frm.View.ResourceLoader, new(g.FileLoader).Init("Res", "local"))
		frm.View.LoadUri("http://local/window.html")
	}
	controls.Run(&frm.Form)
}
