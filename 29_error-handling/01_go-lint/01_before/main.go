package main

import "fmt"

func main() {
	x := 33
	str := eval(x)
	fmt.Println(str)
}

func eval(n int) string {
	if n > 10 {
		return fmt.Sprint("x is greater than 10")
	} else { // this else statement is not necessary; it is just extra verbage that is polluting
		// the code
		return fmt.Sprint("x is not greater than 10")
	}
}

// use golint ./... means this current direction and all of the files folders in it;
// from whereever you are at, it will lint everything from that point down in the directory
// structure recursively

// golint often gives a more readable and understandable description of the error
