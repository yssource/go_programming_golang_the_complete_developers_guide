//--Summary:
//  Create a program that can perform dice rolls. The program should report the
//  results as detailed in the requirements.
//
//--Requirements:
//* Print the sum of the dice roll
//* Print additional information in these cirsumstances:
//  - "Snake eyes": when the total roll is 2, and total dice is 2
//  - "Lucky 7": when the total roll is 7
//  - "Even": when the total roll is even
//  - "Odd": when the total roll is odd
//* The program must handle any number of dice, rolls, and sides
//
//--Notes:
//* Use packages from the standard library to complete the project

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	total := 0
	rand.Seed(time.Now().UnixNano())
	for total <= 50 {
		die1 := rand.Intn(5) + 1
		die2 := rand.Intn(5) + 1
		totalRoll := die1 + die2
		total += totalRoll

		fmt.Println("\nroll1:", die1)
		fmt.Println("roll2:", die2)
		if total == 2 && totalRoll == 2 {
			fmt.Println("Snake eyes")
		} else if totalRoll == 7 {
			fmt.Println("Lucky 7")
		} else if totalRoll%2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}
	}
}
