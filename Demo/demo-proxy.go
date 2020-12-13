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

	frm := new(gm.MiniblinkForm).Init()
	frm.SetTitle("miniblink窗口")
	frm.SetLocation(100, 100)
	frm.SetSize(800, 500)
	frm.View.SetProxy(gm.ProxyInfo{
		Type:     gm.ProxyType_HTTP,
		HostName: "127.0.0.1",
		Port:     58591,
	})
	frm.EvLoad["show"] = func(s cs.GUI) {
		frm.View.LoadUri("https://www.ip.cn")
	}
	cs.Run(&frm.Form)
}
