package main

import "fmt"

func main() {
	greeting := func() { // func() is an anonymous function with no name and it is being assigned
		// to a variable, greeting; when you assign a function to a variable, that is known
		// as a func expression
		fmt.Println("Hola Mundo!")
	}
	greeting()
}
