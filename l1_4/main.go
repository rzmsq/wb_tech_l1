package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

// worker - функция, которая будет выполняться в горутине
func worker(ctx context.Context, id int, wg *sync.WaitGroup, dataCh <-chan int) {
	// Отложенное завершение работы горутины
	defer wg.Done()
	fmt.Printf("worker %d starting\n", id)

	// Обработка данных из канала
	for {
		select {
		// получение данных из канала
		case data, ok := <-dataCh:
			// Проверка на закрытие канала
			if !ok {
				fmt.Printf("worker %d: channel closed, exiting\n", id)
				return
			}
			fmt.Printf("worker %d received data: %d\n", id, data)
			time.Sleep(1 * time.Second)
		case <-ctx.Done():
			// Обработка отмены контекста
			fmt.Printf("worker %d: context cancelled, exiting\n", id)
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
	// Использование WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup

	// Создание контекста для управления жизненным циклом горутин
	// Контекст позволяет отменить выполнение горутин при необходимости
	ctx, cancel := context.WithCancel(context.Background())

	// Запуск N горутин-воркеров
	for i := 0; i < n; i++ {
		// Добавление горутины в WaitGroup
		wg.Add(1)
		go worker(ctx, i, &wg, ch)
	}

	// Обработка сигналов завершения программы
	go func() {
		// Ожидание сигналов SIGINT или SIGTERM для корректного завершения
		sigChan := make(chan os.Signal, 1)
		// Регистрация сигналов для перехвата
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		// Блокировка до получения сигнала
		<-sigChan
		fmt.Println("Received shutdown signal, closing channel and cancelling context")
		// Закрытие канала и отмена контекста
		cancel()
	}()

	fmt.Println("Main goroutine started, generating random numbers")
	// Цикл для отправки случайных чисел в канал
producerLoop:
	for {
		// Генерация случайного числа
		data := rand.Int()

		// Отправка данных в канал с обработкой возможной блокировки
		select {
		// Отправка данных в канал
		case ch <- data:
			fmt.Printf("Main goroutine sent data: %d\n", data)
		// Проверка на закрытие канала
		case <-ctx.Done():
			fmt.Println("Main goroutine exited")
			break producerLoop
		}
		time.Sleep(500 * time.Millisecond)
	}

	// Закрытие канала после завершения генерации данных
	fmt.Println("Close channel")
	close(ch)

	// Ожидание завершения всех горутин
	fmt.Println("Waiting for workers to finish")
	wg.Wait()

	fmt.Println("Main goroutine exited")
}
