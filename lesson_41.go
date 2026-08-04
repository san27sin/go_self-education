/*
41. Литералы слайсов
Литерал слайса похож на литерал массива, но без указания длины.

Это литерал массива:

[3]bool{true, true, false}

А это создает такой же массив, а затем слайс, который ссылается на него:

[]bool{true, true, false}
*/

package main

import "fmt"

func makeTheory() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

func makeHomework() {
	sliceInt := []int{5, 7}
	sliceString := []string{"понедельник", "вторник"}
	fmt.Println("sliceInt = ", sliceInt, " / ", " длина = ", len(sliceInt))
	fmt.Println("sliceString = ", sliceString, " / ", " длина = ", len(sliceString))

	sliceStruct := []struct {
		day  int
		name string
	}{
		{0, "Понедельник"},
		{1, "Вторник"},
	}
	fmt.Println("sliceStruct = ", sliceStruct)
}

func main() {
	fmt.Println("Теория")
	makeTheory()
	fmt.Println("\nПрактика")
	makeHomework()
}
