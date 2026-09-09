/*
51. Мапа (Map)
Мапа (Map) сопоставляет ключи со значениями.

Нулевым значением мапы является nil. Nil мапа не содержит ключей, и в нее нельзя добавить новые ключи.

Функция make возвращает мапу заданного типа, инициализированную и готовую к использованию.
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "go is fun and go is powerful"
	txt := strings.Fields(text)
	counts := make(map[string]int)
	for _, v := range txt {
		counts[v]++
	}
	for key, value := range counts {
		fmt.Println("слово: \"", key, "\" употреблено столько ", value, "раз")
	}
}
