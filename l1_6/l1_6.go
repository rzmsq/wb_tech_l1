package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

// exitForCondition завершает горутину при выполнении условия
func exitForCondition(condition int) {
	i := 0
	for {
		if i == condition {
			fmt.Println("Горутина [1]: условие выполнено, завершение работы...")
			return
		}
		fmt.Println("Горутина [1]: Работаю...")
		time.Sleep(time.Millisecond * 500)
		i++
	}
}

// exitForChan завершает горутину при получении сигнала из канала
func exitForChan(ch chan string) {
	for {
		select {
		case reason := <-ch:
			fmt.Println("Горутина [2]: завершение работы, Причина:", reason)
			return
		default:
			fmt.Println("Горутина [2]: Работаю...")
			time.Sleep(time.Millisecond * 500)
		}
	}
}

// exitForCtx завершает горутину при получении сигнала отмены из контекста
func exitForCtx(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Горутина [3]: Получен сигнал на отмену, завершение работы...")
			return
		default:
			fmt.Println("Горутина [3]: Работаю...")
			time.Sleep(time.Millisecond * 500)
		}
	}
}

// exitForRuntime завершает горутину с помощью runtime.Goexit
func exitForRuntime() {
	i := 0
	for {
		if i == 5 {
			fmt.Println("Горутина [4]: завершение работы Goexit...")
			runtime.Goexit()
		}
		fmt.Println("Горутина [4]: Работаю...")
		time.Sleep(time.Millisecond * 500)
		i++
	}
}

// exitForWaitGroup завершает горутину, используя sync.WaitGroup
func exitForWaitGroup(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Горутина [5]: Работаю...")
	time.Sleep(time.Millisecond * 500)
	fmt.Println("Горутина [5]: завершение работы...")

}

// exitForTimeout завершает горутину по истечении таймаута
func exitForTimeout(ctx context.Context) {
	select {
	case <-time.After(time.Second * 10):
		fmt.Println("Горутина[6] Задача завершена...")
	case <-ctx.Done():
		fmt.Println("Горутина [6]: Превышет таймаут, завершение работы...")
	}
}

// exitForCloseChan завершает горутину при закрытии канала
func exitForCloseChan(stop chan struct{}) {
	for {
		select {
		case <-stop:
			fmt.Println("Горутина [7]: Канал закрыт, завершение работы...")
			return
		default:
			fmt.Println("Горутина [7]: Работаю...")
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func main() {
	// Для всех горутин, чтобы main не завершилась до того, как все горутины закончат работу
	var wg sync.WaitGroup

	// ... [1] Выход по условию
	wg.Add(1)
	go func() {
		defer wg.Done()
		exitForCondition(6)
	}()
	time.Sleep(time.Second * 3)

	// ... [2] Выход через канал
	wg.Add(1)
	exitChan := make(chan string)
	go func() {
		defer wg.Done()
		exitForChan(exitChan)
	}()
	for i := 0; i < 10; i++ {
		if i == 6 {
			exitChan <- "Задание выполнено"
		}
		time.Sleep(time.Millisecond * 500)
	}
	time.Sleep(time.Second * 3)

	// ... [3] Выход через контекст
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go func() {
		defer wg.Done()
		exitForCtx(ctx)
	}()
	time.Sleep(time.Second * 3)
	cancel()
	time.Sleep(time.Second * 3)

	// ... [4] Выход через runtime
	wg.Add(1)
	go func() {
		defer wg.Done()
		exitForRuntime()
	}()
	time.Sleep(time.Second * 3)

	// ... [5] Выход через WaitGroup
	wg.Add(1)
	exitForWaitGroup(&wg)
	time.Sleep(time.Second * 3)

	// ... [6] Выход через таймаут
	ctx, cancel = context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	wg.Add(1)
	go func() {
		defer wg.Done()
		exitForTimeout(ctx)
		time.Sleep(time.Second * 5)
	}()

	// ... [7] Выход при закрытии канала
	stopChan := make(chan struct{})
	wg.Add(1)
	go func() {
		defer wg.Done()
		exitForCloseChan(stopChan)
	}()
	time.Sleep(time.Second * 3)
	close(stopChan)
	time.Sleep(time.Second * 3)

	wg.Wait()
}
