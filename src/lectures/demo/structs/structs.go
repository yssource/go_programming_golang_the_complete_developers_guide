package main

import "fmt"

type Passenger struct {
	Name    string
	Ticket  int
	Boarded bool
}

type Bus struct {
	FrontSeat Passenger
}

func main() {
	mustafa := Passenger{"Mustafa", 1, true}
	taylor := Passenger{Name: "Taylor", Ticket: 2}
	fmt.Println(mustafa, taylor)

	var (
		malena = Passenger{"Malena", 3, true}
	)

	fmt.Println(malena)
	var maya Passenger
	maya.Name = "Maya"
	maya.Boarded = false
	maya.Ticket = 4

	if maya.Boarded {
		fmt.Println("Maya has boarded the bust")
	}

	taylor.Boarded = true
	if taylor.Boarded {
		fmt.Println("Taylor has boarded the bust")
	}

	bus := Bus{mustafa}
	fmt.Println(bus)
	fmt.Println(bus.FrontSeat.Name, "is in the front seat")
}
