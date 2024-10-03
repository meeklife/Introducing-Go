package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x float64
	y float64
	r float64
}

var c Circle

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func rectangleArea(x1, y1, x2, y2 float64) float64 {
	l := distance(x1, y1, x1, y2)
	w := distance(x1, y1, x2, y1)
	return l * w
}

func circleArea(c Circle) float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func main() {
	c := Circle{0, 0, 5}
	fmt.Println(circleArea(c))
	fmt.Println(c.area())
}
