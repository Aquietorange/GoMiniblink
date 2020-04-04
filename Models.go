package goMiniblink

import (
	"image"
)

type JsFunc func(param ...interface{}) interface{}

type GoFuncContext struct {
	Name  string
	State interface{}
	Param []interface{}
}

type GoFuncFn func(context GoFuncContext) interface{}

type GoFunc struct {
	Name  string
	State interface{}
	Fn    GoFuncFn
}

func (_this *GoFunc) OnExecute(param []interface{}) interface{} {
	return _this.Fn(GoFuncContext{
		Name:  _this.Name,
		State: _this.State,
		Param: param,
	})
}

type Point struct {
	X, Y int
}

func (_this Point) IsEqual(point Point) bool {
	return _this.X == point.X && _this.Y == point.Y
}

type Rect struct {
	Width, Height int
}

func (_this Rect) IsEqual(rect Rect) bool {
	return _this.Width == rect.Width && _this.Height == rect.Height
}

func (_this Rect) IsEmpty() bool {
	return _this.Width == 0 || _this.Height == 0
}

type Bound struct {
	Point
	Rect
}

type Bound2 struct {
	Left   int
	Top    int
	Right  int
	Bottom int
}

type Screen struct {
	Full     Rect
	WorkArea Rect
}

type Graphics interface {
	DrawImage(src *image.RGBA, xSrc, ySrc, width, height, xDst, yDst int) Graphics
	Close()
}
