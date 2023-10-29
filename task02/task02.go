package main

import (
	"fmt"
	"sync"
)

/*
Задание 2

Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10)
и выведет их квадраты в stdout.
*/

func main() {
	// Массив чисел
	num := [...]int{2, 4, 6, 8, 10}
	// WaitGroup для ожидания завершения работы горутин
	wg := sync.WaitGroup{}
	// Создание горутин
	for i := 0; i < len(num); i++ {
		wg.Add(1)
		go square(&wg, num[i])
	}
	wg.Wait()
}

// Функция, которая выводит числа возведенные в квадрат
func square(wg *sync.WaitGroup, n int) {
	fmt.Println(n * n)
	wg.Done()
}
