package miniblink

import "time"

type MouseEvArgs struct {
	ButtonIsDown map[MouseButtons]bool
	X, Y, Delta  int
	IsDouble     bool
	Time         time.Time
}

type PaintEvArgs struct {
	Update  Bound
	Context uintptr
}

type KeyEvArgs struct {
	Key        Keys
	KeysIsDown map[Keys]bool
	IsHandle   bool
}

type KeyPressEvArgs struct {
	KeyChar    string
	KeysIsDown map[Keys]bool
	IsHandle   bool
}
