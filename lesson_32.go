/*
32. Накопление отложенных вызовов
Отложенные вызовы функций помещаются в стек. При возврате из функции ее отложенные вызовы выполняются в порядке “последним пришел — первым ушел”.
*/

package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}