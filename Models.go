package GoMiniblink

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

type MouseEvArgs struct {
	Buttons     MouseButtons
	X, Y, Delta int
	IsDBClick   bool
}
