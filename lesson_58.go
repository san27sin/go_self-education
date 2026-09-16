/*
Давайте немного развлечемся с функциями.

Реализуйте функцию fibonacci, которая возвращает функцию (замыкание), возвращающую последовательные числа Фибоначчи (0, 1, 1, 2, 3, 5, …).
*/

package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.

func fibonacci() func() int {
	counter := 0
	slice := make([]int, 1)
	return func() int {
		if counter == 0 {
			slice = append(slice, 1)
			counter++
			return slice[counter-1]
		}
		number := slice[counter-1] + slice[counter]
		slice = append(slice, number)
		counter++
		return slice[counter-1]
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
