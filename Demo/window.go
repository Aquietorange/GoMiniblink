package main

import (
	"fmt"
	fm "gitee.com/aochulai/GoMiniblink/forms"
	cs "gitee.com/aochulai/GoMiniblink/forms/controls"
	gw "gitee.com/aochulai/GoMiniblink/forms/windows"
)

func main() {
	cs.App = new(gw.Provider).Init()
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
