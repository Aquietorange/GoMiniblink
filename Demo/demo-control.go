package main

import (
	"fmt"
	gm "gitee.com/aochulai/GoMiniblink"
	cs "gitee.com/aochulai/GoMiniblink/forms/controls"
	gw "gitee.com/aochulai/GoMiniblink/forms/windows"
	"unsafe"
)

func main() {
	fmt.Println("x64 is", unsafe.Sizeof(uintptr(0)) == 8)
	cs.App = new(gw.Provider).Init()
	cs.App.SetIcon("app.ico")

	frm := new(cs.Form).Init()
	frm.SetTitle("普通窗口")
	frm.SetSize(800, 500)
	frm.SetLocation(100, 100)
	frm.SetBgColor(0x2FAEE3)

	mb := new(gm.MiniblinkBrowser).Init()
	mb.SetBgColor(0x2FAEE3)
	mb.SetSize(700, 400)
	mb.SetLocation(50, 50)
	frm.AddChild(mb)

	frm.EvLoad["show"] = func(s cs.GUI) {
		mb.LoadUri("https://www.baidu.com")
	}
	cs.Run(frm)
}
