// Echo 1 prints its command line arguments
package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args {
		fmt.Printf("%d\t%s\n", i, arg)
	}
}
