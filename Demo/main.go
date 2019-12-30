package main

import (
	"GoMiniblink"
	"GoMiniblink/CrossPlatform/Windows"
	"GoMiniblink/Forms"
	"time"
)

func main() {
	Forms.Provider = new(Windows.Provider).Init()
	Forms.Provider.SetIcon("app.ico")

	var frm = new(Forms.Form).Init()
	frm.SetTitle("miniblink窗口")
	frm.SetSize(GoMiniblink.Rect{Wdith: 300, Height: 500})
	frm.EvLoad["load"] = func(target interface{}) {
		//go func(f *Forms.Form) {
		//	time.Sleep(5 * time.Second)
		//	f.Invoke(func(state interface{}) {
		//		frm.SetMinimizeBox(false)
		//	}, nil)
		//}(target.(*Forms.Form))
	}
	frm.EvMouseClick["show"] = func(target interface{}, e GoMiniblink.MouseEvArgs) {
		println("click\t", time.Now().Format("20060102150405000"), e.IsDBClick)
	}
	frm.EvMouseUp["show"] = func(target interface{}, e GoMiniblink.MouseEvArgs) {
		println("up\t\t", time.Now().Format("20060102150405000"))
	}

	Forms.Run(frm)
}
