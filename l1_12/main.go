package main

import "fmt"

func main() {
	seq := [...]string{"cat", "cat", "dog", "cat", "tree"}
	uniqSeq := make(map[string]interface{})

	for _, v := range seq {
		uniqSeq[v] = nil
	}

	fmt.Println("Исходная последовательность", seq)

	fmt.Print("Последовательность уникальных вхождений ")
	for k, _ := range uniqSeq {
		fmt.Print(k, " ")
	}

}
