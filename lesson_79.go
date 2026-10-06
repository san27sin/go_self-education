/*
79. Ридеры (Readers)
Пакет io определяет интерфейс io.Reader, который представляет конечную точку для чтения потока данных.

Стандартная библиотека Go содержит множество реализаций этого интерфейса, включая файлы, сетевые соединения, компрессоры, 
шифраторы и другие.

Интерфейс io.Reader имеет метод Read:

func (T) Read(b []byte) (n int, err error)

Read заполняет предоставленный байтовый срез (b) данными и возвращает количество записанных байт (n) и значение ошибки (err). 
Он возвращает ошибку io.EOF, когда поток данных заканчивается.

Пример кода создает strings.Reader и потребляет его вывод по 8 байт за раз.


*/

package main

import (
    "fmt"
    "io"
    "strings"
)

func main() {
    r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}