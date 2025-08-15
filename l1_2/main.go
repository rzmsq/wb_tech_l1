package main

import (
	"log"
	"os"
	"strconv"
	"sync"
)

func main() {
	// Создание массива с числами
	arr := [...]int{2, 4, 6, 8, 10}
	// Создание WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup

	// Запуск горутин для каждого элемента массива
	for i := 0; i < len(arr); i++ {
		// Увеличение счетчика WaitGroup
		wg.Add(1)
		// Запуск горутины для обработки числа
		go func() {
			// Уменьшение счетчика WaitGroup при завершении горутины
			defer wg.Done()

			arr[i] *= arr[i]
		}()
	}

	// Ожидание завершения всех горутин
	wg.Wait()

	for _, val := range arr {
		// Печать квадрата числа в стандартный вывод
		_, err := os.Stdout.Write([]byte(strconv.Itoa(val) + " "))
		// Проверка на ошибки при записи в стандартный вывод
		if err != nil {
			log.Fatalf("error writing to stdout: %v", err)
			return
		}
	}
}
