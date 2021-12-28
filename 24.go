package main

import (
	"fmt"
	"math"
)

// Разработать программу нахождения расстояния между двумя точками, которые
// представлены в виде структуры Point с инкапсулированными параметрами x,y и
// конструктором.

// Поля x, y будут доступны в других пакетах
// только через геттеры
type Point struct {
	x float64
	y float64
}

// Конструктор
func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

// Геттеры
func (p *Point) GetX() float64 {
	return p.x
}

func (p *Point) GetY() float64 {
	return p.y
}

// Distance возвращает расстояние между точками a, b *Point
func Distance(a, b *Point) float64 {
	x := a.GetX() - b.GetX()
	y := a.GetY() - b.GetY()
	return math.Sqrt(x*x + y*y)
}

func main() {
	a := NewPoint(-1.57, 5.64)
	b := NewPoint(12.2, 2.01)

	fmt.Println(Distance(a, b))

}
