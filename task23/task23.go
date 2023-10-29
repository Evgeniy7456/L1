package main

import "fmt"

/*
Задание 23

Удалить i-ый элемент из слайса.
*/

func main() {
	slice := []int{1, 2, 3, 4}
	fmt.Println("Начальный слайс:", slice)

	// Если порядок элементов в слайсе не важен
	// В этом случае на место удаляемого индекса копируется последний элемент и делается срез до последнего элемента
	slice, _ = remove1(slice, 3)
	fmt.Println("После удаления 3 элемента:", slice)

	// Если порядок элементов в слайсе важен
	// Смещение элементов на один влево относительно удаляемого и удаление последнего элемента
	slice, _ = remove2(slice, 2)
	fmt.Println("После удаления 2 элемента:", slice)
	// Копирование элементов до удаляемого элемента и после
	slice, _ = remove3(slice, 1)
	fmt.Println("После удаления 1 элемента:", slice)
}

func remove1(slice []int, i int) ([]int, error) {
	// Ошибка, если число выходит за диапазон индексов слайса
	if i > len(slice)-1 || i < 0 {
		return nil, fmt.Errorf("index out of range [%d] with length %d", i, len(slice))
	}
	// Так как функция удаляет элемент по номеру, а индекс начинается с нуля, то для получения индекса нужно элемента
	// происходит уменьшение на 1
	i--
	// Копирование последнего элемента на место удаляемого
	slice[i] = slice[len(slice)-1]
	// Удаление последнего элемента
	slice = slice[:len(slice)-1]
	return slice, nil
}

func remove2(slice []int, i int) ([]int, error) {
	if i > len(slice)-1 || i < 0 {
		return nil, fmt.Errorf("index out of range [%d] with length %d", i, len(slice))
	}
	i--
	// Копирование элементов следующих после удаляемого на место удаляемого и далее
	copy(slice[i:], slice[i+1:])
	slice = slice[:len(slice)-1]
	return slice, nil
}

func remove3(slice []int, i int) ([]int, error) {
	if i > len(slice)-1 || i < 0 {
		return nil, fmt.Errorf("index out of range [%d] with length %d", i, len(slice))
	}
	i--
	// Возвращение копии слайса до удаляемого элемента и после
	return append(slice[:i], slice[i+1:]...), nil
}
