package forms

import (
	"image"
)

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
	GetHandle() uintptr
	DrawImage(src *image.RGBA, xSrc, ySrc, width, height, xDst, yDst int) Graphics
	Close()
}

type MsgBoxParam struct {
	Title  string
	Text   string
	Icon   MsgBoxIcon
	Button MsgBoxButton
}
