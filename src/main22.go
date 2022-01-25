package main

import "fmt"

func main() {
	var myInt int
	var myIntPointer *int
	myIntPointer = &myInt
	fmt.Println(myIntPointer)

	var myFloat float64
	var myFloatPointer *float64
	myFloatPointer = &myFloat
	fmt.Println(myFloatPointer)

	var myString string
	var myStringPointer *string
	myStringPointer = &myString
	fmt.Println(myStringPointer)

	var myBool bool
	var myBoolPointer *bool
	myBoolPointer = &myBool
	fmt.Println(myBoolPointer)

	var myArray [5]int
	var myArrayPointer *[5]int
	myArrayPointer = &myArray
	fmt.Println(myArrayPointer)

	myIntToo := 6
	fmt.Println(myIntToo)
	myIntPointerToo := &myIntToo
	*myIntPointerToo = 7
	fmt.Println(*myIntPointerToo)
	fmt.Println(myIntToo)
}
