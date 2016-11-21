package main

import "fmt"

func makeGreeter() func() string { // func makeGreeter is a func() string, a function that returns a string
	return func() string {
		return "Hola"
	}
}
func main() {
	greet := makeGreeter()
	fmt.Println(greet())
	fmt.Printf("%T \n", greet)
}
