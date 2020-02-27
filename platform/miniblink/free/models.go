package free

type wkeRect struct {
	x, y, w, h int32
}

type wkeNetJob uintptr

type wkeMouseFlags int

const (
	wkeMouseFlags_None    wkeMouseFlags = 0
	wkeMouseFlags_LBUTTON wkeMouseFlags = 0x01
	wkeMouseFlags_RBUTTON wkeMouseFlags = 0x02
	wkeMouseFlags_SHIFT   wkeMouseFlags = 0x04
	wkeMouseFlags_CONTROL wkeMouseFlags = 0x08
	wkeMouseFlags_MBUTTON wkeMouseFlags = 0x10
)
