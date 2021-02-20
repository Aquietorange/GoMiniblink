package main

import (
	gm "github.com/hujun528/GoMiniblink"
	br "github.com/hujun528/GoMiniblink/forms/bridge"
	cs "github.com/hujun528/GoMiniblink/forms/controls"
	gw "github.com/hujun528/GoMiniblink/forms/windows"
)

func main() {
	cs.App = new(gw.Provider).Init()

	frm := new(gm.MiniblinkForm).InitEx(br.FormParam{
		HideInTaskbar: true,
	})
	frm.TransparentMode()
	frm.SetLocation(100, 100)
	frm.SetSize(300, 300)
	//frm.SetTopMost(true)
	frm.View.ResourceLoader = append(frm.View.ResourceLoader, new(gm.FileLoader).Init("Res", "local"))
	frm.EvLoad["show"] = func(s cs.GUI) {
		frm.View.LoadUri("http://local/transparent.html")
	}
	cs.Run(&frm.Form)
}
