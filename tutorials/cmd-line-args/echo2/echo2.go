// Echo 2 prints its command line arguments
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		// sep = " "
		fmt.Print(s, "\n")
	}
	// fmt.Println(s)
}
