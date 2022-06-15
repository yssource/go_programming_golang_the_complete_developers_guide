package main

import "fmt"

func main() {
	var myName = "Mustafa"
	fmt.Println("My name is", myName)

	name := "Malena"
	fmt.Println("username=", name)

	var sum int
	fmt.Println("The sum is", sum)

	part1, other := 1, 5
	fmt.Println("part1 is", part1, "other is", other)

	part2, other := 1, 0
	fmt.Println("part1 is", part2, "other is", other)

	sum = part1 + part2
	fmt.Println("The sum is", sum)

	var (
		lessonName = "variables"
		lessonType = "Demo"
	)

	fmt.Println("lessonName=", lessonName)
	fmt.Println("lessonType=", lessonType)

	word1, word2, _ := "Hello", "world", "!"
	fmt.Println(word1, word2)
}
