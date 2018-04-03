package main

import "fmt"

func main() {
	// These are some of the int values in Go and their value breakdown
	// uint8 : 0 - 255
	// uint16 : 0 - 65535
	// unit32 : 0 - 4294967295
	// uint64 : 0 - 18446744073709
	// int8 : -128 - 127
	// int16 : -32768 - 32767
	// int32 : -2147483648 - 2147483647
	// int64 : -9223372036854775808 - 9223372036854775807

	 var age int = 40 //declares an int
	 var favNum float64 = 1.65844 //declares a float64
	 randomNum := 1 //assignment creates an int, once an int always an int

	 fmt.Println(age, favNum, randomNum)
}
