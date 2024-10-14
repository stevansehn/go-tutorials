package main

import "fmt"
import "unicode/utf8"

func main() {
	var intNum int = 32767
	intNum += 1
	fmt.Println(intNum)

	var floatNum float64 = 12345678.9
	fmt.Println(floatNum)

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)

	var myString string = `Hello
World`
	fmt.Println(myString)

	fmt.Println(utf8.RuneCountInString("a"))

	var myRune rune = 'a'
	fmt.Println(myRune)

	var myBoolean bool = false
	fmt.Println(myBoolean)

	// var myVar string = "text"
	myVar := "text"
	fmt.Println(myVar)

	// var var1, var2 int = 1, 2
	var1, var2 := 1, 2
	fmt.Println(var1, var2)

	const myConst string = "const value"
	fmt.Println(myConst)
}