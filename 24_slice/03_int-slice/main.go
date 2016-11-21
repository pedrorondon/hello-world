package main

import "fmt"

func main() {

	mySlice := []int{1, 3, 5, 7, 9, 11}
	for pos, val := range mySlice { // use the range loop to loop over our slice
		// pos and val are simply variables; range outputs two numbers:
		// the positon/index and the value assigned to the position/index within the slice
		fmt.Println(pos, "-", val)
	}
}
