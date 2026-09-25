/*
67. Интерфейсы

Тип интерфейса определяется как набор сигнатур методов.

Значение типа интерфейса может содержать любое значение, которое реализует эти методы.
*/

package main

/*
1. Интерфейс Shape с двумя методами:
   - Area() float64 — площадь;
   - Perimeter() float64 — периметр.
2. Два типа:
   - Rect с полями W, H float64. Методы объяви с ресивером-значением;
   - Circle с полем R float64. Методы объяви с ресивером-указателем.
3. Функцию describe(s Shape), которая печатает площадь и периметр фигуры.
4. В main создай var s Shape и по очереди положи в неё прямоугольник и круг, каждый раз вызывая describe(s).

Вопросы для самопроверки (ответь на них экспериментом, а не догадкой):
- Скомпилируется ли s = Rect{3, 4}? А s = &Rect{3, 4}?
- Скомпилируется ли s = Circle{1}? А s = &Circle{1}?
- Какой текст ошибки выдаст компилятор в «плохом» случае?
*/

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rect struct {
	W, H float64
}

func (r Rect) Area() float64 {
	return r.H * r.W
}

func (r Rect) Perimeter() float64 {
	return 2 * (r.H + r.W)
}

type Circle struct {
	R float64
}

func (c *Circle) Area() float64 {
	return math.Pi * math.Pow(c.R, 2)
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.R
}

func describe(s Shape) {
	fmt.Println("Площадь = ", s.Area(), "m^2")
	fmt.Println("Периметр = ", s.Perimeter(), "m")
}

func main() {
	var s Shape = Rect{
		W: 5,
		H: 10,
	}
	describe(s)
	s = &Circle{
		R: 7,
	}
	describe(s)
}
