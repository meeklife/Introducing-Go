package main

import (
	"fmt"
)

func main() {
	s := "Elite"

	fmt.Printf("%8T %[1]v %d\n", s, len(s))
	fmt.Printf("%8T %[1]v %d\n", []rune(s), len([]rune(s)))
	fmt.Printf("%8T %[1]v %d\n", []byte(s), len([]byte(s)))
}
