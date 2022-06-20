//--Summary:
//  Create a program to calculate the area and perimeter of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

type coordinate struct {
	x, y int
}

type rectangle struct {
	topLeft     coordinate
	bottomRight coordinate
}

func (r rectangle) width() int {
	return r.bottomRight.x - r.topLeft.x
}

func (r rectangle) length() int {
	return r.topLeft.y - r.bottomRight.y
}

func (r rectangle) area() int {
	return r.width() * r.length()
}

func (r rectangle) perimeter() int {
	return 2*r.width() + 2*r.length()
}

func (r rectangle) info() {
	fmt.Println("\narea:", r.area())
	fmt.Println("perimeter:", r.perimeter())
}

func main() {
	rect := rectangle{
		topLeft:     coordinate{x: 0, y: 7},
		bottomRight: coordinate{x: 10, y: 0},
	}
	rect.info()

	rect.topLeft.y *= 2
	rect.bottomRight.x *= 2
	rect.info()
}
