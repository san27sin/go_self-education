/*
59. Методы
В Go нет классов. Однако вы можете определять методы для типов.

Метод — это функция со специальным аргументом-получателем.

Получатель указывается в собственном списке аргументов между ключевым словом func и именем метода.
*/

package main

import "fmt"

type Counter struct {
	count int
}

func (c *Counter) Increment() {
	c.count++
}

func (c Counter) Value() int {
	return c.count
}

func (c *Counter) Reset() {
	c.count = 0
}

func main() {
	counter := Counter{}
	counter.Increment()
	counter.Increment()
	counter.Increment()
	fmt.Println(counter.Value())
	counter.Reset()
	fmt.Println(counter.Value())
}