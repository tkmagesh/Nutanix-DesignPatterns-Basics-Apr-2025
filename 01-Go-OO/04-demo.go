package main

import (
	"fmt"
)

/* Circle */
type Circle struct {
	Radius float64
}

func (c Circle) Draw() {
	fmt.Printf("Circle (radius : %v) drawn\n", c.Radius)
}

/* Rectangle */
type Rectangle struct {
	Height float64
	Width  float64
}

func (r Rectangle) Draw() {
	fmt.Printf("Rectangle (height : %v, width : %v) drawn\n", r.Height, r.Width)
}

/* Utility function */
/* function to draw a shape */

type Drawable interface {
	Draw()
}

func drawShape(x Drawable) {
	x.Draw()
}

/* Exercise */
/*
Implement an Area() method in both Circle & Rectable
	Circle - Area => math.Pi * radius * radius
	Rectangle - Area => height * width

Implement PrintArea() function that prints the area of the given shape
	"Area : <area>"

Usage:
	main()
		PrintArea(c)
		PrintArea(r)
*/
func main() {
	c := Circle{Radius: 16}
	// c.Draw()
	drawShape(c)

	r := Rectangle{Height: 12, Width: 14}
	// r.Draw()
	drawShape(r)

}
