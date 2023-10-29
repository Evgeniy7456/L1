package main

import (
	"fmt"
	"sync"
)

/*
Задание 9

Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива, во второй — результат
операции x*2, после чего данные из второго канала должны выводиться в stdout.
*/

func main() {
	// Массив чисел
	num := [...]int{1, 2, 3, 4}
	// Первый канал
	firstCh := make(chan int)
	// Второй канал
	secondCh := make(chan int)
	// WaitGroup для ожидания остановки горутин
	wg := sync.WaitGroup{}
	go multiplication(firstCh, secondCh)
	wg.Add(1)
	go out(&wg, secondCh)

	// Перебор всех элементов массива и отправка их в первый канал
	for i := 0; i < len(num); i++ {
		firstCh <- num[i]
	}

	fmt.Println("Ожидание остановки горутин")
	close(firstCh)
	wg.Wait()
}

// Функция, которая умножает числа на 2
func multiplication(firstCh, secondCh chan int) {
	v, ok := <-firstCh
	for ok {
		secondCh <- v * 2
		v, ok = <-firstCh
	}
	fmt.Println("Горутина multiplication остановлена")
	close(secondCh)
}

// Функция, которая выводит результат
func out(wg *sync.WaitGroup, secondCh chan int) {
	v, ok := <-secondCh
	for ok {
		fmt.Println(v)
		v, ok = <-secondCh
	}
	fmt.Println("Горутина out остановлена")
	wg.Done()
}
