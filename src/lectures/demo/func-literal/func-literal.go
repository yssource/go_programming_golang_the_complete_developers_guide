package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func compute(lhs, rhs int, op func(a, b int) int) int {
	fmt.Printf("Running a computation with %v & %v\n", lhs, rhs)
	return op(lhs, rhs)
}

func main() {
	fmt.Println("2 + 2 = ", compute(2, 2, add))

	fmt.Println("13 - 4 = ", compute(13, 4, func(a, b int) int { return a - b }))

	mul := func(a, b int) int {
		fmt.Printf("%v * %v = ", a, b)
		return a * b
	}

	fmt.Println(compute(0, 9, mul))
}
