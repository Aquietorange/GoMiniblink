package main

import (
	gm "gitee.com/aochulai/GoMiniblink"
	br "gitee.com/aochulai/GoMiniblink/forms/bridge"
	cs "gitee.com/aochulai/GoMiniblink/forms/controls"
	gw "gitee.com/aochulai/GoMiniblink/forms/windows"
)

func main() {
	cs.App = new(gw.Provider).Init()

	frm := new(gm.MiniblinkForm).InitEx(br.FormParam{
		HideInTaskbar: true,
	})
	frm.TransparentMode()
	frm.SetLocation(100, 100)
	frm.SetSize(300, 300)
	frm.SetTopMost(true)
	frm.View.ResourceLoader = append(frm.View.ResourceLoader, new(gm.FileLoader).Init("Res", "local"))
	frm.EvLoad["show"] = func(s cs.GUI) {
		frm.View.LoadUri("http://local/transparent.html")
	}
	cs.Run(&frm.Form)
}
