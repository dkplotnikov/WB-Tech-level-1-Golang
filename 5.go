package main

// Разработать программу, которая будет последовательно отправлять значения
// в канал, а с другой стороны канала — читать. По истечению N секунд
// программа должна завершаться.

import (
	"fmt"
	"time"
)

func main() {
	n := 1

	// В канал stop будет записано текущее время по истечению n секунд
	stop := time.After(time.Duration(n) * time.Second)

	c := make(chan int)

	// Запись в канал
	go func() {
		i := 0
		for {
			fmt.Printf("write %d\n", i)
			c <- i
			time.Sleep(100 * time.Millisecond)
			i++
		}
	}()

	// Чтение из канала
	go func() {
		for v := range c {
			fmt.Printf("read %d\n", v)
		}
	}()

	// main в блоке до истечения времени, переданного в time.After
	<-stop

	// Ожидание в течение n секунд
	// time.Sleep(time.Duration(n) * time.Second)
}
