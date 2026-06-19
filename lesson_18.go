/*
18. Константы
Константы объявляются как переменные, но с использованием ключевого слова const.

Константы могут быть символьными, строковыми, логическими или числовыми значениями.

Константы нельзя объявлять с использованием синтаксиса :=.
*/

package main

import "fmt"

const Pi = 3.14

func main() {
	fmt.Println(Pi)
	completeHW()
}

/*
Задача:
Объявите константу SecondsInMinute со значением 60.
В функции main посчитайте и выведите, сколько секунд в 5 минутах,
используя эту константу (без "магических чисел").
*/

func completeHW() {
	const SecondsInMinute = 60
	const fiveMinutes = 5
	var result = SecondsInMinute * fiveMinutes
	fmt.Println("В 5 минутах должно быть секунд - ", result)
}
