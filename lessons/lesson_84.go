/*
84. Параметры типа
Функции в Go могут быть написаны для работы с несколькими типами с использованием параметров типа.
Параметры типа функции указываются в квадратных скобках перед аргументами функции.

func Index[T comparable](s []T, x T) int

Это объявление означает, что s является срезом любого типа T,
который удовлетворяет встроенному ограничению comparable. x также представляет значение того же типа.

comparable — это полезное ограничение, которое позволяет использовать операторы == и != со значениями типа.
В этом примере мы используем его для сравнения значения со всеми элементами среза до тех пор, пока не будет найдено совпадение.
Эта функция Index работает для любого типа, который поддерживает сравнение.
*/

package lessons

import "fmt"

func Contains[T comparable](arr []T, x T) bool {
	for _, value := range arr {
		if value == x {
			return true
		}
	}
	return false
}

func RunLesson84() {
	arrInt := []int{1, 2, 3}
	xInt := 2
	fmt.Println(arrInt, "содержит -", xInt, "ответ", Contains(arrInt, xInt))
	arrStr := []string{"google", "yandex"}
	xStr := "zen"
	fmt.Println(arrStr, "содержит -", xStr, "ответ", Contains(arrStr, xStr))
	/*
		arrArr := [][]int{
		{1, 2, 3},
		{4, 4, 1},
	}
	*/
	//Contains(arrArr, 2)
	// invalid operation: value == x (incomparable types in type set)
}
