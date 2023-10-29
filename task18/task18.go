package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
Задание 18

Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде. По завершению
программа должна выводить итоговое значение счетчика.
*/

// Структура-счетчик
type Counter struct {
	Mutex sync.Mutex
	Value int
}

// Структура-счетчик без использования Mutex
type Counter2 struct {
	Value int64
}

func main() {
	// Инициализация структуры-счетчика
	counter := Counter{}
	// WaitGroup для ожидания остановки горутин
	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			// Использование блокировки Mutex для синхронизации горутин
			counter.Mutex.Lock()
			counter.Value++
			counter.Mutex.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	// Ожидаемый результат 1000
	fmt.Println(counter.Value)

	// Инициализация структуры
	counter2 := Counter2{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			// Использование пакета atomic для синхронизации горутин
			atomic.AddInt64(&counter2.Value, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	// Ожидаемый результат 1000
	fmt.Println(counter2.Value)
}
