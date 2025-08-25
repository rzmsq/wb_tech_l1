package main

import (
	"fmt"
)

func main() {

	decimal := 14684615
	str := "Abcdrd"
	boolean := true
	ch := make(chan interface{})

	detectVar(decimal)
	detectVar(str)
	detectVar(boolean)
	detectVar(ch)

}

func detectVar(t interface{}) {
	switch i := t.(type) {
	case nil:
		fmt.Println("x is nil")
	case int:
		fmt.Printf("Тип: %T, значение: %v\n", i, i)
	case bool:
		fmt.Printf("Тип: %T, значение: %v\n", i, i)
	case string:
		fmt.Printf("Тип: %T, значение: %s\n", i, i)
	case chan interface{}:
		fmt.Printf("Тип: %T, значение: %v\n", i, i)
	default:
		fmt.Println("unknown type")

	}
}
