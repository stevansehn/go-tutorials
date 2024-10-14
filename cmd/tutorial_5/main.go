package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString = "résumé"
	// var indexed = myString[0]
	fmt.Println(myString)

	fmt.Printf("\nThe length of 'myString' is %v", len(myString))

	var myRune = 'a'
	fmt.Printf("\nmyRune = %v", myRune)

	var strSlice = []string{"s", "u", "b", "s", "t", "a", "n", "c", "e"}
	fmt.Println(strSlice)
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	var catStr = strBuilder.String()
	fmt.Printf("\n%v", catStr)
}