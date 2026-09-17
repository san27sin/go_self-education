/*
60. Методы - это функции
Запомните: метод - это просто функция с аргументом-получателем.
*/

package main

import (
	"fmt"	
	"math"
)

type Point struct {
	X, Y float64
}

func (p Point) Distance(o Point) float64 {
	return math.Sqrt(math.Pow(p.X - o.X, 2) + math.Pow(p.Y - o.Y, 2))
}

func DistanceFunc(p, o Point) float64 {
	return math.Sqrt(math.Pow(p.X - o.X, 2) + math.Pow(p.Y - o.Y, 2))
}

func main() {
	p1 := Point{ X: 1.234, Y: 3.624 }
	p2 := Point{ X: 115.234, Y: 343.624 }
	fmt.Println(p1.Distance(p2))
	fmt.Println(DistanceFunc(p1, p2))
}