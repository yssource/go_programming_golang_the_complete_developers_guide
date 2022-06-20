package main

import "fmt"

type Room struct {
	name    string
	cleaned bool
}

func checkCleanliness(rooms [4]Room) {
	for i := 0; i < len(rooms); i++ {
		room := rooms[i]

		if room.cleaned {
			fmt.Println(room.name, "is clean")
		} else {
			fmt.Println(room.name, "is dirty")
		}
	}
}

func main() {
	rooms := [...]Room{
		{name: "Office"},
		{name: "Study room"},
		{name: "Living room"},
		{name: "Bed room"},
	}
	checkCleanliness(rooms)

	fmt.Println("=================")
	fmt.Println("Start cleaning...")
	fmt.Println("=================")

	rooms[1].cleaned = true
	rooms[3].cleaned = true
	checkCleanliness(rooms)
}
