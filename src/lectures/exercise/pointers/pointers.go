//--Summary:
//  Create a program that can activate and deactivate security tags on
//  products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change
package main

import "fmt"

type SecurityTag bool

const (
	Active   SecurityTag = true
	Inactive SecurityTag = false
)

type StoreItem struct {
	tag  SecurityTag
	name string
}

func (i *StoreItem) toggleTag() {
	switch i.tag {
	case Active:
		i.Deactivate()
	default:
		i.Activate()
	}
}

func (i *StoreItem) Deactivate() {
	i.tag = Inactive
}

func (i *StoreItem) Activate() {
	i.tag = Active
}

func (i *StoreItem) Status() string {
	switch i.tag {
	case Active:
		return "Activated"
	default:
		return "Deactivated"
	}
}

func checkout(items []StoreItem) {
	for i := 0; i < len(items); i++ {
		items[i].Deactivate()
	}
}

func main() {
	items := []StoreItem{
		{name: "t-shirt", tag: Active},
		{name: "socks", tag: Active},
		{name: "trousers", tag: Active},
		{name: "tie", tag: Active},
		{name: "hat", tag: Inactive},
	}

	items[0].toggleTag()
	fmt.Printf("%s's tag is %s\n", items[0].name, items[0].Status())

	fmt.Println(items)
	checkout(items)
	fmt.Println(items)

	for _, v := range items {
		fmt.Println(v.name, "is tagged:", v.tag)
	}
}
