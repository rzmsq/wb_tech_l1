package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	count int
	mu    sync.Mutex
}

// Безопасное увеличение счетчика
func (c *Counter) Inc() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func worker(count *Counter, wg *sync.WaitGroup, iteration int) {
	defer wg.Done()

	for i := 0; i < iteration; i++ {
		count.Inc()
	}
}

func main() {
	counter := Counter{0, sync.Mutex{}}
	wg := sync.WaitGroup{}

	numsWorkers := 5
	iterationsWorker := 10000

	wg.Add(numsWorkers)
	for i := 0; i < 5; i++ {
		go worker(&counter, &wg, iterationsWorker)
	}
	wg.Wait()

	exceptedVal := numsWorkers * iterationsWorker
	fmt.Println("End counter:", counter.count)
	fmt.Println("Excepted val:", exceptedVal)
}
