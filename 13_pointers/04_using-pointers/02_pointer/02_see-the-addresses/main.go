package main

import "fmt"

func zero(z *int) {
	fmt.Println(z, "a") // if & were added in fron of z, a memory address for x's memory addres would be displayed
	*z = 0
}

func main() {
	x := 5 // the memory address for x is passed up to func zero as proven in line 6
	fmt.Println(&x, "b")
	zero(&x)
	fmt.Println(x, "c")
}

// notice order of output: b, a, c
