package main

import "fmt"

/*
Задание 1

Дана структура Human (с произвольным набором полей и методов). Реализовать встраивание методов в структуре Action
от родительской структуры Human (аналог наследования).
*/

type (
	// Структура Human
	Human struct {
		Name string
		Age  int
	}

	// Структура Action содержит в себе структуру Human
	Action struct {
		Human
	}
)

// Метод Human
func (h *Human) Eat() {
	fmt.Printf("Human %s eats\n", h.Name)
}

// Конструктор, который возвращает *Action
func NewHuman(name string, age int) *Action {
	return &Action{
		Human{
			Name: name,
			Age:  age,
		},
	}
}

func main() {
	// Вызов функции, которая возвращает *Action
	human := NewHuman("Ivan", 30)
	// Вызов у *Action метода, который был определен для Human
	human.Eat()
}
