package main

import (
	"fmt"
	"strings"
)

func reverse(str string) string {
	// Разделим строку на слова, полагая, что они разделены пробелом.
	words := strings.Split(str, " ")

	// Воспользуемся алгоритмом из предыдущего задания.
	var (
		i = 0
		j = len(words) - 1
	)

	for i < j {
		words[i], words[j] = words[j], words[i]
		i++
		j--
	}

	// Иные варианты реализации см. в предыдущем задании.

	// Соединим все слова обратно.
	return strings.Join(words, " ")
}

func main() {
	fmt.Println(reverse("snow dog sun"))
}
