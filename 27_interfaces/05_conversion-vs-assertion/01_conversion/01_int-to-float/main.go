package main

import "fmt"

func main() {
	var x = 12                  // x is an int
	var y = 12.1230123          // y is a float64
	fmt.Println(y + float64(x)) // converts x from int to float64; this is a widening conversion
	// since a float64 allows for decimal places while int does not

}
