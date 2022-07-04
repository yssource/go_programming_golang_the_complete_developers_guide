//--Summary:
//  Create an interactive command line application that supports arbitrary
//  commands. When the user enters a command, the program will respond
//  with a message. The program should keep track of how many commands
//  have been ran, and how many lines of text was entered by the user.
//
//--Requirements:
//* When the user enters either "hello" or "bye", the program
//  should respond with a message after pressing the enter key.
//* Whenever the user types a "Q" or "q", the program should exit.
//* Upon program exit, some usage statistics should be printed
//  ('Q' and 'q' do not count towards these statistics):
//  - The number of non-blank lines entered
//  - The number of commands entered
//
//--Notes
//* Use any Reader implementation from the stdlib to implement the program

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	var (
		lines = 0
		cmds  = 0
	)

	for s.Scan() {
		v := strings.ToLower(strings.TrimSpace(s.Text()))
		if v == "q" {
			break
		}

		if v != "" {
			lines++
		} else {
			continue
		}

		switch v {
		case "hello":
			cmds++
			fmt.Println("Hiya")
		case "bye":
			cmds++
			fmt.Println("Goodbye to you too")
			// default:
			// 	fmt.Println("not a command")
		}
	}

	fmt.Println("You entered", lines, "lines and", cmds, "commands")
}
