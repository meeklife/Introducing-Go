# Structs

A struct is a type that contains named fields. For example, we could represent a circle struct like this:

```go
type Circle struct {
    x float64
    y float64
    r float64
}
```

The `type` keyword introduces a new type. It’s followed by the name of the type (Circle), the keyword `struct` to indicate that we are defining a struct type, and a list of fields inside of curly braces.

## initialization

We can create an instance of our new Circle type in a variety of ways:

```go
var c Circle
```

Like with other data types, this will create a local Circle variable that is by default set to zero. For a struct, zero means each of the fields is set to their corresponding zero value (0 for ints, 0.0 for floats, "" for strings, nil for pointers, etc.) We can also use the new function:

```go
c := new(Circle)
```

This allocates memory for all the fields, sets each of them to their zero value, and returns a pointer to the struct (\*Circle).

Pointers are often used with structs so that functions can modify their contents.

Using `new` in this way is somewhat uncommon. More typically, we want to give each of the fields an initial value. We can do this in two ways.

The first option looks like this:

```go
c := Circle{x: 0, y: 0, r: 5}
```

The second option is to leave off the field names if we know the order they were defined:

```go
c := Circle{0, 0, 5}
```

This creates the same Circle as the previous example. If you want a pointer to the
struct, use &:

```go
c := &Circle{0, 0, 5}
```

## Fields

We can access fields using the . operator:

```go
fmt.Println(c.x, c.y, c.r)
c.x = 10
c.y = 5
```

One thing to remember is that arguments are always copied in Go. If we attempted to modify one of the fields inside of the circleArea function, it would not modify the original variable. Because of this, we would typically write the function using a pointer to the Circle:

```go
func circleArea(c *Circle) float64 {
    return math.Pi * c.r*c.r
}
```

And change main to use & before c:

```go
c := Circle{0, 0, 5}
fmt.Println(circleArea(&c))
```

## Methods

we can improve it significantly by using a special type of function known as a method:

```go
func (c *Circle) circlArea() float64 {
    return math.Pi * c.r*c.r
}
```

In between the keyword func and the name of the function, we’ve added a receiver. The receiver is like a parameter, it has a name and a type but by creating the function in this way, it allows us to call the function using the . operator:

```go
fmt.Println(c.circlArea())
```

This is much easier to read. We no longer need the & operator (Go automatically knows to pass a pointer to the circle for this method), and because this function can only be used with Circles, we can rename the function to just area.

Let’s do the same thing for the rectangle:

```go
type Rectangle struct {
    x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64 {
    l := distance(r.x1, r.y1, r.x1, r.y2)
    w := distance(r.x1, r.y1, r.x2, r.y1)
    return l * w
}
```

main has:

```go
r := Rectangle{0, 0, 10, 10}
fmt.Println(r.area())
```

## Embedded Types

A struct’s fields usually represent the has-a relationship (e.g., a Circle has a radius). The following snippet shows an example of a person struct:

```go
type Person struct {
    Name string
}

func (p *Person) Talk() {
    fmt.Println("Hi, my name is", p.Name)
}
```

Now suppose we wanted to create a new Android struct. We could do this:

```go
type Android struct {
    Person Person
    Model string
}
```

This would work, but we would rather say an android is a person, rather than an android has a person. Go supports relationships like this by using embedded types (sometimes also referred to as anonymous fields)—they look like this:

```go
type Android struct {
    Person
    Model string
}
```

We use the type (Person) and don’t give it a name. When defined this way, the Person struct can be accessed using the type name:

```go
    a := new(Android)
    a.Person.Talk()
```

But we can also call any Person methods directly on the Android:

```go
a := new(Android)
a.Talk()
```

The is-a relationship works this way intuitively: people can talk, an android is a per‐
son, therefore an android can talk.
