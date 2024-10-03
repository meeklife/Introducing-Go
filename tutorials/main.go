package main

import "fmt"

// Function that halves an integer and returns true if even value or false if odd value:
func OddOrEven(num int) (int, bool) {
	return num / 2, num%2 == 0
}

// Write a function with one variadic parameter that finds the greatest number in a list of numbers.
func greater(nums ...int) int {
	num := nums[0]
	for _, i := range nums {
		if i > num {
			num = i
		}
	}
	return num
}

// Write a program that can swap two integers (x := 1; y := 2; swap(&x, &y) should give you x=2 and y=1).
func swap(x, y *int) {
	*x, *y = *y, *x
}

func main() {
	numbers := []int{1, 2, 3, 4, 51, 2000}
	fmt.Println(OddOrEven(8))
	fmt.Println(greater(numbers...))

}
