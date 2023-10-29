package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
Задание 6

Реализовать все возможные способы остановки выполнения горутины.
*/

func main() {
	// WaitGroup для ожидния остановки выполнения горутины
	wg := sync.WaitGroup{}

	// 1 способ
	// Остановка выполнения горутины после получения сигнала из канала
	ch := make(chan bool)

	fmt.Println("1 способ")
	fmt.Println("Запуск горутины")
	wg.Add(1)
	go worker(&wg, ch)

	time.Sleep(2 * time.Second)
	ch <- true
	wg.Wait()

	// 2 способ
	// Остановка выполнения горутины после закрытия канала из которого горутина получала данные
	fmt.Println("\n2 способ")
	ch2 := make(chan string)
	fmt.Println("Запуск горутины")
	wg.Add(1)
	go worker2(&wg, ch2)

	for i := 0; i < 2; i++ {
		ch2 <- "данные"
	}
	close(ch2)
	wg.Wait()

	// 3 способ
	// Остановка выполнения горутины с помощью контекста
	fmt.Println("\n3 способ")
	// Контекст, который задает время работы в секундах
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Контекст, который задает дату и время до которого будет работать программа
	// t := time.Date(2023, time.October, 18, 15, 00, 0, 0, time.Now().Location())
	// ctx, cancel := context.WithDeadline(context.Background(), t)

	// Контекст, который останавливает горутину после выполнения cancel()
	// ctx, cancel := context.WithCancel(context.Background())

	// Контекст, который останавливает горутину после сигнала от системы
	// ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("Запуск горутины")
	wg.Add(1)
	go worker3(ctx, &wg)
	wg.Wait()

	// 4 способ
	// Можно применять, если горутина делает запрос к серверу. Если запрос приходит за определенное время, то он обрабатывается.
	// Иначе останавливается выполнение горутины.
	fmt.Println("\n4 способ")
	fmt.Println("Запуск горутины")
	// Время ожидания ответа в секундах
	var n time.Duration = 5
	t := n * time.Second
	ch4 := make(chan string)
	wg.Add(1)
	go worker4(&wg, t, ch4)
	ch4 <- "данные"
	wg.Wait()
}

// 1 способ
func worker(wg *sync.WaitGroup, ch chan bool) {
	for {
		select {
		case <-ch:
			fmt.Println("Остановка выполнения горутины")
			wg.Done()
			return
		default:
			fmt.Println("Работа горутины")
			time.Sleep(1 * time.Second)
		}
	}
}

// 2 способ
func worker2(wg *sync.WaitGroup, ch chan string) {
	v, ok := <-ch
	for ok {
		fmt.Printf("Получены данные: %s\n", v)
		fmt.Println("Обработка данных")
		time.Sleep(1 * time.Second)
		v, ok = <-ch
	}
	fmt.Println("Остановка выполнения горутины")
	wg.Done()
}

// 3 способ
func worker3(ctx context.Context, wg *sync.WaitGroup) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Остановка выполнения горутины")
			wg.Done()
			return
		default:
			fmt.Println("Работа горутины")
			time.Sleep(1 * time.Second)
		}
	}
}

// 4 способ
func worker4(wg *sync.WaitGroup, n time.Duration, ch chan string) {
	for {
		fmt.Println("Ожидание ответа сервера")
		select {
		case <-time.After(n):
			fmt.Println("Время истекло")
			fmt.Println("Остановка выполнения горутины")
			wg.Done()
			return
		case data := <-ch:
			fmt.Printf("Получен ответ: %s\n", data)
		}
	}
}
