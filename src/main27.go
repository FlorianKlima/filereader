package datafile

import (
	"dates"
	"fmt"
)

func main() {
	days := 3
	fmt.Println("Your appointment is in", days, "days")
	fmt.Println("With a follow-up in:", days+dates.DaysInWeek, "days")
}
