package main

import (
	"fmt"
	"time"
)

func main() {
	// Go routines helps us run multiple different functions in parallel
	for i := 0; i < 10; i++ {
		go count(i)
	}

	time.Sleep(time.Millisecond * 11000)
}

func count(id int) {
	for i := 0; i < 10; i++ {
		fmt.Println(id, ":", i)
		time.Sleep(time.Millisecond * 1000)
	}
}
