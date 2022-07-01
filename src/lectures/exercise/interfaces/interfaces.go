//--Summary:
//  Create a program that directs vehicles at a mechanic shop to the correct
//  vehicle lift, based on vehicle size.
//
//--Requirements:
//* The shop has lifts for multiple vehicle sizes/types:
//  - Motorcycles: small lifts
//  - Cars: standard lifts
//  - Trucks: large lifts
//* Write a single function to handle all of the vehicles that the shop works
// on.
//* Vehicles have a model name in addition to the vehicle type:
//  - Example: "Truck" is the vehicle type, "Road Devourer" is a model name
//* Direct at least 1 of each vehicle type to the correct
//  lift, and print out the vehicle information.
//
//--Notes:
//* Use any names for vehicle models

package main

import "fmt"

type Lift byte

// enums
const (
	SmallLift Lift = iota
	StandardLift
	LargeLift
)

func (l Lift) String() string {
	switch l {
	case SmallLift:
		return fmt.Sprintf("small lift")
	case StandardLift:
		return fmt.Sprintf("standard lift")
	case LargeLift:
		return fmt.Sprintf("large lift")
	default:
		panic("unknown lift")
	}
}

type LiftPicker interface {
	PickLift() Lift
}

type VehicleType byte

const (
	Motorcycle VehicleType = iota
	Car
	Truck
)

type Vehicle struct {
	Type  VehicleType
	Model string
}

func (v *Vehicle) PickLift() Lift {
	switch v.Type {
	case Motorcycle:
		return SmallLift
	case Car:
		return StandardLift
	case Truck:
		return LargeLift
	default:
		panic("unknown vehicle")
	}
}

func (v Vehicle) String() string {
	switch v.Type {
	case Motorcycle:
		return fmt.Sprintf("Motorcycle: %s", v.Model)
	case Car:
		return fmt.Sprintf("Car: %s", v.Model)
	case Truck:
		return fmt.Sprintf("Truck: %s", v.Model)
	default:
		panic("unknown vehicle")
	}
}

func main() {
	vehicles := []LiftPicker{
		&Vehicle{Type: Truck, Model: "Road Devourer"},
		&Vehicle{Type: Car, Model: "Ferrari"},
		&Vehicle{Type: Motorcycle, Model: "Ducati"},
	}

	for _, v := range vehicles {
		fmt.Println(v)
		fmt.Printf("send to %s\n", v.PickLift())
		fmt.Println()

	}
}
