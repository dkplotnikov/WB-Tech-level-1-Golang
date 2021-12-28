package main

import (
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

// Разработать программу, которая перемножает, делит, складывает,
// вычитает две числовых переменных a, b, значение которых > 2^20.

func sum(a, b *big.Int) *big.Int {
	var c big.Int
	c.Add(a, b)
	return &c
}

func dif(a, b *big.Int) *big.Int {
	var c big.Int
	c.Neg(b)
	return sum(a, &c)
}

func mul(a, b *big.Int) *big.Int {
	var c big.Int
	c.Mul(a, b)
	return &c
}

func div(a, b *big.Int) *big.Int {
	var c big.Int
	c.Quo(a, b)
	return &c
}

func main() {
	rand.Seed(time.Now().UnixMilli())

	// a, b in [-10*2^50, 10*2^50]
	a := big.NewInt((rand.Int63n(21) - 10) * 1 << (rand.Int63n(31) + 20))
	b := big.NewInt((rand.Int63n(21) - 10) * 1 << (rand.Int63n(31) + 20))

	fmt.Printf("a = %v, b = %v\n", a, b)
	fmt.Printf("a + b = %v\n", sum(a, b))
	fmt.Printf("a - b = %v\n", dif(a, b))
	fmt.Printf("a * b = %v\n", mul(a, b))
	fmt.Printf("a / b = %v\n", div(a, b))
}
