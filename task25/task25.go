package main

import (
	"fmt"
	"time"
)

/*
Задание 25

Реализовать собственную функцию sleep.
*/

func main() {
	fmt.Println("Sleep")
	// Запуск функции sleep на 5 секунд
	Sleep(5000)
	fmt.Println("Run")
}

// Функция sleep в миллисекундах
func Sleep(n time.Duration) {
	<-time.After(n * time.Millisecond)
}
