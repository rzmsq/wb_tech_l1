package main

import (
	"fmt"
	"sync"
	"time"
)

// Собственный sleep
func sleep(duration time.Duration) {
	<-time.After(duration)
}

func main() {

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for range 10 {
			fmt.Println("Do something...")
			sleep(1 * time.Second)
		}
	}()

	wg.Wait()
	fmt.Print("Exit")

}
