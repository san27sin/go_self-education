/*
64. Методы и косвенное обращение через указатели
Сравнивая две предыдущие программы, вы могли заметить, что функции с аргументом-указателем должны принимать указатель:

var v Vertex ScaleFunc(v, 5) // Ошибка компиляции! ScaleFunc(&v, 5) // OK

в то время как методы с получателями-указателями принимают либо значение, либо указатель в качестве получателя при вызове:

var v Vertex v.Scale(5) // OK p := &v p.Scale(10) // OK

Для вызова v.Scale(5), даже though v является значением, а не указателем,
метод с получателем-указателем вызывается автоматически. То есть, для удобства,
Go интерпретирует выражение v.Scale(5) как (&v).Scale(5), поскольку метод Scale имеет получатель-указатель.
*/
package main

import "fmt"

type Inner struct {
	Data int
}

func (i *Inner) multiply(number int) {
	i.Data *= number
}

func main() {
	var inner = Inner{
		Data: 2,
	}
	fmt.Println(inner)
	inner.multiply(8)
	fmt.Println(inner)
}
