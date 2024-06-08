package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_FRUIT_PRICE float32 = 5
var MAX_VEG_PRICE float32 = 4

func main() {
	fmt.Println("Channels")
	var c = make(chan int)
	var cbf = make(chan int, 5)
	go process(c)
	go processBuffered(cbf)

	fmt.Println("unBuffered channel ->")
	for i := range c {
		fmt.Println("channle value:", i)
	}
	fmt.Println("buffered channel ->")
	for i := range cbf {
		fmt.Println("channle value:", i)
		time.Sleep(time.Second * 1)
	}

	var fruitChannel = make(chan string)
	var vegChannel = make(chan string)
	var webs = []string{"web1", "web2", "web3", "web4", "web5"}
	for i := range webs {
		go checkFruitPrices(webs[i], fruitChannel)
		go checkVEgPrices(webs[i], vegChannel)
	}
	sendMEssage(fruitChannel, vegChannel)

}

// -> This function runs in an infinite loop, simulating the process of checking fruit prices.
// -> Once a suitable price is found (fruitPrice <= MAX_FRUIT_PRICE), the website name web is sent to fruitChannel using fruitChannel <- web.

func checkFruitPrices(web string, fruitChannel chan string) {

	for {
		time.Sleep(time.Second * 1)
		var fruitPrice = rand.Float32() * 28
		if fruitPrice <= MAX_FRUIT_PRICE {
			fruitChannel <- web
			break
		}
	}
}

func checkVEgPrices(web string, vegChannel chan string) {

	for {
		time.Sleep(time.Second * 1)
		var vegPrice = rand.Float32() * 28
		if vegPrice <= MAX_VEG_PRICE {
			vegChannel <- web
			break
		}
	}
}

// -> The select statement is used to wait on multiple communication operations. It will block until one of its cases can proceed, then it will execute that case.
// -> case web := <-fruitChannel: tries to receive a value from fruitChannel. If a value is available, it is assigned to web and the corresponding block of code is executed.
// -> Similarly, case web := <-vegChannel: tries to receive a value from vegChannel.

func sendMEssage(fruitChannel chan string, vegChannel chan string) {
	select {
	case web := <-fruitChannel:
		fmt.Printf("\n Found a deal on fruit at %s", web)

	case web := <-vegChannel:
		fmt.Printf("\n Found a deal on veg at %s", web)
	}
}

func process(c chan int) {
	defer close(c) //close last before the function executed
	for i := 0; i < 6; i++ {
		c <- i
	}
}

func processBuffered(c chan int) {
	defer close(c) //close last before the function executed
	for i := 0; i < 6; i++ {
		c <- i
	}
	//exist imediately after pass the 5 values and main function do its work
	fmt.Println("exiting processBuffered")
}
