package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 12
	y := "I have this many: " + strconv.Itoa(x)
	fmt.Println(y)
	// fmt.Println("I have this many: ", strconv.Itoa(x), x) shows that the conversion from
	// int to ASCII can be done in the print line
}
