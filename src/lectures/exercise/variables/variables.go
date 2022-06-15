//Summary:
//  Print basic information to the terminal using various variable creation
//  techniques. The information may be printed using any formatting you like.
//
//Requirements:
//* Store your favorite color in a variable using the `var` keyword
//* Store your birth year and age (in years) in two variables using
//  compound assignment
//* Store your first & last initials in two variables using block assignment
//* Declare (but don't assign!) a variable for your age (in days),
//  then assign it on the next line by multiplying 365 with the age
// 	variable created earlier
//
//Notes:
//* Use fmt.Println() to print out information
//* Basic math operations are:
//    Subtraction    -
// 	  Addition       +
// 	  Multiplication *
// 	  Division       /

package main

import "fmt"

func main() {
	var favoriteColor = "green"
	fmt.Println("My favorite color is", favoriteColor)

	birthYear, age := 1992, 29
	fmt.Println("My birth year is", birthYear)
	fmt.Println("My age is", age)

	var (
		firstInitial rune = 'M'
		lastInitial  rune = 'H'
	)
	fmt.Println("My initials are", string(firstInitial), "and", string(lastInitial))

	var myAgeInDays int
	myAgeInDays = age * 365
	fmt.Println("I've lived", myAgeInDays, "days on earth")
}
