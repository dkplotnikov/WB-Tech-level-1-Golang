package main

import (
	"fmt"
	"time"
)

// Реализовать собственную функцию sleep.

func sleep1(sec int) {
	<-time.After(time.Duration(sec) * time.Second)
}

func sleep2(sec int) {
	t := time.Now()
	for time.Now().Sub(t).Seconds() < float64(sec) {
	}
}

func main() {
	sec := 3
	fmt.Printf("start sleeping for %d seconds\n", sec)
	sleep1(sec)
	fmt.Println("awake")

	fmt.Printf("start sleeping for %d seconds\n", sec)
	sleep2(sec)
	fmt.Println("awake")
}
