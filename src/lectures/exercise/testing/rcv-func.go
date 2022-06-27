package main

import "fmt"

type Player struct {
	Health, MaxHealth int
	Energy, MaxEnergy int
	Name              string
}

func (p *Player) ModifyHealth(amount int) {
	if amount >= p.MaxHealth || p.Health+amount >= p.MaxHealth {
		p.Health = p.MaxHealth
		return
	}

	if amount <= -p.MaxHealth || p.Health+amount <= 0 {
		p.Health = 0
		return
	}

	p.Health += amount
}

func (p *Player) ModifyEnergy(amount int) {
	if amount >= p.MaxEnergy || p.Energy+amount >= p.MaxEnergy {
		p.Energy = p.MaxEnergy
		return
	}

	if amount <= -p.MaxEnergy || p.Energy+amount <= 0 {
		p.Energy = 0
		return
	}

	p.Energy += amount
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
	fmt.Println(player)

	player.ModifyEnergy(20)
	player.ModifyEnergy(-40)
	player.ModifyEnergy(50)
	player.ModifyEnergy(-5050)
	player.ModifyEnergy(5050)
	fmt.Println(player)

	fmt.Printf("%#v\n", player)
}
