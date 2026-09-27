/*
70. Значения интерфейсов с nil базовыми значениями
Если конкретное значение внутри интерфейса само является nil, метод будет вызван с nil получателем.

В некоторых языках это вызвало бы исключение нулевого указателя, 
но в Go принято писать методы, которые корректно обрабатывают вызов с nil получателем (как метод M в этом примере).

Обратите внимание, что значение интерфейса, содержащее nil конкретное значение, само по себе не является nil.
*/

package main

import "fmt"

type Logger interface {
	Log() string
}

type FileLogger struct {
	Path string
}

func (f *FileLogger) Log() string {
	if f == nil {
		return "[no-op] <msg>"
	}
    return f.Path
}

func Print234(l Logger) {
	fmt.Println(l.Log())
}

func main() {
	var logger Logger
	var logger213 *FileLogger
	logger = logger213
	Print234(logger)
}