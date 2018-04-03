package main

import (
	"fmt"
	"math"
)

func main() {
	 rect := Rect{ 20, 50 }
	 circle := Circle{ 4 }

	 fmt.Println("Rectangle Area =", getArea(rect))
	 fmt.Println("Circle Area =", getArea(circle))
}

// If you use the shape interface, you must use and implement its methods
 type Shape interface {
 	area() float64
 }

 type Rect struct {
 	height float64
 	width float64
 }

 type Circle struct {
 	radius float64
 }

 func (r Rect) area() float64 {
 	return r.height * r.width
 }

 func (c Circle) area() float64 {
 	return math.Pi * math.Pow(c.radius, 2)
 }

 func getArea(shape Shape) float64 {
 	return shape.area()
 }
