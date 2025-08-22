package main

import "fmt"

func main() {

	// Дано
	temps := [...]float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// map для создания групп
	m := make(map[int][]float64)

	// Проход по массиву, определения группы и сохранение температуры
	for _, v := range temps {
		var key = int(v) / 10 * 10
		m[key] = append(m[key], v)
	}

	fmt.Println(m)
}
