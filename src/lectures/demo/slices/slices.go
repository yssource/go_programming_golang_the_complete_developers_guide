package main

import "fmt"

func printSlices(title string, sl []string) {
	fmt.Println("\n--- ", title, " ---")
	for i := 0; i < len(sl); i++ {
		item := sl[i]
		fmt.Println(item)
	}
}

func main() {
	routes := []string{"Gym", "Home", "Office"}
	printSlices("Route 1", routes)

	routes = append(routes, "Home")
	printSlices("Route 2", routes)

	fmt.Println()
	fmt.Println(routes[0], "visited")
	fmt.Println(routes[1], "visited")

	routes = routes[2:]
	printSlices("Route 3", routes)
}
