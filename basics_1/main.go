package main

import (
	"fmt"
)

type gasEngine struct{
	// ownerInfo owner
	owner
	mpg uint8
	gallons uint8
}

type electricEngine struct{
	mpkwh uint8
	kwh uint8
}

type owner struct {
	name string
}

func main()  {
	fmt.Println("Hello, World!")
	StructandInterface()
}

func (e gasEngine) milesLEft() uint8{
	return e.gallons * e.mpg
}

func (e electricEngine) milesLEft() uint8{
	return e.kwh * e.mpkwh
}

type engine interface {
	milesLEft() uint8
}

func canMakeIt(e engine, miles uint8){
	if miles<=e.milesLEft(){
		fmt.Println("Yes")
	}else {
		fmt.Println("No")
	}
}

func StructandInterface()  {
	// var myEngine gasEngine = gasEngine{mpg: 25, gallons: 15, ownerInfo: owner{name: "John"}}

	var myEngine gasEngine = gasEngine{mpg: 25, gallons: 15, owner: owner{name: "John"}}
	myEngine.mpg = 20 //mutable

	var myEngine2 = struct{
		mpg uint8
		gallons uint8
	}{25,15}

	var electricEng electricEngine = electricEngine{25,15} 
	fmt.Println(myEngine.mpg,myEngine.gallons, myEngine.name)
	fmt.Println(myEngine2.mpg,myEngine2.gallons)

	fmt.Println(myEngine.milesLEft())

	canMakeIt(electricEng, 50)

}


 