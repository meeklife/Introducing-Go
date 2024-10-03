package main

import (
	"fmt"
)

func main() {
	// x := make([]float64, 5)

	// arr := [5]float64{34, 67, 54, 23, 89}
	// y := arr[0:5]
	// z := append(y, 24, 8)
	// fmt.Println(x, y, z)

	// copy(z, x)
	// fmt.Println(z)
	x := []int{48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	smallest := x[0]

	for _, num := range x[1:] {
		if num < smallest {
			smallest = num
		}
	}
	fmt.Println(smallest)
}
