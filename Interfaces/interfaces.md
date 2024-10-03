# Interfaces

Interfaces are named collections of method signatures. An interface is two things: It's a set of methods and also a type

```go
type Shape interface {
    area() float64
}
```

An interface is created using the `type` keyword, followed by a name and the keyword `interface`. But instead of defining fields, we define a method set. A method set is a list of methods that a type must have in order to implement the interface.

example of writing some application where you’re defining Animal datatypes, because that’s a totally realistic situation that happens all the time. The Animal type will be an interface, and we’ll define an Animal as being anything that can speak. This is a core concept in Go’s type system; instead of designing our abstractions in terms of what kind of data our types can hold, we design our abstractions in terms of what actions our types can execute.

We start by defining our Animal interface:

```go
type Animal interface {
    Speak() string
}
```

pretty simple: we define an Animal as being any type that has a method named Speak. The Speak method takes no arguments and returns a string. Any type that defines this method is said to satisfy the Animal interface.

There is no implements keyword in Go; whether or not a type satisfies an interface is determined automatically. Let’s create a couple of types that satisfy this interface:

```go
type Animal interface {
    Speak() string
}

type Dog struct {
}

func (d Dog) Speak() string {
    return "Woof!"
}

type Cat struct {
}

func (c Cat) Speak() string {
    return "Meow!"
}
```

we now have two different kinds of animals: A dog, and a cat
In our main() function, we can create a slice of Animals, and put one of each type into that slice, and see what each animal says.

```go
func main() {
    animals := []Animal{Dog{}, Cat{}, Llama{}, JavaProgrammer{}}
    for \_, animal := range animals {
        fmt.Println(animal.Speak())
    }
}
```

## NB: Another explanation

---

Suppose we want to write a function that calculates the area of several shapes. Using the techniques we’ve discussed so far, we might start to write the function like this:

```go
func totalArea(circles ...Circle) float64 {
    var total float64
    for _, c := range circles {
        total += c.area()
    }
    return total
}
```

And then we’d try to add in Rectangles:
// THIS IS INVALID

```go
func totalArea(circles ...Circle, rectangles ...Rectangle) float64 {
    var total float64
    for _, c := range circles {
        total += c.area()
    }
    for _, r := range rectangles {
        total += r.area()
    }
    return total
}
```

But we can’t write a function that contains two variadic parameters, so we would have to modify the program:

```go
func totalArea(circles []Circle, rectangles []Rectangle) float64 {
    var total float64
    for _, c := range circles {
        total += c.area()
    }
    for _, r := range rectangles {
        total += r.area()
    }
    return total
}
```

This works, but it has a major issue—whenever we define a new shape, we have to change our function to handle it (a third parameter for Polygons, a fourth for Squares, etc.).

This is the problem interfaces are designed to solve. Because both of our shapes have an area method, they both implement the Shape interface and we can change our function to this:

```go
type Shape interface {
    area() float64
}

func totalArea(shapes ...Shape) float64 {
    var area float64
    for _, s := range shapes {
        area += s.area()
    }
    return area
}
```

We would call this function like this:

```go
fmt.Println(totalArea(&c, &r))
```

## The interface{} type

The interface{} type is the interface that has no methods.
That means that if you write a function that takes an interface{} value as a parameter, you can supply that function with any value. So, this function:

```go
func DoSomething(v interface{}) {
   // ...
}
```

will accept any parameter whatsoever.
