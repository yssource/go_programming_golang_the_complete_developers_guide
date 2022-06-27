//--Summary:
//  Copy your rcv-func solution to this directory and write unit tests.
//
//--Requirements:
//* Write unit tests that ensure:
//  - Health & energy can not go above their maximums
//  - Health & energy can not go below 0
//* If any of your  tests fail, make the necessary corrections
//  in the copy of your rcv-func solution file.
//
//--Notes:
//* Use `go test -v ./exercise/testing` to run these specific tests

package main

import (
	"math"
	"testing"
)

func newPlayer() Player {
	return Player{
		Name:      "Mustafa",
		Health:    50,
		Energy:    50,
		MaxHealth: 100,
		MaxEnergy: 100,
	}
}

func TestModifyHealth(t *testing.T) {
	p := newPlayer()

	p.ModifyHealth(-120)
	if p.Health != 0 {
		t.Errorf(
			"should not reduce player's health to less than zero:\n\tExpected: %d\n\tGot: %d",
			0, p.Health,
		)
	}

	p.ModifyHealth(120)
	if p.Health != p.MaxHealth {
		t.Errorf(
			"should not increase player's health more than maximum value:\n\tExpected: %d\n\tGot: %d",
			p.MaxHealth, p.Health,
		)
	}

	p.ModifyHealth(-30)
	if p.Health != 70 {
		t.Errorf(
			"should reduce player's health correctly:\n\tExpected: %d\n\tGot: %d",
			70, p.Health,
		)
	}

	p.ModifyHealth(10)
	if p.Health != 80 {
		t.Errorf(
			"should increase player's health correctly:\n\tExpected: %d\n\tGot: %d",
			80, p.Health,
		)
	}

	p.ModifyHealth(math.MaxInt)
	if p.Health != p.MaxHealth {
		t.Errorf(
			"should not overflow the health value:\n\tExpected: %d\n\tGot: %d",
			p.MaxHealth, p.Health,
		)
	}

	p.ModifyHealth(-math.MaxInt)
	if p.Health != 0 {
		t.Errorf(
			"should not underflow the health value:\n\tExpected: %d\n\tGot: %d",
			0, p.Health,
		)
	}

	p.ModifyHealth(math.MinInt)
	if p.Health != 0 {
		t.Errorf(
			"should not underflow the health value:\n\tExpected: %d\n\tGot: %d",
			0, p.Health,
		)
	}
}

func TestModifyEnergy(t *testing.T) {
	p := newPlayer()

	p.ModifyEnergy(-120)
	if p.Energy != 0 {
		t.Errorf(
			"should not reduce player's energy to less than zero:\n\tExpected: %d\n\tGot: %d",
			0, p.Energy,
		)
	}

	p.ModifyEnergy(120)
	if p.Energy != p.MaxEnergy {
		t.Errorf(
			"should not increase player's energy more than maximum value:\n\tExpected: %d\n\tGot: %d",
			p.MaxEnergy, p.Energy,
		)
	}

	p.ModifyEnergy(-30)
	if p.Energy != 70 {
		t.Errorf(
			"should reduce player's energy correctly:\n\tExpected: %d\n\tGot: %d",
			70, p.Energy,
		)
	}

	p.ModifyEnergy(10)
	if p.Energy != 80 {
		t.Errorf(
			"should increase player's energy correctly:\n\tExpected: %d\n\tGot: %d",
			80, p.Energy,
		)
	}

	p.ModifyEnergy(math.MaxInt)
	if p.Energy != p.MaxEnergy {
		t.Errorf(
			"should not overflow the energy value:\n\tExpected: %d\n\tGot: %d",
			p.MaxEnergy, p.Energy,
		)
	}

	p.ModifyEnergy(-math.MaxInt)
	if p.Energy != 0 {
		t.Errorf(
			"should not underflow the energy value:\n\tExpected: %d\n\tGot: %d",
			0, p.Energy,
		)
	}

	p.ModifyEnergy(math.MinInt)
	if p.Energy != 0 {
		t.Errorf(
			"should not underflow the energy value:\n\tExpected: %d\n\tGot: %d",
			0, p.Energy,
		)
	}
}
