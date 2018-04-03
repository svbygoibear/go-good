package main

import (
	"fmt"
	"os"
	"log"
	"io/ioutil"
)

func main() {
	// Creates a .txt called file, if there's an error then err will be created
	file, err := os.Create("sample.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Write
	// If there's no errors, write to the file and close the file
	file.WriteString("This is some random text")
	file.Close()

	// Read
	stream, err := ioutil.ReadFile("sample.txt")
	if err != nil {
		log.Fatal(err)
	}

	readString := string(stream)
	fmt.Println(readString)
}
