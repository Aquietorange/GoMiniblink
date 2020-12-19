package main

import (
	gm "gitee.com/aochulai/GoMiniblink"
	cs "gitee.com/aochulai/GoMiniblink/forms/controls"
	gw "gitee.com/aochulai/GoMiniblink/forms/windows"
)

func main() {
	cs.App = new(gw.Provider).Init()
	cs.App.SetIcon("app.ico")

	frm := new(gm.MiniblinkForm).Init()
	frm.SetTitle("miniblink窗口")
	frm.SetLocation(100, 100)
	frm.SetSize(800, 500)
	frm.EvLoad["show"] = func(s cs.GUI) {
		frm.View.LoadUri("https://www.baidu.com")
	}
	cs.Run(&frm.Form)
}
