package main

// Дана последовательность чисел: 2,4,6,8,10. Найти сумму их
// квадратов(22+32+42….) с использованием конкурентных вычислений.

import (
	"fmt"
)

func main() {
	var nums = []int{2, 4, 6, 8, 10}

	// Канал для записи результатов вычислений рутин
	c := make(chan int)
	for _, num := range nums {
		go func(num int) {
			c <- num * num
		}(num)
	}

	// Вычисляем сумму, считывая квадраты из канала 5 раз
	sum := 0
	for range nums {
		sum += <-c
	}

	fmt.Println(sum)
}
