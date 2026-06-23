/*
22. Цикл for заменяет while
В Go цикл for заменяет while.

То, что в других языках делают через while, в Go пишут обычным for.

Для этого необходимо убрать начальную и завершающую инструкции, как в примере ниже.
*/

package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
