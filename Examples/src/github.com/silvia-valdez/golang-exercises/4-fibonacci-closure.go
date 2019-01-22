package main

import "fmt"

// fibonacci returns a function that returns an int.
func fibonacci() func() int {
	a, b := 0, 1
	index := 0

	return func() int {
		index++

		if index == 1 {
			return a
		}
		if index == 2 {
			return b
		}

		c := a + b
		a = b
		b = c

		return c
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
