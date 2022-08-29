package main
import "fmt"
 func main() {
	var userNum int
	fmt.Println("entry a number 1-20:")
	fmt.Scanf("%d", &userNum)

	for userNum <= 100 {

		switch {
		case userNum%100 == 0:
			fmt.Println("by one hundred")
		case userNum%5 == 0:
			fmt.Println(userNum, " by 5")


		}
		userNum++
	}
} 

