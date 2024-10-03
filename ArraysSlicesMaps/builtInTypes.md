# Arrays

An array is a numbered sequence of elements of a single type with a fixed length. In Go, they look like this:
var x [5]int
x is an example of an array that is composed of five ints

```go
var x [5]int
x[4] = 100
fmt.Println(x)

[0 0 0 0 100]
```

Go also provides a shorter syntax for creating arrays:

```go
x := [5]float64{ 98, 93, 77, 82, 83 }
```

Go allows you to break it up like this:

```go
x := [5]float64{
98,
93,
77,
82,
83,
}
```

## Slices

Adding or removing elements as we did here requires also changing the length inside the brackets. Because of this and other limitations, you rarely see arrays used directly in Go code. Instead, you will usually use a slice, which is a type built on top of an array.

A slice is a segment of an array. Like arrays, slices are indexable and have a length. Unlike arrays, this length is allowed to change. Here’s an example of a slice:
`var x []float64`
The only difference between this and an array is the missing length between the brackets. In this case, x has been created with a length of zero.

If you want to create a slice, you should use the built-in make function:

```go
x := make([]float64, 5)
```

This creates a slice that is associated with an underlying float64 array of length 5. Slices are always associated with some array, and although they can never be longer than the array, they can be smaller.

The make function also allows a third parameter:

```go
x := make([]float64, 5, 10)
```

(A slice of length 5 with a capacity of 10)

Another way to create slices is to use the [low : high] expression:

```go
arr := [5]float64{1,2,3,4,5}
x := arr[0:5]
```

low is the index of where to start the slice and high is the index of where to end it (but not including the index itself). For example, while arr[0:5] returns [1,2,3,4,5], arr[1:4] returns [2,3,4].
For convenience, we are also allowed to omit low, high, or even both low and high. arr[0:] is the same as arr[0:len(arr)], arr[:5] is the same as arr[0:5], and arr[:] is the same as arr[0:len(arr)].

### append

append adds elements onto the end of a slice. If there’s sufficient capacity in the underlying array, the element is placed after the last element and the length is incre‐ mented. However, if there is not sufficient capacity, a new array is created, all of the existing elements are copied over, the new element is added onto the end, and the new slice is returned.

```go
func main() {
slice1 := []int{1,2,3}
slice2 := append(slice1, 4, 5)
fmt.Println(slice1, slice2)
}
>> [1,2,3] [1,2,3,4,5]
```

### copy

copy takes two arguments: dst and src. All of the entries in src are copied into dst overwriting whatever is there. If the lengths of the two slices are not the same, the smaller of the two will be used.
Here is an example of copy:

```go
func main() {
slice1 := []int{1,2,3}
slice2 := make([]int, 2)
copy(slice2, slice1)
fmt.Println(slice1, slice2)
}
>>[1,2]
```

# Maps

A map is an unordered collection of key-value pairs (maps are also sometimes called associative arrays, hash tables, or dictionaries). Maps are used to look up a value by its associated key. Here’s an example of a map in Go:
`var x map[string]int`

The map type is represented by the keyword map, followed by the key type in brackets and finally the value type. If you were to read this out loud, you would say “x is a map of strings to ints.”
Like arrays and slices, maps can be accessed using brackets. Try running the following program:

```go
var x map[string]int
x["key"] = 10
fmt.Println(x)
```

maps have to be initialized before they can be used. We should have written this:

```go
x := make(map[string]int)
x["key"] = 10
fmt.Println(x["key"])
```

We can also delete items from a map using the built-in delete function:
`delete(x, "key")`

Like we saw with arrays, there is also a shorter way to create maps:

```go
elements := map[string]string{
"H": "Hydrogen",
"He": "Helium",
"Li": "Lithium",
"Be": "Beryllium",
"B":  "Boron",
"C":  "Carbon",
"N":  "Nitrogen",
"O":  "Oxygen",
"F":  "Fluorine",
"Ne": "Neon",
}
```
