/*
52. Литералы мап
Литералы мап похожи на литералы структур, но ключи являются обязательными.

Задание:
1. Объявить структуру Vertex с полями Lat, Long float64.
2. Объявить мапу map[string]Vertex и заполнить её литералом (2-3 записи).
3. Вывести мапу через fmt.Println.
4. Доп: добавить ещё одну запись после объявления (m["..."] = Vertex{...}).
*/

package main

import ("fmt")

type Vertex struct {
	Lat, Long float64
}

func main() {
	var mapa = map[string]Vertex {
		"one": {Lat: 10.123, Long: 11.123 },
		"two": { Lat: 2.123, Long: 3.123 },
	}
	fmt.Println(mapa)
	mapa["three"] = Vertex{
		Lat: 5.123,
		Long: 2.111,
	}
	fmt.Println(mapa)
}
