package main

import "fmt"

type Counter struct {
	hits int
}

func increment(c *Counter) {
	c.hits++
	fmt.Println("incremented to", c.hits)
}

func decrement(c *Counter) {
	c.hits--
	fmt.Println("decremented to", c.hits)
}

func replace(oldVal *string, newVal string, c *Counter) {
	*oldVal = newVal
	increment(c)
}

func main() {
	c := Counter{}

	name := "Mustafa"
	fmt.Println(name)
	replace(&name, "MAXIMUS", &c)
	fmt.Println(name)

	increment(&c)
	increment(&c)
	decrement(&c)

	hello := "hello"
	world := "world"

	replace(&hello, "hi", &c)
	fmt.Println(hello, world)

	phrase := []string{hello, world} // NOTE: makes copies of hello and world
	replace(&phrase[1], "Go!", &c)
	fmt.Println(phrase)
	fmt.Println(hello, world) // NOTE: unaffected
}
