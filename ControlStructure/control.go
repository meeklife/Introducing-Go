package main

import "fmt"

func main() {
	i := 1
	for i <= 100 {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i, "FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println(i, "Fizz")
		} else if i%5 == 0 {
			fmt.Println(i, "Buzz")
		}

		i = i + 1
	}
}
