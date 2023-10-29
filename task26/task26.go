package main

import (
	"fmt"
	"strings"
)

/*
Задание 26

Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
Функция проверки должна быть регистронезависимой.

Например:
abcd — true
abCdefAaf — false
aabcd — false
*/

func main() {
	// Слайс со строками
	s := []string{"abcd", "abCdefAaf", "aabcd"}
	// Перебор всех элементов слайса
	for _, v := range s {
		// Вызов для элемента слайса функции проверяющей строку на уникальность и получение результата
		result := unique(v)
		// Вывод строки и результата
		fmt.Println(v, "-", result)
	}
}

// Функция проверяющая строку на уникальность элементов в этой строке
func unique(s string) bool {
	// Создание map в которой в качестве ключей будут храниться символы строки
	charMap := make(map[string]struct{})
	// Перебор всех символов в строке
	for _, v := range s {
		// Преобразование значения rune в string нижнего регистра и проверка на нахождение в map
		_, ok1 := charMap[strings.ToLower(string(v))]
		// Преобразование значения rune в string верхнего регистра и проверка на нахождение в map
		_, ok2 := charMap[strings.ToUpper(string(v))]
		// Если в map находится данный символ, то строка не уникальна, поэтому возвращается false
		if ok1 || ok2 {
			return false
		}
		// Иначе символ добавляется в map
		charMap[string(v)] = struct{}{}
	}
	// Если все символы в строке уникальны возвращается true
	return true
}
