package main

// Реализовать постоянную запись данных в канал (главный поток). Реализовать
// набор из N воркеров, которые читают произвольные данные из канала и выводят
// в stdout. Необходима возможность выбора количества воркеров при старте.
//
// Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать
// способ завершения работы всех воркеров.

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
)

// enableWorkers запускает n рутин, читающих из канала c
func enableWorkers(n int, c chan int) {
	for i := 0; i < n; i++ {
		go func() {
			// Читает из канала, пока он не закрыт
			for v := range c {
				fmt.Println(v)
			}
		}()
	}
}

func main() {
	c := make(chan int)

	// Канал для записи сигнала ОС
	osSignals := make(chan os.Signal, 1)

	// Если поступил SIGINT, генерируемый Ctrl+C, в канал будет записано значение
	signal.Notify(osSignals, syscall.SIGINT)

	// Задаем количество воркеров и запускаем их
	n := 10
	enableWorkers(n, c)

	// Бесконечный цикл записи в канал случайных чисел
loop:
	for {
		select {
		case <-osSignals:
			close(c)
			fmt.Println("shutdown")
			break loop
		default:
			c <- rand.Int()
		}
	}
}
