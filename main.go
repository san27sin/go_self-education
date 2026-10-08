package main

import (
	"fmt"
	"os"
	"go_self-education/lessons/lesson79"
	"go_self-education/lessons/lesson80"
)

// lessons связывает номер урока с функцией, которая его запускает
var lessons = map[string]func(){
	"79": lesson79.Run,
	"80": lesson80.Run, // раскомментировать, когда у MyReader появится метод Read
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: go run . <номер урока>")
		return
	}
	run, ok := lessons[os.Args[1]]
	if !ok {
		fmt.Println("нет такого урока:", os.Args[1])
		return
	}
	run()
}
