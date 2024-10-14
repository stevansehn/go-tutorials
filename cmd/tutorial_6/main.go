package main

import "fmt"

type gasEngine struct {
	mpg uint8
	gallons uint8
	ownerInfo owner
}

type eletricEngine struct {
	mpkwh uint8
	kwh uint8
	ownerInfo owner
}

type owner struct {
	name string
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

func (e eletricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

type engine interface {
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("\nYou can make it there")
	} else {
		fmt.Println("\nNeed to fuel up first")
	}
}

func main() {
	// var myGasEngine gasEngine = gasEngine{mpg:25, gallons:15, ownerInfo:owner{"Alex"}}
	var myGasEngine gasEngine = gasEngine{25, 15, owner{"Alex"}}
	myGasEngine.mpg = 20
	fmt.Println(myGasEngine)
	fmt.Printf("Total miles left in tank: %v", myGasEngine.milesLeft())
	canMakeIt(myGasEngine, 50)
}