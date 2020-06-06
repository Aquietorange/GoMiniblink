package main

import (
	g "qq2564874169/goMiniblink"
	f "qq2564874169/goMiniblink/forms"
	"qq2564874169/goMiniblink/forms/controls"
	"qq2564874169/goMiniblink/forms/platform/windows"
)

func main() {
	controls.App = new(windows.Provider).Init()
	controls.App.SetIcon("app.ico")
	controls.App.SetBgColor(0x00FF)

	var frm = new(controls.Form).Init()
	frm.SetTitle("miniblink窗口")
	frm.SetSize(800, 500)
	frm.EvLoad["add_child"] = func(target interface{}) {
		mb := new(g.MiniblinkBrowser).Init()
		mb.SetSize(750, 425)
		mb.SetLocation(15, 15)
		mb.SetAnchor(f.AnchorStyle_Top | f.AnchorStyle_Right | f.AnchorStyle_Bottom | f.AnchorStyle_Left)
		mb.ResourceLoader = append(mb.ResourceLoader, new(g.FileLoader).Init("Res", "local"))
		g.BindFunc(g.GoFunc{
			Name: "Func1",
			Func: func(context g.GoFuncContext) interface{} {
				n1 := context.Param[0].(float64)
				n2 := context.Param[1].(float64)
				return n1 * n2
			},
		})
		g.BindFunc(g.GoFunc{
			Name: "Func2",
			Func: func(context g.GoFuncContext) interface{} {
				fn := context.Param[0].(g.JsFunc)
				fn(1.2, 3.4)
				return nil
			},
		})
		g.BindFunc(g.GoFunc{
			Name: "Func3",
			Func: func(context g.GoFuncContext) interface{} {
				data := context.Param[0].(map[string]interface{})
				n1 := data["n1"].(float64)
				n2 := data["n2"].(float64)
				return n1 * n2
			},
		})
		g.BindFunc(g.GoFunc{
			Name: "Func5",
			Func: func(context g.GoFuncContext) interface{} {
				return func(name string) string {
					return "姓名：" + name
				}
			},
		})
		mb.LoadUri("https://local/js_call_net.html")
		frm.AddChild(mb)
	}
	controls.Run(frm)
}
