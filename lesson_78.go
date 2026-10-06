/*
78. Задание: Ошибки
Скопируйте вашу функцию Sqrt из предыдущего задания и измените ее так, чтобы она возвращала значение ошибки.

Sqrt должна возвращать ненулевое значение ошибки, когда получает отрицательное число, так как она не поддерживает комплексные числа.

# Создайте новый тип

type ErrNegativeSqrt float64

и сделайте его ошибкой, добавив метод

func (e ErrNegativeSqrt) Error() string

чтобы вызов ErrNegativeSqrt(-2).Error() возвращал “cannot Sqrt negative number: -2”.

Примечание: Вызов fmt.Sprint(e) внутри метода Error приведет к бесконечному циклу. Вы можете избежать этого, преобразовав e сначала: fmt.Sprint(float64(e)). Почему?

Измените вашу функцию Sqrt, чтобы она возвращала значение ErrNegativeSqrt при получении отрицательного числа.
*/
package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Отрицательное число")
	}

	return math.Sqrt(x), nil
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return "cannot Sqrt negative number: " + strconv.FormatFloat(float64(e), 'f', -1, 64)
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
	var test ErrNegativeSqrt = -1.4
	fmt.Println(ErrNegativeSqrt(2).Error())
	fmt.Println(test.Error())
}
