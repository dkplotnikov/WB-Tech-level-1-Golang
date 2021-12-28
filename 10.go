package main

// Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0,
// 15.5, 24.5, -21.0, 32.5. Объединить данные значения в группы с шагом в 10
// градусов. Последовательность в подмножноствах не важна.
// Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
//
// ! Пример не вполне соответствует условию, так как в группу "0" попадут
// значения температур (-10, 10), т.е. для этой группы шаг получится равным 20.
// Ниже реализованы оба варианта.

import (
	"fmt"
	"math"
)

// Чтобы выдавать данные в порядке возрастания температур,
// при заполнении словаря сразу найдем и запомним min и max
// T - словарь множеств
type Temps struct {
	T    map[int]map[float64]struct{}
	minT int
	maxT int
}

// Hash вычисляет key для словаря (вроде hash-функции для hash-таблицы)
func Hash1(f float64) int {
	return int(f/10) * 10
}

// Определим группы следующим образом T = {t: t in [T, T+10)}.
// Пример: -30:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
func Hash2(f float64) int {
	return int(math.Floor(f/10)) * 10
}

// Fill заполняет map, использую выбранную hash-функцию
func (ts *Temps) Fill(temps []float64, fn func(float64) int) {
	ts.T = make(map[int]map[float64]struct{})
	ts.minT = fn(temps[0])
	ts.maxT = ts.minT

	for _, t := range temps {
		key := fn(t)

		// minimax
		if key < ts.minT {
			ts.minT = key
		}
		if key > ts.maxT {
			ts.maxT = key
		}

		// Если map не содержит key, инициализируется map
		if _, ok := ts.T[key]; !ok {
			ts.T[key] = make(map[float64]struct{})
		}
		ts.T[key][t] = struct{}{}
	}
}

// Запись map в Stdout
func (ts *Temps) Print() {
	for t := ts.minT; t <= ts.maxT; t += 10 {
		if temp, ok := ts.T[t]; ok {
			fmt.Printf("%d:{", t)
			i := 0
			for t := range temp {
				if i > 0 {
					fmt.Print(", ")
				}
				fmt.Printf("%.1f", t)
				i++
			}
			fmt.Print("} ")
		}
	}
	fmt.Println()
}

func main() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, -7.8, 3.5, 0.0}

	ts1 := &Temps{}
	ts1.Fill(temps, Hash1)
	fmt.Print("Hash1: ")
	ts1.Print()

	ts2 := &Temps{}
	ts2.Fill(temps, Hash2)
	fmt.Print("Hash2: ")
	ts2.Print()

}
