package main

import (
	"fmt"
	"math/rand"
	"os"
)

// Запись значений из массива в канал
func producer(arr [100]int) <-chan int {
	c := make(chan int)
	go func() {
		for _, n := range arr {
			c <- n
		}
		close(c)
	}()
	return c
}

// Запись в канал результата операции x*2
func producerConsumer(c <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for n := range c {
			out <- n * 2
		}
		close(out)
	}()
	return out
}

// Вывод данных из канала в stdout
func consumer(ch <-chan int) {
	for n := range ch {
		_, err := os.Stdout.Write([]byte(fmt.Sprintf("%d\n", n)))
		if err != nil {
			panic(err)
		}
	}
}

func main() {

	// Генерация массива
	arr := [100]int{}
	for i := range 100 {
		arr[i] = rand.Intn(1000)
	}

	// pipeline
	consumer(producerConsumer(producer(arr)))

}
