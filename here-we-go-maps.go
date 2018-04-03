package main

import "fmt"

func main() {
	// creates a map with a key of string and value of int
	presidentAge := make(map[string]int)
	presidentAge["President1"] = 31
	presidentAge["President Two"] = 42

	// gives the number of items in the map
	fmt.Println("Length = ", len(presidentAge))
	fmt.Println("Age of PresidentTwo =", presidentAge["President Two"])

	// delete a value
	delete(presidentAge, "President1")
	fmt.Println("Length = ", len(presidentAge))
}
