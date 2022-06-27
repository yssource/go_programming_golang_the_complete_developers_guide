package main

import "fmt"

type Space struct {
	occupied bool
}

type ParkingLot struct {
	spaces []Space
}

func (pl *ParkingLot) occupySpace(index int) {
	pl.spaces[index-1].occupied = true
}

func (pl *ParkingLot) vacateSpace(index int) {
	pl.spaces[index-1].occupied = false
}

func occupySpace(index int, pl *ParkingLot) {
	pl.spaces[index-1].occupied = true
}

func main() {
	lot := ParkingLot{spaces: make([]Space, 4)}
	fmt.Println("initial parking lot:", lot)

	lot.occupySpace(1)
	occupySpace(2, &lot)
	fmt.Println("parking lot after occupation:", lot)

	lot.vacateSpace(2)
	fmt.Println("parking lot after vacation:", lot)

}
