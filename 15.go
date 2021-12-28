package main

import (
	"fmt"
)

/*
К каким негативным последствиям может привести данный фрагмент кода, и как
это исправить? Приведите корректный пример реализации.

var justString string
func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}
func main() {
	someFunc()
}
*/

var justString string

func createHugeString(n int) string {
	return "ФЫВАПРОЛДЖ"
}

func someFunc() {
	v := createHugeString(1 << 10)
	fmt.Println(v)

	// 1. Слайс байт строки v останется в памяти, но не будет доступен,
	//    т.к. v недоступка вне someFunc

	// 2. Если задача в том, чтобы взять подстроку, то проблема может быть в том,
	//    что мы получим не подстроку символов, а строку из слайса байт

	// Взять первые 5 символов
	rjs := make([]rune, 5)
	i := 0
	for _, r := range v {
		if i >= 5 {
			break
		}
		rjs[i] = r
		i++
	}
	justString = string(rjs)
	fmt.Println(justString)

	// Взять слайс первых 4 байт
	fmt.Println(v[:4])

	// Взять слайс первых 5 байт
	fmt.Println(v[:5])
}

func main() {
	someFunc()
}
