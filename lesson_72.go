/*
72. Пустой интерфейс
Тип интерфейса, который определяет нуль методов, известен как пустой интерфейс:

interface{}

Пустой интерфейс может содержать значения любого типа.
(Каждый тип реализует как минимум нуль методов.)

Пустые интерфейсы используются кодом, который обрабатывает значения неизвестного типа.
Например, fmt.Print принимает любое количество аргументов типа interface{}.
*/

package main

import "fmt"

func print77(i interface{}) {
	fmt.Printf("value = %v ; type = %T \n", i, i)
}

func main() {
	var value interface{}
	print77(value)
	value = 5
	print77(value)
	value = "Sasha"
	print77(value)
}
