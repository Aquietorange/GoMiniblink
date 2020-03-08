package free

import (
	mb "qq2564874169/goMiniblink"
	"unsafe"
)

func isExtKey(key mb.Keys) bool {
	switch key {
	case mb.Keys_Insert, mb.Keys_Delete, mb.Keys_Home, mb.Keys_End, mb.Keys_PageUp,
		mb.Keys_PageDown, mb.Keys_Left, mb.Keys_Right, mb.Keys_Up, mb.Keys_Down:
		return true
	default:
		return false
	}
}

func wkePtrToUtf8(ptr uintptr) string {
	var seq []byte
	for {
		b := *((*byte)(unsafe.Pointer(ptr)))
		if b != 0 {
			seq = append(seq, b)
			ptr++
		} else {
			break
		}
	}
	return string(seq)
}
