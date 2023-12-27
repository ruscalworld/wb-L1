package main

import "fmt"

func Intersect[V comparable](a, b *Set[V]) *Set[V] {
	result := NewSet[V]()

	// Проходим по элементам множества a и добавляем в итоговое множество те их них,
	// которые есть также в множестве b.
	a.Walk(func(v V) {
		if b.Has(v) {
			result.Add(v)
		}
	})

	return result
}

func main() {
	// Создадим два множества, которые будут содержать целые числа от 1 до 10 и от 6 до 15.
	a := NewSet(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	b := NewSet(6, 7, 8, 9, 10, 11, 12, 13, 14, 15)

	// Находим пересечение двух множеств.
	result := Intersect(a, b)

	// Выводим элементы множества-пересечения.
	result.Walk(func(v int) {
		fmt.Println(v)
	})
}
