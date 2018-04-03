package main

import "fmt"

func main() {
	var favNums2 [5] float64
	favNums2[0] = 163
	favNums2[1] = 78557
	favNums2[2] = 691
	favNums2[3] = 3.141
	favNums2[4] = 1.618
	fmt.Println("Favourite Number: ", favNums2[3])

	// Another way of declaring an array
	favNums3 := [5]float64{1, 2, 3.141, 5, 0.99 }
	// Want to get the index too?
	for i, value := range favNums3 {
		fmt.Println("Index: ", i, " Value: ", value)
	}

	// Don't need the index value?
	for _, value := range favNums3 {
		fmt.Println("Value: ", value)
	}
}
