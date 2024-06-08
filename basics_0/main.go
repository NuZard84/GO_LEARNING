package main

import (
	"fmt"
	"strings"
	"time"
	"unicode/utf8"
)

func main()  {
	fmt.Println("Hello, World!")
	// variables()
	// var returnFuncString1,returnFuncString2  string = functions("Paramameter passed !,")
	// fmt.Printf("first: %v second: %v ", returnFuncString1, returnFuncString2)
	// arrays()
	// Operation0()
	StringandRune()
}

func variables() {
	var intNum int16 = 32767
	fmt.Println(intNum)

	var floatNum64 float64 = 12345678.9
	var floatNum32 float32 = 12345678.9
	fmt.Println(floatNum32) // 12..8.0000
	fmt.Println(floatNum64) //12..8.9000

	var result float32 = floatNum32 + float32(intNum)
	fmt.Println(result) // 32767 + 12345678.9 = 12378445.9

	var myString string = `helllo ,
	world`
	fmt.Println(myString)
	fmt.Println(len("test")) //gives byte length
	fmt.Println(utf8.RuneCountInString("test")) //gives character length

	var myRune rune = 't' //char
	fmt.Println(myRune)

	var myBoolean bool = true 
	fmt.Println(myBoolean)

	// Fency technique to Declare Vars ...

	myVar := "text 123"
	fmt.Println(myVar)

	// var int1, int2 int  = 1 , 2
	int1, int2 := 1, 2
	fmt.Println(int1, int2)
}

func functions(paramValue string) (string, string){
	// var err error
	// err = errors.New("this is an error")
	// fmt.Println(err)
	fmt.Println(paramValue)
	return paramValue + " added !", paramValue + " extra !"
}

func arrays()  {

	var intArr [4]int32 = [4]int32{1,2,3,4}
	intArr = [...]int32{1,2,3,4}
	fmt.Println(intArr)

	// intSlice := []int32{4,5,6}
	// fmt.Printf("the Length is %v with capacity %v \n", len(intSlice), cap(intSlice))
	// intSlice = append(intSlice, 7) // [4,5,6,7,*,*]
	// fmt.Printf("the Length is %v with capacity %v \n", len(intSlice), cap(intSlice))

	// intSlice0 := []int32{8,9}
	// intSlice = append(intSlice, intSlice0...) // [4,5,6,7,8,9]
	// fmt.Printf("the Length is %v with capacity %v \n", len(intSlice), cap(intSlice))

	// var intSlice3 []int32 = make([]int32, 5, 8)
	// fmt.Printf("the Length is %v with capacity %v \n", len(intSlice3), cap(intSlice3))
	// intSlice3 = append(intSlice3, 1) // [0,0,0,0,0,1,*,*]
	// fmt.Println(intSlice3)

	var myMap map[string]uint8 = map[string]uint8{"a": 8}
	fmt.Println(myMap)
	myMap["b"] = 9
	fmt.Println(myMap["a"])
	delete(myMap, "b")
	var num, ok = myMap["b"]
	if ok{
		fmt.Printf("found %v \n", num)
	}else {
		fmt.Println("not found")
	}
	myMap["c"] = 10

	for name,num := range myMap {
		fmt.Println(name,num)
	}

	for i,v := range intArr{
		fmt.Println(i,v)
	}

	for i := 0; i < 4; i++{
		fmt.Println(i)
	}

	// while loop ...
	// i := 0
	// for (i < 4){
	// 	fmt.Println(i)
	// 	i++
	// }
	
}

func Operation0()  {

	var n int = 1000000
	var testSlice1 = []int{}
	var testSlice2 = make([]int, 0, n)

	fmt.Printf("Total time without prellocation: %v\n", timeLoop(testSlice1,n)) // 15msec
	fmt.Printf("Total time with prellocation: %v\n", timeLoop(testSlice2,n)) //3msec
	
}

func timeLoop(slice []int, n int) time.Duration{
	var t0 = time.Now()
	for len(slice)<n {
		slice = append(slice, 1)
	}
	return time.Since(t0)
}

func StringandRune() {

	var myString string = "rèsumè"
	var runeString = []rune(myString)
	var element = myString[0]
	fmt.Printf("%v, %T \n", element, element) //114, utf8
	for i, v := range myString{
		fmt.Println(i,v)
		// 0 114
		// 1 232 ,è is special char it contain combination of 2 byte codes : 195 & 168
		// 3 115
		// 4 117
		// 5 109
		// 6 232
	}
	for i,v := range runeString{
		fmt.Println(i,v)
		// 0 114
		// 1 232
		// 2 115
		// 3 117
		// 4 109
		// 5 232
	}
	var strSlice = []string{"h","e","t"}
	var catStr = ""

	for i := range strSlice{
		catStr += strSlice[i]
	}
	//it is not modified and every iteration will make new string it is not efficient..
	fmt.Println(catStr)

	var strBuilder strings.Builder
	for i := range strSlice{
		strBuilder.WriteString(strSlice[i])
	}
	var catStrBuild = strBuilder.String()
	fmt.Printf("string builder: %v",catStrBuild)

}

