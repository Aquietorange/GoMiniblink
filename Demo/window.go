package main

import (
	cs "gitee.com/aochulai/GoMiniblink/forms/controls"
	gw "gitee.com/aochulai/GoMiniblink/forms/windows"
)

func main() {
	cs.App = new(gw.Provider).Init()
	cs.App.SetIcon("app.ico")

	frm := new(cs.Form).Init()
	frm.SetTitle("this is form")

	ctrl := new(cs.Control).Init()
	ctrl.SetBgColor(0x2FAEE3)

	frm.AddChild(ctrl)

	cs.Run(frm)
}
