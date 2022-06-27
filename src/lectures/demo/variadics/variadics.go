package main

import "fmt"

func sum(nums ...int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	return sum
}

func main() {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}

	all := append(a, b...)

	answer := sum(1, 2, 3, 4, 5, 6)
	fmt.Println(answer)

	answer = sum(all...)
	fmt.Println(answer)
}
