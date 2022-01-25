package main

import "fmt"

func main() {
	var price int = 100
	fmt.Println("Price is", price, "Dollars")
	var taxRate float64 = 0.08
	var tax float64 = float64(price) * taxRate
	fmt.Println("Tax is", tax, "Dollars")
	var total float64 = float64(price) + tax
	fmt.Println("Total is", total, "Dollars")
	var availlableFunds int = 120
	fmt.Println(availlableFunds, "Dollars available")
	fmt.Println("Within budget?", total <= float64(availlableFunds))
}
