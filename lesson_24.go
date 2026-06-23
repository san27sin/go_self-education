/*
24. Условие if
Условный оператор if в Go похож на цикл for: выражение не нужно заключать в круглые скобки ( ), но фигурные скобки { } обязательны.

Домашнее задание
Классификатор температуры
Напишите функцию classifyTemp(t float64) string, которая принимает температуру в градусах Цельсия и возвращает её описание:

ниже 0 → "мороз"
от 0 до 15 (включительно) → "прохладно"
от 15 до 30 (включительно) → "тепло"
выше 30 → "жара"
В main вызовите функцию для нескольких значений и выведите результаты.
*/

package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func classifyTemp(t float64) string {
	if t < 0 {
		return "мороз"
	} else if t <= 15 {
		return "прохладно"
	} else if t <= 30 {
		return "тепло"
	}
	return "жара"
}

func main() {
	// fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(classifyTemp(-1))
	fmt.Println(classifyTemp(0))
	fmt.Println(classifyTemp(13))
	fmt.Println(classifyTemp(20))
	fmt.Println(classifyTemp(32))
}
