package main

import (
	"flag"
	"fmt"
	"go_self-education/lessons"
)

func main() {
	// go run . -lesson <number>
	lesson := flag.Int("lesson", 0, "номер урока")
	flag.Parse()
	switch *lesson {
	case 82:
		lessons.RunLesson82()
	case 83:
		lessons.RunLesson83()
	case 84:
		lessons.RunLesson84()
	case 85:
		lessons.RunLesson85()
	default:
		fmt.Println("Такого урока нету")
	}
}
