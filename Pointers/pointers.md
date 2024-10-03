# Pointers

Pointers, allowing you to pass references to values and records within your program.

## The \* and & operators

```go
func zero(xPtr *int) {
    *xPtr = 0
}
func main() {
    x := 5
    zero(&x)
    fmt.Println(x) // x is 0
}
```

In Go, a pointer is represented using an asterisk (\*) followed by the type of the stored value. In the zero function, xPtr is a pointer to an int.

An asterisk is also used to dereference pointer variables. Dereferencing a pointer gives us access to the value the pointer points to. When we write *xPtr = 0, we are saying “store the int 0 in the memory location xPtr refers to.” If we try xPtr = 0 instead, we will get a compile-time error because xPtr is not an int; it’s a *int, which can only be given another \*int.

Finally, we use the & operator to find the address of a variable. &x returns a \*int (pointer to an int) because x is an int. This is what allows us to modify the original variable. &x in main and xPtr in zero refer to the same memory location.

## new

Another way to get a pointer is to use the built-in `new` function:

```go
func one(xPtr *int) {
    *xPtr = 1
}

func main() {
    xPtr := new(int)
    one(xPtr)
    fmt.Println(*xPtr) // x is 1
}
```

new takes a type as an argument, allocates enough memory to fit a value of that type, and returns a pointer to it.
