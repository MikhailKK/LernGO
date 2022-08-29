package main

import (
	"fmt"
)

func main() {
	var entrElem string
	// var chartElement map[string]map[string]string
	fmt.Print("Entry a element :")

	chartElement := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},

		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},

		"Li": map[string]string{
			"name":  "Litium",
			"state": "solid",
		},

		"Be": map[string]string{
			"name":  "Berylium",
			"state": "solid",
		},

		"C": map[string]string{
			"name":  "Corbon",
			"state": "solid",
		},

		"N": map[string]string{
			"name":  "Nitrogen",
			"state": "gas",
		},

		"O": map[string]string{
			"name":  "Oxygen",
			"state": "gas",
		},
	}

	if _, err := fmt.Scanln(&entrElem); err != nil {
		fmt.Println("Something wrong ", err)
		return
	} else if entrElem == "end" {
		return
	} else {

		fmt.Println(chartElement[entrElem])
		return
	}
}
