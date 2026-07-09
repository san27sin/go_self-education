/*
29. Порядок выполнения switch
В операторе switch условия case проверяются сверху вниз, выполнение прекращается при первом совпадении.

Например:

switch i { case 0: case f(): }

не вызовет функцию f, если i==0.

Примечание: Время в Go playground всегда начинается с 2009-11-10 23:00:00 UTC.
Значение этой даты оставлено в качестве упражнения для читателя.
*/

package main

import (
	"fmt"
)

func getDay(day int) {
	switch day {
	case 0:
		fmt.Println("Today.")
	case 1:
		fmt.Println("Tomorrow.")
	case 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

func main() {
	fmt.Println("When's Saturday?")
	var day int
	fmt.Scan(&day)
	getDay(day)
}
