package main

import "fmt"

type Creature struct {
	Name     string
	Greeting string
}

func (c Creature) Greet() {
	fmt.Printf("%s says %s\n", c.Name, c.Greeting)
}

func main() {
	sammy := Creature{
		Name:     "Sammy",
		Greeting: "Hello!",
	}

	mikka := Creature{
		Name:     "Mikka",
		Greeting: "Привет",
	}
	Creature.Greet(sammy)
	Creature.Greet(mikka)
}
