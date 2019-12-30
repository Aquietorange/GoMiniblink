package GoMiniblink

type MouseEvArgs struct {
	Buttons     MouseButtons
	X, Y, Delta int
	IsDBClick   bool
}
