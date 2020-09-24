package main

import (
	"fmt"
	"time"
)

func main() {
	var inv uint32 = 5000
	fmt.Println(time.Now())
	time.Sleep(time.Duration(time.Millisecond.Nanoseconds() * int64(inv)))
	fmt.Println(time.Now())
}
