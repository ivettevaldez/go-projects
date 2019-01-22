package main

import (
	"fmt"
	"math"
)

// ErrNegativeSqrt returns a non-nil error value when given a negative number.
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

// Sqrt returns the number s for which z² is most nearly the given number x.
func Sqrt(x float64) (float64, error) {
	z := float64(1.0)
	s := float64(0)

	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	for {
		z = z - (z*z-x)/(2*z)
		if math.Abs(s-z) < 1e-15 {
			break
		}
		s = z
	}

	return s, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
