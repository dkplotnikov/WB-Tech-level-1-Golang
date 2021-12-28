package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/

func RevWords(s string) string {
	w := strings.Fields(s)
	n := len(w)
	for i := 0; i < n/2; i++ {
		w[i], w[n-i-1] = w[n-i-1], w[i]
	}
	return strings.Join(w, " ")
}

func main() {
	s := "snow dog sun moon cat cat cat light day night"
	fmt.Println("String:", s)
	fmt.Println("RevWords:", RevWords(s))
}
