package main

import (
	"image"
	"image/color"
	"math/rand"

	"golang.org/x/tour/pic"
)

type MyImage struct {
	X, Y int
}

func (img MyImage) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.X, img.Y)
}

func (img MyImage) ColorModel() color.Model {
	return color.RGBAModel
}

func (img MyImage) At(x, y int) color.Color {
	v := uint8(rand.Intn(255))
	// returning Red, Green random with blue fixed
	return color.RGBA{v, v, 255, 255}
}

func main() {
	i := MyImage{X: 100, Y: 100}

	pic.ShowImage(i)

}
