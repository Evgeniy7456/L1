package main

import "fmt"

/*
Задание 14

Разработать программу, которая в рантайме способна определить тип переменной:
int, string, bool, channel из переменной типа interface{}.
*/

func main() {
	// Массив типа интерфейс
	array := [...]interface{}{1, "string", true, make(chan int), make(chan string), make(chan bool), make(chan float64)}
	// Перебор всех элементов массива и вызов для них функции определяющей тип переменной
	for _, v := range array {
		typeVariable(v)
	}
}

// Функция, которая с помощью конструкции switch определяет тип переменной
func typeVariable(a interface{}) {
	switch t := a.(type) {
	case int:
		fmt.Printf("Переменная %d типа int\n", t)
	case string:
		fmt.Printf("Переменная %s типа string\n", t)
	case bool:
		fmt.Printf("Переменная %t типа bool\n", t)
	case chan int:
		fmt.Println("Переменная типа chan int")
	case chan string:
		fmt.Println("Переменная типа chan string")
	case chan bool:
		fmt.Println("Переменная типа chan bool")
	case chan float64:
		fmt.Println("Переменная типа chan float64")
	default:
		fmt.Println("Не удалось определить тип переменной")
	}
}
