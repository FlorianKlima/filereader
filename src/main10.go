package datafile

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter a Grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println(input)
}
