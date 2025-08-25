package main

import (
	"fmt"
	"math/rand"
)

// quickSort функция обертка для упрощения вызова
func quickSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}
	return quickSortRecursive(array, 0, len(array)-1)
}

// Рекурсивная реализация быстрой сортировки
func quickSortRecursive(array []int, p, r int) []int {
	if p < r {
		q := Partition(array, p, r)
		quickSortRecursive(array, p, q-1)
		quickSortRecursive(array, q+1, r)
	}
	return array
}

// Partition Переупорядочивание массива A[p...r]
// и поиск опорного элемента
func Partition(array []int, p, r int) int {
	x := array[r]
	i := p - 1
	for j := p; j < r; j++ {
		if array[j] <= x {
			i++
			array[i], array[j] = array[j], array[i]
		}
	}
	array[i+1], array[r] = array[r], array[i+1]
	return i + 1
}

func main() {
	arr := make([]int, 100)
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100)
	}

	arr = quickSort(arr)
	fmt.Println(arr)
}
