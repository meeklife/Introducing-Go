package main

// package tcpServer

import (
	"container/list"
	"errors"
	"fmt"
	"hash/crc32"
	"os"
	"strings"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())

	// Contains(s, substr string) bool
	fmt.Println(strings.Contains("test", "es"))
	// => true

	// func Count(s, sep string) int
	fmt.Println(strings.Count("test", "t"))
	// => 2

	// func HasPrefix(s, prefix string) bool
	fmt.Println(strings.HasPrefix("test", "te"))
	// => true

	// func Index(s, sep string) int
	fmt.Println(strings.Index("test", "e"))
	// => 1

	// func Join(a []string, sep string) string
	fmt.Println(strings.Join([]string{"a", "b"}, "-"))
	// => "a-b"

	// using os package for files/folders
	file, err := os.Open("test.txt")
	if err != nil {
		// handle the error here
		return
	}
	defer file.Close()

	// creating our own error with errors packages
	errs := errors.New("error message")
	fmt.Println(errs)

	// linked list
	var x list.List
	x.PushBack(1)
	x.PushBack(2)
	x.PushBack(3)
	for e := x.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(int))
	}

	// create a hasher
	h := crc32.NewIEEE()
	// write our data to it
	h.Write([]byte("test"))
	// calculate the crc32 checksum
	v := h.Sum32()
	fmt.Println(v)

	// tcp

}
