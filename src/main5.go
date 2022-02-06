package datafile

import "fmt"

func main() {
	var quantity int = 4
	var length, width float64 = 1.2, 2.4
	var customerName string = "Beavis Butthead"

	fmt.Println(customerName)
	fmt.Println("Has ordered", quantity, "bottles of beer")
	fmt.Println("each with a area of")
	fmt.Println(length*width, "square inch")
}
