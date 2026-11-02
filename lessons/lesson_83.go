/*
83. Упражнение: Изображения
Помните генератор изображений, который вы писали ранее?
Давайте создадим еще один, но на этот раз он будет возвращать реализацию image.Image вместо среза данных.

Определите свой собственный тип Image, реализуйте необходимые методы и вызовите pic.ShowImage.

Bounds должен возвращать image.Rectangle, например image.Rect(0, 0, w, h).

ColorModel должен возвращать color.RGBAModel.

At должен возвращать color; значение v из предыдущего генератора изображений соответствует color.RGBA{v, v, 255, 255} в этом.
*/

package lessons

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	w int
	h int
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.w, i.h)
}

func (i Image) At(x, y int) color.Color {
    r := uint8((x + y) / 2)
    g := uint8((x + y) / 2)
	return color.RGBA{r, g, 255, 255}
}

func RunLesson83() {
	m := Image{w: 100, h: 200}
	pic.ShowImage(m)
}
