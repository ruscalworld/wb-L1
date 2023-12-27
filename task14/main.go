package main

import (
	"fmt"
	"reflect"
)

func PrintType(v interface{}) {
	fmt.Println(reflect.TypeOf(v))
}

func main() {
	PrintType(1)
	PrintType(1.0)
	PrintType("hello")
	PrintType([]int64{12})
}
