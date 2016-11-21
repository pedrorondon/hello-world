package main

import "fmt"

func main() {
	greeting := func() {
		fmt.Println("Hola Mundo!")
	}
	greeting()
	fmt.Printf("%T \n", greeting)
}
