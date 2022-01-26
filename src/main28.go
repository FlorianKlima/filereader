package main

import "fmt"

func main() {
	letters := [7]string{"a", "b", "c", "d", "e", "f", "g"}
	for i := 0; i <= len(letters); i++ {
		fmt.Println(i, letters[i])
	}
}
