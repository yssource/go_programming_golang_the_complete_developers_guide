//--Summary:
//  Implement receiver functions to create stat modifications for a video game
//  character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once
package main

import "fmt"

type Player struct {
	Health, MaxHealth int
	Energy, MaxEnergy int
	Name              string
}

func (p *Player) ModifyHealth(amount int) {
	if amount >= 0 {
		// add the positive number to health
		// increase Health
		if p.Health+amount <= p.MaxHealth {
			p.Health += amount
		} else {
			p.Health = p.MaxHealth
		}
	} else {
		// add the negative number to health
		// reduce Health
		if p.Health+amount >= 0 {
			p.Health += amount
		} else {
			p.Health = 0
		}
	}

	fmt.Printf("%s's health is %d\n", p.Name, p.Health)
}

func (p *Player) ModifyEnergy(amount int) {
	if amount >= 0 {
		// add the positive number to Energy
		// increase Energy
		if p.Energy+amount <= p.MaxEnergy {
			p.Energy += amount
		} else {
			p.Energy = p.MaxEnergy
		}
	} else {
		// add the negative number to Energy
		// reduce Energy
		if p.Energy+amount >= 0 {
			p.Energy += amount
		} else {
			p.Energy = 0
		}
	}

	fmt.Printf("%s's Energy is %d\n", p.Name, p.Energy)
}

func main() {
	player := Player{
		Name:      "Mustafa",
		MaxHealth: 100,
		Health:    71,
		MaxEnergy: 100,
		Energy:    92,
	}

	player.ModifyHealth(20)
	player.ModifyHealth(-39)
	player.ModifyHealth(50)
	player.ModifyHealth(5500)
	player.ModifyHealth(-5500)

	player.ModifyEnergy(20)
	player.ModifyEnergy(-40)
	player.ModifyEnergy(50)
	player.ModifyEnergy(-5050)
	player.ModifyEnergy(5050)

	fmt.Printf("%#v\n", player)
}
