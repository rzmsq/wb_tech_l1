package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {

	// Получение n из командной строки
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <N>")
		return
	}

	// Парсинг аргумента n
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	// канал для обмена между горутинами
	ch := make(chan int)
	// канал тайм-аута
	timeout := time.After(time.Duration(n) * time.Second)

	// Горутина, которая отправляет данные в канал
	go func() {
		for i := 0; ; i++ {
			ch <- i
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
	}()

	// Бесконечно слушаем сообщения в канале или сигнал тайм-аута.
	for {
		select {
		// Если пришло значение из канала ch.
		case <-ch:
			fmt.Println("Received", <-ch)
		// Если пришел сигнал из канала тайм-аута.
		case <-timeout:
			fmt.Println("Timeout")
			return
		}
	}

}
