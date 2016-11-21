package main

import "fmt"

func hola() {
	fmt.Print("hello ")

}

func mundo() {
	fmt.Println("world")
}

func main() {
	defer mundo() // defer delays the execution of a function until right before exiting func main;
	// a good use of defer would be to close a file that was opened previously but keeping
	// the lines of code to open and close tightly together
	hola()
}
