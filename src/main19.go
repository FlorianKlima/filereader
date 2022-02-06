package datafile

import (
	"fmt"
	"math"
)

func squareRoot(number float64) (float64, error) {
	if number < 0 {
		return 0, fmt.Errorf("Can't get a square root of negative number: %f", number)
	}
	return math.Sqrt(number), nil
}

func main() {
	root, err := squareRoot(-1.7)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(root)
	}
}
