package main

import (
	"fmt"
	"math/big"
	"math/rand"
)

func main() {
	a, b := big.NewInt(int64(rand.Intn(1000000)*rand.Intn(10000))), big.NewInt(int64(rand.Intn(1000000)*rand.Intn(10000)))

	c := new(big.Int).Mul(a, b)
	fmt.Printf("%d * %d = %d\n", a, b, c)
	e := new(big.Int).Div(c, b)
	fmt.Printf("%d / %d = %d\n", c, b, e)
	f := new(big.Int).Sub(b, a)
	fmt.Printf("%d - %d = %d\n", b, a, f)
	d := new(big.Int).Add(c, a)
	fmt.Printf("%d + %d = %d\n", c, a, d)
}
