package main

import "fmt"

func Find(a []int, v int) bool {
	// Если остался пустой срез, значит искомого элемента нет.
	if len(a) == 0 {
		return false
	}

	// Если искомый элемент больше того, что находится в середине, ищем в правой части.
	mid := len(a) / 2
	if v > a[mid] {
		return Find(a[mid+1:], v)
	}

	// Если искомый элемент меньше того, что находится в середине, ищем в левой части.
	if v < a[mid] {
		return Find(a[:mid], v)
	}

	// Если искомый элемент равен тому, что находится в середине, то получается, что нашли.
	return true
}

func main() {
	// Создаём слайс с 20 нечётными числами
	a := make([]int, 20)
	for i := 0; i < 20; i++ {
		a[i] = 2*i + 1
	}

	fmt.Println(a)

	// Поскольку исходный массив изначально отсортирован, опустим его сортировку.

	fmt.Println(Find(a, 1))
	fmt.Println(Find(a, 2))
	fmt.Println(Find(a, 3))
	fmt.Println(Find(a, 4))
	fmt.Println(Find(a, 5))
}
