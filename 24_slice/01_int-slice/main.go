package main

import "fmt"

func main() { // () are parens; {} are braces
	mySlice := []int{1, 3, 5, 7, 9, 11} // [] are brackets indicating a slice; of type slice of int
	fmt.Printf("%T\n", mySlice)         // []int
	fmt.Println(mySlice)                // [1 3 5 7 9 11]
}

// a slice is simply a list
