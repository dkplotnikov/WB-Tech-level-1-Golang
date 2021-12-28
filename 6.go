package main

// Реализовать все возможные способы остановки выполнения горутины.

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var wg sync.WaitGroup

	// Читающая из канала рутина завершит работу при закрытии канала
	c := make(chan int)
	wg.Add(1)
	go func() {
		fmt.Println("routine started")
		for i := range c {
			fmt.Printf("i've read %d\n", i)
		}
		fmt.Println("routine stopped")
		wg.Done()
	}()
	for i := 0; i < 5; i++ {
		c <- i
	}
	fmt.Println("closing chan")
	close(c)
	wg.Wait()

	// Отдельный канал, при отправке в который рутина завершает работу
	quit := make(chan bool)
	wg.Add(1)
	go func() {
		fmt.Println("routine started")
		i := 0
		for {
			select {
			case <-quit:
				fmt.Println("routine stopped")
				wg.Done()
				return
			default:
				i++
				fmt.Printf("tick: %d\n", i)
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()

	time.Sleep(500 * time.Millisecond)
	fmt.Println("sending to chan")
	quit <- true
	wg.Wait()

	// Отдельный канал, при закрытии которого рутина завершает работу
	// Так можно завершить работу нескольких рутин, в отличии от предыдущего варианта
	done := make(chan bool)
	f := func(n int) {
		fmt.Printf("routine %d started\n", n)
		i := 0
		for {
			select {
			case <-done:
				fmt.Printf("routine %d stopped\n", n)
				wg.Done()
				return
			default:
				i++
				fmt.Printf("routine %d: tick: %d\n", n, i)
			}
			time.Sleep(100 * time.Millisecond)
		}
	}
	wg.Add(2)
	go f(1)
	go f(2)

	time.Sleep(500 * time.Millisecond)
	fmt.Println("closing chan")
	close(done)
	wg.Wait()

	// context
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go func(ctx context.Context) {
		fmt.Println("routine started")
		i := 0
		for {
			select {
			case <-ctx.Done():
				fmt.Println("routine stopped")
				wg.Done()
				return
			default:
				i++
				fmt.Printf("tick: %d\n", i)
			}
			time.Sleep(100 * time.Millisecond)
		}
	}(ctx)
	time.Sleep(500 * time.Millisecond)
	fmt.Println("canceling")
	cancel()
	wg.Wait()

	// Завершение работы при условии (значение переменной)
	var a int64
	wg.Add(1)
	go func() {
		fmt.Println("routine started")
		for a < 1 {
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("a value: %d\n", a)
		}
		fmt.Println("routine stopped")
		wg.Done()
	}()

	time.Sleep(500 * time.Millisecond)
	fmt.Println("incrementing a")
	atomic.AddInt64(&a, 1)
	wg.Wait()

}
