package main

import "fmt"

func makeGreeter() func() string {
	return func() string {
		return "Hola Mundo!"
	}
}
func main() {
	greet := makeGreeter()
	fmt.Println(greet())
}
