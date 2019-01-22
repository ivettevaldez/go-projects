package main

import (
	"fmt"
	"math"
)

// Sqrt returns the number s for which z² is most nearly the given number x.
func Sqrt(x float64) float64 {
	z := float64(1.0)
	s := float64(0)

	fmt.Println("Results...")
	for {
		z = z - (z*z-x)/(2*z)
		if math.Abs(s-z) < 1e-15 {
			break
		}
		s = z
		fmt.Println(s)
	}

	return s
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println("\nLibrary result:", math.Sqrt(2))
}
