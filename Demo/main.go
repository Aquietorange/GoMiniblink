package main

import (
	"qq2564874169/goMiniblink/forms/controls"
	"qq2564874169/goMiniblink/forms/platform/windows"
)

func main() {
	//src := 0xDB7093
	//s := miniblink.IntToRGBA(src)
	//println(uint32(src))
	//r, g, b, _ := s.RGBA()
	//println(r, g, b)
	//rgb := uint32(uint8(r)) | (uint32(uint8(g)) << 8) | uint32(uint8(b))<<16
	//println(rgb)
	controls.App = new(windows.Provider).Init()
	controls.App.SetIcon("app.ico")
	controls.App.SetBgColor(0x00FF)

	var frm = new(controls.Form).Init()
	frm.SetTitle("miniblink窗口")
	frm.SetSize(800, 500)
	frm.EvLoad["add_child"] = func(target interface{}) {
		c := new(controls.Control).Init()
		c.SetSize(100, 100)
		c.SetLocation(15, 15)
		c.SetBgColor(0xCCCCCC)
		frm.AddChild(c)
	}
	controls.Run(frm)
}
