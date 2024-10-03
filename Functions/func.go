package main

import "fmt"

func average(num []float64) float64 {
	total := 0.0
	for _, v := range num {
		total += v
	}
	return total / float64(len(num))
}

// closure function
func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

// Recursion function
func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

// variadic function
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	// xs := []float64{98, 93, 77, 82, 83}
	// fmt.Println(average(xs))
	// nextEven := makeEvenGenerator()
	// fmt.Println(nextEven())
	// fmt.Println(nextEven())
	// fmt.Println(nextEven())
	// fmt.Println(factorial(9))

	numbers := []int{23, 56, 77, 22, 11, 33}
	fmt.Println(sum(numbers...))

}

// makeEvenGenerator returns a function that generates even numbers.
//Each time it’s called, it adds 2 to the local i variable, which—unlike normal local variables—persists between calls.
