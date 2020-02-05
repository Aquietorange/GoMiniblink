package miniblink

import "image"

type Point struct {
	X, Y int
}

func (_this Point) IsEqual(point Point) bool {
	return _this.X == point.X && _this.Y == point.Y
}

type Rect struct {
	Wdith, Height int
}

func (_this Rect) IsEqual(rect Rect) bool {
	return _this.Wdith == rect.Wdith && _this.Height == rect.Height
}

func (_this Rect) IsEmpty() bool {
	return _this.Wdith == 0 || _this.Height == 0
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
	DrawImage(src *image.RGBA, srcXY Point, rect Rect, toXY Point)
	Close()
}
