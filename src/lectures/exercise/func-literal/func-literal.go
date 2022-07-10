//--Summary:
//  Create a program that can create a report of rune information from
//  lines of text.
//
//--Requirements:
//* Create a single function to iterate over each line of text that is
//  provided in main().
//  - The function must return nothing and must execute a closure
//* Using closures, determine the following information about the text and
//  print a report to the terminal:
//  - Number of letters
//  - Number of digits
//  - Number of spaces
//  - Number of punctuation marks
//
//--Notes:
//* The `unicode` stdlib package provides functionality for rune classification

package main

import (
	"fmt"
	"unicode"
)

func report_text_details(lines []string) {
	letters, digits, spaces, punctuation := 0, 0, 0, 0

	extract_details := func(line string) {
		for _, r := range line {
			if unicode.IsDigit(r) {
				digits++
			} else if unicode.IsSpace(r) {
				spaces++
			} else if unicode.IsPunct(r) {
				punctuation++
			} else {
				letters++
			}
		}
	}

	for _, line := range lines {
		extract_details(line)
	}

	fmt.Printf("Number of letters: %v\n", letters)
	fmt.Printf("Number of digits: %v\n", digits)
	fmt.Printf("Number of spaces: %v\n", spaces)
	fmt.Printf("Number of punctuation marks: %v\n", punctuation)
}

func main() {
	lines := []string{
		"There are",
		"68 letters,",
		"five digits,",
		"12 spaces,",
		"and 4 punctuation marks in these lines of text!",
	}

	report_text_details(lines)
}
