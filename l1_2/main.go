package main

import (
	"os"
	"strconv"
	"sync"
)

// wrapper запускает переданную функцию в отдельной горутине и добавляет её в WaitGroup
func wrapper(wg *sync.WaitGroup, fn func()) {
	wg.Add(1)

	go func() {
		defer wg.Done()
		fn()
	}()
}

func main() {
	// Создание массива с числами
	arr := [...]int{2, 4, 6, 8, 10}
	// Создание WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup

	// Запуск горутины для вычисления квадратов чисел в массиве и вывода их на экран
	// Используется wrapper для запуска функции в горутине
	wrapper(&wg, func() {
		for i := 0; i < len(arr); i++ {
			arr[i] *= arr[i]
		}

		for _, el := range arr {
			// Вывод каждого элемента массива в stdout
			_, err := os.Stdout.Write([]byte(strconv.Itoa(el) + " "))
			if err != nil {
				panic(err)
			}
		}
	})

	// Ожидание завершения всех горутин
	wg.Wait()
}
