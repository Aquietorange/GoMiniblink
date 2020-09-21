package main

import "fmt"

func main() {
	rgba := 0x007D9A
	fmt.Printf("%X\n", rgba)
	fmt.Println(0x9A7D00)

	R := uint8(rgba)
	G := uint8(rgba >> 8)
	B := uint8(rgba >> 16)
	A := uint8(rgba >> 24)

	fmt.Printf("%X,%X,%X,%X\n", R, G, B, A)

	res := (int32(R) << 16) | (int32(G) << 8) | int32(B)

	fmt.Printf("%X", res)
}
