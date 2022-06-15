package main

import "fmt"

func greet() {
	fmt.Println("Hello from greet function!")
}

func add(a, b int) int {
	return a + b
}

func double(x int) int {
	return x * 2
}

func main() {
	greet()

	dozen := double(6)
	fmt.Println("A dozen is", dozen)

	bakersDozen := add(dozen, 1)
	fmt.Println("A baker's dozen is", bakersDozen)

	anotherBakersDozen := add(double(6), 1)
	fmt.Println("Another baker's dozen is", anotherBakersDozen)
}
