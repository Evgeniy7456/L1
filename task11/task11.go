package main

import "fmt"

/*
Задание 11

Реализовать пересечение двух неупорядоченных множеств.
*/

// Стуктура множество
type Set struct {
	Elements map[int]struct{}
}

// Конструктор для создания нового множества
func newSet() *Set {
	return &Set{make(map[int]struct{})}
}

// Функция для добавления значения
func (s *Set) Add(item int) {
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
	// Массив с числами для первого множества
	firstArray := [...]int{1, 2, 3, 4}
	// Массив с числами для второго множества
	secondArray := [...]int{3, 4, 5, 6}
	// Инициализация первого множества и заполнение его числами из первого массива
	firstSet := newSet()
	for _, v := range firstArray {
		firstSet.Add(v)
	}
	fmt.Print("Первое множество: ")
	firstSet.List()
	// Инициализация второго множества и заполнение его числами из второго массива
	secondSet := newSet()
	for _, v := range secondArray {
		secondSet.Add(v)
	}
	fmt.Print("Второе множество: ")
	secondSet.List()
	// Инициализация объединенного множества
	unionSet := newSet()
	// Добавление элементов из первого множества
	for v := range firstSet.Elements {
		unionSet.Add(v)
	}
	// Добавление элементов из второго множества
	for v := range secondSet.Elements {
		unionSet.Add(v)
	}
	fmt.Print("Объединенное множество: ")
	unionSet.List()
}
