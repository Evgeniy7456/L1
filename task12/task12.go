package main

import "fmt"

/*
Задание 12

Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
*/

// Стуктура множество
type Set struct {
	Elements map[string]struct{}
}

// Конструктор для создания нового множества
func newSet() *Set {
	return &Set{make(map[string]struct{})}
}

// Функция для добавления значения
func (s *Set) Add(item string) {
	s.Elements[item] = struct{}{}
}

// Функция для вывода всех значений
func (s *Set) List() {
	for v := range s.Elements {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func main() {
	// Массив строк
	str := [...]string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println("Массив строк:", str)
	// Инициализация множества
	set := newSet()
	// Добавление элементов из массива
	for _, v := range str {
		set.Add(v)
	}
	fmt.Print("Множество: ")
	set.List()
}
