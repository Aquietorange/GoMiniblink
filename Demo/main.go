package main

import (
	"qq.2564874169/goMiniblink/forms"
	"qq.2564874169/goMiniblink/platform/windows"
)

func main() {
	//src := 0xDB7093
	//s := miniblink.IntToRGBA(src)
	//println(uint32(src))
	//r, g, b, _ := s.RGBA()
	//println(r, g, b)
	//rgb := uint32(uint8(r)) | (uint32(uint8(g)) << 8) | uint32(uint8(b))<<16
	//println(rgb)
	forms.Provider = new(windows.Provider).Init()
	forms.Provider.SetIcon("app.ico")
	forms.Provider.SetBgColor(0x00FF)

	var frm = new(forms.Form).Init()
	frm.SetTitle("miniblink窗口")
	frm.SetSize(800, 500)
	frm.EvLoad["add_child"] = func(target interface{}) {
		//go func(f *forms.Form) {
		//	time.Sleep(3 * time.Second)
		//	f.Invoke(func(state interface{}) {
		//		f.ShowInTaskbar(false)
		//		//f.SetIconVisable(false)
		//	}, nil)
		//}(target.(*forms.Form))
		//ctrl := new(forms.MiniblinkBrowser).Init()
		//ctrl.SetSize(740, 420)
		//ctrl.SetLocation(20, 20)
		//ctrl.SetBgColor(0xFF0000)
		//ctrl.SetAnchor(goMiniblink.AnchorStyle_Right | goMiniblink.AnchorStyle_Bottom | goMiniblink.AnchorStyle_Top | goMiniblink.AnchorStyle_Left)
		//ctrl.EvLoad["loadUri"] = func(target interface{}) {
		//	ctrl.LoadUri("https://me.csdn.net/iamshuke")
		//}
		//frm.AddChild(ctrl)
	}
	forms.Run(frm)
}
