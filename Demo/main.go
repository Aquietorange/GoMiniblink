package main

import (
	"GoMiniblink/Forms"
	"GoMiniblink/Forms/CrossPlatform/Windows"
)

func main() {
	Forms.Provider = new(Windows.Provider).Init()
	var frm = new(Forms.Form).Init()
	frm.SetSize(300, 500)
	frm.SetLocation(100, 100)
	frm.SetTitle("miniblink窗口")
	frm.ShowInTaskbar(false)
	//frm.SetBorderStyle(Forms.FormBorder_Disable_Resize)
	frm.EvLoad["load"] = func(form *Forms.Form) {
		//frm.SetBorderStyle(Forms.FormBorder_None)
		//go func(f *Forms.Form) {
		//	time.Sleep(5 * time.Second)
		//	f.Invoke(func(state interface{}) {
		//		frm.SetBorderStyle(Forms.FormBorder_None)
		//	}, nil)
		//}(form)
	}
	Forms.Run(frm)
}
