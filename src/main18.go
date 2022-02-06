package datafile

import "fmt"

func status(grade float64) string {
	if grade < 60 {
		return "failing"
	}
	return "passing"
}

func main() {
	fmt.Println(status(60.1))
	fmt.Println(status(59.9))
}
