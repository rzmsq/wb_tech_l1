package main

import (
	"fmt"
	"sync"

	cmap "github.com/orcaman/concurrent-map/v2"
)

/*  Релизация с mutex  */
// writer записывает данные в обычную map с использованием мьютекса для безопасности
func writer(m *map[int]int, locker *sync.RWMutex, wg *sync.WaitGroup) {
	wg.Add(100)
	for i := range 100 {
		go func(val int) {
			defer wg.Done()
			locker.Lock() // Блокировка для записи
			(*m)[val] = val
			locker.Unlock() // Разблокировка
		}(i)
	}
}

// reader читает все данные из обычной map с использованием read-lock
func reader(m *map[int]int, locker *sync.RWMutex) {
	locker.RLock()         // Блокируем для чтения
	defer locker.RUnlock() // Разблокировка при выходе
	for key, value := range *m {
		fmt.Println("[mutex] key:", key, "value:", value)
	}
}

// mutexVersion демонстрирует работу с обычной map + mutex
func mutexVersion() {
	m := make(map[int]int)
	locker := new(sync.RWMutex) // RWMutex для синхронизации
	wg := &sync.WaitGroup{}

	writer(&m, locker, wg) // Запись данных
	wg.Wait()
	reader(&m, locker) // Чтение данных
}

/***********************************************************************/

/* Реализация с concurrent-map */
// writerCmap записывает данные в concurrent-map (потокобезопасную карту)
func writerCmap(wg *sync.WaitGroup, m *cmap.ConcurrentMap[string, int]) {
	wg.Add(100)
	for i := range 100 {
		go func(val int) {
			defer wg.Done()
			// Set() - потокобезопасный метод записи, блокировка не нужна
			m.Set(fmt.Sprintf("%d", val), val)
		}(i)
	}
}

// readerCmap читает данные из concurrent-map
func readerCmap(m *cmap.ConcurrentMap[string, int]) {
	// IterBuffered() возвращает канал для безопасной итерации по карте
	for item := range m.IterBuffered() {
		fmt.Println("[concurrent-map] key:", item.Key, "value:", item.Val)
	}
}

// concMapVersion демонстрирует работу с concurrent-map библиотекой
func concMapVersion() {
	m := cmap.New[int]() // concurrent-map с string ключами и int значениями
	wg := &sync.WaitGroup{}

	writerCmap(wg, &m) // Запись данных
	wg.Wait()
	readerCmap(&m) // Чтение данных

}

/***********************************************************************/

// main функция демонстрирует два подхода к concurrent записи в map:
// 1. Обычная map + RWMutex
// 2. Специализированная concurrent-map библиотека
func main() {
	mutexVersion()

	concMapVersion()
}
