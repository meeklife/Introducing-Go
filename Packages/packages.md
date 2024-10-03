# Packages

functions are the first layer we utilize to allow code reuse. Go also provides another mechanism for code reuse: packages. Nearly every program we’ve seen so far included this line:

```go
import "fmt"
```

## Core Packages

- strings
  Go includes a large number of functions to work with strings in the strings package.
  `import "strings"`
  eg. of string functions `Contains`, `Count`, `HasPrefix`, `HasSuffix`, `Index`, `Join`, etc.

- Input/Output
  The io package consists of a few functions, but mostly interfaces used in other packages. The two main interfaces are Reader and Writer. Readers support reading via the Read method. Writers support writing via the Write method. Many functions in Go take Readers or Writers as arguments.

  For example, the io package has a Copy function that copies data from a Reader to a Writer:

  ```go
    func Copy(dst Writer, src Reader) (written int64, err error)
  ```

  To read or write to a []byte or a string, you can use the Buffer struct found in the bytes package:

  ```go
  var buf bytes.Buffer
  buf.Write([]byte("test"))
  ```

- Files/Folders
  To open a file in Go, use the Open function from the os package.
  `os.open("test.txt")`
  `os.Create()` to create a new file

- Errors
  Go has a built-in type for errors that we have already seen (the error type). We can create our own errors by using the New function in the errors package:

  ```go
  package main

  import "errors"

  func main() {
    err := errors.New("error message")
  }
  ```

- Containers & Sort
  In addition to lists and maps, Go has several more collections available underneath the container package.

  - list
    The container/list package implements a doubly linked list. A linked list is a linked data structure that consists of a set of sequentially linked records called nodes. Each node of the list contains a value (1, 2, or 3, in this case) and a pointer to the next node.
  - sort
    The sort package contains functions for sorting arbitrary data

- Hashes & Cryptography
  A hash function takes a set of data and reduces it to a smaller fixed size. Hash functions in Go are broken into two categories: cryptographic and non-cryptographic. eg: "hash/crc32"

- Servers
  Writing distributed, networked applications in Go is relatively straightforward. Three common approaches to communicating between multiple computers: TCP servers, HTTP servers, and RPC.

  - TCP
    TCP is the primary protocol used for communication over the Internet. Any time you interact with a web page, play a multiplayer computer game, stream a movie, or video chat, there’s a good chance your computer is communicating with a remote server using TCP. we can create a TCP server using the net package’s Listen function. Listen takes a network type (in our case, tcp) and an address and port to bind, and returns a net.Listener

    ```go
    type Listener interface {
    	// Accept waits for and returns the next connection to the listener.
    	Accept() (c Conn, err error)

    	// Close closes the listener.
    	// Any blocked Accept operations will be unblocked and return errors.
    	Close() error

    	// Addr returns the listener's network address.
    	Addr() Addr
    }
    ```

- HTTP
  HTTP servers are even easier to set up and use:

```go
package main

import ("net/http" ; "io")

func hello(res http.ResponseWriter, req *http.Request) {
    res.Header().Set(

            "Content-Type",

            "text/html",
      )

      io.WriteString(
            res,
            `<DOCTYPE html>
            <html>

            <head>
                <title>Hello, World</title>

            </head>
            <body>

                Hello, World!
            </body>
          </html>`,
      )
  }

func main() {

    http.HandleFunc("/hello", hello)

    http.ListenAndServe(":9000", nil)
}
```

- RPC
  The net/rpc (remote procedure call) and net/rpc/jsonrpc packages provide an easy way to expose methods so they can be invoked over a network (rather than just in the program running them).

## Parsing Command-line arguments

When we invoke a command on the terminal, it’s possible to pass that command arguments. We’ve seen this with the go command:
`go run myfile.go`
run and myfile.go are arguments. We can also pass flags to a command:
`go run -v myfile.go`

The flag package allows us to parse arguments and flags sent to our program. Here’s an example program that generates a number between 0 and 6. We can change the max value by sending a flag (-max=100) to the program:

```go
package main
import ("fmt";"flag";"math/rand")
func main() {
    // Define flags
    maxp := flag.Int("max", 6, "the max value")
    // Parse
    flag.Parse()
    // Generate a number between 0 and max
    fmt.Println(rand.Intn(*maxp))
}
```

## Creating Packages

Packages only really make sense in the context of a separate program that uses them. Without this separate program, we have no way of using the package we create.

- Let’s create an application that will use a package we will write. Create a folder in ~/src/ golang-book called chapter8. Inside that folder, create a file called main.go
- Now create another folder inside of the chapter8 folder called math. Inside of this folder, create a file called math.go
