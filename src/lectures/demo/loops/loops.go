package main

import "fmt"

func main() {
	sum := 0

	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println("sum is:", sum)

	for sum > 10 {
		sum -= 8
	}
	fmt.Println("sum is:", sum)
}
