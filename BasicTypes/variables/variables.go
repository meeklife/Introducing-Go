package main

import (
	"fmt"
)

// var x string = "Hello World"

// func main() {
// 	fmt.Println(x)
// }

// func f() {
// 	fmt.Println(x)
// }

func main() {
	fmt.Print("Enter a number: ")
	var num float64
	fmt.Scanf("%f", &num)

	newNum := num * 2
	fmt.Println(newNum)
}
