/*
40. Слайсы являются ссылками на массивы
Слайс не хранит данные, а лишь ссылается на часть базового массива.
Поэтому изменение элементов в сайлсе приводит к изменению соотвествующих элементов в массиве, от которого этот слайс образован.

Другие слайсы, основанные на том же массиве, увидят эти изменения.

В данном случае мы не меняем исходный массив names, а только слайс b, но тем не менее все слайсы (и a и b) видят и работают с этими изменениями.
*/

package main

import "fmt"

func makeTheory() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

func makeHomework() {
	numbers := [6]int{10, 20, 30, 40, 50, 60}
	first := numbers[0:4]
	second := numbers[3:5]
	fmt.Println("array = ", numbers)
	fmt.Println("first = ", first)
	fmt.Println("second = ", second)
	second[0] = 999
	fmt.Println("second[0] = 999")
	fmt.Println("array = ", numbers)
	fmt.Println("first = ", first)
	fmt.Println("second = ", second)
}

func main() {
	fmt.Println("\nТеория")
	makeTheory()
	fmt.Println("\nДомашнее задание")
	makeHomework()
}
