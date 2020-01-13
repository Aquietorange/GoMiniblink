package miniblink

type Point struct {
	X, Y int
}

type Rect struct {
	Wdith, Height int
}

type Bound struct {
	Point
	Rect
}

type Screen struct {
	Full     Rect
	WorkArea Rect
}
