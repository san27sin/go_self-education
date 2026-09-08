/*
48. Итерация Range
Форма range в цикле for выполняет итерацию по слайсу или карте.

При переборе слайса возвращаются два значения для каждой итерации. Первое - это индекс, а второе - копия элемента по этому индексу.
*/

package main

import "fmt"

// Задание:
//  1. Создайте слайс fruits := []string{"apple", "banana", "cherry"}.
//  2. Переберите его через for i, v := range fruits и выведите индекс и
//     значение на каждой итерации.
//  3. Переберите ещё раз, но индекс не нужен — используйте for _, v := range
//     fruits и выведите только значения.
//  4. Переберите в третий раз, используя только индекс (без значения) —
//     for i := range fruits — и выведите fruits[i].

func main() {
	fmt.Println("#1")
	fruits := []string{"apple", "banana", "watermallen"}
	fmt.Println("fruits = ", fruits)
	fmt.Println("#2")
	for i, v := range fruits {
		fmt.Println("индекс = ", i, " ; ", "значение = ", v)
	}
	fmt.Println("#3")
	for _, v := range fruits {
		fmt.Println("значение = ", v)
	}
	fmt.Println("#4")
	for i, _ := range fruits {
		fmt.Println("индекс = ", i)
	}
}
