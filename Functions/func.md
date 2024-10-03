# Functions

the programs we have written in Go have used only one function:
`func main() {}`

Functions start with the keyword `func`, followed by the function’s name. The parameters (inputs) of the function are defined like this: 'name type', 'name type'. After the parameters, we put the return type. Collectively, the parameters and the return type are known as the function’s signature.
eg:

```go
func average(xs []float64) float64 {
    return
}
```

Return types can have names
We can name the return type like this:

```go
func f2() (r int) { r=1
    return
}
```

Multiple values can be returned
Go is also capable of returning multiple values from a function. Here is an example function that returns two integers:

```go
func f() (int, int) {
    return 5, 6
}
func main() {
    x, y := f()
}
```

## Variadic Functions

There is a special form available for the last parameter in a Go function:

```go
func add(args ...int) int {
    total := 0
    for _, v := range args {
        total += v
    }
    return total
}

func main() {
    fmt.Println(add(1,2,3))
}
```

In this example, add is allowed to be called with multiple integers. This is known as a variadic parameter.
By using an ellipsis (...) before the type name of the last parameter, you can indicate that it takes zero or more of those parameters. In this case, we take zero or more ints. We invoke the function like any other function except we can pass as many ints as we want.

We can also pass a slice of ints by following the slice with an ellipsis:

```go
func main() {
    xs := []int{1,2,3}
    fmt.Println(add(xs...))
}
```

## Closure

It is possible to create functions inside of functions. Let’s move the add function we saw before inside of main:

```go
func main() {
    add := func(x, y int) int {
        return x + y
    }
    fmt.Println(add(1,1))
}
```

One way to use closure is by writing a function that returns another function, which when called, can generate a sequence of numbers.

## Recursion

a function is able to call itself. Here is one way to compute the factorial of a number:

```go
func factorial(x uint) uint {
    if x == 0 {
        return 1
    }
    return x * factorial(x-1)
}
```

# defer

Go has a special statement called defer that schedules a function call to be run after the function completes. Consider the following example:

```go
package main
import "fmt"

func first() {
    fmt.Println("1st")
}
func second() {
    fmt.Println("2nd")
}
func main() {
    defer second()
    first()
}
```

This program prints first followed by second. Basically, defer moves the call to second to the end of the function:

```go
func main() {
    first()
    second()
 }
```

defer is often used when resources need to be freed in some way. For example, when we open a file, we need to make sure to close it later.

# panic, and recover

We can handle a runtime panic with the built-in recover function. recover stops the panic and returns the value that was passed to the call to panic. We might be tempted to recover from a panic, Instead, we have to pair it with defer:

```go
package main

import "fmt"

func main() {
    defer func() {
        str := recover()
        fmt.Println(str)
    }()
    panic("PANIC")
}
```
