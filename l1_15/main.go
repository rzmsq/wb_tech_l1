package main

import "fmt"

//Этот код может привести к утечке памяти,
//поскольку он создает строку v большого размера,
//а затем создает срез justString из первых 100 байт этой строки.
//
//Пока существует срез justString, сборщик мусора не сможет освободить память,
//выделенную для всей строки v, даже если используется только её небольшая часть.
//var justString string
//
//func someFunc() {
//	v := createHugeString(1 << 10)
//	justString = v[:100]
//}
//func main() {
//	someFunc()
//}

//Чтобы избежать утечки памяти, необходимо скопировать нужную часть данных
//из большой строки в новую, независимую строку.
//Это позволит сборщику мусора освободить память,
//занимаемую исходной большой строкой v после того, как она перестанет быть нужной.

var justString string

func createHugeString(n int) string {
	s := make([]byte, n)
	for i := 0; i < n; i++ {
		s[i] = 'A'
	}
	return string(s)
}

func someFunc() {
	v := createHugeString(1 << 10)
	// Создаем новую строку, что позволяет сборщику мусора освободить память
	// занимаемую строкой v
	justString = string([]rune(v)[:100])
}

func main() {
	someFunc()
	fmt.Println(justString)
}
