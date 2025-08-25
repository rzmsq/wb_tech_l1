package main

import (
	"fmt"
)

// binarySearch итеративный бинарный поиск
func binarySearch(array []int, find int) int {
	l := 0
	r := len(array) - 1

	for l <= r {
		mid := l + (r-l)/2
		if array[mid] == find {
			return mid
		}
		if array[mid] > find {
			r = mid - 1
		}
		if array[mid] < find {
			l = mid + 1
		}
	}
	return -1
}

func main() {
	arr := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		arr[i] = i
	}
	fmt.Println("Массив:", arr)
	fmt.Println("\nВведите элемент для поиска")
	var find int
	_, err := fmt.Scanf("%d", &find)
	if err != nil {
		return
	}
	fmt.Println("\nИндекс искомого:", binarySearch(arr, find))
}
