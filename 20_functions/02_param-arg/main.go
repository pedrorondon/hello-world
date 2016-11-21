package main

import "fmt"

func main() {
	nombres("Jane") // these three nombres() are arguments
	nombres("John")
	nombres("Pablo")
}

func nombres(naMe string) { // declaring func with keyword func, no receivers, naming the func greet;
	// one parameter given the variable/identifier of naMe and is type string; no returns;
	// this fuction simply receives an argument and assigns it to the var/identifier naMe
	fmt.Println(naMe)
}

// nombres is declared with a parameter
// when calling a function in this case func nombres, you pass in an argument(s)

// functions allow code to be placed in containers that can later be used over and over if needed

// nombres - variable since there are no parens ()
// nombres() - function since there are parens ()
