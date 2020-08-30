package main

import (
	"fmt"
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
		mb.EvLoad["bind_go_func"] = func(target interface{}) {
			mb.EvConsole = append(mb.EvConsole, func(sender *g.MiniblinkBrowser, e g.ConsoleEvArgs) {
				fmt.Println(e.Message())
			})
			mb.BindJsFunc("Func1", func(context g.GoFnContext) interface{} {
				n1 := context.Param[0].(float64)
				n2 := context.Param[1].(float64)
				return n1 * n2
			}, nil)
			mb.BindJsFunc("Func5", func(context g.GoFnContext) interface{} {
				return func(name string) string {
					return "姓名是：" + name
				}
			}, nil)
		}

		//g.JsFunc(g.JsFnBinding{
		//	Name: "Func2",
		//	Fn: func(context g.GoFnContext) interface{} {
		//		fn := context.Param[0].(g.JsFunc)
		//		fn(1.2, 3.4)
		//		return nil
		//	},
		//})
		//g.JsFunc(g.JsFnBinding{
		//	Name: "Func3",
		//	Fn: func(context g.GoFnContext) interface{} {
		//		data := context.Param[0].(map[string]interface{})
		//		n1 := data["n1"].(float64)
		//		n2 := data["n2"].(float64)
		//		return n1 * n2
		//	},
		//})
		//g.JsFunc(g.JsFnBinding{
		//	Name: "Func5",
		//	Fn: func(context g.GoFnContext) interface{} {
		//		return func(name string) string {
		//			return "姓名：" + name
		//		}
		//	},
		//})
		mb.LoadUri("https://local/js_call_net.html")
		frm.AddChild(mb)
	}
	controls.Run(frm)
}
