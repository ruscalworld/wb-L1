package main

import (
	"fmt"
	"math/rand"
)

func partition(a []int, low, high int) int {
	// Распределяем элементы в массиве таким образом, чтобы слева от опорной точки
	// были элементы меньше опорной точки, а справа - элементы больше.
	pivot := a[high]
	i := low - 1

	for j := low; j <= high; j++ {
		if a[j] < pivot {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}

	a[i+1], a[high] = a[high], a[i+1]
	return i + 1
}

func Sort(a []int, low, high int) {
	// Имея два индекса, постепенно увеличиваем один, а другой уменьшаем, проходя от
	// краёв массива к середине.
	if low < high {
		p := partition(a, low, high)
		Sort(a, low, p-1)
		Sort(a, p+1, high)
	}
}

func main() {
	// Создаём срез из 10 случайных целых чисел.
	a := make([]int, 10)
	for i := range a {
		a[i] = rand.Intn(20)
	}

	// Выводим исходный вариант массива, сортируем, выводим итоговый вариант массива.
	fmt.Println(a)
	Sort(a, 0, len(a)-1)
	fmt.Println(a)
}
