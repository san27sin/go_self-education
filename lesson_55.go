/*
55. Упражнение: Мапа
Реализуйте функцию WordCount. Она должна возвращать мапу с количеством вхождений каждого “слова” в строке s. Функция wc.Test запускает набор тестов для предоставленной функции и выводит успех или неудачу.

Вам может пригодиться функция strings.Fields.
*/

package main

import (
	"fmt"
	"strings"
)

func WordCount(str string) map[string]int {
	excludes := map[rune]int{ ',': 0, ' ': 1, '.': 2}
	words := strings.FieldsFunc(str, func(r rune) bool {
		_, has := excludes[r]
		return has
	})
	wordsCounts := make(map[string]int)
	for _, value := range words {
		wordsCounts[value]++
	}
	return wordsCounts
}

func main() {
	fmt.Println(WordCount("go is go, maybe no, or go."))
}