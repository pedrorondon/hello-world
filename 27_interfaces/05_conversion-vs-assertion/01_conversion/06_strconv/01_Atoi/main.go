package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x = "12" // x is in ASCII
	var y = 6

	z, _ := strconv.Atoi(x) // converts ASCII to int; z, _ is used since this a multivalue return
	fmt.Println(z + y)
	fmt.Println(z)
}
