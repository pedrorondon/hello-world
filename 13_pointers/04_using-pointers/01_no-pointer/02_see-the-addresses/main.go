package main

import "fmt"

func zero(z int) {
	fmt.Printf("%p\n", &z) // address in function zero
	fmt.Println(&z, "a")   // alternate method to display address in function zero

	z = 0
}

func main() {
	x := 5
	fmt.Printf("%p\n", &x) // address in main
	fmt.Println(&x, "b")   // alternate method to display address in main
	zero(x)

	fmt.Println(x, "c") // x is still 5
}

// not the order when running the code: b, a, c
