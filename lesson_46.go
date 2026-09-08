/*
46. Слайсы слайсов
Слайсы могут содержать любой тип, включая другие слайсы.
*/
package main

import (
	"fmt"
	"strings"
)

func makeTheory() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

// Задание (5 минут):
//  1. Создайте слайс слайсов names := [][]string{{"Anna", "Bob"}, {"Max"}}.
//  2. Пройдитесь по нему циклом и выведите каждую внутреннюю строку через
//     strings.Join(names[i], ", ").
func makeHW() {
	names := [][]string{
		{"Anna", "Bob"},
		{"Max"},
	}
	for idx_1 := 0; idx_1 < len(names); idx_1++ {
		str := strings.Join(names[idx_1], ", ")
		fmt.Println(str)
	}

}

func main() {
	fmt.Println("Теория")
	makeTheory()
	fmt.Println("Задание")
	makeHW()
}
