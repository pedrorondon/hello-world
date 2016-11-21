package main

import "fmt"

func main() {
	days := []string{"Mon", "Tues"}
	dias := []string{"Wed", "Thurs", "Fri"}

	days = append(days, dias...)
	fmt.Println(days)

	days = append(days[:2], days[3:]...)
	fmt.Println(days)
}
