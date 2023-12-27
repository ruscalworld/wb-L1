package main

import "fmt"

func Intersect[V comparable](a, b *Set[V]) *Set[V] {
	result := NewSet[V]()

	a.Walk(func(v V) {
		if b.Has(v) {
			result.Add(v)
		}
	})

	return result
}

func main() {
	a := NewSet(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	b := NewSet(6, 7, 8, 9, 10, 11, 12, 13, 14)

	result := Intersect(a, b)
	result.Walk(func(v int) {
		fmt.Println(v)
	})
}
