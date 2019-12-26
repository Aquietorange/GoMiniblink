package Windows

import (
	"syscall"
)

func sto16(str string) *uint16 {
	ptr, _ := syscall.UTF16PtrFromString(str)
	return ptr
}
