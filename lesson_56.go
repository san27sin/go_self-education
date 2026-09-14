/*
56. Функции как значения
Функции тоже являются значениями. Их можно передавать так же, как и другие значения.

Значения функций могут использоваться в качестве аргументов функций и возвращаемых значений.
*/

package main

import "fmt"

func compose(f, g func(int) int) func(int) int {
	cFunc := func(a int) int {
		reply_1 := g(a)
		reply_2 := f(reply_1)
		return reply_2
	}
	return cFunc
}

/*
Напиши функцию compose(f, g func(int) int) func(int) int, 
которая возвращает новую функцию, применяющую сначала g, 
затем f к результату (compose(f, g)(x) == f(g(x))). 
В main создай две простые функции int → int (например, удвоение и прибавление единицы), 
скомпонуй их и выведи результат для пары значений.
*/
func main() {
	funcL := func(a int) int {
		reply := a * 2
		fmt.Println(reply)
		return reply
	}
	funcO := func(b int) int {
		reply := b + 1
		fmt.Println(reply)
		return reply
	}
	funcI := compose(funcL, funcO)
	reply := funcI(9)
	fmt.Println(reply)
}