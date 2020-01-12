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
	frm.SetSize(800, 500)
	frm.EvLoad["load"] = func(target interface{}) {
		if f, ok := target.(*Forms.Form); ok {
			f.SetBgColor(0x000000)
		}
		//go func(f *Forms.Form) {
		//	time.Sleep(3 * time.Second)
		//	f.Invoke(func(state interface{}) {
		//		f.ShowInTaskbar(false)
		//		//f.SetIconVisable(false)
		//	}, nil)
		//}(target.(*Forms.Form))
		ctrl := new(Forms.MiniblinkBrowser).Init()
		ctrl.SetSize(740, 420)
		ctrl.SetLocation(20, 20)
		ctrl.SetBgColor(0xCCCCCC)
		ctrl.EvLoad["loadUri"] = func(target interface{}) {
			ctrl.LoadUri("https://www.baidu.com")
		}
		frm.AddChild(ctrl)
	}
	Forms.Run(frm)
}
