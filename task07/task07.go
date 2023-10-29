package main

import (
	"fmt"
	"sync"
)

/*
Задание 7

Реализовать конкурентную запись данных в map.
*/

type Counter struct {
	Mutex sync.Mutex
	Value map[string]int
}

func main() {
	// Инициализация map
	counter := Counter{Value: make(map[string]int)}
	// WaitGroup для ожидания остановки горутин
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			// Использование блокировки Mutex для обеспечения потокобезопасной записи в map
			counter.Mutex.Lock()
			counter.Value["counter"]++
			counter.Mutex.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	// Ожидаемый результат 100
	fmt.Println(counter.Value["counter"])
}
