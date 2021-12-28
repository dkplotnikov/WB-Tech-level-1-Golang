package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

/*
Реализовать бинарный поиск встроенными методами языка.
*/

func OrderedIntSlice(n int, fn func() int) []int {
	s := make([]int, n)
	s[0] = (rand.Intn(2)*2 - 1) * rand.Intn(100)
	for i := 1; i < n; i++ {
		s[i] = s[i-1] + fn()
	}
	return s
}

func PrintSlice(s []int) {
	str := make([]string, len(s))
	for i, x := range s {
		str[i] = strconv.Itoa(x)
	}
	fmt.Println("[" + strings.Join(str, " ") + "]")
}

func BinarySearch(s []int, x int) (c int) {
	l := 0
	r := len(s) - 1
loop:
	for l <= r {
		c = (l + r) / 2
		switch {
		case s[c] == x:
			break loop
		case s[c] < x:
			if s[l] < x {
				l = c + 1
			} else {
				r = c - 1
			}
		case s[c] > x:
			if s[l] > x {
				l = c + 1
			} else {
				r = c - 1
			}
		}
	}
	return c
}

func main() {
	rand.Seed(time.Now().Unix())

	delta := func() int {
		return -rand.Intn(10)
	}
	s := OrderedIntSlice(1_000_000, delta)
	//PrintSlice(s)

	x := s[rand.Intn(len(s))]
	fmt.Printf("search value = %d\n", x)
	i := BinarySearch(s, x)
	fmt.Printf("BinarySearch: found s[%d] = %d\n", i, s[i])
}
