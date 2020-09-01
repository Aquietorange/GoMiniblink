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
			mb.EvConsole["show_msg"] = func(_ interface{}, e g.ConsoleEvArgs) {
				fmt.Println(e.Message())
			}
			mb.JsFuncEx("Func1", func(n1, n2 float64) int {
				return int(n1 * n2)
			})
			mb.JsFuncEx("Func2", func(fn g.JsFunc) {
				fn(5, 6)
			})
			mb.JsFuncEx("Func3", func(param map[string]interface{}) interface{} {
				rs := param["n1"].(float64) * param["n2"].(float64)
				return struct {
					Msg   string
					Value int
				}{
					Msg:   "n1*n2",
					Value: int(rs),
				}
			})
			mb.JsFuncEx("Func5", func() interface{} {
				return func(name string) string {
					return "姓名是：" + name
				}
			})
			mb.LoadUri("https://local/js_call_net.html")
		}
		frm.AddChild(mb)
	}
	controls.Run(frm)
}
