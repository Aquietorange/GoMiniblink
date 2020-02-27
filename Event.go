package goMiniblink

import (
	"time"
)

//todo 应该全是接口
type MouseEvArgs struct {
	Button      MouseButtons
	X, Y, Delta int
	IsDouble    bool
	Time        time.Time
}

type PaintEvArgs struct {
	Clip     Bound
	Graphics Graphics
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
