package main

import (
	"fmt"
	"os"
)

func reverse(str []byte) {
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		str[i], str[j] = str[j], str[i]
	}
}

func reverseWords(str []byte) {
	reverse(str)

	start := 0
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			reverse(str[start:i])
			start = i + 1
		}
	}
	reverse(str[start:])
}

func main() {
	var seq []byte
	if len(os.Args) == 1 {
		seq = []byte("snow dog sun")
	} else {
		seq = []byte(os.Args[1])
	}

	fmt.Println("Исходная строка", string(seq))
	reverseWords(seq)
	fmt.Println("Результат", string(seq))
}
