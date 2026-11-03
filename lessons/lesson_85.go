/*
85. Универсальные типы (Дженерики)
Помимо универсальных функций, Go также поддерживает универсальные типы. Тип может быть параметризован с помощью параметра типа, что полезно для реализации универсальных структур данных.

Этот пример демонстрирует простое объявление типа для односвязного списка, содержащего значения любого типа.

В качестве упражнения добавьте некоторый функционал в эту реализацию списка.
*/

package lessons

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func add[T any](list *List[T], x T) {
	if list.next != nil {
		add(list.next, x)
	} else {
		list.next = &List[T]{
			val:  x,
			next: nil,
		}
	}
}

func printList[T any](list *List[T]) {
	if list == nil {
		return
	}
	fmt.Print(list.val)
	if list.next != nil {
		fmt.Print(" -> ")
		printList(list.next)
	}
}

func RunLesson85() {
	first := List[int]{
		next: nil,
		val:  1,
	}
	add(&first, 2)
	add(&first, 3)
	add(&first, 4)

	printList(&first)
	printList[int](nil)
	fmt.Println()
}
