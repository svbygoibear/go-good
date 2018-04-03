package main

import "fmt"

func main() {
	x := 0
	changeXVal(x) // This passes the value of x, not x itself - to change x use pointers
	fmt.Println("x =", x)

	changeXValNow(&x) // This passes the memory address in to the function
	fmt.Println("x =", x)

	yPtr := new(int)
	changeYValNow(yPtr)
	fmt.Println("y =", *yPtr)
}

func changeXVal(x int) {
	// Changing the value of a variable here will not change it in main
	x = 2
}

func changeXValNow(x *int) {
	// Now stores the value of 2 in the memory address that x refers to
	*x = 2
}

func changeYValNow(yPtr *int) {
	*yPtr = 100
}
