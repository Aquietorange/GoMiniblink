package free

import "unsafe"

func utf8PtrToString(ptr uintptr) string {
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
