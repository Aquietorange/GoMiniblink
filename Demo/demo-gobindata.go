package main

import (
	url2 "net/url"
	"strings"

	gm "github.com/hujun528/GoMiniblink"
	"github.com/hujun528/GoMiniblink/Demo/Res"
	fm "github.com/hujun528/GoMiniblink/forms"
	cs "github.com/hujun528/GoMiniblink/forms/controls"
	gw "github.com/hujun528/GoMiniblink/forms/windows"
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
	rs, err := Res.Asset(strings.TrimLeft(uri.Path, "/"))
	if err != nil {
		return nil
	}
	return rs
}
