package main

import "fmt"

const ctof float64 = 1.8

func main() {
	var temp float64
	fmt.Print("Please enter temperature in Celcius: ")
	fmt.Scan(&temp)
	fahren := temp*ctof + 32
	fmt.Println(temp, "degrees Celcius is", fahren, "degrees Fahrenheit.")
}
