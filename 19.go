package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Разработать программу, которая переворачивает подаваемую на ход строку
(например: «главрыба — абырвалг»). Символы могут быть unicode.
*/

func Reverse(s string) string {
	r := []rune(s)
	n := len(r)
	for i := 0; i < n/2; i++ {
		r[i], r[n-i-1] = r[n-i-1], r[i]
	}
	return string(r)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("String: ")
	in, _ := reader.ReadString('\n')
	s := strings.TrimSuffix(in, "\n")
	fmt.Println("Reversed:", Reverse(s))

}
