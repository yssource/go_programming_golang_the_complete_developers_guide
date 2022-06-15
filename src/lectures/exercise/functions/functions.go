//--Summary:
//  Use functions to perform basic operations and print some information to the
//  terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

func greet(name string) {
	fmt.Println("Hello", name)
}

func message() string {
	return "Sup?"
}

func luckyNumber() int {
	return 13
}

func bestNums() (int, int) {
	return 4, 9
}

func addThreeNums(a, b, c int) int {
	return a + b + c
}

func main() {
	greet("Mustafa")

	fmt.Println(message())

	lucky := luckyNumber()
	firstBest, secondBest := bestNums()
	fmt.Println(addThreeNums(lucky, firstBest, secondBest))
}
