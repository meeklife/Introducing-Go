package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// a type to store homepage sizes
type HomePageSize struct {
	URL  string
	Size int
}

func main() {
	// list of urls
	urls := []string{
		"http://www.apple.com",
		"http://www.amazon.com",
		"http://www.google.com",
		"http://www.microsoft.com",
	}

	// create a channel and start a goroutine for each URL
	results := make(chan HomePageSize)
	for _, url := range urls {
		go func(url string) {
			// make HTTP get request for each url
			res, err := http.Get(url)
			if err != nil {
				panic(err)
			}
			defer res.Body.Close()

			// store the size of response body
			bs, err := ioutil.ReadAll(res.Body)
			if err != nil {
				panic(err)
			}

			results <- HomePageSize{
				URL:  url,
				Size: len(bs),
			}
		}(url)
	}

	var biggest HomePageSize

	for range urls {
		result := <-results
		if result.Size > biggest.Size {
			biggest = result
		}
	}
	fmt.Println("The biggest home page:", biggest.URL)

}
