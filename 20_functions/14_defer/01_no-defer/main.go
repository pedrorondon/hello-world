package main

import "fmt"

func hola() {
	fmt.Print("hello ")
}

func mundo() {
	fmt.Println("world")
}

func main() {
	mundo() // the word func is not needed to call the function within the current function
	hola()
}
