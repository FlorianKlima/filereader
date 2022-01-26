package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("Einkaufsliste.csv")
	if err != nil {
		log.Fatalln(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		log.Println(scanner.Text())
	}
	err = file.Close()
	if err != nil {
		log.Fatalln(err)
	}
	if scanner.Err() != nil {
		log.Fatalln(scanner.Err())
	}
}
