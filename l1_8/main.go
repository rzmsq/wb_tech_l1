package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <N> <1 < i < 65>")
		return
	}

	var n int64
	var i int64
	var offset int64
	var err error

	n, err = strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		fmt.Println("invalid argument")
		return
	}

	i, err = strconv.ParseInt(os.Args[2], 10, 64)
	if err != nil || i < 1 || i > 64 {
		fmt.Println("invalid argument")
		return
	}

	fmt.Printf("Число %d в двоичной форме перед сменой %d бита: %b \n", n, i, n)

	i--
	offset = int64(1 << i) // 01 - сдвигаем единицу на нужную позицию i
	n ^= offset            // xor для смены бита i на противоположное значение

	fmt.Printf("Число %d в двоичной форме после смены %d бита: %b \n", n, i+1, n)
}
