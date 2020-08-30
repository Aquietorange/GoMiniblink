package main

import (
	"fmt"
	g "qq2564874169/goMiniblink"
	f "qq2564874169/goMiniblink/forms"
	"qq2564874169/goMiniblink/forms/controls"
	"qq2564874169/goMiniblink/forms/platform/windows"
	"unsafe"
)

func main() {
	fmt.Println("is x64", unsafe.Sizeof(uintptr(0)) == 8)
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
		mb.EvLoad["bind_go_func"] = func(_ interface{}) {
			mb.EvConsole["show"] = func(_ *g.MiniblinkBrowser, e g.ConsoleEvArgs) {
				fmt.Println(e.Message())
			}
			mb.JsFunc("Func1", func(context g.GoFnContext) interface{} {
				n1 := context.Param[0].(float64)
				n2 := context.Param[1].(float64)
				return n1 * n2
			}, nil)
			mb.JsFunc("Func2", func(context g.GoFnContext) interface{} {
				fn := context.Param[0].(g.JsFunc)
				return fn(5, 6)
			}, nil)
			mb.JsFunc("Func3", func(context g.GoFnContext) interface{} {
				data := context.Param[0].(map[string]interface{})
				n1 := data["n1"].(float64)
				n2 := data["n2"].(float64)
				return n1 * n2
			}, nil)
			mb.JsFunc("Func5", func(context g.GoFnContext) interface{} {
				return func(name string) string {
					return "姓名是：" + name
				}
			}, nil)
		}
		mb.LoadUri("https://local/js_call_net.html")
		frm.AddChild(mb)
	}
	controls.Run(frm)
}
