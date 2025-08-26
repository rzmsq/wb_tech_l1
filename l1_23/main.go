package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var userInput int
	_, err := fmt.Scan(&userInput)
	if err != nil {
		return
	}
	if userInput > 100 || userInput <= 0 {
		return
	}
	userInput--

	arr := make([]int, 100)
	for i := 0; i < 100; i++ {
		arr[i] = rand.Intn(100)
	}

	fmt.Println("срез перед удалением", len(arr), cap(arr), arr)

	newArr := make([]int, len(arr)-1)
	copy(newArr, arr[:userInput])
	copy(newArr[userInput:], arr[userInput+1:])
	arr = newArr

	fmt.Println("срез после удаления", len(arr), cap(arr), arr)
}
