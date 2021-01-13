package main

import (
	"fmt"
	gm "gitee.com/aochulai/GoMiniblink"
	fm "gitee.com/aochulai/GoMiniblink/forms"
	cs "gitee.com/aochulai/GoMiniblink/forms/controls"
	gw "gitee.com/aochulai/GoMiniblink/forms/windows"
)

func main() {
	cs.App = new(gw.Provider).Init()
	cs.App.SetIcon("app.ico")

	frm := new(cs.Form).Init()
	frm.SetTitle("JS互操作")
	frm.SetSize(800, 500)

	mb := new(gm.MiniblinkBrowser).Init()
	mb.SetAnchor(fm.AnchorStyle_Fill)
	mb.ResourceLoader = append(mb.ResourceLoader, new(gm.FileLoader).Init("Res", "local"))
	mb.EvConsole["show"] = func(_ *gm.MiniblinkBrowser, e gm.ConsoleEvArgs) {
		fmt.Println("js console:", e.Message())
	}
	mb.EvDocumentReady["exec"] = func(s *gm.MiniblinkBrowser, e gm.DocumentReadyEvArgs) {
		//调用func_1
		mb.CallJsFunc("func_1", "张三", 18)

		//获取func_2返回的基础数据类型
		f2rs := mb.CallJsFunc("func_2")
		fmt.Println("func_2 result is", f2rs)

		//向func_3传递一个go函数
		mb.CallJsFunc("func_3", func(n1, n2 float64) int {
			//此结果会在js中打印
			return int(n1) * int(n2)
		})

		//获取func_4返回的非基本数据类型
		f4rs := mb.CallJsFunc("func_4").(map[string]interface{})
		fmt.Println("func_4 result is", f4rs)

		//获取并调用func_5返回的js函数
		fn := mb.CallJsFunc("func_5").(gm.JsFunc)
		fn("王老五")
	}
	frm.AddChild(mb)
	frm.EvLoad["show"] = func(s cs.GUI) {
		mb.LoadUri("https://local/call_js.html")
	}
	cs.Run(frm)
}
