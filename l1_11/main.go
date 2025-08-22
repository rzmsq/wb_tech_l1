package main

import (
	"fmt"
	"math/rand"
)

// Intersection Функция проверки пересечения
func Intersection(a, b []int) []int {
	// Наименьший слайс по длине всегда в a
	if len(a) > len(b) {
		a, b = b, a
	}
	// map для хранения вхождений значений слайса a
	m := make(map[int]bool, len(a))
	for _, v := range a {
		m[v] = true
	}

	// Результирующий слайс
	result := make([]int, 0)
	for _, v := range b {
		if m[v] {
			result = append(result, v)
			m[v] = false
		}
	}
	return result
}

func main() {

	// заполнения двух слайсов
	a, b := make([]int, rand.Intn(100)), make([]int, rand.Intn(100))

	for i := range a {
		a[i] = rand.Intn(1000)
	}
	fmt.Println("Слайс a:", a)
	fmt.Println()
	for i := range b {
		b[i] = rand.Intn(1000)
	}
	fmt.Println("Слайс b:", a)
	fmt.Println()

	// результат
	fmt.Println("Пересечение слайсов:", Intersection(a, b))

}
