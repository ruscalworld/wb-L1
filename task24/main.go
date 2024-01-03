package main

import (
	"fmt"
	"math"
)

// Point - структура, хранящая координаты точки на плоскости.
type Point struct {
	x, y float64
}

// NewPoint - конструктор для Point.
func NewPoint(x float64, y float64) *Point {
	return &Point{x: x, y: y}
}

func (p *Point) Distance(other *Point) float64 {
	// Считаем евклидово расстояние - корень из суммы квадратов разностей координат.
	return math.Sqrt(math.Pow(p.x-other.x, 2) + math.Pow(p.y-other.y, 2))
}

func main() {
	// Две вершины прямоугольного треугольника со сторонами 3, 4 и 5,
	// лежащие на гипотенузе.
	a := NewPoint(1, 1)
	b := NewPoint(4, 5)

	fmt.Println(a.Distance(b)) // 5
}
