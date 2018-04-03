package main

import (
	"fmt"
	"strings"
	"sort"
)

func main() {
	sampleString := "Hello World"

	fmt.Println(strings.Contains(sampleString, "lo")) // Checks if the string contains a substring
	fmt.Println(strings.Index(sampleString, "lo")) // Checks what the index of a string is in the string
	fmt.Println(strings.Count(sampleString, "l")) // Counts the occurrence of a substring in the string
	fmt.Println(strings.Replace(sampleString, "l", "x", 2)) // Replaces an old string with a new number n-times

	csvString := "1,2,3,4,5"
	fmt.Println(strings.Split(csvString, ","))

	listOfLetters := []string { "c", "f", "b", "a" }
	sort.Strings(listOfLetters)
	fmt.Println(listOfLetters)

	listOfNums := strings.Join([]string { "3", "2", "1" }, ", ")
	fmt.Println(listOfNums)
}
