package main

import (
	"golang.org/x/tour/pic"
)

// Pic returns a slice of length dy, each element of which is a slice of dx 8-bit unsigned integers.
func Pic(dx, dy int) [][]uint8 {
	result := make([][]uint8, dy)

	for i := range result {
		result[i] = make([]uint8, dx)

		for j := 0; j < dx; j++ {
			result[i][j] = uint8((i + j) / 2)
		}
	}

	return result
}

func main() {
	pic.Show(Pic)
}
