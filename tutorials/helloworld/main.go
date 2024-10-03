package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Println("Hello", os.Args[1])
	} else {
		fmt.Printf("Hello World!")
	}
}
