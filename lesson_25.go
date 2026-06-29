/*
25. Цикл if с краткой инструкцией
Как и for, оператор if может начинаться с краткой инструкции, которая выполняется до проверки условия.

Переменные, объявленные в этой инструкции, доступны только до конца блока if.

(Попробуйте использовать v в последнем операторе return.)

Домашнее задание: конвертер температур

Напиши функцию convertTemp, которая принимает строку вида "100C" или "212F" и конвертирует температуру:
- C → перевести в Фаренгейт
- F → перевести в Цельсий

Требования:

1. Функция должна возвращать два значения: (float64, error)
2. В main используй if с краткой инструкцией для обработки результата
3. Если строка некорректная — вернуть ошибку

Примеры вызовов:
"100C" → 212.00 F
"32F"  → 0.00 C
"abcX" → ошибка
*/

package main

import (
	"fmt"
	"math"
	"strconv"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func completeHW(input string) (float64, error) {
	if n := input[len(input)-1]; n == 'C' || n == 'F' {
		num := input[:len(input)-1]
		C, err := strconv.ParseFloat(num, 64)
		if err != nil {
			return 0, err
		}
		if n == 'C' {
			return C*9/5 + 32, nil
		} else {
			return (C - 32) * 5 / 9, nil
		}
	} else {
		return 0, fmt.Errorf("некорректный ввод")
	}

}

func main() {
	// fmt.Println(
	// 	pow(3, 2, 10),
	// 	pow(3, 3, 20),
	// )

	// homework
	if result, err := completeHW("100C"); err == nil {
		fmt.Println(result)
	} else {
		fmt.Println("Ошибка:", err)
	}
}
