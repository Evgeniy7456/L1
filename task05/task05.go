package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
Задание 5

Разработать программу, которая будет последовательно отправлять значения в канал, а с другой
стороны канала — читать. По истечению N секунд программа должна завершаться.
*/

func main() {
	var N time.Duration
	fmt.Print("Введите время работы программы в секундах: ")
	fmt.Scan(&N)
	ctx, cancel := context.WithTimeout(context.Background(), N*time.Second)
	defer cancel()

	// Создание канала
	ch := make(chan int)
	// WaitGroup для ожидания завершения работы воркером
	wg := sync.WaitGroup{}
	wg.Add(1)
	go worker(&wg, ch)

	i := 0
	for {
		select {
		// Завершение программы по истечению N секунд
		case <-ctx.Done():
			fmt.Println("Ожидание завершения работы воркером")
			close(ch)
			wg.Wait()
			return
		// Отправка данных в канал
		default:
			i++
			ch <- i
		}
	}
}

// Воркер, который завершает работу, если канал закрыт.
func worker(wg *sync.WaitGroup, ch chan int) {
	v, ok := <-ch
	for ok {
		fmt.Println(v)
		v, ok = <-ch
	}
	fmt.Println("Воркер закончил работу")
	wg.Done()
}
