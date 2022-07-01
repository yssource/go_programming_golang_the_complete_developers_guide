package main

import "fmt"

type Preparer interface {
	PrepareDish()
}

type Chicken string

func (c Chicken) PrepareDish() {
	fmt.Println("cook chicken")
}

type Salad string

func (c Salad) PrepareDish() {
	fmt.Println("make salad")
}

func prepareDishes(dishes []Preparer) {
	fmt.Println("Prepare Dishes")

	for i := 0; i < len(dishes); i++ {
		dish := dishes[i]
		fmt.Printf("--Dish: %v--\n", dish)
		dish.PrepareDish()
	}
	fmt.Println()
}

func main() {
	dishes := []Preparer{
		Salad("ceasar salad"),
		Chicken("chicken wings"),
	}

	prepareDishes(dishes)
}
