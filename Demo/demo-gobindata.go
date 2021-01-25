package main

import (
	gm "gitee.com/aochulai/GoMiniblink"
	"gitee.com/aochulai/GoMiniblink/Demo/Res"
	fm "gitee.com/aochulai/GoMiniblink/forms"
	cs "gitee.com/aochulai/GoMiniblink/forms/controls"
	gw "gitee.com/aochulai/GoMiniblink/forms/windows"
	url2 "net/url"
	"strings"
)

func main() {
	cs.App = new(gw.Provider).Init()
	cs.App.SetIcon("app.ico")

	frm := new(gm.MiniblinkForm).Init()
	frm.SetTitle("加载内嵌的静态资源")
	frm.SetSize(800, 500)
	frm.SetStartPosition(fm.FormStart_Screen_Center)
	frm.View.ResourceLoader = append(frm.View.ResourceLoader, new(GobindataLoader))
	frm.EvLoad["show"] = func(s cs.GUI) {
		frm.View.LoadUri("http://local/gobindata.html")
	}
	cs.Run(&frm.Form)
}

type GobindataLoader struct {
}

func (_this *GobindataLoader) Domain() string {
	return "local"
}

func (_this *GobindataLoader) ByUri(uri *url2.URL) []byte {
	//Res.Asset是用调试模式生成的，里面的文件路径是全路径，所以在别的电脑上要重新生成一份，或者自己全局替换一下路径目录
	rs, err := Res.Asset(strings.TrimLeft(uri.Path, "/"))
	if err != nil {
		return nil
	}
	return rs
}
