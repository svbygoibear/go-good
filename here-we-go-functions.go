package main

import "fmt"

func main() {
	// Returns one value
	listNums := []float64{1, 2, 3, 4, 5 }
	fmt.Println("The sum is = ", addThemUp(listNums))

	// Returns two values
	num1, num2 := next2Values(5)
	fmt.Println("Next two values ", num1, num2)

	// Return undefined amount of numbers
	fmt.Println(subtractThem(1, 2, 3, 4, 5))

	// Creating a function inside of a function -> closure
	num3 := 3
	doubleNum := func() int {
		num3 *= 2
		return num3
	}
	fmt.Println(doubleNum())
	fmt.Println(doubleNum())

	// Recursive Functions
	fmt.Println(factorial(8))

	// Defer Functions
	// Allows you to execute one last function as a clean-up operation when needed
	defer printTwo() // After the main finishes, it will execute this function
	printOne() // This means that printOne will be executed first

	// Defer to a recover
	fmt.Println(safeDivision(3, 0))
	fmt.Println(safeDivision(3, 2))

	// Call to Panic
	demPanic()
}

// func:		Declares that it is a function
// AddThemUp:	Name of the function/method
// ():			List of parameters passed through
// float64:	The return type
func addThemUp(numbers []float64) float64 {
	// this variable is not accessible outside of this function
  	sum := 0.0
 	for _, value := range numbers {
 		sum += value
 	}
 	return sum
}

func next2Values(number int) (int, int) {
 	return number + 1, number + 2
}

func subtractThem(args ...int) int {
	finalValue := 0
	for _, value := range args {
		finalValue -= value
	}
	return finalValue
}

// When using recursion, you would need to define an OUT (what is expected out)
func factorial (num int) int {
 	if num == 0 {
 		return 1
 	}
 	return num * factorial(num - 1)
}

func printOne() { fmt.Println(1) }
func printTwo() { fmt.Println(2) }

func safeDivision(number1, number2 int) int {
	defer func() {
		fmt.Println("Error has occured for ", number1, "/", number2, ":", recover())
	}() // recover allows you to catch an error

	solution := number1 / number2 // If a 0 is passed through, it will execute the defer recover
	return solution
}

func demPanic() {
	defer func() {
		fmt.Println("Error occured! ", recover())
	}()
	panic("PANIC") //sending a specific error to the recover function
}