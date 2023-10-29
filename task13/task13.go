package main

import "fmt"

/*
Задание 13

Поменять местами два числа без создания временной переменной.
*/

func main() {
	a := 1
	b := 2

	// С использованием побитовой операции XOR
	fmt.Printf("До замены a = %d, b = %d\n", a, b)
	a = a ^ b
	b = b ^ a
	a = a ^ b
	fmt.Printf("После замены a = %d, b = %d\n", a, b)
}
