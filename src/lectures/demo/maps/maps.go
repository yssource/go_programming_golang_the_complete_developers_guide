package main

import "fmt"

func main() {
	shoppingList := make(map[string]int)
	shoppingList["egg"] = 11
	shoppingList["milk"] = 1
	shoppingList["bread"] += 1

	shoppingList["egg"] += 1
	fmt.Println(shoppingList)

	delete(shoppingList, "milk")
	fmt.Println("After deleting milk from shopping list:", shoppingList)

	for k, v := range shoppingList {
		fmt.Printf("I need %d %s(s)\n", v, k)
	}

	cereal, ok := shoppingList["cereal"]
	fmt.Println("need cereal?")
	if !ok {
		fmt.Println("Nope")
	} else {
		fmt.Printf("yes, %d boxes\n", cereal)
	}

	totalItems := 0
	for _, v := range shoppingList {
		totalItems += v
	}
	fmt.Printf("There are %d items in the shoppingList.\n", totalItems)
}
