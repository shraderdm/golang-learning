package main

import (
	"fmt"
	"math"
)

func main() {
	ints := []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, i := range ints {
		result := math.Mod(i, 2)
		if result != 0 {
			fmt.Println(i, "is an odd number")
		} else {
			fmt.Println(i, "is an even number")
		}

	}
}
