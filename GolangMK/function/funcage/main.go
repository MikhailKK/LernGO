package main

import (
	"fmt"
	"time"
)

func PrintPersonAge(id int64, birthday time.Time) {
	tNow := time.Now()
	hours := tNow.Sub(birthday).Hours()
	days := hours / 24
	age := int64(days / 365)
	fmt.Printf("id: %d age:%d\n", id, age)
}

func main() {
	t := time.Date(1990, 11, 17, 05, 34, 58, 651387237, time.UTC) //1909-11-17 20:34:58.651387237 +0000 UTC
	PrintPersonAge(23, t)

}

// calculates current age using given birthday
// func AgeFromBirthday(birthday time.Time) int64 {
// 	tNow := time.Now()
// 	hours := tNow.Sub(birthday).Hours()
// 	days := hours / 24
// 	return int64(days / 365)
// }
