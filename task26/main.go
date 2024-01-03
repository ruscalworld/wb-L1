package main

import (
	"fmt"
	"strings"
)

func checkUnique(str string) bool {
	// Приведём всё к нижнему регистру, чтобы обеспечить независимость от регистра.
	str = strings.ToLower(str)

	// Преобразуем строку в массив рун и создадим карту, куда будем записывать
	// уже найденные символы.
	runes := []rune(str)
	found := make(map[rune]bool, len(runes))

	// Проходим по каждому символу в исходной строке.
	for _, r := range runes {
		// Проверяем, был ли до этого уже найден такой же символ.
		if _, ok := found[r]; ok {
			// Если был найден, то сразу выходим, сообщая, что строка не подходит.
			return false
		}

		// Запоминаем, что мы встретили этот символ.
		found[r] = true
	}

	// Если одинаковых символов нет, то цикл завершится, и мы дойдём сюда.
	return true
}

func main() {
	fmt.Println(checkUnique("abcd"))
	fmt.Println(checkUnique("abCdefAaf"))
	fmt.Println(checkUnique("aabcd"))
}
