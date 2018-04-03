package main

import "fmt"

func main() {
	// Used to define our own data types
	// rect := Rectangle{leftX: 0, topY: 50, height: 10, width: 10 } // Don't know the order
	rect1 := Rectangle{0, 50, 10, 10 } // Do this if you know the order
	fmt.Println("Rectangle is", rect1.width, "wide and the area is", rect1.area())
}

type Rectangle struct {
	leftX  float64
	topY   float64
	height float64
	width  float64
}

func (rect *Rectangle) area() float64 {
	return rect.width * rect.height
}
