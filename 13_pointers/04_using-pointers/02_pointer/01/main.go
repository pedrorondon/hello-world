package main

import "fmt"

func zero(z *int) {
	*z = 0
	fmt.Println(z, "a") // added this line of code as test
}

func main() {
	x := 5
	zero(&x)
	fmt.Println(x, "b") // x is 0 since it's memory address was dereferenced and assigned a new
	// value through func zero; x has been changed forever from 5 to 0
}
