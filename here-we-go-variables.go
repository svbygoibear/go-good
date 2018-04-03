package main

import "fmt"

func main() {
	// Defining constants
	const pi float64 = 3.14159256

	// Variable with inner values
	var (
		varA = 2
	 	varB = 3
	)

	// Declaring string variables
	var myName string = "Simone"
	var myNameBack string = `Simone`

	fmt.Println(len(myName))
	fmt.Println("I am \n")
	fmt.Println(myNameBack + " is a robot")

	// Bool declerations
	var isOver40 bool = true

	// Using Printf
	fmt.Printf("%f \n", pi) // prints out float
	fmt.Printf("%T \n", pi) // prints out var type
	fmt.Printf("%t \n", isOver40) // prints out bool
	fmt.Printf("%T \n", isOver40)
	fmt.Printf("%d \n", 100) // prints out integer
	fmt.Printf("%b \n", 100) // binary representation
	fmt.Printf("%c \n", 44) // prints out character code
	fmt.Printf("%x \n", 42) // prints out hex codes
	fmt.Printf("%e \n", pi) // scientific notation
}
