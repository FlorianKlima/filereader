package datafile

import (
	"fmt"
	"log"
)

func paintNeeded(width, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("width (%0.2f) must be >= 0", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("height (%0.2f) must be >= 0", height)
	}
	area := width * height
	return area / 10.0, nil
}

func main() {
	amount, err := paintNeeded(10.0, 15.0)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%.2f liters needed\n", amount)
	}
}
