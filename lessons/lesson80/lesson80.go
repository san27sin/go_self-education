/*
80. Задание: Readers
Реализуйте тип Reader, который выдаёт бесконечный поток ASCII символа ‘A’.
*/
package lesson80

import (
	"fmt"

	"golang.org/x/tour/reader"
)

type MyReader struct { 

}

func (r MyReader) Read(b []byte) (n int, err error) {
	for idx := range b {
		b[idx] = 'A'
	}
	return len(b), nil
}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func Run() {
	arr := make([]byte, 5)
	fmt.Println(arr)
	r := MyReader{}
	reader.Validate(r)
	r.Read(arr)
	fmt.Println(arr)
}