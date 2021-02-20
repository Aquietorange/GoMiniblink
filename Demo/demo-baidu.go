package main

import (
	gm "../GoMiniblink"
	fm "../GoMiniblink/forms"

	cs "../GoMiniblink/forms/controls"
	gw "../GoMiniblink/forms/windows"
)

func main() {
	cs.App = new(gw.Provider).Init()
	cs.App.SetIcon("app.ico")

	frm := new(gm.MiniblinkForm).Init()
	frm.SetTitle("miniblink窗口")
	frm.SetSize(800, 500)
	frm.SetStartPosition(fm.FormStart_Screen_Center)
	frm.EvLoad["加载网址"] = func(s cs.GUI) {
		frm.View.LoadUri("https://www.baidu.com")
	}

	cs.Run(&frm.Form)
}
