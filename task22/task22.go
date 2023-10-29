package main

import (
	"fmt"
	"math/big"
)

/*
Задание 22

Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.
*/

func main() {
	// a = 2 * 10^31
	aStr := "20000000000000000000000000000000"
	// Преобразование строки в число a
	a, _ := new(big.Int).SetString(aStr, 0)
	// b = 2 * 10^30
	bStr := "2000000000000000000000000000000"
	// Преобразование строки в число b
	b, _ := new(big.Int).SetString(bStr, 0)
	// a * b = 2 * 10^61
	result := new(big.Int)
	result.Mul(a, b)
	fmt.Println(result)
	// a / b = 10
	result.Div(a, b)
	fmt.Println(result)
	// a + b = 2.2 * 10^31
	result.Add(a, b)
	fmt.Println(result)
	// a - b = 1.8 * 10^31
	result.Sub(a, b)
	fmt.Println(result)
}
