package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Будем использовать готовую встроенную библиотеку для большой арифметики.
	a := big.NewInt(1 << 20)
	b := big.NewInt(2 << 20)
	t := &big.Int{} // Временная переменная для хранения результатов вычислений.

	fmt.Printf("%d + %d = %d\n", a, b, t.Add(a, b))
	fmt.Printf("%d - %d = %d\n", b, a, t.Sub(b, a))
	fmt.Printf("%d * %d = %d\n", a, b, t.Mul(a, b))
	fmt.Printf("%d / %d = %d\n", b, a, t.Div(b, a))
}
