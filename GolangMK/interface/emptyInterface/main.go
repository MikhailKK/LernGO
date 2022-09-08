package main

import "fmt"

type Animal struct {
	species string
}

func (a Animal) String() string {
	return fmt.Sprintf("I am %s", a.species)
}

func main() {
	a := 5
	b := "some string"
	c := 12.5
	f := Animal{species: "Medved"}

	PrintEveryThing(a)
	PrintEveryThing(b)
	PrintEveryThing(c)
	PrintEveryThing(f)

	PrintCaseEvery(a)
	PrintCaseEvery(f)
}

func PrintEveryThing(everything interface{}) {
	fmt.Println(everything)
}

func PrintCaseEvery(everything interface{}) {
	switch everything.(type) {
	case int:
		fmt.Println(everything, "it is int")
	case float64:
		fmt.Println(everything, "it is float")
	case string:
		fmt.Println(everything, "it is string")
	case Animal:
		fmt.Println(fmt.Sprintf("animal: %s", everything.(Animal).String()))
	}
}
