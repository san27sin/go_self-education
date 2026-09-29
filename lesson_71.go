/*
71. Nil значения интерфейсов
Nil значение интерфейса не содержит ни значения, ни конкретного типа.

Вызов метода на nil интерфейсе вызывает ошибку времени выполнения,
потому что в кортеже интерфейса нет типа,
указывающего, какой конкретный метод вызывать.
*/

package main

import (
	"fmt"
	"slices"
)

type test interface {
	m()
}

type jo struct {
	i string
}

func (j *jo) m() {
	fmt.Printf("(%v, %T)\n", j.i, j.i)
}

func main() {
	var l = jo{i: "salam"}
	l.m()
	var w test
	if w == nil {
		fmt.Println("<nil>")
	} else {
		w.m()
	}

	// работы со слайсами
	testSlice := []int{9, 3, 8}
	slices.Sort(testSlice)
	fmt.Println(testSlice)
}
