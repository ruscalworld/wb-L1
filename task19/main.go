package main

import (
	"fmt"
)

func reverse(str string) string {
	// Превратим строку в массив rune, чтобы не было проблем.
	c := []rune(str)

	var (
		i = 0
		j = len(c) - 1
	)

	// Пойдём равномерно с двух сторон - с начала и с конца.
	// Будем менять местами элемент на позиции i с элементом на позиции j.
	// После этого i будет увеличиваться, а j - уменьшаться.
	// Повторяем до тех пор, пока не дойдём до середины (i >= j).
	for i <= j {
		c[i], c[j] = c[j], c[i]
		i++
		j--
	}

	// Можно обойтись без второй переменной и хорошенько ужать код,
	// правда тогда это читается несколько хуже:
	//
	// 	for i := 0; i < len(c)/2; i++ {
	// 		c[i], c[len(c)-i-1] = c[len(c)-i-1], c[i]
	// 	}

	// Можно с доп. памятью.
	// Перекладываем по порядку все элементы из исходного массива во временный,
	// только во временный начинаем их складывать с конца.
	//
	// 	temp := make([]rune, len(c))
	// 	for i := 0; i < len(c); i++ {
	// 		temp[len(c)-1-i] = c[i]
	// 	}
	// 	return string(temp)

	// Топ-1 по извращённости - обход двусвязного списка в разных направлениях.
	// Проще было бы сделать то же самое со стеком, но в стандартной библиотеке
	// Go стека нет, поэтому поразвлекаемся со списком.
	//
	// 	l := list.New()
	// 	for i := 0; i < len(c); i++ {
	// 		l.PushBack(c[i])
	// 	}
	//
	// 	i := 0
	// 	for p := l.Back(); p != nil; p = p.Prev() {
	// 		c[i] = p.Value.(rune)
	// 		i++
	// 	}

	// На этом мысли закончились...

	_ = j
	_ = i

	return string(c)
}

func main() {
	fmt.Println(reverse("главрыба"))
}
