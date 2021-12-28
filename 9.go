package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из
// массива, во второй — результат операции x*2, после чего данные из второго
// канала должны выводиться в stdout.

// []int -> c1 -> /x*2/ -> c2 -> stdout

func main() {
	rand.Seed(time.Now().Unix())

	nums := make([]int, 10)
	for i := 0; i < 10; i++ {
		nums[i] = rand.Intn(10)
	}
	fmt.Println(nums)

	c1 := make(chan int)
	c2 := make(chan int)

	var wg sync.WaitGroup
	wg.Add(3)

	// Запись в c1 из слайса
	go func() {
		for _, num := range nums {
			c1 <- num
		}
		close(c1)
		wg.Done()
	}()

	// Чтение из c1 и запись нового значения в c2
	go func() {
		for num := range c1 {
			c2 <- num * 2
		}
		close(c2)
		wg.Done()
	}()

	// Чтение из c2 и запись в os.Stdout
	go func() {
		for num := range c2 {
			fmt.Printf("%d\n", num)
		}
		wg.Done()
	}()

	wg.Wait()

	// Читающие рутины завершат работу после закрытия каналов пишущими рутинами
}
