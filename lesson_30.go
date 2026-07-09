/*
30. Switch без условия
Switch без условия эквивалентен switch true.

Такая конструкция может быть удобным способом записи длинных цепочек if-then-else.
*/

package main

import (
    "fmt"
    "time"
)

func main() {
    t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Println("Good morning!")
    case t.Hour() < 17:
        fmt.Println("Good afternoon.")
    default:
        fmt.Println("Good evening.")
    }
}