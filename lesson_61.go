/*
61. Методы, продолжение
Вы также можете объявлять методы для неструктурных типов.

В этом примере мы видим числовой тип MyFloat с методом Abs.

Вы можете объявить метод только с получателем,
тип которого определен в том же пакете, что и метод.
Вы не можете объявить метод с получателем,
тип которого определен в другом пакете (включая встроенные типы, такие как int).
*/

package main

import "fmt"

type Celsius float64

func (c Celsius) ToFahrenheit() float64 {
	F := float64(c*9/5 + 32)
	return F
}

func (f float64) ToFahrenheit() float64 {
	return f*9/5 + 32
}

func main() {
	fmt.Println(Celsius(9.231).ToFahrenheit())
}
