package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	var result strings.Builder
	for i := 0; i < length; i++ {
		result.WriteByte(charset[rand.Intn(len(charset))])
	}
	return result.String()
}

func isUnique(str string) bool {
	str = strings.ToLower(str)
	m := make(map[rune]bool)

	for _, v := range str {
		if _, ok := m[v]; ok {
			return false
		}
		m[v] = true
	}
	return true
}

func main() {

	randomStr := generateRandomString(10)
	fmt.Println(randomStr)

	if isUnique(randomStr) {
		fmt.Println("Строка состоит из уникальных символов")
	} else {
		fmt.Println("Строка содержит повторяющиеся символы")
	}
}
