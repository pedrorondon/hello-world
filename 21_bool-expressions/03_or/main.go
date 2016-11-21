package main

import "fmt"

func main() {
	if true || false { // || is known as double pipe
		fmt.Println("This ran")
	}
}

// an expression produces a value and is read horizontally
// a statement performs an action and is read vertically; statements are often made up of expressions
// but not vice versa
