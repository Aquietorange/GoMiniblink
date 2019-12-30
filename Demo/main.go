package main

import (
	"GoMiniblink"
	"GoMiniblink/CrossPlatform/Windows"
	"GoMiniblink/Forms"
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

	Forms.Run(frm)
}
