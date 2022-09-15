package main

import (
	"errors"
	"fmt"
	"log"
)

type Man struct {
	Name     string
	LastName string
	Age      int
	Gender   string
	Crimes   int
}

func FCrime(m Man) int {
	return m.Crimes
}

func FName(m Man) string {
	return m.Name
}

func main() {

	people := make(map[string]Man)
	var suspections []string

	Kiril := Man{"Kiril", "Bagrov", 28, "male", 0}
	Malvina := Man{"Malvina", "Buratinovna", 15, "female", 0}
	Andrei := Man{"Andrei", "Spasitel", 33, "male", 0}
	Max := Man{"Max", "Wild", 37, "male", 0}

	people["Kiril"] = Kiril
	people["Malvina"] = Malvina
	people["Andrei"] = Andrei
	people["Max"] = Max

	suspections = append(suspections, "Kiril", "Malvina", "Andrei", "Max")

	f := func(s []string) (Man, error) {
		var n string
		var mc int

		for _, m := range s {
			n2 := people[m]
			c := FCrime(n2)
			if mc < c {
				mc = c
				n = m
			}
			log.Print(m)
			if mc == 0 {
				return people[n], errors.New("got zero crimes")

			}
		}

		return people[n], nil
	}

	fmt.Println(f(suspections))

}
