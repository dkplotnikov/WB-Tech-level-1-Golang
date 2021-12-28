package main

import (
	"fmt"
	"math/rand"
)

// Удалить i-ый элемент из слайса.

func RandIntSlice(k int, n int) []int {
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = (rand.Intn(2)*2 - 1) * rand.Intn(k)
	}
	return s
}

func RemoveElement(s []int, i int) []int {
	ns := make([]int, 0, len(s)-1)
	ns = append(ns, s[:i]...)
	return append(ns, s[i+1:]...)
}

func main() {

	s := RandIntSlice(10, 10)
	fmt.Printf("%v\n", s)

	ns := RemoveElement(s, 4)
	fmt.Printf("%v\n", ns)

}
