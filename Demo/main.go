package main

import (
	"GoMiniblink/CrossPlatform/Windows"
	"GoMiniblink/Forms"
)

func main() {

	Forms.Provider = new(Windows.Provider).Init()
	Forms.Provider.SetIcon("app.ico")

	var frm = new(Forms.Form).Init()
	frm.SetTitle("miniblink窗口")
	frm.SetSize(300, 500)
	frm.EvLoad["load"] = func(target interface{}) {
		//go func(f *Forms.Form) {
		//	time.Sleep(3 * time.Second)
		//	f.Invoke(func(state interface{}) {
		//		f.ShowInTaskbar(false)
		//		//f.SetIconVisable(false)
		//	}, nil)
		//}(target.(*Forms.Form))
		ctrl := new(Forms.Control).Init()
		ctrl.SetSize(100, 100)
		ctrl.SetLocation(20, 20)
		ctrl.SetBgColor(0xCCCCCC)
		frm.AddControl(ctrl)
	}

	Forms.Run(frm)
}
