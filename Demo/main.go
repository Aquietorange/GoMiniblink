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
	//frm.SetBorderStyle(Forms.FormBorder_None)
	frm.EvLoad = append(frm.EvLoad, func(form *Forms.Form) {
		//go func(f *Forms.Form) {
		//	time.Sleep(2 * time.Second)
		//	f.Invoke(func(state interface{}) {
		//		frm.ShowInTaskbar(false)
		//	}, nil)
		//	time.Sleep(2 * time.Second)
		//	f.Invoke(func(state interface{}) {
		//		frm.Show()
		//	}, nil)
		//}(form)
	})
	Forms.Run(frm)
}
