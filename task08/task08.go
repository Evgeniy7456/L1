package main

import (
	"fmt"
	"math"
)

/*
Задание 8

Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
*/

func main() {
	// 4 = 100
	var N int64 = 4
	// Замена 0 бита на 1. Получаем 101 = 5
	bitOperation(N, 0)
	// Замена 1 бита на 1. Получаем 110 = 6
	bitOperation(N, 1)
	// Замена 2 бита на 0. Получаем 000 = 0
	bitOperation(N, 2)
	// 4 можно представить как 0100. Замена 3 бита на 1. Получаем 1100 = 12
	bitOperation(N, 3)
}

func bitOperation(N int64, i float64) {
	fmt.Println(N ^ int64(math.Pow(2, i)))
}
