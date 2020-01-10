package Core

import (
	"GoMiniblink/Core/free"
	"GoMiniblink/Forms"
)

type MiniblinkBrowser struct {
	Forms.Control

	miniblink IMiniblink
}

func (_this *MiniblinkBrowser) Init() *MiniblinkBrowser {
	_this.Control.Init()
	_this.EvLoad["__initWke"] = _this.initWke
	return _this
}

func (_this *MiniblinkBrowser) LoadUri(uri string) {
	_this.miniblink.LoadUri(uri)
}

func (_this *MiniblinkBrowser) initWke(target interface{}) {
	defer func() {
		recover()
	}()
	_this.miniblink = new(free.Miniblink).Init(_this.Handle)
}
