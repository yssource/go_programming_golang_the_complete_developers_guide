package main

import (
	"coursecontent/demo/pkg/display"
	"coursecontent/demo/pkg/msg"
	"fmt"
)

func main() {
	msg.Hi()
	display.Display("Hello from display")
	fmt.Println(msg.Exciting("an exciting message"))
}
