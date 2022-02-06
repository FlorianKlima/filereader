package datafile

import (
	"fmt"
	"keyboard"
	"log"
)

func main() {
	fmt.Println("Enter a grade: ")
	grade, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	var status string
	if grade >= 60 {
		status = "Passing"
	} else {
		status = "Failing"
	}
	fmt.Printf("You have a %s grade.\n", status)
}
