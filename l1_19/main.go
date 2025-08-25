package main

import "fmt"

func reverseStr(str []rune) []rune {
	l, r := 0, len(str)-1
	for l < r {
		str[l], str[r] = str[r], str[l]
		l++
		r--
	}
	return str
}

func main() {

	str := []rune("Hello Мир😁")

	fmt.Println(string(reverseStr(str)))

}
