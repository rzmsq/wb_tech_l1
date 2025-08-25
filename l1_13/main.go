package main

import "fmt"

func main() {
	a, b := 10, 50

	fmt.Printf("a=%d, b=%d\n", a, b)
	// xor swap
	a ^= b
	b ^= a
	a ^= b

	fmt.Printf("a=%d, b=%d", a, b)
}
