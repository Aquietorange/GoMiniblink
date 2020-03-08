package goMiniblink

import (
	"time"
)

//todo 应该全是接口

type RequestEvArgs interface {
	Url() string
	Method() string
	SetData([]byte)
	GetData() []byte
	SetCancel(b bool)
	IsCancel() bool
}

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
	Key      Keys
	Value    uintptr
	IsHandle bool
	IsSys    bool
}

type KeyPressEvArgs struct {
	KeyChar  string
	Value    uintptr
	IsHandle bool
	IsSys    bool
}
