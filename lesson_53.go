/*
53. Литералы мап, продолжение
Если вышестоящий тип является просто именем типа, вы можете опустить его в элементах литерала.
*/

package main

import "fmt"

type Vertex struct {
    Lat, Long float64
}

var m = map[string]Vertex{
    "Bell Labs": {40.68433, -74.39967},
    "Google":    {37.42202, -122.08408},
}

func main() {
    fmt.Println(m)
}