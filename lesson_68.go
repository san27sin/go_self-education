/*
68. Интерфейсы реализуются неявно
Тип реализует интерфейс путем реализации его методов. 
Нет явного объявления намерения, нет ключевого слова “implements”.

Неявные интерфейсы разделяют определение интерфейса от его реализации, 
которая может затем появиться в любом пакете без предварительной договоренности.
*/

package main

import "fmt"

type Namer interface {
	Name() string
}

type Cat struct {
	Nick string
}

func (c Cat) Name() string {
	return "Привет, " + c.Nick + "!"
}

func Print(n Namer) {
	fmt.Println(n.Name())
}

func main() {
	c := Cat{
		Nick: "Петух",
	}
	Print(c)
}