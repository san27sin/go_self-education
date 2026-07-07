/*
26. Условие if с else
Переменные, объявленные в краткой инструкции if, также доступны внутри всех блоков else.

(Оба вызова pow возвращают свои результаты до того, как начинается вызов fmt.Println в функции main.)
*/

package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// переменная v не доступна здесь
	// попытка вывести переменную v приведет к ошибке undefined: v
	// fmt.Println(v)

	return lim
}

func classify(temp float64) string {
	if diff := temp - 22; diff < -5 {
		return "холодно"
	} else if diff < 0 {
		return "прохладно"
	} else if diff == 0 {
		return "комфортно"
	} else if diff <= 5 {
		return "тепло"
	} else {
		return "жарко"
	}
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
	fmt.Println("Домашнее задание")
	fmt.Println(classify(10))
	fmt.Println(classify(18))
	fmt.Println(classify(22))
	fmt.Println(classify(26))
	fmt.Println(classify(35))
}
