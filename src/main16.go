package main

import "fmt"

var metersPerLitre float64

func paintNeeded(width, height float64) float64 {
	area := width * height
	return area / metersPerLitre
}

func main() {
	metersPerLitre = 10.0
	fmt.Printf("%.2f liters\n", paintNeeded(3.0, 2.0))
}
