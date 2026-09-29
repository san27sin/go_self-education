/*
74. Переключатели типов
Переключатель типов - это конструкция, которая позволяет выполнять несколько утверждений типа последовательно.

Переключатель типов похож на обычный оператор switch, но case в переключателе типов указывают типы (не значения), 
и эти значения сравниваются с типом значения, содержащегося в данном значении интерфейса.

switch v := i.(type) {
case T:
// здесь v имеет тип T
case S:й
// здесь v имеет тип S
default:
// нет совпадения; здесь v имеет тот же тип, что и i
}
Объявление в переключателе типов имеет тот же синтаксис, что и утверждение типа i.(T), 
но конкретный тип T заменяется ключевым словом type.

Этот оператор switch проверяет, содержит ли значение интерфейса i значение типа T или S. 
В каждом из случаев T и S переменная v будет иметь тип T или S соответственно и содержать значение из i. 
В случае default (когда нет совпадения) переменная v имеет тот же тип интерфейса и значение, что и i.
*/

package main

import (
	"errors"
	"fmt"
	"strings"
)

type Temp float64

func (t Temp) String() string {
	return fmt.Sprintf("%.1f°C", float64(t))
}

func classify(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("int:", v, "квадрат:", v*v)
	case string:
		fmt.Printf("string: %q, в верхнем регистре: %s\n", v, strings.ToUpper(v))
	case fmt.Stringer:
		fmt.Println("Stringer:", v.String())
	case float32, float64:
		// Здесь v имеет тип interface{}, поэтому v * 2 не скомпилируется
		fmt.Println("дробное число:", v)
	case nil:
		fmt.Println("пустой интерфейс")
	case error:
		fmt.Println("ошибка:", v.Error())
	default:
		fmt.Printf("неизвестный тип: %T\n", v)
	}
}

func main() {
	classify(7)
	classify("go")
	classify(3.5)
	classify(float32(1.5))
	var nilInterface interface{}
	classify(nilInterface)
	classify(errors.New("упс"))
	classify([]int{1, 2})
	classify(true)
	classify(Temp(21.5))
}
