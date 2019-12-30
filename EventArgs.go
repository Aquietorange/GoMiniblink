package GoMiniblink

import "time"

type MouseEvArgs struct {
	Buttons     MouseButtons
	X, Y, Delta int
	IsDouble    bool
	Time        time.Time
}
