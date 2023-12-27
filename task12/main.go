package main

import "fmt"

func main() {
	// Создаём множество из данных в задании элементов.
	s := NewSet("cat", "cat", "dog", "cat", "tree")

	// Выводим элементы итогового множества.
	s.Walk(func(v string) {
		fmt.Println(v)
	})
}
