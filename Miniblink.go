package GoMiniblink

import (
	"image"
)

type ResponseCallback func(args ResponseEvArgs)
type RequestBeforeCallback func(args RequestBeforeEvArgs)
type JsReadyCallback func(args JsReadyEvArgs)
type ConsoleCallback func(args ConsoleEvArgs)
type DocumentReadyCallback func(args DocumentReadyEvArgs)
type PaintUpdatedCallback func(args PaintUpdatedEvArgs)

type Miniblink interface {
	SetBmpPaintMode(b bool)
	SetProxy(info ProxyInfo)
	MouseIsEnable() bool
	MouseEnable(b bool)
	ToBitmap() *image.RGBA
	CallJsFunc(name string, param []interface{}) interface{}
	JsFunc(name string, fn GoFn, state interface{})
	RunJs(script string) interface{}
	SetOnConsole(callback ConsoleCallback)
	SetOnJsReady(callback JsReadyCallback)
	SetOnRequestBefore(callback RequestBeforeCallback)
	SetOnDocumentReady(callback DocumentReadyCallback)
	SetOnPaintUpdated(callback PaintUpdatedCallback)
	LoadUri(uri string)
	GetHandle() wkeHandle
}
