package datafile

import "fmt"

func main() {
	var quantity int
	var length, width float64
	var customerName string

	quantity = 4
	length, width = 1.2, 2.4
	customerName = "Beavis Butthead"

	fmt.Println(customerName)
	fmt.Println("Has ordered", quantity, "bottles of beer")
	fmt.Println("each with a area of")
	fmt.Println(length*width, "square inch")
}
