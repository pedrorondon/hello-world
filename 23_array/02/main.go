package main

import "fmt"

func main() {
	var x [58]string
	// declaring a variable x which is of type array that stores a string and is of length 58
	fmt.Println(x) // the zero value of a string is an empty string; the code outputs emptiness
	// [                                            ] is the output
	fmt.Println(len(x)) // length of the array
	// outputs length of 58
	fmt.Println(x[42]) // access an item at index position 42; zero based index, meaning that 0 gives first
	// item, 1 gives second item, etc.
	// outputs emptiness since nothing is at that position
	for i := 65; i <= 122; i++ { // here is a loop
		x[i-65] = string(i)
	}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[42])
}
