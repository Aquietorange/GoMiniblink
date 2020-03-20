package main

import (
	"fmt"
	"unsafe"
)

func main() {
	num := float64(3.333)
	p := uintptr(unsafe.Pointer(&num))
	low := *((*int32)(unsafe.Pointer(p)))
	high := *((*int32)(unsafe.Pointer(p + 4)))
	fmt.Println(low, high)

	var l = int64(high)<<32 + int64(low)
	d := *((*float64)(unsafe.Pointer(&l)))
	fmt.Println(d)

	//high := 2
	//low := 2
	//num := float64(high<<32) + float64(low)
	//fmt.Println(num)
}
