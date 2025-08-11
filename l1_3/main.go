package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// worker - функция, которая будет выполняться в горутине
func worker(id int, dataCh <-chan int) {
	// Обработка данных из канала
	for data := range dataCh {
		// Печать данных в стандартный вывод
		_, err := os.Stdout.Write([]byte("Worker " + strconv.Itoa(id) + ": " + strconv.Itoa(data) + "\n"))
		// Проверка на ошибки при записи в стандартный вывод
		if err != nil {
			fmt.Printf("Worker %d: error writing to stdout: %v\n", id, err)
			return
		}
	}
}

func main() {
	// Проверка аргументов командной строки
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <N>")
		return
	}

	// Преобразование аргумента в число и проверка на корректность
	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n <= 0 {
		fmt.Printf("Invalid number: %s\n", os.Args[1])
		return
	}

	// Канал для передачи данных между горутинами
	ch := make(chan int)

	// Запуск N горутин-воркеров
	for i := 0; i < n; i++ {
		go worker(i, ch)
	}

	// Генерация случайных чисел и отправка их в канал
	for {
		// Отправка случайного числа в канал
		ch <- rand.Int()
		time.Sleep(500 * time.Millisecond)
	}
}
