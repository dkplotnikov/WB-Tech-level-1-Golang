package main

// Дана переменная int64.
// Разработать программу которая устанавливает i-й бит в 1 или 0.

import (
	"fmt"
	"math/rand"
	"time"
	"unsafe"
)

// On устанавливает i-й бит в 1, используя побитовое ИЛИ
// с маской типа 00100 с единицей в позиции i
func On(n int64, i int) int64 {
	return n | 1<<(i-1)
}

// Off устанавливает i-й бит в 0, используя побитовое И
// с маской типа 11011 с нулем в позиции i
func Off(n int64, i int) int64 {
	return n & ^(1 << (i - 1))
}

func main() {
	rand.Seed(time.Now().Unix())

	// Случайное число типа int64
	// Первый rand генерирует -1 или 1, отображая [0,1] на [-1,1]
	p := (rand.Int63n(2)*2 - 1) * rand.Int63()

	on := On(p, 5)
	off := Off(p, 5)

	fmt.Printf("%b\n", *(*uint64)(unsafe.Pointer(&p)))
	fmt.Printf("%b\n", *(*uint64)(unsafe.Pointer(&on)))
	fmt.Printf("%b\n", *(*uint64)(unsafe.Pointer(&off)))

}
