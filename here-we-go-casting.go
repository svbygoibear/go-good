package main

import (
	"fmt"
	"strconv"
)

func main() {
	randInt := 5
	randFloat := 10.5
	randString := "100"
	randString2 := "200.5"

	// Casting a int to a float
	fmt.Println(float64(randInt))

	// Casting a float into an int
	fmt.Println(int(randFloat))

	// Casting a string to int
	// ParseInt Parameters: string, base int, bit size
	newInt, _ := strconv.ParseInt(randString, 0, 64)
	fmt.Println(newInt)

	newFloat, _ := strconv.ParseFloat(randString2, 64)
	fmt.Println(newFloat)
}
