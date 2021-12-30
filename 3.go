package main

// Дана последовательность чисел: 2,4,6,8,10. Найти сумму их
// квадратов(2^2+3^2+4^2….) с использованием конкурентных вычислений.

import (
	"fmt"
	"sync"
	"sync/atomic"
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

	// Вычисление суммы в рутинах
	var wg sync.WaitGroup
	var sum1 int64
	for _, num := range nums {
		wg.Add(1)
		go func(num int) {
			atomic.AddInt64(&sum1, int64(num*num))
			wg.Done()
		}(num)
	}
	wg.Wait()
	fmt.Println(sum1)

}
