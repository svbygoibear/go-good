package main

import "fmt"

func main() {
	yourAge := 18
	if yourAge >= 16 {
		fmt.Println("You can drive.")
	} else if yourAge >= 18 {
		// this will not print out because after the first condition has been met
		// it breaks out of the if statements
		fmt.Println("You can vote.")
	} else {
		fmt.Println("You can have fun.")
	}

	switch yourAge {
	case 16:
		fmt.Println("Go Drive")
	case 18:
		fmt.Println("Go Drink")
	default:
		fmt.Println("Do Nothing")
	}
}
