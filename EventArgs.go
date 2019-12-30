package GoMiniblink

import "time"

type MouseEvArgs struct {
	Buttons     MouseButtons
	X, Y, Delta int
	IsDBClick   bool
	Time        time.Time
}
