package main

import (
	"fmt"
	"time"
	"strconv"
)

var pizzaNum = 0
var pizzaName = ""

func makeDough(stringChan chan string) {
	pizzaNum++
	pizzaName = "Pizza #" + strconv.Itoa(pizzaNum)
	fmt.Println("Make Dough and Send for Sauce")

	// Now pass the values on to another channel for a go routine to use
	stringChan <- pizzaName
	time.Sleep(time.Millisecond * 10)
}

func addSauce(stringChan chan string) {
	// Assign the value passed through the channel to pizza
	pizza := <- stringChan
	fmt.Println("Added Sauce and Send", pizza, "for toppings")

	// Now pass the values on to another channel for a go routine to use
	stringChan <- pizzaName
	time.Sleep(time.Millisecond * 10)
}

func addToppings(stringChan chan string) {
	// Assign the value passed through the channel to pizza
	pizza := <- stringChan
	fmt.Println("Addeded Toppings to", pizza, "and ship")
	time.Sleep(time.Millisecond * 10)
}

func main() {
	// Channels allows you to pass values back and forth between go routines
	stringChan := make(chan string)
	for i := 0; i < 3; i++ {
		go makeDough(stringChan)
		go addSauce(stringChan)
		go addToppings(stringChan)
		time.Sleep(time.Millisecond * 5000)
	}
}
