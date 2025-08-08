package main

import (
	"log"
	"os"
	"strconv"
	"sync"
)

func main() {
	arr := [...]int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup

	for _, i := range arr {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			_, err := os.Stdout.Write([]byte(strconv.Itoa(val*val) + "\n"))
			if err != nil {
				log.Fatalf("error writing to stdout: %v", err)
				return
			}
		}(i)
	}

	wg.Wait()
}
