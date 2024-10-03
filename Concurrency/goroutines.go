package main

import (
	"fmt"
)

// goroutines
func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}
}

func main() {
	// goroutines
	go f(0)
	var input string
	fmt.Scanln(&input)
}
