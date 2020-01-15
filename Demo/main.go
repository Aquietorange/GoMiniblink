package main

import (
	"qq.2564874169/miniblink/forms"
	"qq.2564874169/miniblink/platform/windows"
)

func main() {
	forms.Provider = new(windows.Provider).Init()
	forms.Provider.SetIcon("app.ico")

	var frm = new(forms.Form).Init()
	frm.SetTitle("miniblink窗口")
	frm.SetSize(800, 500)
	frm.EvLoad["add_child"] = func(target interface{}) {
		if f, ok := target.(*forms.Form); ok {
			f.SetBgColor(0x000000)
		}
		//go func(f *forms.Form) {
		//	time.Sleep(3 * time.Second)
		//	f.Invoke(func(state interface{}) {
		//		f.ShowInTaskbar(false)
		//		//f.SetIconVisable(false)
		//	}, nil)
		//}(target.(*forms.Form))
		ctrl := new(forms.MiniblinkBrowser).Init()
		ctrl.SetSize(740, 420)
		ctrl.SetLocation(20, 20)
		ctrl.SetBgColor(0xCCCCCC)
		ctrl.EvLoad["loadUri"] = func(target interface{}) {
			ctrl.LoadUri("https://www.baidu.com")
		}
		frm.AddChild(ctrl)
	}
	forms.Run(frm)
}
