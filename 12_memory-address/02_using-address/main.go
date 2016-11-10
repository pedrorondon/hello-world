package main

import "fmt"

const meterstoyards float64 = 1.09362

func main() {
	var meters float64 // this line of code creates a variable, meters, and its type, float 64
	fmt.Print("Enter meters swam: ")
	fmt.Scan(&meters)
	yards := meters * meterstoyards
	fmt.Println(meters, " meters is ", yards, " yards. ")
}
