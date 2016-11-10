package main

import "fmt"

func main() {
	var n int
	fmt.Print("Please enter a number: ")
	fmt.Scan(&n)

	x := n % 2

	if x == 1 {
		fmt.Println("ODD")

	} else {
		fmt.Println("EVEN")
	}
}
