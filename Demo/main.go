package main

import (
	"fmt"
	m "qq2564874169/goMiniblink"
	"qq2564874169/goMiniblink/forms"
	"qq2564874169/goMiniblink/platform/windows"
	"unsafe"
)

func main() {
	//src := 0xDB7093
	//s := miniblink.IntToRGBA(src)
	//println(uint32(src))
	//r, g, b, _ := s.RGBA()
	//println(r, g, b)
	//rgb := uint32(uint8(r)) | (uint32(uint8(g)) << 8) | uint32(uint8(b))<<16
	//println(rgb)
	fmt.Println("is x64 ：", unsafe.Sizeof(uintptr(0)) == 8)
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
		mb.SetAnchor(m.AnchorStyle_Right | m.AnchorStyle_Bottom | m.AnchorStyle_Top | m.AnchorStyle_Left)
		resDir := forms.Provider.AppDir() + "\\Res"
		mb.ResourceLoader = append(mb.ResourceLoader, new(forms.FileLoader).Init(resDir, "loc.res"))
		mb.BindJsFunc("Func1", func(context m.GoFnContext) interface{} {
			n1 := context.Param[0].(float64)
			n2 := context.Param[1].(float64)
			return int(n1) * int(n2)
		}, nil)
		mb.BindJsFunc("Func2", func(context m.GoFnContext) interface{} {
			fn := context.Param[0].(m.JsFunc)
			fn(5, 6)
			return nil
		}, nil)
		mb.BindJsFunc("Func3", func(context m.GoFnContext) interface{} {
			rs := context.Param[0].(map[string]interface{})
			n1 := rs["n1"].(float64)
			n2 := rs["n2"].(float64)
			return int(n1) * int(n2)
		}, nil)
		mb.BindJsFunc("Func5", func(context m.GoFnContext) interface{} {
			return func(name string) string {
				return "姓名是：" + name
			}
		}, nil)
		mb.LoadUri("http://loc.res/js_call_net.html")
		//mb.LoadUri("https://www.baidu.com")
		frm.AddChild(mb)
	}
	forms.Run(frm)
}
