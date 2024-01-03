package main

import "fmt"

func remove[T any](i int, a []T) []T {
	// Создаём новый слайс размером на 1 меньше.
	result := make([]T, len(a)-1)

	// Копируем в новый слайс всё, что до удаляемого элемента,
	// и всё, что после него.
	copy(result[:i], a[:i])
	copy(result[i:], a[i+1:])

	// Возвращаем результат.
	return result
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	r := remove(2, a)

	fmt.Println(r) // [1 2 4 5]
}
