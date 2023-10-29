package main

import (
	"L1/task21/exampleLib"
	"fmt"
)

/*
Задание 21

Реализовать паттерн «адаптер» на любом примере.
*/

// Интерфейс под который необходимо адаптировать метод из библиотеки
type OurInterface interface {
	Action() string
}

// Структура адаптер, включающая в себя структуру библиотеки, которая реализует необходимый метод
type Adapter struct {
	*exampleLib.LibStruct
}

// Реализация интерфеса с вызовом нужного метода из библиотеки
func (a *Adapter) Action() string {
	return a.ExampleAction()
}

// Конструктор адаптера
func NewAdapter(exampleLib *exampleLib.LibStruct) OurInterface {
	return &Adapter{exampleLib}
}

func main() {
	// Создание экземпляра адаптера
	adapter := NewAdapter(&exampleLib.LibStruct{})
	// Вызов метода у адаптера
	s := adapter.Action()
	// Вывод результата
	fmt.Println(s)
}
