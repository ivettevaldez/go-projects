package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

// Image returns an implementation of image.Image instead of a slice of data.
type Image struct {
	Width, Height int
	colr          uint8
}

// Bounds should return a image.Rectangle, like image.Rect(0, 0, w, h).
func (r *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, r.Width, r.Height)
}

// ColorModel should return color.RGBAModel.
func (r *Image) ColorModel() color.Model {
	return color.RGBAModel
}

// At should return a color; the value v in the last picture generator corresponds to color.RGBA{v, v, 255, 255} in this one.
func (r *Image) At(x, y int) color.Color {
	return color.RGBA{r.colr + uint8(x), r.colr + uint8(y), 255, 255}
}

func main() {
	m := Image{100, 100, 128}
	pic.ShowImage(&m)
}
