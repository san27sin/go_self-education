/*
82. Интерфейс Image
Пакет image определяет интерфейс Image:

package image

type Image interface { ColorModel() color.Model Bounds() Rectangle At(x, y int) color.Color }

Примечание: возвращаемое значение Rectangle метода Bounds на самом деле является image.Rectangle, 
поскольку объявление находится внутри пакета image.

(Подробности смотрите в документации.)

Типы color.Color и color.Model также являются интерфейсами, но мы проигнорируем это, 
используя предопределённые реализации color.RGBA и color.RGBAModel. Эти интерфейсы и типы определены в пакете image/color.
*/

package lessons

import (
	"fmt"
	"image"
)

func RunLesson82() {
    m := image.NewRGBA(image.Rect(0, 0, 100, 100))
    fmt.Println(m.Bounds())
    fmt.Println(m.At(0, 0).RGBA())
}