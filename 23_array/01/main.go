package main

import "fmt"

func main() {
	var x [58]int
	// declaring a variable x which is of type array that stores an int and is of length 58
	fmt.Println(x)
	fmt.Println(len(x)) // length of the array
	fmt.Println(x[42])  // access an item at index position 42; zero based index, meaning that 0 gives first
	// item, 1 gives second item, etc.
	x[42] = 333
	fmt.Println(x[42])
}
