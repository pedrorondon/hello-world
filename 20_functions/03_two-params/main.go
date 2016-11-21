package main

import "fmt"

func main() {
	nombres("Juan", "Pablo")
}

func nombres(a string, b string) { // can also be written as func nombres(a,b string) {
	fmt.Println(a, b)
}
