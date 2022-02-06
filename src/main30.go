package datafile

import "fmt"

func main() {
	numbers := [3]float64{111.5, 32.8, 300.1}
	var sum float64 = 0
	var number float64 = 0
	for _, number = range numbers {
		sum += number
	}
	fmt.Println(sum)
	sampleCount := float64(len(numbers))
	fmt.Println(sum / sampleCount)
}
