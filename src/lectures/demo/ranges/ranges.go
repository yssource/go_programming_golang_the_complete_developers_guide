package main

import "fmt"

func main() {
	s := []string{"Hello", "world", "!"}
	for i, v := range s {
		fmt.Println(i, v, ":")

		for _, ch := range v {
			fmt.Printf("\t%q\n", ch)
		}
	}
}
