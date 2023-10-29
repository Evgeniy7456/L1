package main

import (
	"fmt"
	"strings"
)

/*
Задание 20

Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/

func main() {
	// Исходная строка
	s := "snow dog sun"
	// Вызов функции возвращающей перевернутую строку
	r := reverse(s)
	// Вывод результата
	fmt.Printf("Исходная строка: %s\n", s)
	fmt.Printf("Перевернутая строка: %s", r)
}

// Функция принимает в качестве параметра строку и возвращает перевернутую строку
func reverse(s string) string {
	// Получение слайса слов строки
	sliceWord := strings.Split(s, " ")
	// Создание пустого слайса
	var reverse []string
	// Цикл, который проходит по слайсу слов строки в обратном порядке и добавляет их в другой слайс
	for i := len(sliceWord) - 1; i >= 0; i-- {
		reverse = append(reverse, sliceWord[i])
	}
	// Вызов функции, которая преобразует слайс из слов с строку
	return strings.Join(reverse, " ")
}
