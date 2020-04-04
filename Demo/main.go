package main

import (
	"qq2564874169/goMiniblink"
	"qq2564874169/goMiniblink/forms"
	"qq2564874169/goMiniblink/platform/windows"
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
		mb := new(forms.MiniblinkBrowser).Init()
		mb.SetSize(740, 420)
		mb.SetLocation(20, 20)
		mb.SetAnchor(goMiniblink.AnchorStyle_Right | goMiniblink.AnchorStyle_Bottom | goMiniblink.AnchorStyle_Top | goMiniblink.AnchorStyle_Left)
		resDir := forms.Provider.AppDir() + "\\Res"
		mb.ResourceLoader = append(mb.ResourceLoader, new(forms.FileLoader).Init(resDir, "loc.res"))
		mb.BindFunc("Func1", func(context goMiniblink.GoFuncContext) interface{} {
			n1 := context.Param[0].(float64)
			n2 := context.Param[1].(float64)
			return int(n1) * int(n2)
		}, nil)
		mb.LoadUri("http://loc.res/js_call_net.html")
		//mb.LoadUri("https://www.baidu.com")
		frm.AddChild(mb)
	}
	forms.Run(frm)
}
