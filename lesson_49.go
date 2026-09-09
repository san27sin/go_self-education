/*
49. Range продолжение
Вы можете пропустить индекс или значение, присвоив их символу _.

for i, _ := range pow for _, value := range pow

Если вам нужен только индекс, вы можете опустить вторую переменную.

for i := range pow
*/

package main

import ("fmt")

func main() {
	pow := make([]int, 10)
    for i := range pow {
    	pow[i] = 1 << uint(i) // == 2**i
    }
    for _, value := range pow {
        fmt.Printf("%d\n", value)
    }
}
