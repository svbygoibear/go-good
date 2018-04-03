package main

import "fmt"

func main() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	h := 10
	for h >= 1 {
		fmt.Println(h)
		h--
	}


	for j := 0; j < 5; j++ {
		fmt.Println(j + 1)
	}
}
