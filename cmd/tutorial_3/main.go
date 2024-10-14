package main

import (
	"errors" 
	"fmt"
)

func main() {
	var printValue string = "Hello, world!"
	printMe(printValue)

	var numerator int = 11
	var denominator int = 0
	var result, remainder, err = intDivision(numerator, denominator)
	switch {
		case err != nil:
			fmt.Println(err.Error())
		case remainder == 0:
			fmt.Printf("The result of the integer division is %v", result)
		default:
			fmt.Printf("The result of the integer division is %v with remainder %v", result, remainder)
	}
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("Cannot divide by zero")
		return 0, 0, err
	}

	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}