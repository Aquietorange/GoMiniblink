package main

import "fmt"

type A interface {
	Run()
}

type B interface {
	Stop()
}

type C struct {
}

func (_this C) Run() {
	fmt.Println("running")
}
func (_this C) Stop() {
	fmt.Println("stoping")
}

func main() {
	c := C{}
	var a A = c

	if b, ok := a.(B); ok {
		b.Stop()
	}
}
