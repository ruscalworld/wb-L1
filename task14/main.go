package main

import "fmt"

func PrintType(v interface{}) {
	fmt.Printf("%T\n", v)
}

func main() {
	PrintType(1)
	PrintType(1.0)
	PrintType("hello")
	PrintType([]int64{12})
}
