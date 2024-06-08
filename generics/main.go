package main

import (
	"fmt"
)

//example ...

type gasEngine struct {
	gallons float32
	mpg     float32
}

type electricEngine struct {
	mpkwh float32
	kwh   float32
}

type car[T gasEngine | electricEngine] struct {
	carMAke  string
	carModel string
	engine   T
}

func main() {
	fmt.Println("Generics")

	var intSlice = []int{1, 2, 3}
	fmt.Println(sumIntSlice(intSlice))

	var floatSlice = []float32{1.1, 2.2, 3.3}
	fmt.Println(sumFloatSlice(floatSlice))

	var intSlice2 = []int{1, 2, 3}
	fmt.Println(sumSlices(intSlice2))

	var floatSlice2 = []float32{1.1, 2.2, 3.3}
	fmt.Println(sumSlices(floatSlice2))

	//example ...

	var gasCar = car[gasEngine]{
		carMAke:  "honda",
		carModel: "civic",
		engine: gasEngine{
			gallons: 10,
			mpg:     30,
		},
	}

	var electricCar = car[electricEngine]{
		carMAke:  "tesla",
		carModel: "model 3",
		engine: electricEngine{
			mpkwh: 4,
			kwh:   60,
		},
	}

	fmt.Println(gasCar, electricCar)
}

func sumIntSlice(slice []int) int {
	var sum int
	for _, v := range slice {
		sum += v
	}
	return sum
}

func sumFloatSlice(slice []float32) float32 {
	var sum float32
	for _, v := range slice {
		sum += v
	}
	return sum
}

// generic function
//can be write like this too :
//func sumSlices[T any](slice []T) T {

func sumSlices[T int | float32 | float64](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}
